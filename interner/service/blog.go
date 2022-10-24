package service

import (
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

	//查传入数组的相应标签
	var tags []model.Tag
	global.DB.Find(&tags, params.Tags)

	var blog model.Blog
	global.DB.Find(&blog, params.Id).Updates(&model.Blog{
		BasicModel: model.BasicModel{
			Id: params.Id,
		},
		Title:    params.Title,
		Content:  params.Content,
		Src:      params.Src,
		Overview: params.Overview,
		Hidden:   params.Hidden,
	})

	//替换关联
	// 先查后关联
	global.DB.Model(&blog).Association("Tags").Clear()
	global.DB.Model(&blog).Association("Tags").Replace(&tags)

	//

	return nil
}

func (b *BlogService) CreateBlog(params request.CreateBlog) (uint, error) {
	newParams := serializer.BuildCreateBlogParams(params)
	res := global.DB.Create(&newParams)
	// 返回id
	return newParams.Id, res.Error
}

func (b *BlogService) DeleteBlog(id uint) error {
	err := global.DB.Select(clause.Associations).Delete(&model.Blog{BasicModel: model.BasicModel{Id: id}}).Error
	return err
}
