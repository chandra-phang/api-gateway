package orderservice

import (
	"api-gateway/api/controllers"
	"api-gateway/config"
	"api-gateway/dto/response"
	"api-gateway/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartController struct {
}

func InitCartController() *cartController {
	return &cartController{}
}

func (c *cartController) AddToCart(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/carts", config.OrderSvcHost)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.Post(url, reqBody, authorization)
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

	// Return the cart data as JSON
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}
