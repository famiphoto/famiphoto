package di

import "github.com/famiphoto/famiphoto/api/infrastructures/repositories"

func NewPhotoRepository() repositories.PhotoRepository {
	return repositories.NewPhotoRepository(NewMySQLClient())
}

func NewPhotoFileRepository() repositories.PhotoFileRepository {
	return repositories.NewPhotoFileRepository(NewMySQLClient())
}

func NewPhotoExifRepository() repositories.PhotoExifRepository {
	return repositories.NewPhotoExifRepository(NewMySQLClient())
}

func NewPhotoStorageRepository() repositories.PhotoStorageRepository {
	return repositories.NewPhotoStorageRepository(NewLocalStorage())
}

func NewTransactionRepository() repositories.TransactionRepository {
	return repositories.NewTransactionRepository(NewMySQLClient())
}
