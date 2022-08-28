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
	Id int64 `json:"id"`
}

func (b *BlogDetail) BlogDetailService() {
	dao.SelectBlogDetail(b.Id)
}
