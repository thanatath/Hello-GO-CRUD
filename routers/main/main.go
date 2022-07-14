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
	//router.GET("/", Index)
	//router.POST("/create_user", Create_user)
	router.GET("/get_user", controllers.Get_user)
	//router.PATCH("/update_user", Update_user)
	//router.DELETE("/del_user", Del_user)
	return router
}
