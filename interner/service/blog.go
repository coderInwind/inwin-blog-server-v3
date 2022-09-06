package service

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils/dto"
)

type BlogService struct{}

func (b *BlogService) GetBlogList(params request.PageRequest) (blogs []model.Blog, total int64, err error) {

	var pageSize, pageIndex = params.PageSize, params.PageIndex
	err = global.DB.Preload("Tag").Scopes(dto.Paginate(pageIndex, pageSize)).Find(&blogs).Count(&total).Error
	return blogs, total, err
}

func (b *BlogService) GetBlogDetail(params request.DetailRequest) (blog model.Blog, err error) {
	var id = params.Id
	err = global.DB.Where("id = ?", id).First(&blog).Error
	return blog, err
}
