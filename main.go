package main

import (
	"api-gateway/config"
	v1 "api-gateway/request/v1"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.InitConfig()

	// ProductService API
	e.GET("apigw/v1/products", v1.ListProducts)
	e.GET("apigw/v1/products/:id", v1.GetProduct)
	e.POST("apigw/v1/products", v1.CreateProduct)
	e.PUT("apigw/v1/products/:id", v1.UpdateProduct)
	e.PUT("apigw/v1/products/:id/disable", v1.DisableProduct)
	e.PUT("apigw/v1/products/:id/enable", v1.EnableProduct)
	e.PUT("apigw/v1/products/:id/increase-booked-quota", v1.IncreaseBookedQuota)
	e.PUT("apigw/v1/products/:id/decrease-booked-quota", v1.DecreaseBookedQuota)

	log.Println("Server is running at 8000 port.")
	e.Logger.Fatal(e.Start(":8000"))
}
