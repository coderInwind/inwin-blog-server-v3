package service

import (
	"inwind-blog-server-v3/dao"
	"inwind-blog-server-v3/serializer"
)

func RoleListService() serializer.Response {
	roleList := dao.SelectRoleList()

	return serializer.BuildListResponse(roleList, 1)
}

func RoleOptionsService() serializer.Response {
	roleList := dao.SelectRoleList()
	roleOptions := serializer.BuildRoleOptions(roleList)
	return serializer.BuildListResponse(roleOptions, 1)
}
