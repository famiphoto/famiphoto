package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/cast"
)

type UserAdapter interface {
	IsAlreadyUsedMyID(ctx context.Context, myID string) (bool, error)
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
}

type userAdapter struct {
	userRepo repositories.UserRepository
}

func NewUserAdapter(userRepo repositories.UserRepository) UserAdapter {
	return &userAdapter{userRepo: userRepo}
}

func (a *userAdapter) IsAlreadyUsedMyID(ctx context.Context, myID string) (bool, error) {
	return a.userRepo.ExistMyID(ctx, myID)
}

func (a *userAdapter) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	dst, err := a.userRepo.Insert(ctx, &dbmodels.User{
		UserID:  user.UserID,
		Status:  int(user.Status),
		IsAdmin: cast.BoolToInt8(user.IsAdmin),
	})
	if err != nil {
		return nil, err
	}
	return a.toEntity(dst), nil
}

func (a *userAdapter) toEntity(row *dbmodels.User) *entities.User {
	return &entities.User{
		UserID:  row.UserID,
		Status:  entities.UserStatus(row.Status),
		IsAdmin: cast.IntToBool(row.IsAdmin),
	}
}
