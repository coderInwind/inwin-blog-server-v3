package service

import (
	"fmt"
	"gorm.io/gorm/clause"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/serializer"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils/dto"
)

type BlogService struct{}

func (b *BlogService) GetBlogList(params request.PageRequest) (blogs []model.Blog, total int64, err error) {

	var pageSize, pageIndex = params.PageSize, params.PageIndex
	// limit -1 offset -1 表示查询总的条数
	err = global.DB.Order("id desc").Scopes(dto.Paginate(pageIndex, pageSize)).Preload("Tags").Where("hidden = ?", 0).Find(&blogs).Limit(-1).Offset(-1).Count(&total).Error

	return blogs, total, err
}

func (b *BlogService) GetBlogDetail(params request.SelectBlogRequest) (blog model.Blog, err error) {
	var id = params.Id
	err = global.DB.Where("id = ?", id).Preload("Tags").First(&blog).Error
	return blog, err
}

func (b *BlogService) EditBlog(params request.EditBlog) error {
	//查询所有标签

	//// 关联模式 的更新
	var tag model.Tag
	var blog model.Blog
	global.DB.Find(&blog, params.Id)

	//替换关联
	// 先查后关联
	global.DB.Model(&blog).Association("Tags").Find(&tag)
	fmt.Println(tag)

	return nil
}

func (b *BlogService) CreateBlog(params request.CreateBlog) error {
	newParams := serializer.BuildCreateBlogParams(params)
	err := global.DB.Create(&newParams).Error

	return err
}

func (b *BlogService) DeleteBlog(id uint) error {
	err := global.DB.Select(clause.Associations).Delete(&model.Blog{BasicModel: model.BasicModel{ID: id}}).Error
	return err
}
