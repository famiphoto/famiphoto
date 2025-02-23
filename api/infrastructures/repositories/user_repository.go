package repositories

import (
	"context"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepository interface {
	ExistMyID(ctx context.Context, myID string) (bool, error)
	Insert(ctx context.Context, user *dbmodels.User) (*dbmodels.User, error)
}

type userRepository struct {
	cluster db.Cluster
}

func (r *userRepository) ExistMyID(ctx context.Context, myID string) (bool, error) {
	return dbmodels.Users(qm.Where("my_id = ?", myID)).Exists(ctx, r.cluster.GetTxnOrExecutor(ctx))
}

func (r *userRepository) Insert(ctx context.Context, user *dbmodels.User) (*dbmodels.User, error) {
	if err := user.Insert(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Infer()); err != nil {
		return nil, err
	}
	return user, nil
}
