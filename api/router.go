package api

import (
	v1AuthSvc "api-gateway/api/controllers/v1/auth_service"
	v1OrderSvc "api-gateway/api/controllers/v1/order_service"
	v1ProductSvc "api-gateway/api/controllers/v1/product_service"
	"api-gateway/api/middleware"
	"log"

	"github.com/labstack/echo/v4"
)

func InitRoutes() {
	e := echo.New()

	productV1 := v1ProductSvc.InitProductController()
	authV1 := v1AuthSvc.InitAuthController()
	orderV1 := v1OrderSvc.InitOrderController()
	cartV1 := v1OrderSvc.InitCartController()

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

	// OrderService API
	v1Api.POST("/carts", cartV1.AddToCart)
	v1Api.POST("/orders", orderV1.CreateOrder)
	v1Api.PUT("/orders/:id/cancel", orderV1.CancelOrder)
	v1Api.GET("/orders", orderV1.ListOrders)

	log.Println("Server is running at 8000 port.")
	e.Logger.Fatal(e.Start(":8000"))
}
