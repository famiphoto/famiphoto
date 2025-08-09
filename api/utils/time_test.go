package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// LocationFromOffset と LocationOrDefaultFromOffset の単体テスト
// それぞれ複数のパターンを t.Run で分割して検証します。
// 値の比較には testify/assert を使用します。

func TestLocationFromOffset(t *testing.T) {
	// 正常系: 正のオフセット（+09:00）
	// 期待: ロケーション名が+09:00で、UTCから+32400秒のオフセット
	t.Run("正のオフセット +09:00", func(t *testing.T) {
		loc, err := LocationFromOffset("+09:00")
		assert.NoError(t, err)
		name, offset := time.Date(2000, 1, 1, 0, 0, 0, 0, loc).Zone()
		assert.Equal(t, "+09:00", loc.String(), "ロケーション名")
		assert.Equal(t, "+09:00", name, "Zone() が返す名前")
		assert.Equal(t, 9*3600, offset, "UTCオフセット秒")
	})

	// 正常系: 符号なし（09:00）
	// 期待: +09:00 と等価に扱われる
	t.Run("符号なし 09:00 は +09:00 と等価", func(t *testing.T) {
		loc, err := LocationFromOffset("09:00")
		assert.NoError(t, err)
		_, offset := time.Date(2000, 1, 1, 0, 0, 0, 0, loc).Zone()
		assert.Equal(t, "+09:00", loc.String())
		assert.Equal(t, 9*3600, offset)
	})

	// 正常系: 負のオフセット（-05:30）
	// 期待: ロケーション名が-05:30で、UTCから-19800秒のオフセット
	t.Run("負のオフセット -05:30", func(t *testing.T) {
		loc, err := LocationFromOffset("-05:30")
		assert.NoError(t, err)
		name, offset := time.Date(2000, 1, 1, 0, 0, 0, 0, loc).Zone()
		assert.Equal(t, "-05:30", loc.String())
		assert.Equal(t, "-05:30", name)
		assert.Equal(t, -(5*3600+30*60), offset)
	})

	// 異常系: コロンがない（フォーマット不正）
	// 期待: エラー
	t.Run("不正フォーマット: '09'", func(t *testing.T) {
		loc, err := LocationFromOffset("09")
		assert.Error(t, err)
		assert.Nil(t, loc)
	})

	// 異常系: 数字でない
	// 期待: エラー
	t.Run("不正フォーマット: 'aa:bb'", func(t *testing.T) {
		loc, err := LocationFromOffset("aa:bb")
		assert.Error(t, err)
		assert.Nil(t, loc)
	})

	// 異常系: 空文字や符号のみ
	// 期待: エラー
	t.Run("不正フォーマット: 空や符号のみ", func(t *testing.T) {
		for _, s := range []string{"", "+", "-", ":"} {
			loc, err := LocationFromOffset(s)
			assert.Error(t, err, s)
			assert.Nil(t, loc, s)
		}
	})
}

func TestLocationOrDefaultFromOffset(t *testing.T) {
	// デフォルトロケーションを用意
	defaultLoc := time.FixedZone("DEFAULT", 0)

	// 正常系: 正しいオフセットの場合はそのロケーションが返る
	t.Run("有効なオフセットならデフォルトではなく計算結果を返す", func(t *testing.T) {
		loc := LocationOrDefaultFromOffset("+09:00", defaultLoc)
		assert.NotNil(t, loc)
		assert.NotEqual(t, defaultLoc, loc, "デフォルトではないこと")
		name, offset := time.Date(2000, 1, 1, 0, 0, 0, 0, loc).Zone()
		assert.Equal(t, "+09:00", loc.String())
		assert.Equal(t, "+09:00", name)
		assert.Equal(t, 9*3600, offset)
	})

	// 異常系: 不正なフォーマットの場合はデフォルトを返す
	t.Run("不正なオフセットならデフォルトを返す", func(t *testing.T) {
		for _, s := range []string{"", "+", "-", "09", ":", "aa:bb"} {
			loc := LocationOrDefaultFromOffset(s, defaultLoc)
			assert.Equal(t, defaultLoc, loc, s)
		}
	})
}
