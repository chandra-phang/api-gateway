package httpconnector

import (
	"api-gateway/config"
	"api-gateway/dto/request/v1/auth"
	"api-gateway/dto/response"
	"api-gateway/request"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

var authServiceCon *AuthServiceConnector

type AuthServiceConnector struct {
	Host            string
	LoginUri        string
	LogoutUri       string
	AuthenticateUri string
}

func InitAuthServiceConnector(cfg config.Config) {
	authServiceCon = &AuthServiceConnector{
		Host:            cfg.AuthSvcHost,
		LoginUri:        cfg.LoginUri,
		LogoutUri:       cfg.LogoutUri,
		AuthenticateUri: cfg.AuthenticateUri,
	}
}

func GetAuthServiceConnector() *AuthServiceConnector {
	return authServiceCon
}

func (con AuthServiceConnector) Login(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.LoginUri
	reqBody, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return nil, err
	}

	resp, statusCode, err := request.Post(url, reqBody, "")
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

func (con AuthServiceConnector) Logout(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.LogoutUri

	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.Post(url, []byte{}, authorization)
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

func (con AuthServiceConnector) Authenticate(ctx echo.Context) (interface{}, error) {
	url := con.Host + con.AuthenticateUri

	dto := auth.AuthenticateDTO{
		SourceUri: ctx.Request().Host + ctx.Request().RequestURI,
	}

	data, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	authorization := ctx.Request().Header.Get("Authorization")
	resp, statusCode, err := request.Post(url, data, authorization)
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
