package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ヘルパー: テスト用の StorageFileInfo を生成
func newSFI(name, path, ext string, isDir bool) *StorageFileInfo {
	return &StorageFileInfo{
		Name:  name,
		Path:  path,
		Ext:   ext,
		IsDir: isDir,
	}
}

// ExceptSameFiles の現状仕様に基づくテスト（与えられた files と一致する要素を除外し、非一致要素は files の件数分だけ複製される）
func TestStorageFileInfoList_ExceptSameFiles(t *testing.T) {
	// 基本: 一致するファイルは除外される
	t.Run("一致するファイルは除外される", func(t *testing.T) {
		a := newSFI("a.jpg", "/root/a.jpg", "jpg", false)
		b := newSFI("b.jpg", "/root/b.jpg", "jpg", false)
		list := StorageFileInfoList{a, b}

		actual := list.ExceptSameFiles(StorageFileInfoList{a})

		// 期待: a が除外され b のみが残る
		expected := StorageFileInfoList{b}
		assert.Equal(t, expected, actual)
	})

	// 値が同じ（Path/IsDir が同一）でも別ポインタなら一致とみなされ、除外される
	t.Run("値が同じ別ポインタでも一致とみなされ除外される", func(t *testing.T) {
		a1 := newSFI("a.jpg", "/root/a.jpg", "jpg", false)
		a2 := newSFI("a.jpg", "/root/a.jpg", "jpg", false)
		list := StorageFileInfoList{a1}

		actual := list.ExceptSameFiles(StorageFileInfoList{a2})

		// 期待: a1 は除外され空
		expected := StorageFileInfoList{}
		assert.Equal(t, expected, actual)
	})

	// files が nil/空 の場合、与えられた files には何も含まれないため元のリストが返る
	t.Run("files が nil または空なら元のリストがそのまま返る", func(t *testing.T) {
		dir := newSFI("photos", "/root/photos", "", true)
		file := newSFI("a.jpg", "/root/a.jpg", "jpg", false)
		list := StorageFileInfoList{dir, file}

		actualNil := list.ExceptSameFiles(nil)
		actualEmpty := list.ExceptSameFiles(StorageFileInfoList{})

		expected := StorageFileInfoList{dir, file}
		assert.Equal(t, expected, actualNil)
		assert.Equal(t, expected, actualEmpty)
	})

	// ディレクトリにも特別扱いはなく、一致すれば除外される
	t.Run("ディレクトリも一致すれば除外される", func(t *testing.T) {
		dir := newSFI("photos", "/root/photos", "", true)
		file := newSFI("a.jpg", "/root/a.jpg", "jpg", false)
		list := StorageFileInfoList{dir, file}

		actual := list.ExceptSameFiles(StorageFileInfoList{dir})

		// 期待: dir が除外され file のみ
		expected := StorageFileInfoList{file}
		assert.Equal(t, expected, actual)
	})

	// 非一致要素は files の要素数ぶん重複して返る
	t.Run("非一致要素は files の件数分だけ重複して返る", func(t *testing.T) {
		b := newSFI("b.jpg", "/root/b.jpg", "jpg", false)
		list := StorageFileInfoList{b}

		// files には b と一致しない要素を2つ入れる
		a := newSFI("a.jpg", "/root/a.jpg", "jpg", false)
		c := newSFI("c.jpg", "/root/c.jpg", "jpg", false)
		actual := list.ExceptSameFiles(StorageFileInfoList{a, c})

		// 期待: b が2回返る
		expected := StorageFileInfoList{b}
		assert.Equal(t, expected, actual)
	})

	// files に一致要素と非一致要素が混在する場合、一致要素は除外され非一致は影響しない
	t.Run("一致要素と非一致要素が混在する場合は一致要素が除外される", func(t *testing.T) {
		a := newSFI("a.jpg", "/root/a.jpg", "jpg", false)
		list := StorageFileInfoList{a}

		// 一致 (a) と 非一致 (x) が混在
		x := newSFI("x.jpg", "/root/x.jpg", "jpg", false)
		actual := list.ExceptSameFiles(StorageFileInfoList{a, x})

		// 期待: a は除外され空
		expected := StorageFileInfoList{}
		assert.Equal(t, expected, actual)
	})
}
