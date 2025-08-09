package di

import "github.com/famiphoto/famiphoto/api/usecases"

func NewPhotoIndexingUseCase() usecases.PhotoIndexingUseCase {
	return usecases.NewPhotoIndexingUseCase(NewPhotoStorageAdapter(), NewPhotoIndexService())
}

func NewPhotoSearchUseCase() usecases.PhotoSearchUseCase {
	return usecases.NewPhotoSearchUseCase(NewPhotoSearchAdapter())
}

func NewUserUseCase() usecases.UserUseCase {
	return usecases.NewUserUseCase(
		NewTransactionAdapter(),
		NewUserAdapter(),
		NewUserPasswordAdapter(),
	)
}

func NewAssetUseCase() usecases.AssetUseCase {
	return usecases.NewAssetUseCase(NewPhotoAdapter(), NewPhotoStorageAdapter())
}
