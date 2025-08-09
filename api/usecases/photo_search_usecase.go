package usecases

import (
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
)

type PhotoSearchUseCase interface {
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
