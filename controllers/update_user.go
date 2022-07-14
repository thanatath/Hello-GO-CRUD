package controllers

import (
	"hello-go-crud/configs"

	"github.com/gin-gonic/gin"
)

func Update_user(c *gin.Context) {

	configs.DB.First(&user, c.Param("id"))
	user.SetName(c.PostForm("name"))
	user.SetLastname(c.PostForm("lastname"))
	user.SetAge(c.PostForm("age"))
	configs.DB.Save(&user)
	c.JSON(200, gin.H{
		"User_updated": &user,
	})

}
