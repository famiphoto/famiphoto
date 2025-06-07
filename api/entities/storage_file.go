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

func (l StorageFileInfoList) FilterDirs() StorageFileInfoList {
	dst := make(StorageFileInfoList, 0)
	for _, v := range l {
		if v.IsDir {
			dst = append(dst, v)
		}
	}
	return dst
}

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

// GroupByBaseFileName 拡張子のみが異なるファイル群でグループ化した配列を返します。
// 例: [file1.jpg, file1.arw], [file2.jpg, file2.arw]
func (l StorageFileInfoList) GroupByBaseFileName() []StorageFileInfoList {
	// ディレクトリを除外したファイルのみのリストを作成
	files := make(StorageFileInfoList, 0)
	for _, v := range l {
		if !v.IsDir {
			files = append(files, v)
		}
	}

	// 結果を格納する配列
	result := make([]StorageFileInfoList, 0)
	// 処理済みのファイルを記録するマップ
	processed := make(map[string]bool)

	// 各ファイルについて処理
	for _, file := range files {
		// 既に処理済みのファイルはスキップ
		if processed[file.Path] {
			continue
		}

		// 同じベース名を持つファイルをグループ化
		group := make(StorageFileInfoList, 0)
		basePathHash := file.FilePathExceptExtHash()

		for _, f := range files {
			if f.FilePathExceptExtHash() == basePathHash {
				group = append(group, f)
				processed[f.Path] = true
			}
		}

		// グループが1つ以上のファイルを含む場合のみ結果に追加
		if len(group) > 0 {
			result = append(result, group)
		}
	}

	return result
}

type StorageFileData []byte

func (e StorageFileData) FileHash() string {
	return fmt.Sprintf("%x", md5.Sum(e))
}
