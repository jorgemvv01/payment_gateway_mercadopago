package routes

import (
	"github.com/gin-gonic/gin"
	"payment_gateway_mercadopago/adapter/http/mp_payment_http"
)

func RegisterMPPaymentRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("/mp-payment")
	controller := mp_payment_http.NewMPPaymentHTTP()
	orderRouter.GET("/:id", controller.GetMPPaymentByOrder)
}
