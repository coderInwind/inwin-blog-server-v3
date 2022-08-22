package dao

import (
	"gorm.io/gorm"
	"inwind-blog-server-v3/db"
	"inwind-blog-server-v3/model"
)

type CreateUserDao struct {
	Username string
	Password string
	RoleId   int
}

func SelectUserList(s int, i int) ([]model.User, int64) {
	// preload嵌套预加载 - 结构体中嵌套有结构体使用
	var users []model.User
	var user model.User
	var total int64

	db.DbEngine.Preload("Role").Limit(s).Offset((i - 1) * s).Find(&users)
	db.DbEngine.Model(user).Count(&total)
	return users, total
}

func CreateUser(user model.User) *gorm.DB {
	result := db.DbEngine.Create(&user)
	return result
}
