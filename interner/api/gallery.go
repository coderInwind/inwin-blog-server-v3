package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/interner/service"
)

type GalleryApi struct {
}

func (*GalleryApi) Upload(c *gin.Context) {
	params := model.Gallery{}
	res := response.NewResponse(c)
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	if err := service.ServiceGroupApp.Upload(params); err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.Ok()
}

func (*GalleryApi) GetGalleryList(c *gin.Context) {
	res := response.NewResponse(c)
	list, total, err := service.ServiceGroupApp.GalleryService.GetGalleryList()
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}
	res.OkWithList(list, total)
}

func (*GalleryApi) DeletePicture(c *gin.Context) {
	params := request.DeletePicture{}
	res := response.NewResponse(c)

	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	if err := service.ServiceGroupApp.DeletePicture(params); err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.Ok()
}
