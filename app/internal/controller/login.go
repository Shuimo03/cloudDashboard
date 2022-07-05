package controller

import (
	"dashboard/app/internal/dto"
	"dashboard/app/internal/services/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var user dto.User
	user.Name = "test"
	user.Email = "test@test.com"
	var createUser services.services
	createUser.Create(&user)

	context.JSON(http.StatusOK, gin.H{
		"name":     user.Name,
		"password": user.Password,
		"email:":   user.Email,
	})
}
