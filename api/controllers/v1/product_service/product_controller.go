package productservice

import (
	"api-gateway/api/controllers"
	"api-gateway/httpconnector"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productController struct {
	productSvcCon *httpconnector.ProductServiceConnector
}

func InitProductController() *productController {
	return &productController{
		productSvcCon: httpconnector.GetProductServiceConnector(),
	}
}

func (c *productController) GetProduct(ctx echo.Context) error {
	result, err := c.productSvcCon.GetProduct(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *productController) ListProducts(ctx echo.Context) error {
	result, err := c.productSvcCon.ListProducts(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *productController) CreateProduct(ctx echo.Context) error {
	result, err := c.productSvcCon.CreateProduct(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *productController) UpdateProduct(ctx echo.Context) error {
	result, err := c.productSvcCon.UpdateProduct(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *productController) DisableProduct(ctx echo.Context) error {
	result, err := c.productSvcCon.DisableProduct(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *productController) EnableProduct(ctx echo.Context) error {
	result, err := c.productSvcCon.EnableProduct(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *productController) IncreaseBookedQuota(ctx echo.Context) error {
	result, err := c.productSvcCon.IncreaseBookedQuota(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}

func (c *productController) DecreaseBookedQuota(ctx echo.Context) error {
	result, err := c.productSvcCon.DecreaseBookedQuota(ctx)
	if err != nil {
		return controllers.WriteError(ctx, http.StatusInternalServerError, err)
	}
	return controllers.WriteSuccess(ctx, http.StatusOK, result)
}
