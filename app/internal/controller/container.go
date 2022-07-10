package controller

import (
	"dashboard/app/internal/services/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetContainerName(context *gin.Context) {

	containerName := services.GetContainerNameList()
	context.JSON(http.StatusOK, gin.H{
		"containerName": containerName,
	})
}
