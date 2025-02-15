package di

import "github.com/famiphoto/famiphoto/api/infrastructures/repositories"

func NewPhotoRepository() repositories.PhotoRepository {
	return repositories.NewPhotoRepository(NewMySQLCluster())
}

func NewPhotoFileRepository() repositories.PhotoFileRepository {
	return repositories.NewPhotoFileRepository(NewMySQLCluster())
}

func NewPhotoExifRepository() repositories.PhotoExifRepository {
	return repositories.NewPhotoExifRepository(NewMySQLCluster())
}

func NewPhotoStorageRepository() repositories.PhotoStorageRepository {
	return repositories.NewPhotoStorageRepository(NewLocalStorage())
}

func NewTransactionRepository() repositories.TransactionRepository {
	return repositories.NewTransactionRepository(NewMySQLCluster())
}

func NewPhotoElasticSearchRepository() repositories.PhotoElasticSearchRepository {
	return repositories.NewPhotoElasticSearchRepository(NewElasticSearchClient(), NewTypesElasticSearchClient())
}

func NewSessionRepository() repositories.SessionRepository {
	return repositories.NewSessionRepository(NewSessionDB())
}
