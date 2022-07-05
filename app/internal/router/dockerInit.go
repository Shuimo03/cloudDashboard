package router

import (
	controller2 "dashboard/app/internal/controller"
	"github.com/gin-gonic/gin"
)

func DokcerInitRouter() *gin.Engine {

	router := gin.Default()

	PodrouterGroup := router.Group("/v1")
	{
		PodrouterGroup.GET("/images", controller2.GetImagesID)
		PodrouterGroup.GET("/containernamelist", controller2.GetContainerName)
		PodrouterGroup.POST("/images/:name")
		PodrouterGroup.DELETE("/images/:id", controller2.DeleteImages)
	}
	return router
}
