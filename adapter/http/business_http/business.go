package business_http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_gateway_mercadopago/domain/models/business_model"
	"payment_gateway_mercadopago/domain/models/response_model"
	"payment_gateway_mercadopago/domain/services/business_service"
	"payment_gateway_mercadopago/infrastructure/repository/business_repository"
	"payment_gateway_mercadopago/storage"
)

type BusinessHTTP struct{}

func NewBusinessHTTP() *BusinessHTTP {
	return &BusinessHTTP{}
}

func (controller *BusinessHTTP) GetAllBusiness(c *gin.Context) {

	businessRepository := business_repository.NewRepository(storage.DB)
	businessService := business_service.NewService(businessRepository)
	err, business := businessService.GetAllBusiness()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response_model.Response{
			Status:  "error",
			Message: `unable to get business` + err.Error(),
		})
	}
	if len(*business) == 0 {
		c.JSON(http.StatusNotFound, response_model.Response{
			Status:  "error",
			Message: "no business found",
		})
		return
	}
	var businessResponse []business_model.BusinessResponse
	for _, businessItem := range *business {
		businessResponse = append(businessResponse, business_model.NewBusinessResponse(businessItem))
	}
	c.JSON(http.StatusOK, response_model.Response{
		Status:  "success",
		Message: "business found",
		Data:    businessResponse,
	})
}
