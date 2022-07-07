package controller

import (
	"dashboard/app/internal/dto"
	"dashboard/app/internal/services/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO bug修复,第一次可以正常展示,但是第二次就会出现panic
func GetPods(context *gin.Context) {

	parm := new(dto.Pod)
	podList := services.GetPodsNamesList(parm.Name)
	if err := context.ShouldBind(parm.Name); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"PodName": podList,
		})
	}
}

func GetNamespace(context *gin.Context) {
	nameList := services.GetNamespace()
	context.JSON(http.StatusOK, gin.H{
		"NameSpace:": nameList,
	})
}
