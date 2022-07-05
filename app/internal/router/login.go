package router

import (
	"dashboard/app/internal/controller"
	"github.com/gin-gonic/gin"
)

func LoginUserRouter() *gin.Engine {
	router := gin.Default()
	LoginUser := router.Group("/v1/user")
	{
		LoginUser.POST("/login", controller.Login) //注册
		//	LoginUser.GET("/login")                    //登录
	}
	return router
}
