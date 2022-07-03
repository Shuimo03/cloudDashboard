package dao

import (
	"dashboard/src/app/internal/dto"
	"dashboard/src/app/internal/model"

	"k8s.io/apimachinery/pkg/util/uuid"
)

type User model.User

type UserInterface interface {
	CreateUser(user *dto.User) error
	//UserByID(id uint) ([]model.User, error)
}

func (userModel *User) CreateUser(user *dto.User) (*model.User, error) {
	t := &userModel{
		Id:   uuid.New().String(),
		Name: user.Name,
	}
}
