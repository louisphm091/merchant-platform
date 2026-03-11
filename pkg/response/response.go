package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
		"data":    err,
	})
}
