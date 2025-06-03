package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *handler) AuthPostSignUp(ctx echo.Context) error {
	req := new(schema.AuthSignUpRequest)
	if err := h.bind(ctx, req); err != nil {
		return err
	}

	isAdmin := req.GetIsAdminOrDefault(false)

	user, err := h.authUseCase.SignUp(ctx.Request().Context(), req.UserId, req.Password, isAdmin, time.Now())
	if err != nil {
		return err
	}

	res := &schema.AuthSignUpResponse{
		IsAdmin: user.IsAdmin,
		UserId:  user.UserID,
	}
	return ctx.JSON(http.StatusOK, res)
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
