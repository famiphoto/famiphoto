package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/famiphoto/famiphoto/api/drivers/db"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepository interface {
	ExistUserID(ctx context.Context, userID string) (bool, error)
	Insert(ctx context.Context, user *dbmodels.User) (*dbmodels.User, error)
	Get(ctx context.Context, userID string) (*dbmodels.User, error)
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

func (r *userRepository) Get(ctx context.Context, userID string) (*dbmodels.User, error) {
	user, err := dbmodels.FindUser(ctx, r.cluster.GetTxnOrExecutor(ctx), userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBNotFoundError, err)
		}
		return nil, err
	}
	return user, nil
}
