package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	//	DokcerInitRouter()
	//	PodInit()
	LoginUserRouter()
	return LoginUserRouter()
}
