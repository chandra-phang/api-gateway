package orderservice

import (
	"api-gateway/api/controllers"
	"api-gateway/httpconnector"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartController struct {
	orderSvcCon *httpconnector.OrderServiceConnector
}

func InitCartController() *cartController {
	return &cartController{
		orderSvcCon: httpconnector.GetOrderServiceConnector(),
	}
}

func (c *cartController) AddToCart(ctx echo.Context) error {
	result, err := c.orderSvcCon.AddToCart(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}
