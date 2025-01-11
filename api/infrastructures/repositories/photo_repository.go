package repositories

import (
	"context"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PhotoRepository interface {
	Insert(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error)
	Update(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error)
}

func NewPhotoRepository(client db.Client) PhotoRepository {
	return &photoRepository{db: client}
}

type photoRepository struct {
	db db.Client
}

func (r *photoRepository) Insert(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	if err := photo.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *photoRepository) Update(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	if _, err := photo.Update(ctx, r.db, boil.Blacklist(
		dbmodels.PhotoColumns.CreatedAt,
		dbmodels.PhotoColumns.ImportedAt,
	)); err != nil {
		return nil, err
	}
	return photo, nil
}
