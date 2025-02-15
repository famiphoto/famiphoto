package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/valkey-io/valkey-go"
	"time"
)

type SessionRepository interface {
	Save(ctx context.Context, sessionID string, data map[any]any, expire time.Duration) error
	Get(ctx context.Context, sessionID string) (map[any]any, error)
	Delete(ctx context.Context, sessionID string) error
	AddSessionID(ctx context.Context, userID int64, sessionID string) error
	RemoveSessionID(ctx context.Context, userID int64, sessionID string) error
	GetSessionIDs(ctx context.Context, userID int64) ([]string, error)
}

func NewSessionRepository(client valkey.Client) SessionRepository {
	return &sessionRepository{
		client: client,
	}
}

type sessionRepository struct {
	client valkey.Client
}

func (r *sessionRepository) Save(ctx context.Context, sessionID string, data map[any]any, expire time.Duration) error {
	serialized, err := json.Marshal(data)
	if err != nil {
		return err
	}
	stmt := r.client.B().Set().Key(r.genSessionKey(sessionID)).Value(string(serialized)).Ex(expire).Build()
	resp := r.client.Do(ctx, stmt)
	return resp.Error()
}

func (r *sessionRepository) Get(ctx context.Context, sessionID string) (map[any]any, error) {
	stmt := r.client.B().Get().Key(r.genSessionKey(sessionID)).Build()
	resp := r.client.Do(ctx, stmt)
	if resp.Error() != nil {
		if valkey.IsValkeyNil(resp.Error()) {
			return nil, errors.New(errors.DBNotFoundError, resp.Error())
		}
		return nil, resp.Error()
	}

	var val map[any]any
	if err := json.Unmarshal([]byte(resp.String()), &val); err != nil {
		return nil, err
	}
	return val, nil
}

func (r *sessionRepository) Delete(ctx context.Context, sessionID string) error {
	stmt := r.client.B().Del().Key(r.genSessionKey(sessionID)).Build()
	resp := r.client.Do(ctx, stmt)
	if resp.Error() != nil {
		if valkey.IsValkeyNil(resp.Error()) {
			return nil
		}
		return resp.Error()
	}
	return nil
}

func (r *sessionRepository) AddSessionID(ctx context.Context, userID int64, sessionID string) error {
	stmt := r.client.B().Sadd().Key(r.genSessionSetKey(userID)).Member(sessionID).Build()
	resp := r.client.Do(ctx, stmt)
	return resp.Error()
}

func (r *sessionRepository) RemoveSessionID(ctx context.Context, userID int64, sessionID string) error {
	stmt := r.client.B().Srem().Key(r.genSessionSetKey(userID)).Member(sessionID).Build()
	resp := r.client.Do(ctx, stmt)
	if resp.Error() != nil {
		if valkey.IsValkeyNil(resp.Error()) {
			return nil
		}
		return resp.Error()
	}
	return nil
}

func (r *sessionRepository) GetSessionIDs(ctx context.Context, userID int64) ([]string, error) {
	stmt := r.client.B().Smembers().Key(r.genSessionSetKey(userID)).Build()
	resp := r.client.Do(ctx, stmt)
	if resp.Error() != nil {
		if valkey.IsValkeyNil(resp.Error()) {
			return []string{}, nil
		}
		return nil, resp.Error()
	}
	return resp.AsStrSlice()
}

func (r *sessionRepository) genSessionKey(sessionID string) string {
	return fmt.Sprintf("famiphoto:session:%s", sessionID)
}

func (r *sessionRepository) genSessionSetKey(userID int64) string {
	return fmt.Sprintf("famiphoto:session_set:%d", userID)
}
