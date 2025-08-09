package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/random"
)

type PhotoAdapter interface {
	InsertIfNotExist(ctx context.Context, photo *entities.StorageFileInfo) (string, error)
	FindByID(ctx context.Context, photoID string) (*entities.StorageFileInfo, error)
}

func NewPhotoAdapter(photoRepo repositories.PhotoRepository) PhotoAdapter {
	return &photoAdapter{
		photoRepo: photoRepo,
	}
}

type photoAdapter struct {
	photoRepo repositories.PhotoRepository
}

func (a *photoAdapter) InsertIfNotExist(ctx context.Context, photo *entities.StorageFileInfo) (string, error) {
	m := &dbmodels.Photo{
		PhotoID:      random.GenerateUUID(),
		Name:         photo.Name,
		FileNameHash: photo.FilePathExceptExtHash(),
	}

	row, err := a.photoRepo.GetPhotoByFileNameHash(ctx, m.FileNameHash)
	if err != nil {
		if !errors.IsErrCode(err, errors.DBNotFoundError) {
			return "", err
		}

		// DBに存在しない場合のみ新規追加する
		dst, err := a.photoRepo.Insert(ctx, m)
		if err != nil {
			return "", err
		}
		return dst.PhotoID, nil
	}

	return row.PhotoID, nil
}

func (a *photoAdapter) FindByID(ctx context.Context, photoID string) (*entities.StorageFileInfo, error) {
	row, err := a.photoRepo.GetPhotoByID(ctx, photoID)
	if err != nil {
		return nil, err
	}
	return &entities.StorageFileInfo{
		Name: row.Name,
	}, nil
}
