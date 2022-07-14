package main

import (
	"hello-go-crud/configs"
	routers "hello-go-crud/routers/main"
)

func main() {
	configs.InitDB()
	router := routers.SetupRouter()
	router.Run(":9999")

}
