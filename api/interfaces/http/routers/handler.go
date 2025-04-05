package routers

import "github.com/labstack/echo/v4"

func NewHandler() ServerInterface {
	return &handler{}
}

type handler struct {
}

func (h *handler) Health(ctx echo.Context) error {
	panic("")
}
