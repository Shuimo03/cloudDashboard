package v1

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserOperate interface {
	Create(db *gorm.DB) (*User, error)
}

type User struct {
	Id        string    `gorm:"type:char(45);primaryKey"`
	Password  string    `json:"password"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (userModel *User) Create(db *gorm.DB) (*User, error) {
	userModel.Id = uuid.New().String()
	userModel.CreatedAt = time.Now()
	userModel.UpdatedAt = time.Now()
	hash, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	userModel.Password = string(hash)
	dberr := db.Create(&userModel).Error
	if err != nil {
		log.Fatalln(dberr)
	}
	return userModel, dberr
}
