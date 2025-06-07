package services

import (
	"github.com/famiphoto/famiphoto/api/utils/exif"
	"github.com/famiphoto/famiphoto/api/utils/image"
)

func NewImageProcessService() ImageProcessService {
	return &imageProcessService{}
}

type ImageProcessService interface {
	CreateThumbnail(data []byte, orientation int64) ([]byte, error)
	CreatePreview(data []byte, orientation int64) ([]byte, error)
}

type imageProcessService struct {
}

func (s *imageProcessService) CreateThumbnail(data []byte, orientation int64) ([]byte, error) {
	rotateData, err := s.rotateByOrientation(data, orientation)
	if err != nil {
		return nil, err
	}

	return s.resizeHeight(rotateData, 400)
}

func (s *imageProcessService) CreatePreview(data []byte, orientation int64) ([]byte, error) {
	rotateData, err := s.rotateByOrientation(data, orientation)
	if err != nil {
		return nil, err
	}

	return s.resizeWidth(rotateData, 1920)
}

func (s *imageProcessService) resizeWidth(data []byte, dstWidth int64) ([]byte, error) {
	srcWidth, srcHeight, err := image.GetSize(data)
	if err != nil {
		return nil, err
	}

	thumbData := data
	if dstWidth <= srcWidth {
		dstHeight := image.CalcToResizeWidth(srcWidth, srcHeight, dstWidth)
		thumbData, err = image.ResizeJPEG(data, dstWidth, dstHeight)
		if err != nil {
			return nil, err
		}
	}

	return thumbData, nil
}

func (s *imageProcessService) resizeHeight(data []byte, dstHeight int64) ([]byte, error) {
	srcWidth, srcHeight, err := image.GetSize(data)
	if err != nil {
		return nil, err
	}

	thumbData := data
	if dstHeight <= dstHeight {
		dstWidth := image.CalcToResizeHeight(srcWidth, srcHeight, dstHeight)
		thumbData, err = image.ResizeJPEG(data, dstWidth, dstHeight)
		if err != nil {
			return nil, err
		}
	}

	return thumbData, nil
}

func (s *imageProcessService) rotateByOrientation(data []byte, orientation int64) ([]byte, error) {
	switch orientation {
	case exif.ExifOrientationHorizontal:
		return image.FlipHJPEG(data)
	case exif.ExifOrientationRotate180:
		return image.Rotate180JPEG(data)
	case exif.ExifOrientationVertical:
		return image.FlipVJPEG(data)
	case exif.ExifOrientationHorizontalRotate270:
		dst, err := image.FlipHJPEG(data)
		if err != nil {
			return nil, err
		}
		return image.Rotate90JPEG(dst)
	case exif.ExifOrientationRotate90:
		return image.Rotate270JPEG(data)
	case exif.ExifOrientationHorizontalRotate90:
		dst, err := image.FlipHJPEG(data)
		if err != nil {
			return nil, err
		}
		return image.Rotate270JPEG(dst)
	case exif.ExifOrientationRotate270:
		return image.Rotate90JPEG(data)
	}
	return data, nil
}
