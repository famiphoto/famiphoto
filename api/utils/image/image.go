package image

import (
	"bytes"
	"github.com/disintegration/imaging"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
)

func ResizeJPEG(srcData []byte, width, height int64) ([]byte, error) {
	src, _, err := image.Decode(bytes.NewReader(srcData))
	if err != nil {
		return nil, err
	}

	dst := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	draw.CatmullRom.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	return toJPEG(dst)
}

func GetSize(srcData []byte) (int64, int64, error) {
	src, _, err := image.Decode(bytes.NewReader(srcData))
	if err != nil {
		return 0, 0, err
	}

	bounds := src.Bounds()
	return int64(bounds.Dx()), int64(bounds.Dy()), nil
}

func CalcToResizeWidth[T ~int64 | ~int](srcWidth, srcHeight, dstWidth T) T {
	return T((float64(dstWidth) / float64(srcWidth)) * float64(srcHeight))
}

func CalcToResizeHeight[T ~int64 | ~int](srcWidth, srcHeight, dstHeight T) T {
	return T((float64(dstHeight) / float64(srcHeight)) * float64(srcWidth))
}

func Rotate90JPEG(srcDate []byte) ([]byte, error) {
	src, err := imaging.Decode(bytes.NewReader(srcDate))
	if err != nil {
		return nil, err
	}

	dst := imaging.Rotate90(src)
	return toJPEG(dst)
}

func Rotate180JPEG(srcDate []byte) ([]byte, error) {
	src, err := imaging.Decode(bytes.NewReader(srcDate))
	if err != nil {
		return nil, err
	}

	dst := imaging.Rotate180(src)
	return toJPEG(dst)
}

func Rotate270JPEG(srcDate []byte) ([]byte, error) {
	src, err := imaging.Decode(bytes.NewReader(srcDate))
	if err != nil {
		return nil, err
	}

	dst := imaging.Rotate270(src)
	return toJPEG(dst)
}

func FlipHJPEG(srcDate []byte) ([]byte, error) {
	src, err := imaging.Decode(bytes.NewReader(srcDate))
	if err != nil {
		return nil, err
	}

	dst := imaging.FlipH(src)
	return toJPEG(dst)
}

func FlipVJPEG(srcDate []byte) ([]byte, error) {
	src, err := imaging.Decode(bytes.NewReader(srcDate))
	if err != nil {
		return nil, err
	}

	dst := imaging.FlipV(src)
	return toJPEG(dst)
}

func toJPEG(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
