package service

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

type BlogService struct{}

func (b *BlogService) GetBlogList(params request.PageRequest) ([]model.Blog, int64, error) {
	var blogs []model.Blog
	var total int64
	var size, index = params.PageSize, params.PageIndex
	err := global.DB.Preload("Tag").Limit(size).Offset((index - 1) * size).Find(&blogs).Count(&total).Error
	return blogs, total, err
}

func (b *BlogService) GetBlogDetail(params request.DetailRequest) (model.Blog, error) {
	var blog model.Blog
	var id = params.Id
	err := global.DB.Where("id = ?", id).First(&blog).Error
	return blog, err
}
