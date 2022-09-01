package dao

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

func SelectBlogList(s int, i int) []model.Blog {

	var blogs []model.Blog

	//result := global.DbEngine.Preload("Tag").Limit(s).Offset((i - 1) * s).Find(&blogs)

	return blogs
}

func SelectBlogDetail(id uint) (model.Blog, error) {
	var blog model.Blog
	result := global.DbEngine.Where("id = ?", id).First(&blog)
	return blog, result.Error
}

func UpdateBlog[T any](e T, id uint) {
	//var blog model.Blog
	//global.DbEngine.Model(&blog).Where("id=?", id).Update()
}
