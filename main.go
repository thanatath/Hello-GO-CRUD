package main

import routers "hello-go-crud/routers/main"

func main() {

	router := routers.SetupRouter()
	router.Run(":3000")

}
