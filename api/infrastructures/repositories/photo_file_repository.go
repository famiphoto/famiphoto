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

type PhotoFileRepository interface {
	Insert(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
	Update(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
	GetPhotoFileByFilePath(ctx context.Context, filePath string) (*dbmodels.PhotoFile, error)
	GetPhotoFilesByPhotoID(ctx context.Context, photoID string) (dbmodels.PhotoFileSlice, error)
	GetPhotoFileByPhotoFileID(ctx context.Context, photoFileID string) (*dbmodels.PhotoFile, error)
}

func NewPhotoFileRepository(cluster db.Cluster) PhotoFileRepository {
	return &photoFileRepository{cluster: cluster}
}

type photoFileRepository struct {
	cluster db.Cluster
}

func (r *photoFileRepository) Insert(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error) {
	if err := photoFile.Insert(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Infer()); err != nil {
		return nil, err
	}
	return photoFile, nil
}

func (r *photoFileRepository) Update(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error) {
	if _, err := photoFile.Update(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Infer()); err != nil {
		return nil, err
	}
	return photoFile, nil
}

func (r *photoFileRepository) GetPhotoFileByFilePath(ctx context.Context, filePath string) (*dbmodels.PhotoFile, error) {
	row, err := dbmodels.PhotoFiles(qm.Where("file_path = ?", filePath)).One(ctx, r.cluster.GetTxnOrExecutor(ctx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBNotFoundError, err)
		}
		return nil, err
	}
	return row, nil
}

func (r *photoFileRepository) GetPhotoFilesByPhotoID(ctx context.Context, photoID string) (dbmodels.PhotoFileSlice, error) {
	rows, err := dbmodels.PhotoFiles(qm.Where("photo_id = ?", photoID)).All(ctx, r.cluster.GetTxnOrExecutor(ctx))
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *photoFileRepository) GetPhotoFileByPhotoFileID(ctx context.Context, photoFileID string) (*dbmodels.PhotoFile, error) {
	return dbmodels.FindPhotoFile(ctx, r.cluster.GetTxnOrExecutor(ctx), photoFileID)
}
