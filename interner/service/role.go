package service

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

type RoleService struct {
}

func (*RoleService) GetRoleList() ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64
	err := global.DB.Find(&roles).Count(&total).Error
	return roles, total, err
}
