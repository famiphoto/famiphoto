package db

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Client interface {
	boil.ContextExecutor
	boil.ContextBeginner
}

type Executor interface {
	boil.ContextExecutor
}

func NewMySQLClient(hostname, port, user, password, dbName string) Client {
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		hostname,
		port,
		dbName,
	)
	newDB, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	if err := newDB.Ping(); err != nil {
		panic(fmt.Sprintf("%s %+v", hostname, err))
	}

	return newDB
}
