package adapters

import (
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"path"
	"path/filepath"
)

type PhotoStorageAdapter interface {
	OpenPhoto(filePath string) (entities.StorageFileData, error)
	ReadDir(dirPath string) (entities.StorageFileInfoList, error)
	SavePreviewImage(photoID string, data []byte) error
	SaveThumbnailImage(photoID string, data []byte) error
}

func NewPhotoStorageAdapter(photoStorageRepo repositories.PhotoStorageRepository) PhotoStorageAdapter {
	return &photoStorageAdapter{
		assetRootPath:    config.Env.AssetRootPath,
		photoStorageRepo: photoStorageRepo,
	}
}

type photoStorageAdapter struct {
	assetRootPath    string
	photoStorageRepo repositories.PhotoStorageRepository
}

func (a *photoStorageAdapter) OpenPhoto(filePath string) (entities.StorageFileData, error) {
	return a.photoStorageRepo.ReadFile(filePath)
}

func (a *photoStorageAdapter) ReadDir(dirPath string) (entities.StorageFileInfoList, error) {
	list, err := a.photoStorageRepo.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	files := make([]*entities.StorageFileInfo, len(list))
	for i, v := range list {
		files[i] = &entities.StorageFileInfo{
			Name:  v.Name(),
			Path:  path.Join(dirPath, v.Name()),
			Ext:   filepath.Ext(v.Name()),
			IsDir: v.IsDir(),
		}
	}
	return files, nil
}

func (a *photoStorageAdapter) SavePreviewImage(photoID string, data []byte) error {
	filePath := path.Join(a.assetRootPath, "previews", photoID)
	_, err := a.photoStorageRepo.SaveContent(filePath, data)
	if err != nil {
		return err
	}
	return nil
}

func (a *photoStorageAdapter) SaveThumbnailImage(photoID string, data []byte) error {
	filePath := path.Join(a.assetRootPath, "thumbnail", photoID)
	_, err := a.photoStorageRepo.SaveContent(filePath, data)
	if err != nil {
		return err
	}
	return nil
}
