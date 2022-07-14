package routers

import (
	"fmt"
	"hello-go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	router := SetupRouter()
	router.Run(port)

	fmt.Println("Server running on port " + port)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.Index)
	router.POST("/create_user", controllers.Create_user)
	router.GET("/get_user", controllers.Get_user)
	router.PATCH("/update_user/:id", controllers.Update_user)
	router.DELETE("/del_user/:id", controllers.Del_user)
	return router
}
