package httpconnector

import (
	"api-gateway/config"
	"api-gateway/dto/response"
	"api-gateway/dto/response/v1/product"
	"api-gateway/request"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var productServiceCon *ProductServiceConnector

type ProductServiceConnector struct {
	Host                   string
	ListProductUri         string
	CreateProductUri       string
	GetProductUri          string
	UpdateProductUri       string
	DisableProductUri      string
	EnableProductUri       string
	IncreaseBookedQuotaUri string
	DecreaseBookedQuotaUri string
}

func InitProductServiceConnector(cfg config.Config) {
	productServiceCon = &ProductServiceConnector{
		Host:                   cfg.ProductSvcHost,
		ListProductUri:         cfg.ListProductUri,
		CreateProductUri:       cfg.CreateProductUri,
		GetProductUri:          cfg.GetProductUri,
		UpdateProductUri:       cfg.UpdateProductUri,
		DisableProductUri:      cfg.DisableProductUri,
		EnableProductUri:       cfg.EnableProductUri,
		IncreaseBookedQuotaUri: cfg.IncreaseBookedQuotaUri,
		DecreaseBookedQuotaUri: cfg.DecreaseBookedQuotaUri,
	}
}

func GetProductServiceConnector() *ProductServiceConnector {
	return productServiceCon
}

func (con ProductServiceConnector) GetProduct(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")
	url := con.Host + con.GetProductUri
	url = strings.Replace(url, ":id", id, 1)

	resp, statusCode, err := request.Get(url, "")
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
	var response product.GetProductResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (con ProductServiceConnector) ListProducts(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.ListProductUri

	resp, statusCode, err := request.Get(url, "")
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
	var response product.ListProductsResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (con ProductServiceConnector) CreateProduct(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.CreateProductUri

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	resp, statusCode, err := request.Post(url, reqBody, "")
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

func (con ProductServiceConnector) UpdateProduct(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")
	url := con.Host + con.UpdateProductUri
	url = strings.Replace(url, ":id", id, 1)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	resp, statusCode, err := request.Put(url, reqBody, "")
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

func (con ProductServiceConnector) DisableProduct(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")
	url := con.Host + con.DisableProductUri
	url = strings.Replace(url, ":id", id, 1)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	resp, statusCode, err := request.Put(url, reqBody, "")
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

func (con ProductServiceConnector) EnableProduct(ctx echo.Context) (interface{}, error) {
	id := ctx.Param("id")
	url := con.Host + con.EnableProductUri
	url = strings.Replace(url, ":id", id, 1)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	resp, statusCode, err := request.Put(url, reqBody, "")
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

func (con ProductServiceConnector) IncreaseBookedQuota(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.IncreaseBookedQuotaUri

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	resp, statusCode, err := request.Put(url, reqBody, "")
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

func (con ProductServiceConnector) DecreaseBookedQuota(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.DecreaseBookedQuotaUri

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	resp, statusCode, err := request.Put(url, reqBody, "")
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
