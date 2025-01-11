package usecases

import (
	"context"
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/labstack/gommon/log"
	"sync"
)

type PhotoIndexingUseCase interface {
	IndexPhotos(ctx context.Context, extensions []string, maxParallels int64) error
}

func NewPhotoIndexingUseCase(photoStorageAdapter adapters.PhotoStorageAdapter) PhotoIndexingUseCase {
	return &photoIndexingUseCase{
		photoStorageAdapter: photoStorageAdapter,
	}
}

type photoIndexingUseCase struct {
	photoFiles          chan *entities.StorageFileInfo
	isInProcess         bool
	photoStorageAdapter adapters.PhotoStorageAdapter
	mutex               sync.Mutex
}

func (u *photoIndexingUseCase) IndexPhotos(ctx context.Context, extensions []string, maxParallels int64) error {
	u.photoFiles = make(chan *entities.StorageFileInfo, maxParallels)
	u.isInProcess = true

	fmt.Println("basePath", config.Env.StorageRootPath)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := u.findPhotosRecursive(ctx, "/", extensions); err != nil {
			panic(err)
		}
		u.mutex.Lock()
		u.isInProcess = false
		u.mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		u.indexPhotoProcess()
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
			fmt.Println("append", c.Path)
			u.photoFiles <- c
		}
	}

	return nil
}

func (u *photoIndexingUseCase) indexPhotoProcess() {
	var wg sync.WaitGroup
	for {
		if !u.isInProcess {
			fmt.Println("isInProcess", u.isInProcess)
			wg.Wait()
			break
		}

		photoFile := <-u.photoFiles
		wg.Add(1)

		go func(pf *entities.StorageFileInfo) {
			defer func() {
				wg.Done()
			}()

			data, err := u.photoStorageAdapter.OpenPhoto(pf.Path)
			if err != nil {
				log.Error(err)
			}

			// 登録処理
			fmt.Println("process", pf.Path, data.FileHash())
		}(photoFile)

	}
}
