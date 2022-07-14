package controllers

import (
	"hello-go-crud/configs"
	"hello-go-crud/models"

	"github.com/gin-gonic/gin"
)

func Get_user(c *gin.Context) {
	var users []models.User
	rs := configs.DB.Find(&users)
	if rs.Error != nil {
		c.JSON(400, gin.H{
			"Error": "Error while getting users",
		})
	} else {
		c.JSON(200, gin.H{
			"Count":    rs.RowsAffected,
			"All_User": users,
		})
	}
}
