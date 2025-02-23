package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/handlers"
	"github.com/famiphoto/famiphoto/api/interfaces/http/middlewares"
	"github.com/famiphoto/famiphoto/api/interfaces/http/validators"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
	"net/http"
)

type Router interface {
	Start(address string) error
}

func NewAPIRouter(sessionStore sessions.Store) Router {
	r := &apiRouter{
		echo:         echo.New(),
		sessionStore: sessionStore,
	}
	return r
}

type apiRouter struct {
	echo         *echo.Echo
	sessionStore sessions.Store
	authHandler  handlers.AuthHandler
}

func (r *apiRouter) Start(address string) error {
	return r.echo.Start(address)
}

func (r *apiRouter) route() {
	r.echo.HTTPErrorHandler = middlewares.HandleError
	r.echo.Validator = validators.NewValidator()
	r.echo.Pre(middleware.RemoveTrailingSlash())
	r.echo.Use(echotrace.Middleware())
	r.echo.Use(middleware.Logger())
	r.echo.Use(session.Middleware(r.sessionStore))
	r.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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
	r.echo.Use(middleware.Recover())

	r.echo.GET("status", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	auth := r.echo.Group("/auth")
	auth.POST("sign_up", r.authHandler.SignUp)
	auth.POST("sign_in", r.authHandler.SignIn)
	auth.POST("sign_out", r.authHandler.SignOut)
}
