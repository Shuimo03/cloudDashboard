package v1

import (
	"gorm.io/gorm"
	"time"
)

type UserOperate interface {
	Create(db *gorm.DB) (*User, error)
	Where(db *gorm.DB) (*User, error)
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
	err := db.Create(userModel).Error
	return userModel, err
}

func (userModel *User) Where(db *gorm.DB) (*User, error) {
	err := db.Where(userModel).Error
	return userModel, err
}
