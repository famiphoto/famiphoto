package di

import (
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/drivers/db"
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
