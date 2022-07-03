package router

import (
	"dashboard/src/app/internal/controller"
	"github.com/gin-gonic/gin"
)

func DokcerInitRouter() *gin.Engine {

	router := gin.Default()

	PodrouterGroup := router.Group("/v1")
	{
		PodrouterGroup.GET("/images", controller.GetImagesID)
		PodrouterGroup.GET("/containernamelist", controller.GetContainerName)
		PodrouterGroup.POST("/images/:name")
		PodrouterGroup.DELETE("/images/:id", controller.DeleteImages)
	}
	return router
}
