package di

import "github.com/famiphoto/famiphoto/api/usecases"

func NewPhotoIndexingUseCase() usecases.PhotoIndexingUseCase {
	return usecases.NewPhotoIndexingUseCase(NewPhotoStorageAdapter(), NewPhotoIndexService())
}

func NewUserUseCase() usecases.UserUseCase {
	return usecases.NewUserUseCase(
		NewTransactionAdapter(),
		NewUserAdapter(),
		NewUserPasswordAdapter(),
	)
}
