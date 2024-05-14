package routes

import (
	"github.com/gin-gonic/gin"
	"payment_gateway_mercadopago/adapter/http/business_http"
)

func RegisterBusinessRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("/business")
	controller := business_http.NewBusinessHTTP()
	orderRouter.GET("", controller.GetAllBusiness)
}
