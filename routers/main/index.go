package routers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Go lang with Gin for CRUD API",
	})
}
