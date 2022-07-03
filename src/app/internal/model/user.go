package model

import (
	"time"
)

type User struct {
	Id        string    `json:"Id"`
	Password  string    `json:"Password"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	Salt      string    `json:"Salt"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
