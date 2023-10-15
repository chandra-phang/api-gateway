package orderservice

import (
	"api-gateway/api/controllers"
	"api-gateway/httpconnector"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderController struct {
	orderSvcCon *httpconnector.OrderServiceConnector
}

func InitOrderController() *orderController {
	return &orderController{
		orderSvcCon: httpconnector.GetOrderServiceConnector(),
	}
}

func (c *orderController) CreateOrder(ctx echo.Context) error {
	result, err := c.orderSvcCon.CreateOrder(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *orderController) CancelOrder(ctx echo.Context) error {
	result, err := c.orderSvcCon.CancelOrder(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *orderController) ListOrders(ctx echo.Context) error {
	result, err := c.orderSvcCon.ListOrders(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}
