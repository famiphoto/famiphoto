package exif

import (
	"github.com/famiphoto/famiphoto/api/utils"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
	"time"
)

func TestParseDatetime(t *testing.T) {
	actual, err := ParseDatetime("2022:07:19 21:53:00", "Asia/Tokyo")
	assert.NoError(t, err)
	expected := time.Date(2022, 7, 19, 21, 53, 0, 0, utils.MustLoadLocation("Asia/Tokyo"))
	assert.Equal(t, expected, actual)
}

func TestExifItemList(t *testing.T) {
	f, err := os.Open("../../testing/resources/exif_tester.jpg")
	defer f.Close()
	assert.NoError(t, err)
	data, err := io.ReadAll(f)
	assert.NoError(t, err)

	exifData, err := ParseExifItemsAll(data)
	assert.NoError(t, err)

	t.Run("Make", func(t *testing.T) {
		actual := exifData.Make()
		expected := "SONY"
		assert.Equal(t, expected, actual)
	})

	t.Run("Model", func(t *testing.T) {
		actual := exifData.Model()
		expected := "ILCE-7CM2"
		assert.Equal(t, expected, actual)
	})

	t.Run("SerialNumber", func(t *testing.T) {
		actual := exifData.SerialNumber()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("FNumber", func(t *testing.T) {
		actual := exifData.FNumber()
		expected := float64(5)
		assert.Equal(t, expected, actual)
	})

	t.Run("DateTimeOriginal", func(t *testing.T) {
		actual := exifData.DateTimeOriginal()
		expected := "2023:11:24 14:49:38"
		assert.Equal(t, expected, actual)
	})

	t.Run("DateTimeDigitized", func(t *testing.T) {
		actual := exifData.DateTimeDigitized()
		expected := "2023:11:24 14:49:38"
		assert.Equal(t, expected, actual)
	})

	t.Run("CreateDate", func(t *testing.T) {
		actual := exifData.CreateDate()
		expected := "2023:11:24 14:49:38"
		assert.Equal(t, expected, actual)
	})

	t.Run("SubsecTimeOriginal", func(t *testing.T) {
		actual := exifData.SubsecTimeOriginal()
		expected := "247"
		assert.Equal(t, expected, actual)
	})

	t.Run("TimezoneOffset", func(t *testing.T) {
		actual := exifData.TimezoneOffset()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("ExposureTime", func(t *testing.T) {
		actual := exifData.ExposureTime()
		expected := float64(0.0015625)
		assert.Equal(t, expected, actual)
	})

	t.Run("ISO", func(t *testing.T) {
		actual := exifData.ISO()
		expected := int64(100)
		assert.Equal(t, expected, actual)
	})

	t.Run("FocalLength", func(t *testing.T) {
		actual := exifData.FocalLength()
		expected := float64(24)
		assert.Equal(t, expected, actual)
	})

	t.Run("FocalLengthIn35mm", func(t *testing.T) {
		actual := exifData.FocalLengthIn35mm()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("ExposureProgram", func(t *testing.T) {
		actual := exifData.ExposureProgram()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("ExposureCompensation", func(t *testing.T) {
		actual := exifData.ExposureCompensation()
		expected := float64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("MeteringMode", func(t *testing.T) {
		actual := exifData.MeteringMode()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("Flash", func(t *testing.T) {
		actual := exifData.Flash()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("LensMake", func(t *testing.T) {
		actual := exifData.LensMake()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("LensModel", func(t *testing.T) {
		actual := exifData.LensModel()
		expected := "FE 24-70mm F2.8 GM II"
		assert.Equal(t, expected, actual)
	})

	t.Run("LensSerialNumber", func(t *testing.T) {
		actual := exifData.LensSerialNumber()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("Width", func(t *testing.T) {
		actual := exifData.Width()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("Height", func(t *testing.T) {
		actual := exifData.Height()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("ColorSpace", func(t *testing.T) {
		actual := exifData.ColorSpace()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("WhiteBalance", func(t *testing.T) {
		actual := exifData.WhiteBalance()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("Orientation", func(t *testing.T) {
		actual := exifData.Orientation()
		expected := int64(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("Software", func(t *testing.T) {
		actual := exifData.Software()
		expected := "ILCE-7CM2 v1.01"
		assert.Equal(t, expected, actual)
	})

	t.Run("Firmware", func(t *testing.T) {
		actual := exifData.Firmware()
		expected := ""
		assert.Equal(t, expected, actual)
	})

	t.Run("OffsetTimeOriginal", func(t *testing.T) {
		actual := exifData.OffsetTimeOriginal()
		expected := "+09:00"
		assert.Equal(t, expected, actual)
	})
}
