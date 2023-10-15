package authservice

import (
	"api-gateway/api/controllers"
	"api-gateway/httpconnector"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
	authServiceCon *httpconnector.AuthServiceConnector
}

func InitAuthController() *authController {
	return &authController{
		authServiceCon: httpconnector.GetAuthServiceConnector(),
	}
}

func (c *authController) Login(ctx echo.Context) error {
	result, err := c.authServiceCon.Login(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *authController) Logout(ctx echo.Context) error {
	result, err := c.authServiceCon.Logout(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *authController) Authenticate(ctx echo.Context) error {
	result, err := c.authServiceCon.Authenticate(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}
