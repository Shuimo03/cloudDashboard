package router

import (
	"dashboard/app/internal/controller"
	"github.com/gin-gonic/gin"
)

func LoginUserRouter() *gin.Engine {
	router := gin.Default()
	LoginUser := router.Group("/v1/user")
	{
		LoginUser.POST("/register", controller.CreateUser) //注册
	}
	return router
}
