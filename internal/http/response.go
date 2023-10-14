package http

import "github.com/gin-gonic/gin"

func sendError(c *gin.Context, code int, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, gin.H{
		"error": message,
	})
}

func sendSuccess(c *gin.Context, code int, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, data)
}
