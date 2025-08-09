package usecases

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
)

type PhotoSearchUseCase interface {
	Search(ctx context.Context, photoSearchQuery *entities.PhotoSearchQuery) (*entities.PhotoSearchResult, error)
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
