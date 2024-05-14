package mp_payment_http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_gateway_mercadopago/domain/models/mp_payment_model"
	"payment_gateway_mercadopago/domain/models/response_model"
	"payment_gateway_mercadopago/domain/services/mp_payment_service"
	"payment_gateway_mercadopago/infrastructure/repository/mp_payment_repository"
	"payment_gateway_mercadopago/storage"
	"strconv"
)

type MPPaymentHTTP struct{}

func NewMPPaymentHTTP() *MPPaymentHTTP {
	return &MPPaymentHTTP{}
}

func (controller *MPPaymentHTTP) GetMPPaymentByOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response_model.Response{
			Status:  "error",
			Message: "invalid order ID",
		})
		return
	}

	mpPaymentRepository := mp_payment_repository.NewRepository(storage.DB)
	mpPaymentService := mp_payment_service.NewService(mpPaymentRepository)
	err, mpPayment := mpPaymentService.GetMPPaymentByOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response_model.Response{
			Status:  "error",
			Message: `unable to get mp payment ` + err.Error(),
		})
		return
	}
	if mpPayment.ID == 0 {
		c.JSON(http.StatusNotFound, response_model.Response{
			Status:  "error",
			Message: "no mp payment found",
		})
		return
	}

	c.JSON(http.StatusOK, response_model.Response{
		Status:  "success",
		Message: "mp payment found",
		Data:    mp_payment_model.NewMPPaymentResponse(*mpPayment),
	})
}
