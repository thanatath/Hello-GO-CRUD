package controllers

import (
	"hello-go-crud/configs"

	"github.com/gin-gonic/gin"
)

func Update_user(c *gin.Context) {

	rs := configs.DB.First(&user, c.Param("id"))
	user.SetName(c.PostForm("name"))
	user.SetLastname(c.PostForm("lastname"))
	user.SetAge(c.PostForm("age"))
	configs.DB.Save(&user)
	if rs.Error != nil {
		c.JSON(400, gin.H{
			"Error": "Error while updating user",
		})
	} else {
		c.JSON(200, gin.H{
			"User_updated": &user,
		})

	}
}
