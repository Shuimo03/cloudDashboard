package v1

import (
	mysql "dashboard/app/internal/db"
	model "dashboard/app/internal/model/v1"
)

type UserDao interface {
	CreateUser(user *model.User) (*model.User, error)
}

func CreateUser(user *model.User) (*model.User, error) {
	db := mysql.DBcommon()
	return user.Create(db)
}
