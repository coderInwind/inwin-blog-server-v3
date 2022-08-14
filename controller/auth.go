package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/service"
	"net/http"
)

//服务入口，负责处理路由，参数校验，请求转发

func Login(c *gin.Context) {
	var loginParams common.LoginParams
	//校验参数
	if err := c.ShouldBind(&loginParams); err != nil {
		common.ResponseError(c, http.StatusUnprocessableEntity, 422, err.Error())
		return
	}

	fmt.Printf("%v", loginParams)
	service.Login(c)
}

func CreateUser(c *gin.Context) {
	var userParams common.UserParams
	//校验参数
	if err := c.ShouldBind(&userParams); err != nil {
		common.ResponseError(c, http.StatusUnprocessableEntity, 422, err.Error())
		return
	}

	service.CreateUser(c, userParams)
}
