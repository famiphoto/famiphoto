package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"time"
)

type PhotoSearchAdapter interface {
	Index(ctx context.Context, photo *entities.Photo, meta entities.PhotoMeta) error
	Get(ctx context.Context, photoID string) (*entities.Photo, error)
	Search(ctx context.Context, limit, offset int) ([]*entities.Photo, int64, error)
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

func (r *photoSearchAdapter) Get(ctx context.Context, photoID string) (*entities.Photo, error) {
	item, err := r.esRepo.Get(ctx, photoID)
	if err != nil {
		return nil, err
	}

	return r.toEntity(item), nil
}

func (r *photoSearchAdapter) Search(ctx context.Context, limit, offset int) ([]*entities.Photo, int64, error) {
	rows, total, err := r.esRepo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return r.toEntities(rows), total, nil
}

func (r *photoSearchAdapter) toEntity(item *models.PhotoIndex) *entities.Photo {
	return &entities.Photo{
		PhotoID:       item.PhotoID,
		Name:          item.Name,
		ImportedAt:    time.Unix(item.ImportedAt, 0),
		DescriptionJa: item.DescriptionJa,
		DescriptionEn: item.DescriptionEn,
	}
}

func (r *photoSearchAdapter) toEntities(items []*models.PhotoIndex) []*entities.Photo {
	photos := make([]*entities.Photo, len(items))
	for i, item := range items {
		photos[i] = r.toEntity(item)
	}
	return photos
}
