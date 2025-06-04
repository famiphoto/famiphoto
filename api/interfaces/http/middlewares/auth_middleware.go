package middlewares

import (
	"github.com/famiphoto/famiphoto/api/errors"
	"github.com/famiphoto/famiphoto/api/interfaces/http/sessions"
	"github.com/famiphoto/famiphoto/api/usecases"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	AuthUser(next echo.HandlerFunc) echo.HandlerFunc
	AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
	userUseCase usecases.UserUseCase
}

func NewAuthMiddleware(userUseCase usecases.UserUseCase) AuthMiddleware {
	return &authMiddleware{
		userUseCase: userUseCase,
	}
}

func (m *authMiddleware) AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := sessions.GetUserID(c)
		if userID == "" {
			return errors.New(errors.UserAuthorizeError, nil)
		}

		if err := m.userUseCase.VerifyUser(c.Request().Context(), userID); err != nil {
			return err
		}

		return next(c)
	}
}

func (m *authMiddleware) AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := sessions.GetUserID(c)
		if userID == "" {
			return errors.New(errors.UserAuthorizeError, nil)
		}

		if err := m.userUseCase.VerifyAdminUser(c.Request().Context(), userID); err != nil {
			return err
		}

		return next(c)
	}
}
