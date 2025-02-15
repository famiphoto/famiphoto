package repositories

import (
	"context"
	"database/sql"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserPasswordRepository interface {
	Upsert(ctx context.Context, m *dbmodels.UserPassword) error
	Get(ctx context.Context, userID int64) (*dbmodels.UserPassword, error)
}

func NewUserPasswordRepository(cluster db.Cluster) UserPasswordRepository {
	return &userPasswordRepository{
		cluster: cluster,
	}
}

type userPasswordRepository struct {
	cluster db.Cluster
}

func (r *userPasswordRepository) Upsert(ctx context.Context, m *dbmodels.UserPassword) error {
	return m.Upsert(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Whitelist(
		dbmodels.UserPasswordColumns.Password,
		dbmodels.UserPasswordColumns.IsInitialized,
		dbmodels.UserPasswordColumns.LastModifiedAt,
		dbmodels.UserPasswordColumns.UpdatedAt,
	), boil.Infer())
}

func (r *userPasswordRepository) Get(ctx context.Context, userID int64) (*dbmodels.UserPassword, error) {
	row, err := dbmodels.FindUserPassword(ctx, r.cluster.GetTxnOrExecutor(ctx), userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBNotFoundError, err)
		}
		return nil, err
	}
	return row, nil
}
