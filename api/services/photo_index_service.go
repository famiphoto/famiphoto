package services

import (
	"context"
	"fmt"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/utils/exif"
	"time"
)

type PhotoIndexService interface {
	RegisterPhotoToMasterData(ctx context.Context, files entities.StorageFileInfoList) (string, error)
	RegisterPhotoToSearchEngine(ctx context.Context, photoID string) error
	CreatePreviewImages(ctx context.Context, photoID string) error
}

func NewPhotoIndexService(
	photoAdapter adapters.PhotoAdapter,
	photoFileAdapter adapters.PhotoFileAdapter,
	photoStorageAdapter adapters.PhotoStorageAdapter,
	photoSearchAdapter adapters.PhotoSearchAdapter,
	transactionAdapter adapters.TransactionAdapter,
	imageProcessService ImageProcessService,
) PhotoIndexService {
	return &photoIndexService{
		photoAdapter:        photoAdapter,
		photoFileAdapter:    photoFileAdapter,
		photoStorageAdapter: photoStorageAdapter,
		photoSearchAdapter:  photoSearchAdapter,
		transactionAdapter:  transactionAdapter,
		imageProcessService: imageProcessService,
		nowFunc:             time.Now,
	}
}

type photoIndexService struct {
	photoAdapter        adapters.PhotoAdapter
	photoFileAdapter    adapters.PhotoFileAdapter
	photoStorageAdapter adapters.PhotoStorageAdapter
	photoSearchAdapter  adapters.PhotoSearchAdapter
	transactionAdapter  adapters.TransactionAdapter
	imageProcessService ImageProcessService
	nowFunc             func() time.Time
}

func (s *photoIndexService) RegisterPhotoToMasterData(ctx context.Context, files entities.StorageFileInfoList) (string, error) {
	if len(files) == 0 {
		return "", fmt.Errorf("photo files are empty")
	}

	var photoID string
	var err error
	err = s.transactionAdapter.BeginTxn(ctx, func(ctx2 context.Context) error {
		photoID, err = s.photoAdapter.InsertIfNotExist(ctx2, files[0])
		if err != nil {
			return err
		}

		for _, pf := range files {
			data, err := s.photoStorageAdapter.OpenPhoto(pf.Path)
			if err != nil {
				return err
			}

			if _, err := s.photoFileAdapter.Upsert(ctx2, &entities.PhotoFile{
				PhotoID:  photoID,
				FileHash: data.FileHash(),
				File:     *pf,
			}); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return photoID, nil
}

func (s *photoIndexService) RegisterPhotoToSearchEngine(ctx context.Context, photoID string) error {
	if photoID == "" {
		return fmt.Errorf("invalid photoID")
	}

	photoFiles, err := s.photoFileAdapter.FindByPhotoID(ctx, photoID)
	if err != nil {
		return err
	}
	if len(photoFiles) == 0 {
		return fmt.Errorf("photo files are empty")
	}

	data, err := s.photoStorageAdapter.OpenPhoto(photoFiles[0].File.Path)
	if err != nil {
		return err
	}
	exifData, err := exif.ParseExifItemsAll(data)
	if err != nil {
		return err
	}

	return s.photoSearchAdapter.Index(ctx, photoID, photoFiles, exifData, s.nowFunc())
}

func (s *photoIndexService) CreatePreviewImages(ctx context.Context, photoID string) error {
	if photoID == "" {
		return fmt.Errorf("invalid photoID")
	}

	photoFiles, err := s.photoFileAdapter.FindByPhotoID(ctx, photoID)
	if err != nil {
		return err
	}
	if len(photoFiles) == 0 {
		return fmt.Errorf("photo files are empty")
	}

	jpegImage := photoFiles.FindFileByFileType(photoID, entities.PhotoFileTypeJPEG)
	if jpegImage == nil {
		// JPEG画像が無ければプレビュー画像を作成しない
		return nil
	}

	data, err := s.photoStorageAdapter.OpenPhoto(jpegImage.File.Path)
	if err != nil {
		return err
	}
	exifData, err := exif.ParseExifItemsAll(data)
	if err != nil {
		return err
	}

	previewData, err := s.imageProcessService.CreatePreview(data, exifData.Orientation())
	if err != nil {
		return err
	}
	if err := s.photoStorageAdapter.SavePreviewImage(photoID, previewData); err != nil {
		return err
	}

	thumbnailData, err := s.imageProcessService.CreateThumbnail(data, exifData.Orientation())
	if err != nil {
		return err
	}
	if err := s.photoStorageAdapter.SaveThumbnailImage(photoID, thumbnailData); err != nil {
		return err
	}

	return nil
}
