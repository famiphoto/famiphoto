package middlewares

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	AuthUser(next echo.HandlerFunc) echo.HandlerFunc
	AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc
}
