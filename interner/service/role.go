package service

import (
	"inwind-blog-server-v3/interner/dao"
	serializer2 "inwind-blog-server-v3/interner/serializer"
)

func RoleListService() serializer2.Response {
	roleList := dao.SelectRoleList()

	return serializer2.BuildListResponse(roleList, 1)
}

func RoleOptionsService() serializer2.Response {
	roleList := dao.SelectRoleList()
	roleOptions := serializer2.BuildRoleOptions(roleList)
	return serializer2.BuildListResponse(roleOptions, 1)
}
