package v1

import (
	"dashboard/app/internal/db"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	Id        string    `gorm:"type:char(45);primaryKey"`
	Password  string    `json:"password"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserOperate interface {
	db.DBService
	CreateUser(user *UserModel, config gorm.Config) error
}
type userService struct {
	db.MySQLDBService
}

func (service *userService) CreateUser(user *UserModel, config gorm.Config) error {
	res := service.DBcommon(config)
	user.Id = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = string(hash)
	res.Create(user)
	return nil
}
