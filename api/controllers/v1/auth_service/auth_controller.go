package authservice

import (
	"api-gateway/api/controllers"
	"api-gateway/config"
	"api-gateway/dto/request/v1/auth"
	"api-gateway/dto/response"
	"api-gateway/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
}

func InitAuthController() *authController {
	return &authController{}
}

func (c *authController) Login(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/login", config.AuthSvcHost)

	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	resp, statusCode, err := request.Post(url, reqBody)
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

	// Return the auth data as JSON
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *authController) Logout(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/logout", config.AuthSvcHost)

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.PostWithAuthorization(url, []byte{}, authorization)
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

	// Return the auth data as JSON
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}

func (c *authController) Authenticate(ctx echo.Context) error {
	config := config.GetConfig()
	url := fmt.Sprintf("%s/v1/authenticate", config.AuthSvcHost)

	dto := auth.AuthenticateDTO{
		SourceUri: ctx.Request().Host + ctx.Request().RequestURI,
	}
	data, err := json.Marshal(dto)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.PostWithAuthorization(url, data, authorization)
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

	// Return the auth data as JSON
	return controllers.WriteSuccess(ctx, statusCode, response.Result)
}
