package api

import (
	v1 "api-gateway/api/controllers/v1"
	"api-gateway/api/middleware"
	"log"

	"github.com/labstack/echo/v4"
)

func InitRoutes() {
	e := echo.New()

	productV1 := v1.InitProductController()
	authV1 := v1.InitAuthController()

	v1Api := e.Group("apigw/v1")
	v1Api.POST("/login", authV1.Login)
	v1Api.POST("/logout", authV1.Logout)
	v1Api.POST("/authenticate", authV1.Authenticate)

	v1Api.Use(middleware.AuthMiddleware)

	// ProductService API
	v1Api.GET("/products", productV1.ListProducts)
	v1Api.GET("/products/:id", productV1.GetProduct)
	v1Api.POST("/products", productV1.CreateProduct)
	v1Api.PUT("/products/:id", productV1.UpdateProduct)
	v1Api.PUT("/products/:id/disable", productV1.DisableProduct)
	v1Api.PUT("/products/:id/enable", productV1.EnableProduct)
	v1Api.PUT("/products/:id/increase-booked-quota", productV1.IncreaseBookedQuota)
	v1Api.PUT("/products/:id/decrease-booked-quota", productV1.DecreaseBookedQuota)

	log.Println("Server is running at 8000 port.")
	e.Logger.Fatal(e.Start(":8000"))
}
