package handlers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/requests"
	"github.com/famiphoto/famiphoto/api/usecases"
	"github.com/labstack/echo/v4"
	"time"
)

type AuthHandler interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	SignOut(c echo.Context) error
}

type authHandler struct {
	authUseCase usecases.AuthUseCase
}

func (h *authHandler) SignUp(c echo.Context) error {
	req := new(requests.SignUpRequest)
	if err := req.Bind(c); err != nil {
		return err
	}

	h.authUseCase.SignUp(c.Request().Context(), req.MyID, req.Password, req.IsAdmin, time.Now())

	panic("")
}

func (h *authHandler) SignIn(c echo.Context) error {
	req := new(requests.SignInRequest)
	if err := req.Bind(c); err != nil {
		return err
	}

	panic("")
}

func (h *authHandler) SignOut(c echo.Context) error {
	panic("")
}
