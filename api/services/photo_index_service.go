package services

import (
	"context"
	"fmt"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/utils"
	"time"
)

type PhotoIndexService interface {
	RegisterPhotoToMasterData(ctx context.Context, files entities.StorageFileInfoList) (string, error)
	RegisterPhotoToSearchEngine(ctx context.Context, photoID string) error
}

func NewPhotoIndexService(
	photoAdapter adapters.PhotoAdapter,
	photoFileAdapter adapters.PhotoFileAdapter,
	photoStorageAdapter adapters.PhotoStorageAdapter,
	photoSearchAdapter adapters.PhotoSearchAdapter,
	transactionAdapter adapters.TransactionAdapter,
) PhotoIndexService {
	return &photoIndexService{
		photoAdapter:        photoAdapter,
		photoFileAdapter:    photoFileAdapter,
		photoStorageAdapter: photoStorageAdapter,
		photoSearchAdapter:  photoSearchAdapter,
		transactionAdapter:  transactionAdapter,
		nowFunc:             time.Now,
	}
}

type photoIndexService struct {
	photoAdapter        adapters.PhotoAdapter
	photoFileAdapter    adapters.PhotoFileAdapter
	photoStorageAdapter adapters.PhotoStorageAdapter
	photoSearchAdapter  adapters.PhotoSearchAdapter
	transactionAdapter  adapters.TransactionAdapter
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
	if len(photoFiles) == 0 {
		return fmt.Errorf("photo files are empty")
	}

	data, err := s.photoStorageAdapter.OpenPhoto(photoFiles[0].File.Path)
	if err != nil {
		return err
	}
	exif, err := utils.ParseExifItemsAll(data)
	if err != nil {
		return err
	}
	photoMeta := entities.NewPhotoMeta(exif)

	return s.photoSearchAdapter.Index(ctx, photoID, photoFiles, photoMeta, s.nowFunc())
}
