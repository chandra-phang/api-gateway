package controller

import (
	"api-gateway/api/controllers"
	"api-gateway/config"
	"api-gateway/dto/response"
	"api-gateway/dto/response/v1/product"
	"api-gateway/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productController struct {
}

func InitProductController() *productController {
	return &productController{}
}

func (c *productController) GetProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s", config.ProductSvcHost, id)

	resp, statusCode, err := request.Get(url)
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
	var response product.GetProductResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *productController) ListProducts(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products", config.ProductSvcHost)

	resp, statusCode, err := request.Get(url)
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
	var response product.ListProductsResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return controllers.WriteError(ctx, statusCode, err)
	}

	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *productController) CreateProduct(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products", config.ProductSvcHost)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := request.Post(url, reqBody)
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

	// Return the product data as JSON
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *productController) UpdateProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s", config.ProductSvcHost, id)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := request.Put(url, reqBody)
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
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *productController) DisableProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s/disable", config.ProductSvcHost, id)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := request.Put(url, reqBody)
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
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *productController) EnableProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s/enable", config.ProductSvcHost, id)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := request.Put(url, reqBody)
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
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *productController) IncreaseBookedQuota(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s/increase-booked-quota", config.ProductSvcHost, id)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := request.Put(url, reqBody)
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
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *productController) DecreaseBookedQuota(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s/decrease-booked-quota", config.ProductSvcHost, id)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := request.Put(url, reqBody)
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
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}
