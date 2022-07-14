package controllers

import (
	"hello-go-crud/configs"
	"hello-go-crud/models"

	"github.com/gin-gonic/gin"
)

func Create_user(c *gin.Context) {
	var new_user = &models.User{NAME: c.PostForm("name"), LASTNAME: c.PostForm("lastname"), AGE: c.PostForm("age")}
	rs := configs.DB.Create(&new_user)

	if rs.Error != nil {
		c.JSON(400, gin.H{
			"Error": "Error while creating user",
		})
	} else {
		c.JSON(200, gin.H{
			"User_created": &new_user,
		})
	}

}
