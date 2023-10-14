package http

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/user/:id/balance", GetUserBalance)
		v1.POST("/user", CreateUser)
	}
}
