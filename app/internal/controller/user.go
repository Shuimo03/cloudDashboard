package controller

import (
	"dashboard/app/internal/dto"
	"dashboard/app/internal/services/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	user := &dto.User{
		Name:     "userTest",
		Password: "userTestPassword",
		Email:    "user@email.com",
	}
	res, _ := services.UserCrud.SignUp(user)
	context.JSON(http.StatusOK, gin.H{
		"User": res,
	})
}
