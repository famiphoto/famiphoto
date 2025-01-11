package repositories

import (
	"context"
	"github.com/famiphoto/famiphoto/api/drivers/db"
)

type TransactionRepository interface {
	RunInTxn(ctx context.Context, fn func(txn db.Executor) error) error
}

func NewTransactionRepository(client db.Client) TransactionRepository {
	return &transactionRepository{
		db: client,
	}
}

type transactionRepository struct {
	db db.Client
}

func (r *transactionRepository) RunInTxn(ctx context.Context, fn func(txn db.Executor) error) error {
	txn, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(txn)
	if err != nil {
		if rollBackErr := txn.Rollback(); rollBackErr != nil {
			return rollBackErr
		}
		return err
	}

	return txn.Commit()
}
