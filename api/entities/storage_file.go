package entities

import (
	"crypto/md5"
	"fmt"
	"github.com/famiphoto/famiphoto/api/utils/array"
	"strings"
)

type StorageFileInfo struct {
	Name  string
	Path  string
	Ext   string
	IsDir bool
}

func (f StorageFileInfo) IsMatchExt(extensions []string) bool {
	return array.IsContain(strings.ToLower(f.Ext), array.Map(extensions, strings.ToLower))
}

func (f StorageFileInfo) FilePathHash() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(f.Path)))
}

type StorageFileData []byte

func (e StorageFileData) FileHash() string {
	return fmt.Sprintf("%x", md5.Sum(e))
}
