package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
)

type PhotoAdapter interface {
	Upsert(ctx context.Context, photo *entities.Photo) error
}

type photoAdapter struct {
	photoRepo repositories.PhotoRepository
}

func (a *photoAdapter) Upsert(ctx context.Context, photo *entities.Photo) error {

}
