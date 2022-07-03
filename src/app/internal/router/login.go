package router

import (
	"dashboard/src/app/internal/controller"

	"github.com/gin-gonic/gin"
)

func LoginUserRouter() *gin.Engine {
	router := gin.Default()
	LoginUser := router.Group("/v1/user")
	{
		LoginUser.POST("/sign-up", controller.Login) //注册
		LoginUser.GET("/login")                      //登录
	}
	return router
}
