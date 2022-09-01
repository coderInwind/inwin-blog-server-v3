package service

import (
	"fmt"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	serializer2 "inwind-blog-server-v3/interner/serializer"
	"inwind-blog-server-v3/utils"
)

// 逻辑（服务）层，一般是业务逻辑的入口，可以认为从这里开始，所有的请求参数一定是合法的。业务逻辑和业务流程也都在这一层中。

type LoginService struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (l *LoginService) Login() serializer2.Response {
	var user model.User
	// 验证用户名密码
	if err := global.DbEngine.Where("username = ?", l.Username).First(&user).Error; err != nil {
		fmt.Println("错误")
	}

	if !user.CheckPassword(l.Password) {
		fmt.Println("密码错误")
		return serializer2.Response{
			Code: utils.ErrorAuthCheckPasswordFail,
			Msg:  utils.GetMsg(utils.ErrorAuthCheckPasswordFail),
		}
	}

	//颁发token
	token, err := utils.GenerateToken(user.ID, l.Username)
	if err != nil {
		fmt.Println("错误")
	}

	return serializer2.Response{
		Code: 200,
		Data: serializer2.TokenData{
			Data:  serializer2.BuildLoginUser(user),
			Token: token,
		},
		Msg: "成功",
	}

	//dao.Login(c)
}
