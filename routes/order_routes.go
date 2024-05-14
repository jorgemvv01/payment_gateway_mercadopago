package routes

import (
	"github.com/gin-gonic/gin"
	"payment_gateway_mercadopago/adapter/http/order_http"
)

func RegisterOrderRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("/order")
	controller := order_http.NewOrderHTTP()
	orderRouter.POST("", controller.CreateOrder)
	orderRouter.GET("/user/:id", controller.GetOrdersByUser)
	orderRouter.GET("/details/:id", controller.GetOrderDetail)
	orderRouter.POST("/feedback", controller.MPFeedback)
}
