package usecases

import (
	"context"
	"github.com/famiphoto/famiphoto/api/entities"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/infrastructures/adapters"
	"time"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userID, pw string, isAdmin bool, now time.Time) (*entities.User, error)
	DisableUser(ctx context.Context, userID string) error
	VerifyToSignIn(ctx context.Context, userID, pw string) (*entities.User, error)
}

func NewUserUseCase(
	txnAdapter adapters.TransactionAdapter,
	userAdapter adapters.UserAdapter,
	userPasswordAdapter adapters.UserPasswordAdapter,
) UserUseCase {
	return &userUseCase{
		txnAdapter:          txnAdapter,
		userAdapter:         userAdapter,
		userPasswordAdapter: userPasswordAdapter,
	}
}

type userUseCase struct {
	txnAdapter          adapters.TransactionAdapter
	userAdapter         adapters.UserAdapter
	userPasswordAdapter adapters.UserPasswordAdapter
}

func (u *userUseCase) CreateUser(ctx context.Context, userID, pw string, isAdmin bool, now time.Time) (*entities.User, error) {
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

func (u *userUseCase) DisableUser(ctx context.Context, userID string) error {
	if _, err := u.userAdapter.GetAvailableUser(ctx, userID); err != nil {
		return err
	}

	return u.userAdapter.UpdateStatus(ctx, userID, entities.UserStatusDisabled)
}

func (u *userUseCase) VerifyToSignIn(ctx context.Context, userID, pw string) (*entities.User, error) {
	isValid, err := u.userPasswordAdapter.VerifyPassword(ctx, userID, pw)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New(errors.UserAuthorizeError, nil)
	}

	user, err := u.userAdapter.GetAvailableUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
