package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	DokcerInitRouter()
	PodInit()
	return PodInit()
	//LoginUserRouter()
}
