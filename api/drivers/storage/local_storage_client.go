package storage

import (
	native "errors"
	"github.com/famiphoto/famiphoto/api/errors"
	"os"
)

type Client interface {
	CreateFile(filePath string, data []byte) error
	CreateDir(dirPath string, perm os.FileMode) error
	Rename(old, file string) error
	ReadDir(dirPath string) ([]os.FileInfo, error)
	ReadFile(filePath string) ([]byte, error)
	Glob(pattern string) ([]string, error)
	Exist(filePath string) bool
	Stat(filePath string) (os.FileInfo, error)
	Delete(filePath string) error
	DeleteAll(path string) error
}

func NewLocalStorage() Client {
	return &localStorageDriver{}
}

type localStorageDriver struct {
}

func (d *localStorageDriver) CreateFile(filePath string, data []byte) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

func (d *localStorageDriver) CreateDir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

func (d *localStorageDriver) Rename(old, file string) error {
	return os.Rename(old, file)
}

func (d *localStorageDriver) ReadDir(dirPath string) ([]os.FileInfo, error) {
	res, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	fileInfoList := make([]os.FileInfo, len(res))
	for i, v := range res {
		fileInfoList[i], err = v.Info()
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return nil, err
		}
	}
	return fileInfoList, nil
}

func (d *localStorageDriver) ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (d *localStorageDriver) Delete(filePath string) error {
	return os.Remove(filePath)
}
func (d *localStorageDriver) DeleteAll(p string) error {
	return os.Remove(p)
}

func (d *localStorageDriver) Glob(pattern string) ([]string, error) {
	panic("Not implemented")
}

func (d *localStorageDriver) Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func (d *localStorageDriver) Stat(filePath string) (os.FileInfo, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		if native.Is(err, os.ErrNotExist) {
			return nil, errors.New(errors.FileNotFoundError, err)
		}
		return nil, err
	}
	return stat, nil
}
