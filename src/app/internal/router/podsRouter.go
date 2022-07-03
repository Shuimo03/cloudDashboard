package router

import (
	"dashboard/src/app/internal/controller"

	"github.com/gin-gonic/gin"
)

func PodInit() *gin.Engine {

	router := gin.Default()

	PodrouterGroup := router.Group("/v1")
	{
		PodrouterGroup.GET("/hello", controller.TestHello)
		PodrouterGroup.GET("/podNames", controller.GetPods)
		PodrouterGroup.GET("/nameSpaces", controller.GetNamespace)
	}
	return router
}
