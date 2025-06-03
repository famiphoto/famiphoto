package repositories

import (
	"context"
	"fmt"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepository interface {
	ExistUserID(ctx context.Context, userID string) (bool, error)
	Insert(ctx context.Context, user *dbmodels.User) (*dbmodels.User, error)
}

type userRepository struct {
	cluster db.Cluster
}

func NewUserRepository(cluster db.Cluster) UserRepository {
	return &userRepository{cluster: cluster}
}

func (r *userRepository) ExistUserID(ctx context.Context, userID string) (bool, error) {
	return dbmodels.Users(
		qm.Where(fmt.Sprintf("%s = ?", dbmodels.UserColumns.UserID), userID),
	).Exists(ctx, r.cluster.GetTxnOrExecutor(ctx))
}

func (r *userRepository) Insert(ctx context.Context, user *dbmodels.User) (*dbmodels.User, error) {
	if err := user.Insert(ctx, r.cluster.GetTxnOrExecutor(ctx), boil.Infer()); err != nil {
		return nil, err
	}
	return user, nil
}
