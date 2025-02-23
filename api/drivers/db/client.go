package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Client interface {
	boil.ContextExecutor
	boil.ContextBeginner
}

type Executor interface {
	boil.ContextExecutor
}

type Transactor interface {
	boil.Transactor
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

type Cluster interface {
	NewTxn(ctx context.Context) (context.Context, Transactor, error)
	DeleteTxn(ctx context.Context) context.Context
	GetTxnOrExecutor(ctx context.Context) Executor
}

func NewCluster(name string, client Client) Cluster {
	return &cluster{
		name:   name,
		client: client,
	}
}

type cluster struct {
	name   string
	client Client
}

func (c *cluster) NewTxn(ctx context.Context) (context.Context, Transactor, error) {
	newTxn, err := c.client.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	newCtx := context.WithValue(ctx, c.name, newTxn)
	return newCtx, newTxn, nil
}

func (c *cluster) DeleteTxn(ctx context.Context) context.Context {
	return context.WithValue(ctx, c.name, nil)
}

func (c *cluster) GetTxnOrExecutor(ctx context.Context) Executor {
	txn, ok := ctx.Value(c.name).(Executor)
	if ok {
		return txn
	}
	return c.client
}
