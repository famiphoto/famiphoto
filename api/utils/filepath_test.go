package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// FileNameExceptExt のテスト
// テストパターン毎に t.Run で分割し、assert を使って検証します。
func TestFileNameExceptExt(t *testing.T) {
	// 拡張子ありのシンプルなファイル名
	t.Run("拡張子ありのシンプルなファイル名", func(t *testing.T) {
		actual := FileNameExceptExt("name.jpg")
		assert.Equal(t, "name", actual)
	})

	// ディレクトリを含むパス
	t.Run("ディレクトリを含むパス", func(t *testing.T) {
		actual := FileNameExceptExt("dir/subdir/name.png")
		assert.Equal(t, "name", actual)
	})

	// 複数ドットを含むファイル名
	t.Run("複数ドットを含むファイル名", func(t *testing.T) {
		actual := FileNameExceptExt("archive.tar.gz")
		assert.Equal(t, "archive.tar", actual)
	})

	// 拡張子なし
	t.Run("拡張子なし", func(t *testing.T) {
		actual := FileNameExceptExt("filename")
		assert.Equal(t, "filename", actual)
	})

	// ドット始まりの隠しファイル(拡張子なし)
	t.Run("ドット始まりの隠しファイル(拡張子なし)", func(t *testing.T) {
		actual := FileNameExceptExt(".gitignore")
		assert.Equal(t, "", actual)
	})

	// ドット始まりで拡張子様の接尾辞を持つ
	t.Run("ドット始まりで拡張子様の接尾辞を持つ", func(t *testing.T) {
		actual := FileNameExceptExt(".env.local")
		assert.Equal(t, ".env", actual)
	})

	// 末尾スラッシュ付きのディレクトリパス
	t.Run("末尾スラッシュ付きのディレクトリパス", func(t *testing.T) {
		actual := FileNameExceptExt("path/to/dir/")
		assert.Equal(t, "dir", actual)
	})
}
