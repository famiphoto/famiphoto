package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/labstack/echo/v4"
)

func (h *handler) AuthPostSignUp(ctx echo.Context) error {
	req := new(schema.AuthSignInRequest)
	if err := h.bind(ctx, req); err != nil {
		return err
	}
	panic("")
}

func (h *handler) AuthPostSignIn(ctx echo.Context) error {
	panic("")
}

func (h *handler) AuthPostSignOut(ctx echo.Context) error {
	panic("")
}

func (h *handler) AuthGetMe(ctx echo.Context) error {
	panic("")
}
