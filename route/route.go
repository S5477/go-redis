package route

import (
	"github.com/S5477/go-redis/controller"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()

	publicRoutes := router.Group("/")
	publicRoutes.POST("/add", controller.Add)
	publicRoutes.GET("/get", controller.Get)
	publicRoutes.POST("/hadd", controller.HAdd)
	publicRoutes.GET("/hget", controller.HGet)

	return router
}
