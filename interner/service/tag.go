package service

import (
	"inwind-blog-server-v3/interner/dao"
	"inwind-blog-server-v3/interner/serializer"
	"inwind-blog-server-v3/utils"
)

func TagListService() serializer.Response {
	tags, total, err := dao.SelectTagList()
	if err != nil {
		return serializer.BuildListErrorResponse(utils.ERROR, utils.GetMsg(utils.ERROR))
	}

	return serializer.BuildListResponse(tags, total)
}
