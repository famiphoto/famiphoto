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

type PhotoExifRepository interface {
	Insert(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error)
	Update(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error)
	GetPhotoExifByPhotoIDTagID(ctx context.Context, photoID string, tagID int64) (*dbmodels.PhotoExif, error)
}

func NewPhotoExifRepository(cluster db.Cluster) PhotoExifRepository {
	return &photoExifRepository{cluster: cluster}
}

type photoExifRepository struct {
	cluster db.Cluster
}

func (r *photoExifRepository) Insert(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error) {
	if err := exif.Insert(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Infer()); err != nil {
		return nil, err
	}
	return exif, nil
}

func (r *photoExifRepository) Update(ctx context.Context, exif *dbmodels.PhotoExif) (*dbmodels.PhotoExif, error) {
	if _, err := exif.Update(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Infer()); err != nil {
		return nil, err
	}
	return exif, nil
}

func (r *photoExifRepository) GetPhotoExifByPhotoIDTagID(ctx context.Context, photoID string, tagID int64) (*dbmodels.PhotoExif, error) {
	m, err := dbmodels.PhotoExifs(qm.Where("photo_id = ?", photoID), qm.Where("tag_id = ?", tagID)).One(ctx, r.cluster.GetTxnOrExecutor(ctx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBNotFoundError, err)
		}
		return nil, err
	}
	return m, nil
}
