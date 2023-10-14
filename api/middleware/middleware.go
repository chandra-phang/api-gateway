package middleware

import (
	"api-gateway/dto/request/v1/auth"
	"api-gateway/request"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := auth.AuthenticateDTO{
			SourceUri: c.Request().Host + c.Request().RequestURI,
		}

		errResp := map[string]interface{}{"message": "Authentication failed"}
		data, err := json.Marshal(dto)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errResp)
		}

		authHeader := c.Request().Header.Get("Authorization")

		// pass the Authorization header to auth service API
		_, statusCode, err := request.PostWithAuthorization("http://127.0.0.1:8081/v1/authenticate", data, authHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errResp)
		}
		if statusCode != http.StatusOK {
			return c.JSON(statusCode, errResp)
		}

		return next(c)
	}
}
