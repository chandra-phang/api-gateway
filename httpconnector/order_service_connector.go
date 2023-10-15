package httpconnector

import (
	"api-gateway/config"
	"api-gateway/dto/response"
	"api-gateway/request"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var orderServiceCon *OrderServiceConnector

type OrderServiceConnector struct {
	Host           string
	AddToCartUri   string
	CreateOrderUri string
	CancelOrderUri string
	ListOrdersUri  string
}

func InitOrderServiceConnector(cfg config.Config) {
	orderServiceCon = &OrderServiceConnector{
		Host:           cfg.OrderSvcHost,
		AddToCartUri:   cfg.AddToCartUri,
		CreateOrderUri: cfg.CreateOrderUri,
		CancelOrderUri: cfg.CancelOrderUri,
		ListOrdersUri:  cfg.ListOrdersUri,
	}
}

func GetOrderServiceConnector() *OrderServiceConnector {
	return orderServiceCon
}

func (con OrderServiceConnector) CreateOrder(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.CreateOrderUri

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.Post(url, reqBody, authorization)
	if err != nil {
		return nil, err
	}

	// Early return if statusCode is not CREATED
	if statusCode != http.StatusCreated {
		var response response.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return nil, err
		}
		return nil, errors.New(response.Failure)
	}

	// Deserialize the response
	var response response.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (con OrderServiceConnector) CancelOrder(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")
	url := con.Host + con.CancelOrderUri
	url = strings.Replace(url, ":id", id, 1)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.Put(url, reqBody, authorization)
	if err != nil {
		return nil, err
	}

	// Early return if statusCode is not OK
	if statusCode != http.StatusOK {
		var response response.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return nil, err
		}
		return nil, errors.New(response.Failure)
	}

	// Deserialize the response
	var response response.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (con OrderServiceConnector) ListOrders(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.ListOrdersUri

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.Get(url, authorization)
	if err != nil {
		return nil, err
	}

	// Early return if statusCode is not OK
	if statusCode != http.StatusOK {
		var response response.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return nil, err
		}
		return nil, errors.New(response.Failure)
	}

	// Deserialize the response
	var response response.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (con OrderServiceConnector) AddToCart(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.AddToCartUri

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.Post(url, reqBody, authorization)
	if err != nil {
		return nil, err
	}

	// Early return if statusCode is not CREATED
	if statusCode != http.StatusCreated {
		var response response.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			return nil, err
		}
		return nil, errors.New(response.Failure)
	}

	// Deserialize the response
	var response response.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
