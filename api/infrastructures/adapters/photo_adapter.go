package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/random"
)

type PhotoAdapter interface {
	Upsert(ctx context.Context, photo *entities.Photo) (*entities.Photo, error)
}

func NewPhotoAdapter(photoRepo repositories.PhotoRepository) PhotoAdapter {
	return &photoAdapter{
		photoRepo: photoRepo,
	}
}

type photoAdapter struct {
	photoRepo repositories.PhotoRepository
}

func (a *photoAdapter) Upsert(ctx context.Context, photo *entities.Photo) (*entities.Photo, error) {
	m := &dbmodels.Photo{
		PhotoID:       random.GenerateUUID(),
		Name:          photo.Name,
		ImportedAt:    photo.ImportedAt,
		DescriptionJa: photo.DescriptionJa,
		DescriptionEn: photo.DescriptionEn,
		FileNameHash:  photo.FileNameHash,
	}

	row, err := a.photoRepo.GetPhotoByFileNameHash(ctx, photo.FileNameHash)
	if err != nil {
		if !errors.IsErrCode(err, errors.DBNotFoundError) {
			return nil, err
		}
		dst, err := a.photoRepo.Insert(ctx, m)
		if err != nil {
			return nil, err
		}
		return a.toEntity(dst), nil
	}

	m.PhotoID = row.PhotoID
	dst, err := a.photoRepo.Update(ctx, m)
	if err != nil {
		return nil, err
	}
	return a.toEntity(dst), nil
}

func (a *photoAdapter) toEntity(row *dbmodels.Photo) *entities.Photo {
	return &entities.Photo{
		PhotoID:       row.PhotoID,
		Name:          row.Name,
		DescriptionJa: row.DescriptionJa,
		DescriptionEn: row.DescriptionEn,
		ImportedAt:    row.ImportedAt,
	}
}
