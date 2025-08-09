package usecases

import (
	"context"
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/services"
	"github.com/labstack/gommon/log"
	"strings"
	"sync"
)

type PhotoIndexingUseCase interface {
	IndexPhotos(ctx context.Context, extensions []string, maxParallels int64) error
}

func NewPhotoIndexingUseCase(
	photoStorageAdapter adapters.PhotoStorageAdapter,
	photoIndexService services.PhotoIndexService,

) PhotoIndexingUseCase {
	return &photoIndexingUseCase{
		photoStorageAdapter: photoStorageAdapter,
		photoIndexService:   photoIndexService,
	}
}

type photoIndexingUseCase struct {
	photoFileSet        chan entities.StorageFileInfoList
	isFinishSearching   bool
	photoStorageAdapter adapters.PhotoStorageAdapter
	photoIndexService   services.PhotoIndexService
	mutex               sync.Mutex
}

func (u *photoIndexingUseCase) IndexPhotos(ctx context.Context, extensions []string, maxParallels int64) error {

	if err := u.photoIndexService.CreateIndexIfNotExist(ctx); err != nil {
		return err
	}
	log.Info("Create Search Index")

	u.photoFileSet = make(chan entities.StorageFileInfoList, maxParallels)
	u.isFinishSearching = false

	fmt.Println("basePath", config.Env.StorageRootPath)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := u.findPhotosRecursive(ctx, "/", extensions); err != nil {
			panic(err)
		}
		u.mutex.Lock()
		u.isFinishSearching = true
		u.mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		u.indexPhotoProcess(ctx)
	}()
	wg.Wait()
	close(u.photoFileSet)
	fmt.Println("search done")
	return nil
}

func (u *photoIndexingUseCase) findPhotosRecursive(ctx context.Context, dirPath string, extensions []string) error {
	fmt.Println("search:", dirPath)
	contents, err := u.photoStorageAdapter.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, dir := range contents.FilterDirs() {
		if err := u.findPhotosRecursive(ctx, dir.Path, extensions); err != nil {
			return err
		}
	}

	for _, c := range contents {
		if c.IsDir {
			continue
		}
		if !c.IsMatchExt(extensions) {
			continue
		}
		sameFiles := contents.FilterSameNameFiles(c, extensions)
		fileNames := make([]string, len(sameFiles))
		for i, v := range sameFiles {
			fileNames[i] = v.Path
		}
		fmt.Println("enqueue", strings.Join(fileNames, ","))
		contents = contents.ExceptSameFiles(sameFiles)
		u.photoFileSet <- sameFiles
	}

	return nil
}

func (u *photoIndexingUseCase) indexPhotoProcess(ctx context.Context) {
	var wg sync.WaitGroup
	for {
		if u.isFinishSearching {
			wg.Wait()
			break
		}

		photoFiles := <-u.photoFileSet
		wg.Add(1)

		go func(pfList entities.StorageFileInfoList) {
			defer func() {
				wg.Done()
			}()

			if len(pfList) == 0 {
				return
			}

			if err := u.registerPhoto(ctx, pfList); err != nil {
				log.Error(err)
			}
		}(photoFiles)

	}
}

func (u *photoIndexingUseCase) registerPhoto(ctx context.Context, pfList entities.StorageFileInfoList) error {
	if len(pfList) == 0 {
		return nil
	}

	list := make([]string, len(pfList))
	for i, v := range pfList {
		list[i] = v.Path
	}

	// データベースへの登録
	photoID, err := u.photoIndexService.RegisterPhotoToMasterData(ctx, pfList)
	if err != nil {
		return err
	}

	// 検索エンジンへの登録
	err = u.photoIndexService.RegisterPhotoToSearchEngine(ctx, photoID)
	if err != nil {
		log.Error(err)
		return err
	}

	// サムネイル画像、プレビュー画像の登録
	if err := u.photoIndexService.CreatePreviewImages(ctx, photoID); err != nil {
		return err
	}

	// 登録処理
	fmt.Println("process", pfList[0].NameExceptExt())

	return nil
}
