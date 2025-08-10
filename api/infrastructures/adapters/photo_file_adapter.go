package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/random"
	"path/filepath"
)

type PhotoFileAdapter interface {
	Upsert(ctx context.Context, photoFile *entities.PhotoFile) (string, error)
	FindByPhotoID(ctx context.Context, photoID string) (entities.PhotoFileList, error)
	FindByPhotoFileID(ctx context.Context, photoFIleID string) (*entities.PhotoFile, error)
}

func NewPhotoFileAdapter(photoFileRepo repositories.PhotoFileRepository) PhotoFileAdapter {
	return &photoFileAdapter{photoFileRepo: photoFileRepo}
}

type photoFileAdapter struct {
	photoFileRepo repositories.PhotoFileRepository
}

func (a *photoFileAdapter) Upsert(ctx context.Context, photoFile *entities.PhotoFile) (string, error) {
	dbModel := &dbmodels.PhotoFile{
		PhotoFileID:  random.GenerateUUID(),
		PhotoID:      photoFile.PhotoID,
		FileType:     photoFile.FileType().ToString(),
		FilePath:     photoFile.File.Path,
		FilePathHash: photoFile.File.FilePathHash(),
		FileHash:     photoFile.FileHash,
	}

	row, err := a.photoFileRepo.GetPhotoFileByFilePath(ctx, photoFile.File.Path)
	if err != nil {
		if errors.IsErrCode(err, errors.DBNotFoundError) {
			if _, err := a.photoFileRepo.Insert(ctx, dbModel); err != nil {
				return "", err
			}
			return dbModel.PhotoFileID, nil
		}
		return "", err
	}

	dbModel.PhotoFileID = row.PhotoFileID
	if _, err := a.photoFileRepo.Update(ctx, dbModel); err != nil {
		return "", err
	}
	return dbModel.PhotoFileID, nil
}

func (a *photoFileAdapter) FindByPhotoID(ctx context.Context, photoID string) (entities.PhotoFileList, error) {
	rows, err := a.photoFileRepo.GetPhotoFilesByPhotoID(ctx, photoID)
	if err != nil {
		return nil, err
	}

	return a.toEntities(rows), nil
}

func (a *photoFileAdapter) FindByPhotoFileID(ctx context.Context, photoFIleID string) (*entities.PhotoFile, error) {
	row, err := a.photoFileRepo.GetPhotoFileByPhotoFileID(ctx, photoFIleID)
	if err != nil {
		return nil, err
	}

	return a.toEntity(row), nil
}

func (a *photoFileAdapter) toEntity(row *dbmodels.PhotoFile) *entities.PhotoFile {
	if row == nil {
		return nil
	}

	return &entities.PhotoFile{
		PhotoFileID: row.PhotoFileID,
		PhotoID:     row.PhotoID,
		FileHash:    row.FileHash,
		File: entities.StorageFileInfo{
			Name:  filepath.Base(row.FilePath),
			Path:  row.FilePath,
			IsDir: false,
			Ext:   filepath.Ext(row.FilePath),
		},
	}
}

func (a *photoFileAdapter) toEntities(rows dbmodels.PhotoFileSlice) []*entities.PhotoFile {
	if rows == nil {
		return nil
	}

	result := make([]*entities.PhotoFile, len(rows))
	for i, row := range rows {
		result[i] = a.toEntity(row)
	}
	return result
}
