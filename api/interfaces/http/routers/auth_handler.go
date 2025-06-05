package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/famiphoto/famiphoto/api/interfaces/http/sessions"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) AuthPostSignIn(ctx echo.Context) error {
	req := new(schema.AuthSignInRequest)
	if err := h.bind(ctx, req); err != nil {
		return err
	}

	user, err := h.userUseCase.VerifyToSignIn(ctx.Request().Context(), req.UserId, req.Password)
	if err != nil {
		return err
	}

	// Set user ID and isAdmin in session
	if err := sessions.SetUserID(ctx, user.UserID); err != nil {
		return err
	}
	if err := sessions.SetIsAdmin(ctx, user.IsAdmin); err != nil {
		return err
	}

	res := &schema.AuthSignInResponse{
		IsAdmin: user.IsAdmin,
		UserId:  user.UserID,
	}
	return ctx.JSON(http.StatusOK, res)
}

func (h *handler) AuthPostSignOut(ctx echo.Context) error {
	if err := sessions.ExpireSession(ctx); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (h *handler) AuthGetMe(ctx echo.Context) error {
	// GetAvailableUser user ID from session
	userID := sessions.GetUserID(ctx)
	if userID == "" {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	// GetAvailableUser isAdmin from session
	isAdmin := sessions.GetIsAdmin(ctx)

	res := &schema.AuthMeResponse{
		UserId:  userID,
		IsAdmin: isAdmin,
	}
	return ctx.JSON(http.StatusOK, res)
}
