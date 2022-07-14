package controllers

import (
	"hello-go-crud/models"

	"github.com/gin-gonic/gin"
)

var user = models.User{}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Go CRUD",
	})
}
