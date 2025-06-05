package di

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/middlewares"
	"github.com/famiphoto/famiphoto/api/interfaces/http/routers"
	"github.com/famiphoto/famiphoto/api/interfaces/http/schema"
	"github.com/famiphoto/famiphoto/api/interfaces/http/sessions"
)

func NewAPIRouter() routers.Router {
	return routers.NewAPIRouter(sessions.NewStore(NewSessionAdapter()), NewHandler(), NewAuthMiddleware())
}

func NewHandler() schema.ServerInterface {
	return routers.NewHandler(NewUserUseCase())
}

func NewAuthMiddleware() middlewares.AuthMiddleware {
	return middlewares.NewAuthMiddleware(NewUserUseCase())
}
