package usecases

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"time"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, userID, pw string, isAdmin bool, now time.Time) (*entities.User, error)
	SignIn(ctx context.Context, userID, pw string) (*entities.User, error)
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

	var user *entities.User
	var err error
	err = u.txnAdapter.BeginTxn(ctx, func(ctx2 context.Context) error {
		user, err = u.userAdapter.Create(ctx2, entities.NewInitUser(userID, isAdmin))
		if err != nil {
			return err
		}
		return u.userPasswordAdapter.SetPassword(ctx2, user.UserID, pw, true, now)
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *authUseCase) SignIn(ctx context.Context, userID, pw string) (*entities.User, error) {
	isValid, err := u.userPasswordAdapter.VerifyPassword(ctx, userID, pw)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New(errors.UserAuthorizeError, nil)
	}

	user, err := u.userAdapter.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
