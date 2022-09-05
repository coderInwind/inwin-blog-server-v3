package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/common/serializer"
	"inwind-blog-server-v3/interner/service"
)

type BlogApi struct{}

func (b *BlogApi) GetBlogList(c *gin.Context) {
	var params request.PageRequest
	res := response.NewResponse(c)
	//校验入参
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	list, total, err := service.ServiceGroupApp.GetBlogList(params)
	//查询
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}
	res.OkWithList(serializer.BuildBlogs(list), total)
}

func (b *BlogApi) GetBlogDetail(c *gin.Context) {
	var params request.DetailRequest
	res := response.NewResponse(c)

	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	detail, err := service.ServiceGroupApp.GetBlogDetail(params)
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.OkWithData(serializer.BuildBlog(detail))
}
