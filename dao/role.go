package dao

import (
	"inwind-blog-server-v3/db"
	"inwind-blog-server-v3/model"
)

func SelectRoleList() []model.Role {

	var role []model.Role
	db.DbEngine.Find(&role)

	return role
}
