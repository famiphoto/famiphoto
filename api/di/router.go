package di

import (
	"github.com/famiphoto/famiphoto/api/interfaces/http/routers"
	"github.com/famiphoto/famiphoto/api/interfaces/http/sessions"
)

func NewAPIRouter() routers.Router {
	return routers.NewAPIRouter(sessions.NewStore(NewSessionAdapter()))
}
