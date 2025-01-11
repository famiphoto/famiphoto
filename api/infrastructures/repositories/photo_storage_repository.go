package repositories

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/drivers/storage"
	"github.com/famiphoto/famiphoto/api/errors"
	"os"
	"path"
)

type PhotoStorageRepository interface {
	ReadDir(dirPath string) ([]os.FileInfo, error)
	ReadFile(path string) ([]byte, error)
}

func NewPhotoStorageRepository(driver storage.Client) PhotoStorageRepository {
	return &photoStorageRepository{
		driver:  driver,
		baseDir: config.Env.StorageRootPath,
	}
}

type photoStorageRepository struct {
	driver  storage.Client
	baseDir string
}

func (r *photoStorageRepository) ReadDir(dirPath string) ([]os.FileInfo, error) {
	return r.driver.ReadDir(path.Join(r.baseDir, dirPath))
}

func (r *photoStorageRepository) ReadFile(filePath string) ([]byte, error) {
	filePath = path.Join(r.baseDir, filePath)
	if exist := r.driver.Exist(filePath); !exist {
		return nil, errors.New(errors.FileNotFoundError, fmt.Errorf(filePath))
	}
	return r.driver.ReadFile(filePath)
}
