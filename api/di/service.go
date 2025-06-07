package di

import (
	"github.com/famiphoto/famiphoto/api/services"
)

func NewPhotoIndexService() services.PhotoIndexService {
	return services.NewPhotoIndexService(NewPhotoAdapter(), NewPhotoFileAdapter(), NewPhotoStorageAdapter(), NewPhotoSearchAdapter(), NewTransactionAdapter(), NewImageProcessService())
}

func NewImageProcessService() services.ImageProcessService {
	return services.NewImageProcessService()
}
