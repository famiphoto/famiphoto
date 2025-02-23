package adapters

import (
	"context"
	"github.com/famiphoto/famiphoto/api/infrastructures/repositories"
	"time"
)

type SessionAdapter interface {
	LoadSession(ctx context.Context, sessionID string) (map[any]any, error)
	SaveSession(ctx context.Context, sessionID string, userID int64, values map[any]any, age int) error
	DeleteSession(ctx context.Context, sessionID string, userID int64) error
	DeleteSessionAll(ctx context.Context, userID int64) error
}

func NewSessionAdapter(sessionRepo repositories.SessionRepository) SessionAdapter {
	return &sessionAdapter{sessionRepo: sessionRepo}
}

type sessionAdapter struct {
	sessionRepo repositories.SessionRepository
}

func (a *sessionAdapter) LoadSession(ctx context.Context, sessionID string) (map[any]any, error) {
	return a.sessionRepo.Get(ctx, sessionID)
}

func (a *sessionAdapter) SaveSession(ctx context.Context, sessionID string, userID int64, values map[any]any, age int) error {
	if err := a.sessionRepo.Save(ctx, sessionID, values, time.Duration(age)*time.Second); err != nil {
		return err
	}
	if userID > 0 {
		if err := a.sessionRepo.AddSessionID(ctx, userID, sessionID); err != nil {
			return err
		}
	}
	return nil
}

func (a *sessionAdapter) DeleteSession(ctx context.Context, sessionID string, userID int64) error {
	if err := a.sessionRepo.Delete(ctx, sessionID); err != nil {
		return err
	}
	if userID > 0 {
		if err := a.sessionRepo.RemoveSessionID(ctx, userID, sessionID); err != nil {
			return err
		}
	}
	return nil
}

func (a *sessionAdapter) DeleteSessionAll(ctx context.Context, userID int64) error {
	sessionIDs, err := a.sessionRepo.GetSessionIDs(ctx, userID)
	if err != nil {
		return err
	}
	for _, sessionID := range sessionIDs {
		err = a.DeleteSession(ctx, sessionID, userID)
		if err != nil {
			return err
		}
	}
	return nil
}
