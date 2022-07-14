package controllers

import (
	"hello-go-crud/configs"
	"hello-go-crud/models"

	"github.com/gin-gonic/gin"
)

var user = models.User{}

func Get_user(c *gin.Context) {
	rs := configs.DB.Find(&user)
	c.JSON(200, gin.H{
		"Count":    rs.RowsAffected,
		"All_User": user,
	})
}
