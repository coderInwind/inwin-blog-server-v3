package service

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

type MenuService struct {
}

func (*MenuService) GetMenuMap() (list []*model.Menu, err error) {
	err = global.DB.Find(&list).Error
	return list, err
}
