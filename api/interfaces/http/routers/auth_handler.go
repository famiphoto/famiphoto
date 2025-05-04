package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/labstack/echo/v4"
)

func (h *handler) SignUp(ctx echo.Context) error {
	req := new(schema.SignUpJSONRequestBody)
	if err := h.bind(ctx, req); err != nil {
		return err
	}
	panic("")
}

func (h *handler) SignIn(ctx echo.Context) error {
	panic("")
}

func (h *handler) SignOut(ctx echo.Context) error {
	panic("")
}

func (h *handler) Me(ctx echo.Context) error {
	panic("")
}
