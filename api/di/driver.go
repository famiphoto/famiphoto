package di

import (
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/drivers/storage"
)

var mySQLClient db.Client

func NewMySQLClient() db.Client {
	if mySQLClient != nil {
		return mySQLClient
	}
	c := db.NewMySQLClient(
		config.Env.MySQLHostName,
		config.Env.MySQLPort,
		config.Env.MySQLUser,
		config.Env.MySQLPassword,
		config.Env.MySQLDatabase,
	)
	mySQLClient = c
	return mySQLClient
}

var localStorage storage.Client

func NewLocalStorage() storage.Client {
	if localStorage != nil {
		return localStorage
	}

	localStorage = storage.NewLocalStorage()
	return localStorage
}
