package usecases

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
)

type PhotoSearchUseCase interface {
	Search(ctx context.Context, photoSearchQuery *entities.PhotoSearchQuery) (*entities.PhotoSearchResult, error)
	GetByPhotoID(ctx context.Context, photoID string) (*entities.PhotoSearchResult, error)
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

func (u *photoSearchUseCase) Search(ctx context.Context, photoSearchQuery *entities.PhotoSearchQuery) (*entities.PhotoSearchResult, error) {
	// TODO 検索クエリの内容精査など

	return u.photoSearchAdapter.Search(ctx, photoSearchQuery)
}

func (u *photoSearchUseCase) GetByPhotoID(ctx context.Context, photoID string) (*entities.PhotoSearchResult, error) {
	result, err := u.photoSearchAdapter.Search(ctx, &entities.PhotoSearchQuery{
		PhotoID: photoID,
		Limit:   1,
		Offset:  0,
	})
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, errors.New(errors.PhotoNotFoundError, nil)
	}

	return result, nil
}
