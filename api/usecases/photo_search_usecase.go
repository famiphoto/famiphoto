package usecases

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
)

type PhotoSearchUseCase interface {
	GetPhotoList(ctx context.Context, limit, offset int) ([]*entities.Photo, int64, error)
}

func NewPhotoSearchUseCase(
	photoSearchAdapter adapters.PhotoSearchAdapter,
) PhotoSearchUseCase {
	return &photoSearchUseCase{
		photoSearchAdapter: photoSearchAdapter,
	}
}

type photoSearchUseCase struct {
	photoSearchAdapter adapters.PhotoSearchAdapter
}

func (u *photoSearchUseCase) GetPhotoList(ctx context.Context, limit, offset int) ([]*entities.Photo, int64, error) {
	photos, total, err := u.photoSearchAdapter.Search(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return photos, total, nil
}
