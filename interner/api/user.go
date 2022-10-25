package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/common/serializer"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/interner/service"
)

type UserApi struct{}

func (UserApi) GetUserList(c *gin.Context) {
	params := request.PageRequest{}
	res := response.NewResponse(c)
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	list, total, err := service.ServiceGroupApp.GetUserList(params)
	if err != nil {
		res.FailWithMsg(errcode.ServerError.WithDetail(err.Error()))
		return
	}

	res.OkWithList(serializer.BuildUsers(list), total)
}

// Login 登录
func (UserApi) Login(c *gin.Context) {
	params := request.LoginRequest{}
	res := response.NewResponse(c)
	// 校验参数
	if err := c.ShouldBind(&params); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	user, token, err := service.ServiceGroupApp.UserService.Login(params)

	if err != nil {
		res.FailWithMsg(errcode.UsernameOrPasswordError.WithDetail(err.Error()))
		return
	}

	res.OkWithData(serializer.BuildLogin(user, token))
}

func (UserApi) Logout(c *gin.Context) {
	res := response.NewResponse(c)
	service.ServiceGroupApp.UserService.Logout()

	res.OkWithMsg()
}

func (UserApi) CreateUser(c *gin.Context) {
	user := model.User{}
	res := response.NewResponse(c)

	// 校验参数
	if err := c.ShouldBind(&user); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	if err := service.ServiceGroupApp.CreateUser(user); err != nil {
		res.FailWithMsg(errcode.UsernameOrPasswordError.WithDetail(err.Error()))
		return
	}

	res.OkWithMsg()
}
