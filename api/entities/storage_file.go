package entities

import (
	"crypto/md5"
	"fmt"
	"github.com/famiphoto/famiphoto/api/utils"
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

func (f StorageFileInfo) FilePathExceptExtHash() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(utils.FileNameExceptExt(f.Path))))
}

func (f StorageFileInfo) NameExceptExt() string {
	return utils.FileNameExceptExt(f.Name)
}

type StorageFileInfoList []*StorageFileInfo

// FilterSameNameFiles 与えられたファイルと拡張子を除いたパスが同じファイル一覧を取得します。
func (l StorageFileInfoList) FilterSameNameFiles(f *StorageFileInfo, extensions []string) StorageFileInfoList {
	dst := make(StorageFileInfoList, 0)
	for _, v := range l {
		if v.IsDir {
			continue
		}
		if v.FilePathExceptExtHash() == f.FilePathExceptExtHash() {
			if v.IsMatchExt(extensions) {
				dst = append(dst, v)
			}
		}
	}
	return dst
}

// ExceptSameFiles 与えられたファイル群に一致するものを除いた配列を返します。
func (l StorageFileInfoList) ExceptSameFiles(files StorageFileInfoList) StorageFileInfoList {
	dst := make(StorageFileInfoList, 0)
	for _, v := range l {
		if v.IsDir {
			dst = append(dst, v)
		}
		if !array.IsContain(v, files) {
			dst = append(dst, v)
		}
	}
	return dst
}

type StorageFileData []byte

func (e StorageFileData) FileHash() string {
	return fmt.Sprintf("%x", md5.Sum(e))
}
