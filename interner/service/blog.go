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
	// limit -1 offset -1 表示查询总的条数
	err = global.DB.Order("id desc").Scopes(dto.Paginate(pageIndex, pageSize)).Preload("Tag").Where("hidden = ?", 0).Find(&blogs).Limit(-1).Offset(-1).Count(&total).Error

	return blogs, total, err
}

func (b *BlogService) GetBlogDetail(params request.SelectBlogRequest) (blog model.Blog, err error) {
	var id = params.Id
	err = global.DB.Where("id = ?", id).Preload("Tag").First(&blog).Error
	return blog, err
}

func (b *BlogService) EditBlog(params *model.Blog) error {

	err := global.DB.Model(&params).Where("id=?", &params.ID).Updates(&params).Error

	return err
}

func (b *BlogService) CreateBlog(params request.CreateBlog) error {
	blog := model.Blog{
		BasicModel: model.BasicModel{},
		Title:      params.Title,
		Content:    params.Content,
		Src:        params.Src,
		Tag: model.Tag{
			Id: params.TagId,
		},
		Overview: params.Overview,
		Pv:       0,
		Like:     0,
	}
	err := global.DB.Create(&blog).Error

	return err
}

func (b *BlogService) DeleteBlog(id int) error {
	err := global.DB.Where("id = ?", id).Delete(&model.Blog{}).Error

	return err
}
