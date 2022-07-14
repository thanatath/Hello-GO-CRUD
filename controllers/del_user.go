package controllers

import (
	"hello-go-crud/configs"

	"github.com/gin-gonic/gin"
)

func Del_user(c *gin.Context) {

	configs.DB.First(&user, c.Param("id"))
	configs.DB.Delete(&user)
	c.JSON(200, gin.H{
		"User_deleted": &user,
	})

}
