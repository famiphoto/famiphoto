package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/infrastructures/dbmodels"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"github.com/famiphoto/famiphoto/api/utils/cast"
	"time"
)

type UserPasswordAdapter interface {
	SetPassword(ctx context.Context, userID int64, hashedPassword string, isInitialized bool, now time.Time) error
}

func NewUserPasswordAdapter(userPasswordRepo repositories.UserPasswordRepository) UserPasswordAdapter {
	return &userPasswordAdapter{userPasswordRepo: userPasswordRepo}
}

type userPasswordAdapter struct {
	userPasswordRepo repositories.UserPasswordRepository
}

func (a *userPasswordAdapter) SetPassword(ctx context.Context, userID int64, hashedPassword string, isInitialized bool, now time.Time) error {
	return a.userPasswordRepo.Upsert(ctx, &dbmodels.UserPassword{
		UserID:         userID,
		Password:       hashedPassword,
		LastModifiedAt: now,
		IsInitialized:  cast.BoolToInt8(isInitialized),
	})
}
