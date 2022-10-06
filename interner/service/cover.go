package service

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

type CoverService struct {
}

func (*CoverService) GetCoverService() (list []model.Cover, err error) {
	err = global.DB.Find(&list).Error

	return list, err
}
