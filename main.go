package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Go lang with simple CRUD API",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
