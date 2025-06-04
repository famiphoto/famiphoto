package sessions

import (
	"github.com/famiphoto/famiphoto/api/utils/cast"
	"github.com/labstack/echo/v4"
)

const appSessionName = "famiphoto_session"

const (
	userID  = "user_id"
	isAdmin = "is_admin"
)

func SetUserID(ctx echo.Context, v string) error {
	return set(ctx, appSessionName, userID, v)
}

func GetUserID(ctx echo.Context) string {
	return getString(ctx, appSessionName, userID, "")
}

func SetIsAdmin(ctx echo.Context, v bool) error {
	return set(ctx, appSessionName, isAdmin, cast.BoolToInt8(v))
}

func GetIsAdmin(ctx echo.Context) bool {
	return cast.IntToBool(getInt64(ctx, appSessionName, isAdmin, 0))
}

// ExpireSession セッションを削除してサインアウト状態にします。
func ExpireSession(ctx echo.Context) error {
	return expireSession(ctx, appSessionName)
}
