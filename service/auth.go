package service

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/dao"
)

// 逻辑（服务）层，一般是业务逻辑的入口，可以认为从这里开始，所有的请求参数一定是合法的。业务逻辑和业务流程也都在这一层中。

type LoginParams struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginService(c *gin.Context) {
	dao.Login(c)
}
