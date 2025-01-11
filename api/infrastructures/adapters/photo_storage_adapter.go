package adapters

import (
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"path"
	"path/filepath"
)

type PhotoStorageAdapter interface {
	OpenPhoto(filePath string) (entities.StorageFileData, error)
	ReadDir(dirPath string) ([]*entities.StorageFileInfo, error)
}

func NewPhotoStorageAdapter(photoStorageRepo repositories.PhotoStorageRepository) PhotoStorageAdapter {
	return &photoStorageAdapter{
		photoStorageRepo: photoStorageRepo,
	}
}

type photoStorageAdapter struct {
	photoStorageRepo repositories.PhotoStorageRepository
}

func (a *photoStorageAdapter) OpenPhoto(filePath string) (entities.StorageFileData, error) {
	return a.photoStorageRepo.ReadFile(filePath)
}

func (a *photoStorageAdapter) ReadDir(dirPath string) ([]*entities.StorageFileInfo, error) {
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
