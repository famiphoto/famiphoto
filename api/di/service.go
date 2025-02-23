package di

import (
	"github.com/famiphoto/famiphoto/api/services"
)

func NewPhotoIndexService() services.PhotoIndexService {
	return services.NewPhotoIndexService(NewPhotoAdapter(), NewPhotoFileAdapter(), NewPhotoStorageAdapter(), NewPhotoMetaAdapter(), NewPhotoSearchAdapter(), NewTransactionAdapter())
}
