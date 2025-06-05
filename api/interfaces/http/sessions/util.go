package sessions

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/errors"

	"github.com/famiphoto/famiphoto/api/utils/cast"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func getInt64(ctx echo.Context, sessionName, key string, defaultValue int64) int64 {
	val, err := get(ctx, sessionName, key)
	if err != nil {
		return defaultValue
	}
	if val == nil {
		return defaultValue
	}

	dst, err := cast.ToInt64(val)
	if err != nil {
		log.Error("failed to cast session data, key: ", key, err)
		return defaultValue
	}
	return dst
}

func getString(ctx echo.Context, sessionName, key string, defaultValue string) string {
	val, err := get(ctx, sessionName, key)
	if err != nil {
		return defaultValue
	}
	if val == nil {
		return defaultValue
	}

	dst, err := cast.ToString(val)
	if err != nil {
		log.Error("failed to cast session data, key: ", key, err)
		return defaultValue
	}
	return dst
}

func get(c echo.Context, sessionName, key string) (any, error) {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return nil, err
	}
	val, ok := sess.Values[key]
	if !ok {
		return nil, nil
	}
	return val, nil
}

func set[T any](c echo.Context, sessionName, key string, value T) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return err
	}
	sess.Values[key] = value
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return errors.New(errors.SessionFatal, fmt.Errorf("fatal to set common session, key: %s", key))
	}
	return nil
}

func del[T any](c echo.Context, sessionName, key string) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return err
	}

	_, exist := sess.Values[key]
	if !exist {
		return nil
	}

	delete(sess.Values, key)
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return errors.New(errors.SessionFatal, fmt.Errorf("fatal to delete session, key: %s", key))
	}
	return nil
}

func expireSession(ctx echo.Context, sessionName string) error {
	sess, err := session.Get(sessionName, ctx)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		return err
	}
	return nil
}
