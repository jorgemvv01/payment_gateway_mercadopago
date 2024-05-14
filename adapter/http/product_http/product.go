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

func (controller *ProductHTTP) GetAllProductsByBusiness(c *gin.Context) {
	businessID, exists := c.GetQuery("business_id")
	if !exists {
		c.JSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "business id not found in query",
		})
		return
	}
	businessIDUINT, err := strconv.ParseUint(businessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "the business id sent is invalid",
		})
		return
	}
	productRepository := product_repository.NewRepository(storage.DB)
	productService := product_service.NewService(productRepository)
	err, products := productService.GetAllProductsByBusiness(uint(businessIDUINT))
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

func (controller *ProductHTTP) GetPromotionalProductsByBusiness(c *gin.Context) {
	businessID, exists := c.GetQuery("business_id")
	if !exists {
		c.JSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "business id not found in query",
		})
		return
	}
	businessIDUINT, err := strconv.ParseUint(businessID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "the business id sent is invalid",
		})
		return
	}
	productRepository := product_repository.NewRepository(storage.DB)
	productService := product_service.NewService(productRepository)
	err, products := productService.GetPromotionalProductsByBusiness(uint(businessIDUINT))
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
