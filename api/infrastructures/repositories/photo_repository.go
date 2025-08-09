package repositories

import (
	"context"
	"database/sql"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PhotoRepository interface {
	Insert(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error)
	Update(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error)
	GetPhotoByFileNameHash(ctx context.Context, filePathHash string) (*dbmodels.Photo, error)
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
	)); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *photoRepository) GetPhotoByFileNameHash(ctx context.Context, filePathHash string) (*dbmodels.Photo, error) {
	row, err := dbmodels.Photos(qm.Where("file_name_hash = ?", filePathHash)).One(ctx, r.cluster.GetTxnOrExecutor(ctx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBNotFoundError, err)
		}
		return nil, err
	}
	return row, nil
}
