package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/service"
)

func CreateUser(c *gin.Context) {

	var createUserService service.CreateUserService
	if err := c.ShouldBind(&createUserService); err != nil {
		c.JSON(200, ResponseError(err))
	} else {
		result := createUserService.CreateUser()
		c.JSON(200, result)
	}
}

func GetUserList(c *gin.Context) {
	// 获取参数 校验参数
	var userListService service.UserListService

	if err := c.ShouldBind(&userListService); err != nil {
		c.JSON(200, ResponseError(err))
	} else {
		result := userListService.GetUserList()
		c.JSON(200, result)
	}

}
