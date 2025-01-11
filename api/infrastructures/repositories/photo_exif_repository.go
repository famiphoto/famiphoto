package repositories

import (
	"context"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PhotoExifRepository interface {
	Insert(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error)
	Update(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error)
}

func NewPhotoExifRepository(client db.Client) PhotoExifRepository {
	return &photoExifRepository{db: client}
}

type photoExifRepository struct {
	db db.Client
}

func (r *photoExifRepository) Insert(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error) {
	if err := exif.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return exif, nil
}

func (r *photoExifRepository) Update(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error) {
	if _, err := exif.Update(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return exif, nil
}
