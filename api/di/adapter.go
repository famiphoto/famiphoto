package di

import "github.com/famiphoto/famiphoto/api/infrastructures/adapters"

func NewPhotoStorageAdapter() adapters.PhotoStorageAdapter {
	return adapters.NewPhotoStorageAdapter(NewPhotoStorageRepository())
}
