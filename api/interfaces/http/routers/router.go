package routers

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/validators"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

type Router interface {
	Start(address string) error
}

func NewAPIRouter(sessionStore sessions.Store, handler ServerInterface) Router {
	r := &apiRouter{
		echo:         echo.New(),
		sessionStore: sessionStore,
		handler:      handler,
	}
	return r
}

type apiRouter struct {
	echo         *echo.Echo
	sessionStore sessions.Store
	handler      ServerInterface
}

func (r *apiRouter) Start(address string) error {
	r.setMiddleware(r.echo)
	r.route(r.echo, r.handler)
	return r.echo.Start(address)
}

func (r *apiRouter) setMiddleware(e *echo.Echo) {
	e.HTTPErrorHandler = ErrorHandle
	e.Validator = validators.NewValidator()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echotrace.Middleware())
	e.Use(middleware.Logger())
	e.Use(session.Middleware(r.sessionStore))
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
}

func (r *apiRouter) route(e EchoRouter, si ServerInterface) {

	w := ServerInterfaceWrapper{
		Handler: si,
	}

	e.POST("sign_up", w.SignUp)
	e.POST("sign_in", w.SignIn)
	e.POST("sign_out", w.SignOut)
}
