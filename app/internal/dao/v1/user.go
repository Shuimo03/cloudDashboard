package v1

import (
	mysql "dashboard/app/internal/db"
	model "dashboard/app/internal/model/v1"
)

var db = mysql.DBcommon()

func CreateUser(user *model.User) (*model.User, error) {
	return user.Create(db)
}

func GetUserByUserName(user *model.User) (*model.User, error) {
	res := db.Where("name ? =", user.Name)
	return user.Where(res)
}
