package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/interner/service"
)

type RoleApi struct {
}

func (*RoleApi) GetRoleList(c *gin.Context) {
	res := response.NewResponse(c)

	list, total, err := service.ServiceGroupApp.RoleService.GetRoleList()
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.OkWithList(list, total)
}
