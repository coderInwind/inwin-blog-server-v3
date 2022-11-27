package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/common/serializer"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/interner/service"
	"inwind-blog-server-v3/utils"
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

	token, err := service.ServiceGroupApp.UserService.Login(params)

	if err != nil {
		res.FailWithMsg(errcode.UsernameOrPasswordError.WithDetail(err.Error()))
		return
	}

	res.OkWithData(token)
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

func (UserApi) GetInfo(c *gin.Context) {
	var req request.GetInfoRequest
	res := response.NewResponse(c)
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	jwt := utils.NewJWT()
	id := jwt.ParseToken(req.Token).Id

	info, err := service.ServiceGroupApp.GetUserById(id)
	if err != nil {
		res.FailWithMsg(errcode.InvalidParams.WithDetail(err.Error()))
	}
	// 整理
	var resInfo = make(map[string]interface{})
	resInfo["power"] = info.Role.Power
	resInfo["username"] = info.Username
	resInfo["role"] = info.Role.Name

	res.OkWithData(resInfo)
}
