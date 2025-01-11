package main

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/interfaces/http/routers"
	"github.com/labstack/gommon/log"
)

func main() {
	app := routers.Router
	if err := app.Start(fmt.Sprintf(":%d", config.Env.Port)); err != nil {
		app.Logger.Fatal(err)
		log.Error(err)
	}
}
