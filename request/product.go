package request

import (
	"api-gateway/config"
	"api-gateway/dto"
	"api-gateway/dto/product"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s", config.ProductSvcHost, id)

	resp, statusCode, err := GetRequest(url)
	if err != nil {
		fmt.Println("Error when getting product", err.Error())
		return ctx.JSON(statusCode, map[string]string{"error": err.Error()})
	}

	// Early return if statusCode is not OK
	if statusCode != http.StatusOK {
		var response dto.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			fmt.Println("Error when unmarshalling response", err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		return ctx.JSON(statusCode, response)
	}

	// Deserialize the response
	var response product.GetProductResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Error when unmarshalling response", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return ctx.JSON(statusCode, response)
}

func ListProducts(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products", config.ProductSvcHost)

	resp, statusCode, err := GetRequest(url)
	if err != nil {
		fmt.Println("Error when getting products", err.Error())
		return ctx.JSON(statusCode, map[string]string{"error": err.Error()})
	}

	// Early return if statusCode is not OK
	if statusCode != http.StatusOK {
		var response dto.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			fmt.Println("Error when unmarshalling response", err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		return ctx.JSON(statusCode, response)
	}

	// Deserialize the response
	var response product.ListProductsResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Error when unmarshalling response", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return ctx.JSON(statusCode, response)
}

func CreateProduct(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products", config.ProductSvcHost)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := PostRequest(url, reqBody)
	if err != nil {
		fmt.Println("Error when getting products", err.Error())
		return ctx.JSON(statusCode, map[string]string{"error": err.Error()})
	}

	// Early return if statusCode is not CREATED
	if statusCode != http.StatusCreated {
		var response dto.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			fmt.Println("Error when unmarshalling response", err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		return ctx.JSON(statusCode, response)
	}

	// Deserialize the response
	var response dto.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Error when unmarshalling response", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	// Return the product data as JSON
	return ctx.JSON(statusCode, response)
}

func UpdateProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/products/%s", config.ProductSvcHost, id)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := PutRequest(url, reqBody)
	if err != nil {
		fmt.Println("Error when update product", err.Error())
		return ctx.JSON(statusCode, map[string]string{"error": err.Error()})
	}

	// Early return if statusCode is not OK
	if statusCode == http.StatusOK {
		var response dto.FailureResponse
		if err := json.Unmarshal(resp, &response); err != nil {
			fmt.Println("Error when unmarshalling response", err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		return ctx.JSON(statusCode, response)
	}

	// Deserialize the response
	var response dto.SuccessResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		fmt.Println("Error when unmarshalling response", err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return ctx.JSON(http.StatusOK, response)
}
