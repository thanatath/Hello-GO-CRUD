package controllers

import (
	"hello-go-crud/configs"
	"hello-go-crud/models"

	"github.com/gin-gonic/gin"
)

func Update_user(c *gin.Context) {
	var user = models.User{NAME: c.PostForm("name"), LASTNAME: c.PostForm("lastname"), AGE: c.PostForm("age")}

	user.SetId(c.Param("id"))
	configs.DB.Save(&user)
	c.JSON(200, gin.H{
		"User_updated": &user,
	})

}
