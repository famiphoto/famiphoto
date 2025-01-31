package services

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/utils"
	"time"
)

type photoIndexService struct {
	photoAdapter        adapters.PhotoAdapter
	photoFileAdapter    adapters.PhotoFileAdapter
	photoStorageAdapter adapters.PhotoStorageAdapter
	photoMetaAdapter    adapters.PhotoMetaAdapter
	photoSearchAdapter  adapters.PhotoSearchAdapter
	transactionAdapter  adapters.TransactionAdapter
	nowFunc             func() time.Time
}

func (s *photoIndexService) RegisterNewPhoto(ctx context.Context, photoFile *entities.StorageFileInfo) error {
	data, err := s.photoStorageAdapter.OpenPhoto(photoFile.Path)
	if err != nil {
		return err
	}
	exif, err := utils.ParseExifItemsAll(data)
	if err != nil {
		return err
	}
	if err := s.photoMetaAdapter.Upsert(ctx, entities.NewPhotoMeta(exif)); err != nil {
		return err
	}

	if err := s.photoFileAdapter.Upsert(ctx, &entities.PhotoFile{
		FileHash: data.FileHash(),
		File:     *photoFile,
	}); err != nil {
		return err
	}

	if err := s.photoAdapter.Upsert(ctx, &entities.Photo{
		Name:       photoFile.Name,
		ImportedAt: s.nowFunc(),
	}); err != nil {
		return err
	}

	// TODO 検索エンジン

	return nil
}
