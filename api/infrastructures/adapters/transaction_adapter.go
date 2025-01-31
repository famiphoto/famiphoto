package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
)

type TransactionAdapter interface {
	BeginTxn(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransactionAdapter(transactionRepo repositories.TransactionRepository) TransactionAdapter {
	return &transactionAdapter{
		transactionRepo: transactionRepo,
	}
}

type transactionAdapter struct {
	transactionRepo repositories.TransactionRepository
}

func (a *transactionAdapter) BeginTxn(ctx context.Context, fn func(ctx context.Context) error) error {
	return a.transactionRepo.RunInTxn(ctx, fn)
}
