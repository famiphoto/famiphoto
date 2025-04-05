package sessions

import (
	"fmt"

	"github.com/famiphoto/famiphoto/api/utils/cast"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

func getInt64(values map[any]any, key string, defaultValue int64) int64 {
	val, ok := values[key]
	if !ok {
		return defaultValue
	}
	dst, err := cast.ToInt64(val)
	if err != nil {
		fmt.Println("Failed to cast session data. key: ", key, err)
		return defaultValue
	}
	return dst
}

func getString(values map[any]any, key string, defaultValue string) string {
	val, ok := values[key]
	if !ok {
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
		return errors.Wrap(err, fmt.Sprintf("Fatal to set session, key: %s", key))
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
		return errors.Wrap(err, fmt.Sprintf("Fatal to delete session, key: %s", key))
	}
	return nil
}
