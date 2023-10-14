package middleware

import (
	"api-gateway/api/controllers"
	"api-gateway/config"
	"api-gateway/dto/request/v1/auth"
	"api-gateway/request"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		dto := auth.AuthenticateDTO{
			SourceUri: c.Request().Host + c.Request().RequestURI,
		}

		data, err := json.Marshal(dto)
		if err != nil {
			return controllers.WriteError(c, http.StatusInternalServerError, err)
		}

		authHeader := c.Request().Header.Get("Authorization")

		// pass the Authorization header to auth service API
		authUrl := fmt.Sprintf("%s/v1/authenticate", config.GetConfig().AuthSvcHost)
		_, statusCode, err := request.PostWithAuthorization(authUrl, data, authHeader)
		if err != nil {
			return controllers.WriteError(c, http.StatusInternalServerError, err)
		}
		if statusCode != http.StatusOK {
			return controllers.WriteErrorMsg(c, statusCode, "Authentication failed")
		}

		return next(c)
	}
}
