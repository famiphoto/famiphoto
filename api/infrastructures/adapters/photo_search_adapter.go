package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/models"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"time"
)

type PhotoSearchAdapter interface {
	Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta entities.PhotoMeta, now time.Time) error
}

func NewPhotoSearchAdapter(esRepo repositories.PhotoElasticSearchRepository) PhotoSearchAdapter {
	return &photoSearchAdapter{
		esRepo: esRepo,
	}
}

type photoSearchAdapter struct {
	esRepo repositories.PhotoElasticSearchRepository
}

func (r *photoSearchAdapter) Index(ctx context.Context, photoID string, photoFiles entities.PhotoFileList, meta entities.PhotoMeta, now time.Time) error {
	// TODO 登録処理
	doc := &models.PhotoIndex{
		PhotoID:               photoID,
		Name:                  "",
		ImportedAt:            now.Unix(),
		DateTimeOriginal:      meta.DateTimeOriginal(),
		DateTimeOriginalParts: models.DateTimeOriginalParts{},
		Orientation:           0,
		Location:              nil,
		ImageUrls:             models.ImageUrls{},
		OriginalImageFiles:    nil,
		Exif:                  models.ExifData{},
		DescriptionJa:         "",
		DescriptionEn:         "",
	}

	return r.esRepo.Index(ctx, doc)
}
