package repositories

import (
	"context"
	"github.com/famiphoto/famiphoto/api/drivers/db"
)

type TransactionRepository interface {
	RunInTxn(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransactionRepository(cluster db.Cluster) TransactionRepository {
	return &transactionRepository{
		cluster: cluster,
	}
}

type transactionRepository struct {
	cluster db.Cluster
}

func (r *transactionRepository) RunInTxn(ctx context.Context, fn func(ctx context.Context) error) error {
	txnCtx, txn, err := r.cluster.NewTxn(ctx)
	if err != nil {
		return err
	}

	defer func() {
		r.cluster.DeleteTxn(txnCtx)
	}()

	err = fn(txnCtx)
	if err != nil {
		if rollBackErr := txn.Rollback(); rollBackErr != nil {
			return rollBackErr
		}
		return err
	}

	return txn.Commit()
}
