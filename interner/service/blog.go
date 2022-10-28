package service

import (
	"gorm.io/gorm/clause"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils/dto"
	"time"
)

type BlogService struct{}

// GetBlogList 所有的博客，显示作者名称
func (b *BlogService) GetBlogList(params request.PageRequest) (blogs []model.Blog, total int64, err error) {
	var pageSize, pageIndex = params.PageSize, params.PageIndex

	// limit -1 offset -1 表示查询总的条数
	err = global.DB.Order("id desc").Scopes(dto.Paginate(pageIndex, pageSize)).Preload("Tags").Preload("Author").Where("hidden = ?", 0).Find(&blogs).Limit(-1).Offset(-1).Count(&total).Error

	return blogs, total, err
}

// GetOwnBlogList 自己的博客
type Blog struct {
	Id        uint        `form:"id" gorm:"PrimaryKey" json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Title     string      `form:"title" json:"title"`
	Content   string      `form:"content" json:"content"`
	Src       string      `form:"src" json:"src"`
	Tags      []model.Tag `gorm:"many2many:blog_tags" json:"tags"`
	Overview  string      `form:"overview" json:"overview"`
	Pv        int         `form:"pv" json:"pv"`
	Like      int         `form:"like" json:"like"`
	Hidden    int         `json:"hidden"`
}

func (b *BlogService) GetOwnBlogList(params request.PageRequest) (ownBlogs []Blog, total int64, err error) {
	var pageSize, pageIndex = params.PageSize, params.PageIndex
	// 查询部分字段

	err = global.DB.Model(&model.Blog{}).Order("id desc").Scopes(dto.Paginate(pageIndex, pageSize)).Preload("Tags").Where("hidden = ?", 0).Find(&ownBlogs).Limit(-1).Offset(-1).Count(&total).Error

	return ownBlogs, total, err
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

func (b *BlogService) CreateBlog(blog model.Blog) (uint, error) {

	res := global.DB.Create(&blog)
	// 返回id
	return blog.Id, res.Error
}

func (b *BlogService) DeleteBlog(id uint) error {
	err := global.DB.Select(clause.Associations).Delete(&model.Blog{BasicModel: model.BasicModel{Id: id}}).Error
	return err
}
