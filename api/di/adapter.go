package di

import "github.com/famiphoto/famiphoto/api/infrastructures/adapters"

func NewPhotoAdapter() adapters.PhotoAdapter {
	return adapters.NewPhotoAdapter(NewPhotoRepository())
}

func NewPhotoFileAdapter() adapters.PhotoFileAdapter {
	return adapters.NewPhotoFileAdapter(NewPhotoFileRepository())
}

func NewPhotoStorageAdapter() adapters.PhotoStorageAdapter {
	return adapters.NewPhotoStorageAdapter(NewPhotoStorageRepository())
}

func NewPhotoMetaAdapter() adapters.PhotoMetaAdapter {
	return adapters.NewPhotoMetaAdapter(NewPhotoExifRepository())
}

func NewPhotoSearchAdapter() adapters.PhotoSearchAdapter {
	return adapters.NewPhotoSearchAdapter(NewPhotoElasticSearchRepository())
}

func NewTransactionAdapter() adapters.TransactionAdapter {
	return adapters.NewTransactionAdapter(NewTransactionRepository())
}

func NewSessionAdapter() adapters.SessionAdapter {
	return adapters.NewSessionAdapter(NewSessionRepository())
}
