package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/cast"
	"github.com/famiphoto/famiphoto/api/utils/password"
	"time"
)

type UserPasswordAdapter interface {
	SetPassword(ctx context.Context, userID, pw string, isInitialized bool, now time.Time) error
	VerifyPassword(ctx context.Context, userID, password string) (bool, error)
}

func NewUserPasswordAdapter(userPasswordRepo repositories.UserPasswordRepository) UserPasswordAdapter {
	return &userPasswordAdapter{userPasswordRepo: userPasswordRepo}
}

type userPasswordAdapter struct {
	userPasswordRepo repositories.UserPasswordRepository
}

func (a *userPasswordAdapter) SetPassword(ctx context.Context, userID, pw string, isInitialized bool, now time.Time) error {
	hashedPw, err := password.HashPassword(pw, config.Env.PasswordSecretKey)
	if err != nil {
		return err
	}

	return a.userPasswordRepo.Upsert(ctx, &dbmodels.UserPassword{
		UserID:         userID,
		Password:       hashedPw,
		LastModifiedAt: now,
		IsInitialized:  cast.BoolToInt8(isInitialized),
	})
}

func (a *userPasswordAdapter) VerifyPassword(ctx context.Context, userID, pw string) (bool, error) {
	userPassword, err := a.userPasswordRepo.Get(ctx, userID)
	if err != nil {
		return false, err
	}

	return password.MatchPassword(pw, userPassword.Password, config.Env.PasswordSecretKey)
}
