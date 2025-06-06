package usecases

import (
	"context"
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/services"
	"github.com/labstack/gommon/log"
	"sync"
)

type PhotoIndexingUseCase interface {
	IndexPhotos(ctx context.Context, extensions []string, maxParallels int64) error
}

func NewPhotoIndexingUseCase(photoStorageAdapter adapters.PhotoStorageAdapter, photoIndexService services.PhotoIndexService) PhotoIndexingUseCase {
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
	contents, err := u.photoStorageAdapter.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, c := range contents {
		if c.IsDir {
			return u.findPhotosRecursive(ctx, c.Path, extensions)
		} else if c.IsMatchExt(extensions) {
			// 同じファイル名で拡張子違いを探す
			sameFiles := contents.FilterSameNameFiles(c, extensions)
			fmt.Println("enqueue", c.Path)
			contents = contents.ExceptSameFiles(sameFiles)
			u.photoFileSet <- sameFiles
		}
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

			// 登録処理
			fmt.Println("process", pfList[0].NameExceptExt())
		}(photoFiles)

	}
}

func (u *photoIndexingUseCase) registerPhoto(ctx context.Context, pfList entities.StorageFileInfoList) error {
	list := make([]string, len(pfList))
	for i, v := range pfList {
		list[i] = v.Path
	}

	// TODO サムネイル画像、プレビュー画像の登録

	// データベースへの登録
	photoID, err := u.photoIndexService.RegisterPhotoToMasterData(ctx, pfList)
	if err != nil {
		return err
	}

	// 検索エンジンへの登録
	err = u.photoIndexService.RegisterPhotoToSearchEngine(ctx, photoID)
	if err != nil {
		return err
	}

	return nil
}
