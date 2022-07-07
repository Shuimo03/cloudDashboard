package services

import (
	dao "dashboard/app/internal/dao/v1"
	"dashboard/app/internal/dto"
	model "dashboard/app/internal/model/v1"
	"fmt"
)

type UserCrud struct {
	dao dao.UserDao
}

func (u *UserCrud) SignUp(user *dto.User) (*model.User, error) {
	res := &model.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	createUser, err := u.dao.CreateUser(res)
	if err != nil {
		fmt.Println(err)
	}
	return createUser, nil

}
