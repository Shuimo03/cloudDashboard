package controller

import (
	"dashboard/app/internal/services/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetImagesID(context *gin.Context) {

	imagesID := services.ImagesID()
	context.JSON(http.StatusOK, gin.H{
		"ImagesID:": imagesID,
	})
}

func DeleteImages(context *gin.Context) {
	imageID := context.Param("id")
	deleteImages := services.DelteImages(imageID)
	context.BindJSON(&deleteImages)
	context.JSON(http.StatusOK, gin.H{
		"Context": "删除成功!",
	})
}
