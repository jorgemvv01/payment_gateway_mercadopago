package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	apiRouterGroup := router.Group("/api")
	{
		apiRouterGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		RegisterOrderRouter(apiRouterGroup)
		RegisterBusinessRouter(apiRouterGroup)
		RegisterProductRouter(apiRouterGroup)
		RegisterMPPaymentRouter(apiRouterGroup)
	}
	return router
}
