package dao

import (
	"inwind-blog-server-v3/db"
	"inwind-blog-server-v3/model"
)

func SelectBlogList(s int, i int) ([]model.Blog, int64, error) {
	var blogs []model.Blog

	result := db.DbEngine.Limit(s).Offset((i - 1) * s).Find(&blogs)

	return blogs, result.RowsAffected, result.Error
}

func SelectBlogDetail(id int64) (model.Blog, error) {
	var blog model.Blog

	resule := db.DbEngine.Where("name = ?", id).First(&blog)

	return blog, resule.Error
}
