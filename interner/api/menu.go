package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/common/serializer"
	"inwind-blog-server-v3/interner/service"
)

type MenuApi struct {
}

func (*MenuApi) GetMenuMap(c *gin.Context) {
	res := response.NewResponse(c)

	list, err := service.ServiceGroupApp.MenuService.GetMenuMap()

	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	routeTrees := serializer.MergeMenu(list, 0)

	res.OkWithData(routeTrees)
}
