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

func NewPhotoRepository(cluster db.Cluster) PhotoRepository {
	return &photoRepository{cluster: cluster}
}

type photoRepository struct {
	cluster db.Cluster
}

func (r *photoRepository) Insert(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	if err := photo.Insert(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Infer()); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *photoRepository) Update(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	if _, err := photo.Update(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Blacklist(
		dbmodels.PhotoColumns.CreatedAt,
		dbmodels.PhotoColumns.ImportedAt,
	)); err != nil {
		return nil, err
	}
	return photo, nil
}
