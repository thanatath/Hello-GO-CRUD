package controllers

import (
	"hello-go-crud/configs"

	"github.com/gin-gonic/gin"
)

func Get_user(c *gin.Context) {
	rs := configs.DB.Find(&user)
	c.JSON(200, gin.H{
		"Count":    rs.RowsAffected,
		"All_User": user,
	})
}
