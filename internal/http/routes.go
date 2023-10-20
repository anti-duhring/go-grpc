package http

import (
	"github.com/anti-duhring/go-grpc/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRoutes(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user/:id/balance", GetUserBalance)
		v1.POST("/user", CreateUser)
		v1.POST("/user/:id/transfer", TransferMoney)
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
