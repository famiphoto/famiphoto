package di

import (
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/drivers/storage"
)

var mySQLCluster db.Cluster

func NewMySQLCluster() db.Cluster {
	if mySQLCluster != nil {
		return mySQLCluster
	}
	c := db.NewMySQLClient(
		config.Env.MySQLHostName,
		config.Env.MySQLPort,
		config.Env.MySQLUser,
		config.Env.MySQLPassword,
		config.Env.MySQLDatabase,
	)
	mySQLCluster = db.NewCluster("famiphoto_mysql", c)
	return mySQLCluster
}

var localStorage storage.Client

func NewLocalStorage() storage.Client {
	if localStorage != nil {
		return localStorage
	}

	localStorage = storage.NewLocalStorage()
	return localStorage
}
