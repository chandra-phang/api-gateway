package main

import "api-gateway/app"

func main() {
	application := app.NewApplication()
	application.InitApplication()
}
