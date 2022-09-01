package dao

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

func SelectRoleList() []model.Role {

	var role []model.Role
	global.DbEngine.Find(&role)

	return role
}
