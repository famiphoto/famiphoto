package services

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/utils"
	"time"
)

type PhotoIndexService interface {
	RegisterPhotoToMasterData(ctx context.Context, photoFile *entities.StorageFileInfo) (photoID int64, err error)
}

func NewPhotoIndexService(
	photoAdapter adapters.PhotoAdapter,
	photoFileAdapter adapters.PhotoFileAdapter,
	photoStorageAdapter adapters.PhotoStorageAdapter,
	photoMetaAdapter adapters.PhotoMetaAdapter,
	photoSearchAdapter adapters.PhotoSearchAdapter,
	transactionAdapter adapters.TransactionAdapter,
) PhotoIndexService {
	return &photoIndexService{
		photoAdapter:        photoAdapter,
		photoFileAdapter:    photoFileAdapter,
		photoStorageAdapter: photoStorageAdapter,
		photoMetaAdapter:    photoMetaAdapter,
		photoSearchAdapter:  photoSearchAdapter,
		transactionAdapter:  transactionAdapter,
		nowFunc:             time.Now,
	}
}

type photoIndexService struct {
	photoAdapter        adapters.PhotoAdapter
	photoFileAdapter    adapters.PhotoFileAdapter
	photoStorageAdapter adapters.PhotoStorageAdapter
	photoMetaAdapter    adapters.PhotoMetaAdapter
	photoSearchAdapter  adapters.PhotoSearchAdapter
	transactionAdapter  adapters.TransactionAdapter
	nowFunc             func() time.Time
}

func (s *photoIndexService) RegisterPhotoToMasterData(ctx context.Context, photoFile *entities.StorageFileInfo) (int64, error) {
	data, err := s.photoStorageAdapter.OpenPhoto(photoFile.Path)
	if err != nil {
		return 0, err
	}
	exif, err := utils.ParseExifItemsAll(data)
	if err != nil {
		return 0, err
	}

	photoID := int64(0)
	err = s.transactionAdapter.BeginTxn(ctx, func(ctx2 context.Context) error {
		photo, err := s.photoAdapter.Upsert(ctx2, &entities.Photo{
			Name:         photoFile.Name,
			ImportedAt:   s.nowFunc(),
			FileNameHash: utils.FileNameExceptExt(photoFile.Path),
		})
		if err != nil {
			return err
		}

		if err := s.photoMetaAdapter.Upsert(ctx2, photo.PhotoID, entities.NewPhotoMeta(exif)); err != nil {
			return err
		}

		if err := s.photoFileAdapter.Upsert(ctx2, &entities.PhotoFile{
			PhotoID:  photo.PhotoID,
			FileHash: data.FileHash(),
			File:     *photoFile,
		}); err != nil {
			return err
		}

		photoID = photo.PhotoID
		return nil
	})
	if err != nil {
		return 0, err
	}

	return photoID, nil
}

func (s *photoIndexService) RegisterPhotoToSearchEngine(ctx context.Context, photoID int64) error {
	panic("")
}
