package product_http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_gateway_mercadopago/domain/models/product_model"
	"payment_gateway_mercadopago/domain/models/response_model"
	"payment_gateway_mercadopago/domain/services/product_service"
	"payment_gateway_mercadopago/infrastructure/repository/product_repository"
	"payment_gateway_mercadopago/storage"
	"strconv"
)

type ProductHTTP struct{}

func NewProductHTTP() *ProductHTTP {
	return &ProductHTTP{}
}

// GetAllProductsByBusiness
// @Summary Get all products by business
// @Description Get all products by business.
// @Param business_id path string true "The business ID is required in the query"
// @Produce application/json
// @Tags Products
// @Success 200 {object} response_model.Response{}
// @Failure 400 {object} response_model.Response{}
// @Failure 404 {object} response_model.Response{}
// @Failure 500 {object} response_model.Response{}
// @Router /product/by-business/{business_id} [get]
func (controller *ProductHTTP) GetAllProductsByBusiness(c *gin.Context) {
	businessID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "the business id sent is invalid",
		})
		return
	}
	productRepository := product_repository.NewRepository(storage.DB)
	productService := product_service.NewService(productRepository)
	err, products := productService.GetAllProductsByBusiness(uint(businessID))
	if len(*products) == 0 {
		c.JSON(http.StatusNotFound, response_model.Response{
			Status:  "error",
			Message: "no products found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, response_model.Response{
			Status:  "error",
			Message: `unable to get products` + err.Error(),
		})
	}
	var productsResponse []product_model.ProductResponse
	for _, productItem := range *products {
		productsResponse = append(productsResponse, product_model.NewProductResponse(productItem))
	}
	c.JSON(http.StatusOK, response_model.Response{
		Status:  "success",
		Message: "products found",
		Data:    productsResponse,
	})
}

// GetPromotionalProductsByBusiness
// @Summary Get promotional products by business
// @Description Get promotional products by business.
// @Param business_id path string true "The business ID is required in the query"
// @Produce application/json
// @Tags Products
// @Success 200 {object} response_model.Response{}
// @Failure 400 {object} response_model.Response{}
// @Failure 404 {object} response_model.Response{}
// @Failure 500 {object} response_model.Response{}
// @Router /product/promotional/{business_id} [get]
func (controller *ProductHTTP) GetPromotionalProductsByBusiness(c *gin.Context) {
	businessID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "the business id sent is invalid",
		})
		return
	}
	productRepository := product_repository.NewRepository(storage.DB)
	productService := product_service.NewService(productRepository)
	err, products := productService.GetPromotionalProductsByBusiness(uint(businessID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response_model.Response{
			Status:  "error",
			Message: `unable to get products` + err.Error(),
		})
		return
	}
	if len(*products) == 0 {
		c.JSON(http.StatusNotFound, response_model.Response{
			Status:  "error",
			Message: "no products found",
		})
		return
	}
	var productsResponse []product_model.ProductResponse
	for _, productItem := range *products {
		productsResponse = append(productsResponse, product_model.NewProductResponse(productItem))
	}
	c.JSON(http.StatusOK, response_model.Response{
		Status:  "success",
		Message: "products found",
		Data:    productsResponse,
	})
}
