package service

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils/dto"
)

type UserService struct{}

func (UserService) GetUserList(params request.PageRequest) (users []model.User, total int64, err error) {

	pageIndex, pageSize := params.PageIndex, params.PageSize

	err = global.DB.Scopes(dto.Paginate(pageIndex, pageSize)).Find(&users).Count(&total).Error

	return users, total, err
}
