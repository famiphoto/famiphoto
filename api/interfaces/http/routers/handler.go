package routers

import (
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/famiphoto/famiphoto/api/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewHandler(userUseCase usecases.UserUseCase) schema.ServerInterface {
	return &handler{
		userUseCase: userUseCase,
	}
}

type handler struct {
	userUseCase usecases.UserUseCase
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

func (h *handler) HealthGetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &schema.HealthResponse{Status: "OK"})
}
