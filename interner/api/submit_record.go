package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/interner/service"
)

type SubmitRecordApi struct {
}

func (*SubmitRecordApi) GetSubmitRecord(c *gin.Context) {
	params := request.GetRecordParams{}
	res := response.NewResponse(c)
	if err := c.ShouldBindJSON(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}
	records, err := service.ServiceGroupApp.GetSubmitRecord(params.Date)
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
	}

	res.OkWithData(records)
}
