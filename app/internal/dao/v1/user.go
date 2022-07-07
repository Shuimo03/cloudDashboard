package v1

import (
	model "dashboard/app/internal/model/v1"
	"gorm.io/gorm"
)

type UserDao interface {
	CreateUser(user *model.User) (*model.User, error)
}

type userDB struct {
	db *gorm.DB
}

func (u *userDB) CreateUser(user *model.User) (*model.User, error) {
	return user.Create(u.db)
}
