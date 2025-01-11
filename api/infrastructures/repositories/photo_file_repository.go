package repositories

import (
	"context"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PhotoFileRepository interface {
	Insert(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
	Update(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
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
