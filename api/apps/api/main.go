package main

import (
	"fmt"
	"github.com/famiphoto/famiphoto/api/config"
	"github.com/famiphoto/famiphoto/api/di"
	"github.com/labstack/gommon/log"
)

func main() {
	app := di.NewAPIRouter()
	if err := app.Start(fmt.Sprintf(":%d", config.Env.Port)); err != nil {
		log.Error(err)
	}
}
