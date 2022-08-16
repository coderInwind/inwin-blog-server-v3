package service

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/dao"
)

// 逻辑（服务）层，一般是业务逻辑的入口，可以认为从这里开始，所有的请求参数一定是合法的。业务逻辑和业务流程也都在这一层中。

func Login(c *gin.Context) {
	dao.Login(c)
}
