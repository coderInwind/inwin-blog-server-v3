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
	var params request.PageInfo
	res := response.NewResponse(c)
	//校验入参
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	list := service.ServiceGroupApp.GetBlogList(params)
	res.OkWithList(serializer.BuildBlogs(list))
}
