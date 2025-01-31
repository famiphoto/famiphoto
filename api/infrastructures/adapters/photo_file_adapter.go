package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
)

type PhotoFileAdapter interface {
	Upsert(ctx context.Context, photoFile *entities.PhotoFile) error
}

func NewPhotoFileAdapter(photoFileRepo repositories.PhotoFileRepository) PhotoFileAdapter {
	return &photoFileAdapter{photoFileRepo: photoFileRepo}
}

type photoFileAdapter struct {
	photoFileRepo repositories.PhotoFileRepository
}

func (a *photoFileAdapter) Upsert(ctx context.Context, photoFile *entities.PhotoFile) error {
	dbModel := &dbmodels.PhotoFile{
		PhotoFileID: 0,
		PhotoID:     photoFile.PhotoID,
		FileType:    photoFile.FileType().ToString(),
		FilePath:    photoFile.File.Path,
		FileHash:    photoFile.FileHash,
	}

	row, err := a.photoFileRepo.GetPhotoFileByFilePath(ctx, photoFile.File.Path)
	if err != nil {
		if errors.IsErrCode(err, errors.DBNotFoundError) {
			if _, err := a.photoFileRepo.Insert(ctx, dbModel); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	dbModel.PhotoFileID = row.PhotoFileID
	if _, err := a.photoFileRepo.Update(ctx, dbModel); err != nil {
		return err
	}
	return nil
}
