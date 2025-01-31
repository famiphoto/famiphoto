package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
)

type TransactionAdapter interface {
	BeginTxn(ctx context.Context, fn func(executor db.Executor) error) error
}

type transactionAdapter struct {
	transactionRepo repositories.TransactionRepository
}

func (a *transactionAdapter) BeginTxn(ctx context.Context, fn func(executor db.Executor) error) error {
	return a.transactionRepo.RunInTxn(ctx, fn)
}
