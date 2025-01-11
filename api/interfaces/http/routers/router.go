package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/middlewares"
	"github.com/famiphoto/famiphoto/api/interfaces/http/validators"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

var Router *echo.Echo

func init() {
	e := echo.New()
	e.HTTPErrorHandler = middlewares.HandleError
	e.Validator = validators.NewValidator()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echotrace.Middleware())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:        nil,
		BeforeNextFunc: nil,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Error(v.Method, v.URIPath, v.Error)
			}
			return nil
		},
		LogLatency:   true,
		LogProtocol:  true,
		LogRemoteIP:  false,
		LogHost:      false,
		LogMethod:    true,
		LogURI:       true,
		LogURIPath:   true,
		LogRoutePath: true,
		LogStatus:    true,
		LogError:     true,
	}))
	e.Use(middleware.Recover())

	Router = e
}
