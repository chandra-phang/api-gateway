package app

import (
	"api-gateway/api"
	"api-gateway/config"
	"api-gateway/httpconnector"
)

type Application struct {
}

// Returns a new instance of the application
func NewApplication() Application {
	return Application{}
}

func (a Application) InitApplication() {
	cfg := config.InitConfig()

	httpconnector.InitAuthServiceConnector(*cfg)

	api.InitRoutes()
}
