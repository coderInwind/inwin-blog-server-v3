package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/interner/service"
)

type bannerApi struct {
}

func (*bannerApi) GetCoverList(c *gin.Context) {

	res := response.NewResponse(c)
	list, err := service.ServiceGroupApp.GetCoverService()
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}
	res.OkWithData(list)
}
