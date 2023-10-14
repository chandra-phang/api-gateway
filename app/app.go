package app

import (
	"api-gateway/api"
	"api-gateway/config"
)

type Application struct {
}

// Returns a new instance of the application
func NewApplication() Application {
	return Application{}
}

func (a Application) InitApplication() {
	config.InitConfig()

	api.InitRoutes()
}
