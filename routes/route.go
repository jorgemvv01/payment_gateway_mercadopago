package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	apiRouterGroup := router.Group("/api")
	{
		RegisterOrderRouter(apiRouterGroup)
		RegisterBusinessRouter(apiRouterGroup)
		RegisterProductRouter(apiRouterGroup)
		RegisterMPPaymentRouter(apiRouterGroup)
	}
	return router
}
