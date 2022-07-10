package services

import (
	dao "dashboard/app/internal/dao/v1"
	"dashboard/app/internal/dto"
	model "dashboard/app/internal/model/v1"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(user *dto.User) (*model.User, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatalf("PassWord: %v", err)
	}

	res := &model.User{
		Id:        uuid.New().String(),
		Name:      user.Name,
		Password:  string(password),
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := dao.CreateUser(res)
	if err != nil {
		logrus.Fatalf("Create User %v", err)
	}
	return createUser, nil

}
