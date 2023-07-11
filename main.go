package main

import (
	"fmt"

	"github.com/S5477/go-redis/controller"

	"github.com/gin-gonic/gin"
)

const PORT = ":8000"

func main() {
	serveApplication()
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/")
	publicRoutes.GET("/add", controller.Add)
	publicRoutes.GET("/get", controller.Get)

	router.Run(PORT)
	fmt.Println("Server running on port" + PORT)
}
