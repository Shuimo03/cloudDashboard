package controller

import (
	"dashboard/app/internal/dto"
	"dashboard/app/internal/services/v1"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateUser(context *gin.Context) {

	user := &dto.User{
		Name:     context.Query("name"),
		Password: context.Query("password"),
		Email:    context.Query("email"),
	}
	res, err := services.Register(user)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
	}

	context.JSON(http.StatusOK, gin.H{
		"User": res,
	})
}

func GetByUserName(context *gin.Context) {

	user := &dto.User{
		Name: context.Query("name"),
	}
	res, err := services.IsValidUsername(user)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
	}

	context.JSON(http.StatusOK, gin.H{
		"User": res,
	})
}
