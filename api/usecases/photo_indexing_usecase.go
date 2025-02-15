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
	photoFiles          chan *entities.StorageFileInfo
	isFinishSearching   bool
	photoStorageAdapter adapters.PhotoStorageAdapter
	photoIndexService   services.PhotoIndexService
	mutex               sync.Mutex
}

func (u *photoIndexingUseCase) IndexPhotos(ctx context.Context, extensions []string, maxParallels int64) error {
	u.photoFiles = make(chan *entities.StorageFileInfo, maxParallels)
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
	close(u.photoFiles)
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
			fmt.Println("enqueue", c.Path)
			u.photoFiles <- c
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

		photoFile := <-u.photoFiles
		wg.Add(1)

		go func(pf *entities.StorageFileInfo) {
			defer func() {
				wg.Done()
			}()

			if err := u.registerPhoto(ctx, pf); err != nil {
				log.Error(err)
			}

			// 登録処理
			fmt.Println("process", pf.Path)
		}(photoFile)

	}
}

func (u *photoIndexingUseCase) registerPhoto(ctx context.Context, pf *entities.StorageFileInfo) error {
	photo, meta, err := u.photoIndexService.RegisterPhotoToMasterData(ctx, pf)
	if err != nil {
		return err
	}
	err = u.photoIndexService.RegisterPhotoToSearchEngine(ctx, photo, meta)
	if err != nil {
		return err
	}
	return nil
}
