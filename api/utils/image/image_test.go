package image

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestResizeJPEG(t *testing.T) {
	imageFile, err := os.Open("../../testing/resources/test_image001.jpg")
	if err != nil {
		t.Fatal("Failed to os.Open", err)
	}
	defer imageFile.Close()

	data, err := ioutil.ReadAll(imageFile)
	if err != nil {
		t.Fatal("Failed to ioutil.ReadAll", err)
	}

	actual, err := ResizeJPEG(data, 600, 400)
	if err != nil {
		t.Fatal(err)
	}

	exceptFile, err := os.Open("../../testing/resources/test_image001_dst.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer exceptFile.Close()

	excepted, err := ioutil.ReadAll(exceptFile)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, excepted, actual)
}

func TestGetSize(t *testing.T) {
	imageFile, err := os.Open("../../testing/resources/test_image001.jpg")
	if err != nil {
		t.Fatal("Failed to os.Open", err)
	}
	defer imageFile.Close()

	data, err := ioutil.ReadAll(imageFile)
	if err != nil {
		t.Fatal("Failed to ioutil.ReadAll", err)
	}

	actualWdith, actualHeight, err := GetSize(data)
	assert.NoError(t, err)
	assert.Equal(t, int64(2000), actualWdith)
	assert.Equal(t, int64(1333), actualHeight)
}

func TestCalcToResizeWidth(t *testing.T) {
	srcWidth := int64(600)
	srcHeight := int64(400)
	dstWidth := int64(300)

	actual := CalcToResizeWidth(srcWidth, srcHeight, dstWidth)
	expected := int64(200)
	assert.Equal(t, expected, actual)
}

func TestCalcToResizeHeight(t *testing.T) {
	srcWidth := int64(600)
	srcHeight := int64(400)
	dstHeight := int64(200)

	actual := CalcToResizeHeight(srcWidth, srcHeight, dstHeight)
	expected := int64(300)
	assert.Equal(t, expected, actual)
}
