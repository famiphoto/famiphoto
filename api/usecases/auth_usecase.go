package usecases

import (
	"context"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"github.com/famiphoto/famiphoto/api/utils/password"
	"time"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, userID, pw string, isAdmin bool, now time.Time) (*entities.User, error)
}

func NewAuthUseCase(
	txnAdapter adapters.TransactionAdapter,
	userAdapter adapters.UserAdapter,
	userPasswordAdapter adapters.UserPasswordAdapter,
) AuthUseCase {
	return &authUseCase{
		txnAdapter:          txnAdapter,
		userAdapter:         userAdapter,
		userPasswordAdapter: userPasswordAdapter,
	}
}

type authUseCase struct {
	txnAdapter          adapters.TransactionAdapter
	userAdapter         adapters.UserAdapter
	userPasswordAdapter adapters.UserPasswordAdapter
}

func (u *authUseCase) SignUp(ctx context.Context, userID, pw string, isAdmin bool, now time.Time) (*entities.User, error) {
	if exist, err := u.userAdapter.IsAlreadyUsedUserID(ctx, userID); err != nil {
		return nil, err
	} else if exist {
		return nil, errors.New(errors.UserIDAlreadyUsedError, nil)
	}

	hashedPw, err := password.HashPassword(pw, config.Env.PasswordSecretKey)
	if err != nil {
		return nil, err
	}

	var user *entities.User
	err = u.txnAdapter.BeginTxn(ctx, func(ctx2 context.Context) error {
		user, err = u.userAdapter.Create(ctx, entities.NewInitUser(userID, isAdmin))
		if err != nil {
			return err
		}
		return u.userPasswordAdapter.SetPassword(ctx, user.UserID, hashedPw, true, now)
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
