package repositories

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/drivers/storage"
	"github.com/famiphoto/famiphoto/api/errors"
	"os"
	"path"
	"path/filepath"
)

type PhotoStorageRepository interface {
	ReadDir(dirPath string) ([]os.FileInfo, error)
	ReadFile(path string) ([]byte, error)
	GetFileInfo(filePath string) (os.FileInfo, string, error)
	SaveContent(filePath string, data []byte) (os.FileInfo, error)
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

func (r *photoStorageRepository) GetFileInfo(filePath string) (os.FileInfo, string, error) {
	filePath = path.Join(r.baseDir, filePath)
	stat, err := r.driver.Stat(filePath)
	if err != nil {
		return nil, "", errors.New(errors.FileNotFoundError, fmt.Errorf(filePath))
	}

	return stat, filePath, nil
}

func (r *photoStorageRepository) SaveContent(filePath string, data []byte) (os.FileInfo, error) {
	if err := r.createDirIfNotExist(filepath.Dir(filePath)); err != nil {
		return nil, err
	}
	if err := r.driver.CreateFile(filePath, data); err != nil {
		return nil, err
	}

	info, err := r.driver.Stat(filePath)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (r *photoStorageRepository) createDirIfNotExist(p string) error {
	stat, err := r.driver.Stat(p)
	if err != nil && !errors.IsErrCode(err, errors.FileNotFoundError) {
		return err
	}
	if stat == nil {
		return r.driver.CreateDir(p, os.ModePerm)
	}
	if !stat.IsDir() {
		return errors.New(errors.UnExpectedFileAlreadyExistError, nil)
	}
	return nil
}
