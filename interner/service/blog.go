package service

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

type BlogService struct{}

func (b *BlogService) GetBlogList(info request.PageInfo) []model.Blog {
	var blogs []model.Blog
	var size, index = info.PageSize, info.PageIndex

	global.DB.Preload("Tag").Limit(size).Offset((index - 1) * size).Find(&blogs)
	return blogs
}
