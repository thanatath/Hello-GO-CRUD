package routers

import "github.com/gin-gonic/gin"

func Update_user(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update User API Mockup",
	})
}
