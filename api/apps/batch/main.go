package main

import (
	"context"
	"github.com/famiphoto/famiphoto/api/di"
	"github.com/labstack/gommon/log"
)

func main() {
	uc := di.NewPhotoIndexingUseCase()
	err := uc.IndexPhotos(context.Background(), []string{".jpg", ".jpeg", ".arw"}, 3)
	if err != nil {
		log.Error(err)
	}
}
