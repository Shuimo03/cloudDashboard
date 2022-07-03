package controller

import (
	"dashboard/src/app/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContainerName(context *gin.Context) {

	containerName := services.GetContainerNameList()
	context.JSON(http.StatusOK, gin.H{
		"containerName": containerName,
	})
}
