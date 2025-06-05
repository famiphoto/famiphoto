package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *handler) AdminUserManagementCreateUser(ctx echo.Context) error {
	req := new(schema.AdminCreateUserRequest)
	if err := h.bind(ctx, req); err != nil {
		return err
	}

	user, err := h.userUseCase.CreateUser(ctx.Request().Context(), req.UserId, req.Password, req.GetIsAdminOrDefault(false), time.Now())
	if err != nil {
		return err
	}

	res := &schema.AdminCreateUserResponse{
		UserId:  user.UserID,
		IsAdmin: user.IsAdmin,
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *handler) AdminUserManagementDeleteUser(ctx echo.Context, userId string) error {
	if err := h.userUseCase.DisableUser(ctx.Request().Context(), userId); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
