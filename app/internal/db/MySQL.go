package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBcommon() *gorm.DB {
	dsn := "root:WUliang1998102@tcp(localhost)/clouddashboard?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
