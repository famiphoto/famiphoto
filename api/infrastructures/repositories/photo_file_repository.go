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
}

func NewPhotoFileRepository(client db.Client) PhotoFileRepository {
	return &photoFileRepository{db: client}
}

type photoFileRepository struct {
	db db.Client
}

func (r *photoFileRepository) Insert(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error) {
	if err := photoFile.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return photoFile, nil
}

func (r *photoFileRepository) Update(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error) {
	if _, err := photoFile.Update(ctx, r.db, boil.Blacklist(
		dbmodels.PhotoColumns.ImportedAt,
	)); err != nil {
		return nil, err
	}
	return photoFile, nil
}

func (r *photoFileRepository) GetPhotoFileByFilePath(ctx context.Context, filePath string) (*dbmodels.PhotoFile, error) {
	row, err := dbmodels.PhotoFiles(qm.Where("file_path = ?", filePath)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBNotFoundError, err)
		}
		return nil, err
	}
	return row, nil
}
