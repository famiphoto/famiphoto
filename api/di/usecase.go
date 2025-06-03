package di

import "github.com/famiphoto/famiphoto/api/usecases"

func NewPhotoIndexingUseCase() usecases.PhotoIndexingUseCase {
	return usecases.NewPhotoIndexingUseCase(NewPhotoStorageAdapter(), NewPhotoIndexService())
}

func NewAuthUseCase() usecases.AuthUseCase {
	return usecases.NewAuthUseCase(
		NewTransactionAdapter(),
		NewUserAdapter(),
		NewUserPasswordAdapter(),
	)
}
