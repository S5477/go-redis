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
	publicRoutes.POST("/add", controller.Add)
	publicRoutes.GET("/get", controller.Get)
	publicRoutes.POST("/hadd", controller.HAdd)
	publicRoutes.GET("/hget", controller.HGet)

	router.Run(PORT)
	fmt.Println("Server running on port 8000")
}
