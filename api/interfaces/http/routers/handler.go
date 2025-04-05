package routers

import (
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/labstack/echo/v4"
)

func NewHandler() ServerInterface {
	return &handler{}
}

type handler struct {
}

func (h *handler) bind(ctx echo.Context, req any) error {
	if err := ctx.Bind(req); err != nil {
		return err
	}
	if err := ctx.Validate(req); err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	return nil
}

func (h *handler) Health(ctx echo.Context) error {
	panic("")
}
