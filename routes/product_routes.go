package routes

import (
	"github.com/gin-gonic/gin"
	"payment_gateway_mercadopago/adapter/http/product_http"
)

func RegisterProductRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("/product")
	controller := product_http.NewProductHTTP()
	orderRouter.GET("/promotional/:id", controller.GetPromotionalProductsByBusiness)
	orderRouter.GET("/by-business/:id", controller.GetAllProductsByBusiness)
}
