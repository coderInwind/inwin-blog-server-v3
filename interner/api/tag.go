package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/interner/service"
)

type TagApi struct {
}

func (TagApi) GetTagList(c *gin.Context) {
	res := response.NewResponse(c)
	params := request.AllowEmptyPageRequest{}
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	list, total, err := service.ServiceGroupApp.TagService.GetTagList(params)

	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.OkWithList(list, total)
}

func (TagApi) EditTag(c *gin.Context) {
	res := response.NewResponse(c)
	params := request.EditTag{}
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	if err := service.ServiceGroupApp.TagService.EditTag(params); err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}
	res.Ok()
}
