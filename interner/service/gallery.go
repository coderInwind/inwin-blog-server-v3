package service

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

type GalleryService struct{}

func (g *GalleryService) Upload(params model.Gallery) error {

	err := global.DB.Create(&params).Error

	return err
}

func (g *GalleryService) GetGalleryList() ([]model.Gallery, int64, error) {
	var list []model.Gallery
	var total int64
	err := global.DB.Find(&list).Count(&total).Error

	return list, total, err
}

func (g *GalleryService) DeletePicture(params request.DeletePicture) error {

	err := global.DB.Where("id = ?", params.Id).Delete(&model.Gallery{}).Error

	return err
}
