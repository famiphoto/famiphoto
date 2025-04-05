package middlewares

import (
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	AuthUser(next echo.HandlerFunc) echo.HandlerFunc
	AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

func (m *authMiddleware) AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func (m *authMiddleware) AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
