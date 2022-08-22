package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/service"
)

//服务入口，负责处理路由，参数校验，请求转发

func Login(c *gin.Context) {
	var loginParams service.LoginParams
	//校验参数
	if err := c.ShouldBind(&loginParams); err != nil {
		c.JSON(200, ResponseError(err))
	} else {
		service.LoginService(c)
	}

}
