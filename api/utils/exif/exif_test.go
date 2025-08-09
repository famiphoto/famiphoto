package exif

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"time"
)

func TestParseDatetime(t *testing.T) {
	// 基本のEXIFフォーマットとオフセットの組み合わせを検証する
	dateStr := "2023:11:24 14:49:38"

	t.Run("有効なオフセット +09:00 (FixedZone)", func(t *testing.T) {
		tm, err := ParseDatetime(dateStr, "+09:00")
		assert.NoError(t, err)
		assert.Equal(t, 2023, tm.Year())
		assert.Equal(t, time.November, tm.Month())
		assert.Equal(t, 24, tm.Day())
		assert.Equal(t, 14, tm.Hour())
		assert.Equal(t, 49, tm.Minute())
		assert.Equal(t, 38, tm.Second())
		name, offset := tm.Zone()
		assert.Equal(t, "+09:00", name)
		assert.Equal(t, 9*3600, offset)
		assert.Equal(t, "+09:00", tm.Location().String())
	})

	t.Run("有効なオフセット -05:30 (FixedZone)", func(t *testing.T) {
		tm, err := ParseDatetime(dateStr, "-05:30")
		assert.NoError(t, err)
		name, offset := tm.Zone()
		assert.Equal(t, "-05:30", name)
		assert.Equal(t, -(5*3600+30*60), offset)
		assert.Equal(t, "-05:30", tm.Location().String())
	})

	t.Run("空のオフセットはデフォルトの Asia/Tokyo を使用", func(t *testing.T) {
		tm, err := ParseDatetime(dateStr, "")
		assert.NoError(t, err)
		// Asia/Tokyo は通常 JST(+09:00)
		name, offset := tm.Zone()
		assert.Equal(t, 9*3600, offset)
		assert.Equal(t, "Asia/Tokyo", tm.Location().String())
		// Zone名は"JST"になるはず（環境依存のため厳密には offset と Location を優先）
		assert.Equal(t, "JST", name)
	})

	t.Run("不正なオフセットはデフォルトの Asia/Tokyo にフォールバック", func(t *testing.T) {
		tm, err := ParseDatetime(dateStr, "aa:bb")
		assert.NoError(t, err)
		_, offset := tm.Zone()
		assert.Equal(t, 9*3600, offset)
		assert.Equal(t, "Asia/Tokyo", tm.Location().String())
	})

	t.Run("不正な日時フォーマットはエラーを返す", func(t *testing.T) {
		_, err := ParseDatetime("2023-11-24 14:49:38", "+09:00")
		assert.Error(t, err)
	})
}

func TestExifItemList(t *testing.T) {
	f, err := os.Open("../../testing/resources/exif_tester.jpg")
	defer f.Close()
	assert.NoError(t, err)
	data, err := io.ReadAll(f)
	assert.NoError(t, err)

	exifData, err := ParseExifItemsAll(data)
	assert.NoError(t, err)

	t.Run("メーカー (Make)", func(t *testing.T) {
		actual := exifData.Make()
		expected := "SONY"
		assert.Equal(t, expected, actual)
	})

	t.Run("モデル (Model)", func(t *testing.T) {
		actual := exifData.Model()
		expected := "ILCE-7CM2"
		assert.Equal(t, expected, actual)
	})

	t.Run("シリアル番号 (SerialNumber)", func(t *testing.T) {
		actual := exifData.SerialNumber()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("F値 (FNumber)", func(t *testing.T) {
		actual := exifData.FNumber()
		expected := float64(5)
		assert.Equal(t, expected, actual)
	})

	t.Run("撮影日時 (DateTimeOriginal)", func(t *testing.T) {
		actual := exifData.DateTimeOriginal()
		expected := "2023:11:24 14:49:38"
		assert.Equal(t, expected, actual)
	})

	t.Run("デジタイズ日時 (DateTimeDigitized)", func(t *testing.T) {
		actual := exifData.DateTimeDigitized()
		expected := "2023:11:24 14:49:38"
		assert.Equal(t, expected, actual)
	})

	t.Run("作成日時 (CreateDate)", func(t *testing.T) {
		actual := exifData.CreateDate()
		expected := "2023:11:24 14:49:38"
		assert.Equal(t, expected, actual)
	})

	t.Run("サブ秒 (SubsecTimeOriginal)", func(t *testing.T) {
		actual := exifData.SubsecTimeOriginal()
		expected := "247"
		assert.Equal(t, expected, actual)
	})

	t.Run("タイムゾーンオフセット (TimezoneOffset)", func(t *testing.T) {
		actual := exifData.TimezoneOffset()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("露出時間 (ExposureTime)", func(t *testing.T) {
		actual := exifData.ExposureTime()
		expected := float64(0.0015625)
		assert.Equal(t, expected, actual)
	})

	t.Run("ISO感度 (ISO)", func(t *testing.T) {
		actual := exifData.ISO()
		expected := int64(100)
		assert.Equal(t, expected, actual)
	})

	t.Run("焦点距離 (FocalLength)", func(t *testing.T) {
		actual := exifData.FocalLength()
		expected := float64(24)
		assert.Equal(t, expected, actual)
	})

	t.Run("35mm換算焦点距離 (FocalLengthIn35mm)", func(t *testing.T) {
		actual := exifData.FocalLengthIn35mm()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("露出プログラム (ExposureProgram)", func(t *testing.T) {
		actual := exifData.ExposureProgram()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("露出補正 (ExposureCompensation)", func(t *testing.T) {
		actual := exifData.ExposureCompensation()
		expected := float64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("測光モード (MeteringMode)", func(t *testing.T) {
		actual := exifData.MeteringMode()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("フラッシュ (Flash)", func(t *testing.T) {
		actual := exifData.Flash()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("レンズメーカー (LensMake)", func(t *testing.T) {
		actual := exifData.LensMake()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("レンズモデル (LensModel)", func(t *testing.T) {
		actual := exifData.LensModel()
		expected := "FE 24-70mm F2.8 GM II"
		assert.Equal(t, expected, actual)
	})

	t.Run("レンズシリアル番号 (LensSerialNumber)", func(t *testing.T) {
		actual := exifData.LensSerialNumber()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("幅 (Width)", func(t *testing.T) {
		actual := exifData.Width()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("高さ (Height)", func(t *testing.T) {
		actual := exifData.Height()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("色空間 (ColorSpace)", func(t *testing.T) {
		actual := exifData.ColorSpace()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("ホワイトバランス (WhiteBalance)", func(t *testing.T) {
		actual := exifData.WhiteBalance()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("回転情報 (Orientation)", func(t *testing.T) {
		actual := exifData.Orientation()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("ソフトウェア (Software)", func(t *testing.T) {
		actual := exifData.Software()
		expected := "ILCE-7CM2 v1.01"
		assert.Equal(t, expected, actual)
	})

	t.Run("ファームウェア (Firmware)", func(t *testing.T) {
		actual := exifData.Firmware()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("オフセット時間 (OffsetTimeOriginal)", func(t *testing.T) {
		actual := exifData.OffsetTimeOriginal()
		expected := "+09:00"
		assert.Equal(t, expected, actual)
	})
}
