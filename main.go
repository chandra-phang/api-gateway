package main

import (
	"api-gateway/config"
	"api-gateway/request"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.InitConfig()

	// ProductService API
	e.GET("apigw/v1/products", request.ListProducts)
	e.GET("apigw/v1/products/:id", request.GetProduct)
	e.POST("apigw/v1/products", request.CreateProduct)
	e.PUT("apigw/v1/products/:id", request.UpdateProduct)
	// e.PUT("apigw/v1/products/:id/disable", request.DisableProduct)
	// e.PUT("apigw/v1/products/:id/enable", request.EnableProduct)
	// e.PUT("apigw/v1/products/:id/increase-booked-quota", request.IncreaseBookedQuota)
	// e.PUT("apigw/v1/products/:id/decrease-booked-quota", request.DecreaseBookedQuota)

	log.Println("Server is running at 8000 port.")
	e.Logger.Fatal(e.Start(":8000"))
}
