package orderservice

import (
	"api-gateway/api/controllers"
	"api-gateway/config"
	"api-gateway/dto/response"
	"api-gateway/dto/response/v1/order"
	"api-gateway/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderController struct {
}

func InitOrderController() *orderController {
	return &orderController{}
}

func (c *orderController) CreateOrder(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/orders", config.OrderSvcHost)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.PostWithAuthorization(url, reqBody, authorization)
	if err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}

	// Early return if statusCode is not CREATED
	if statusCode != http.StatusCreated {
		var response response.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return controllers.WriteError(ctx, statusCode, err)
		}
		return controllers.WriteErrorMsg(ctx, statusCode, response.Failure)
	}

	// Deserialize the response
	var response response.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}

	// Return the order data as JSON
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *orderController) CancelOrder(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/orders/%s/cancel", config.OrderSvcHost, id)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.PutWithAuthorization(url, reqBody, authorization)
	if err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}

	// Early return if statusCode is not OK
	if statusCode != http.StatusOK {
		var response response.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return controllers.WriteError(ctx, statusCode, err)
		}
		return controllers.WriteErrorMsg(ctx, statusCode, response.Failure)
	}

	// Deserialize the response
	var response response.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}

	// Return the order data as JSON
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *orderController) ListOrders(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/orders", config.OrderSvcHost)

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.GetWithAuthorization(url, authorization)
	if err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}

	// Early return if statusCode is not OK
	if statusCode != http.StatusOK {
		var response response.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return controllers.WriteError(ctx, statusCode, err)
		}
		return controllers.WriteErrorMsg(ctx, statusCode, response.Failure)
	}

	// Deserialize the response
	var response order.ListOrdersResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}

	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}
