package service

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils/dto"
)

type TagService struct {
}

func (TagService) GetTagList(params request.PageRequest) (list []model.Tag, total int64, err error) {
	pageSize, pageIndex := params.PageSize, params.PageIndex

	err = global.DB.Scopes(dto.Paginate(pageIndex, pageSize)).Find(&list).Count(&total).Error
	return list, total, err
}
