package routers

import "github.com/gin-gonic/gin"

func Get_user(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get User API Mockup",
	})
}
