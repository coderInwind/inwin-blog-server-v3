package service

import (
	"inwind-blog-server-v3/dao"
	"inwind-blog-server-v3/serializer"
	"inwind-blog-server-v3/utils"
)

//博客列表业务

type BlogListSerivce struct {
	PageSize  int `form:"pageSize"  json:"pageSize" binding:"required"`
	PageIndex int `form:"pageIndex" json:"pageIndex" binding:"required"`
}

func (b *BlogListSerivce) GetBlogList() serializer.Response {

	blogs, total, err := dao.SelectBlogList(b.PageSize, b.PageIndex)

	if err != nil {
		return serializer.BuildListErrorResponse(utils.ERROR, utils.GetMsg(utils.ERROR))
	}

	return serializer.BuildListResponse(blogs, total)
}

type BlogDetail struct {
	Id int64 `form:"id" json:"id" binding:"required"`
}

func (b *BlogDetail) BlogDetailService() serializer.Response {
	blog, err := dao.SelectBlogDetail(b.Id)
	if err != nil {
		serializer.BuildErrorResponse(utils.ERROR, utils.GetMsg(utils.ERROR))
	}
	return serializer.Response{
		Code: utils.SUCCESS,
		Data: blog,
		Msg:  utils.GetMsg(utils.SUCCESS),
	}
}
