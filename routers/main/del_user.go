package routers

import "github.com/gin-gonic/gin"

func Del_user(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Del User API Mockup",
	})
}
