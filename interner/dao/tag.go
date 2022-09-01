package dao

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
)

func SelectTagList() ([]model.Tag, int64, error) {
	var tags []model.Tag
	result := global.DbEngine.Find(&tags)
	return tags, result.RowsAffected, result.Error
}
