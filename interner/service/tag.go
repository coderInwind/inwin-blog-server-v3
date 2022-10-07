package service

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils/dto"
)

type TagService struct {
}

func (*TagService) GetTagList(params request.AllowEmptyPageRequest) (list []model.Tag, total int64, err error) {
	pageSize, pageIndex := params.PageSize, params.PageIndex

	err = global.DB.Scopes(dto.Paginate(pageIndex, pageSize)).Find(&list).Limit(-1).Offset(-1).Count(&total).Error
	return list, total, err
}

func (*TagService) EditTag(params request.EditTag) error {

	err := global.DB.Model(&model.Tag{}).Where("id = ?", &params.Id).Updates(&params).Error
	return err
}

func (*TagService) CreateTag(params request.CreateTag) error {
	err := global.DB.Create(&model.Tag{Name: params.Name}).Error

	return err
}

func (*TagService) DeleteTag(params request.DeleteTag) error {
	err := global.DB.Delete(&model.Tag{Id: params.Id}).Error

	return err
}
