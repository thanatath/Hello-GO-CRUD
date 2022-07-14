package controllers

import (
	"hello-go-crud/configs"

	"github.com/gin-gonic/gin"
)

func Del_user(c *gin.Context) {

	rs := configs.DB.First(&user, c.Param("id"))
	configs.DB.Delete(&user)
	if rs.Error != nil {
		c.JSON(400, gin.H{
			"Error": "Error while deleting user",
		})
	} else {
		c.JSON(200, gin.H{
			"User_deleted": &user,
		})

	}
}
