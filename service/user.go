package service

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/dao"
)

func CreateUser(c *gin.Context, params common.UserParams) {
	dao.CreateUser(c, params)
}

func GetUserList(c *gin.Context) {
	dao.GetUserList(c)
}
