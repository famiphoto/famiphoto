package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
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
	dst, err := a.photoRepo.Insert(ctx, &dbmodels.Photo{
		PhotoID:       photo.PhotoID,
		Name:          photo.Name,
		ImportedAt:    photo.ImportedAt,
		DescriptionJa: photo.DescriptionJa,
		DescriptionEn: photo.DescriptionEn,
	})
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
