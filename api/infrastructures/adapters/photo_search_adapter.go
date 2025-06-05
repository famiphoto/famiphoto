package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
)

type PhotoSearchAdapter interface {
	Index(ctx context.Context, photo *entities.Photo, meta entities.PhotoMeta) error
	Search(ctx context.Context, limit, offset int) ([]*models.PhotoIndex, error)
}

func NewPhotoSearchAdapter(esRepo repositories.PhotoElasticSearchRepository) PhotoSearchAdapter {
	return &photoSearchAdapter{
		esRepo: esRepo,
	}
}

type photoSearchAdapter struct {
	esRepo repositories.PhotoElasticSearchRepository
}

func (r *photoSearchAdapter) Index(ctx context.Context, photo *entities.Photo, meta entities.PhotoMeta) error {
	doc := &models.PhotoIndex{
		PhotoID:          photo.PhotoID,
		Name:             photo.Name,
		ImportedAt:       photo.ImportedAt.Unix(),
		DateTimeOriginal: meta.DateTimeOriginal(),
		DescriptionJa:    photo.DescriptionJa,
		DescriptionEn:    photo.DescriptionEn,
	}

	return r.esRepo.Index(ctx, doc)
}

func (r *photoSearchAdapter) Search(ctx context.Context, limit, offset int) ([]*models.PhotoIndex, error) {
	return r.esRepo.List(ctx, limit, offset)
}
