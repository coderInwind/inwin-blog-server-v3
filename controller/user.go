package controller

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/service"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var userParams common.UserParams
	//校验参数
	if err := c.ShouldBind(&userParams); err != nil {
		common.ResponseError(c, http.StatusUnprocessableEntity, 422, err.Error())
		return
	}

	service.CreateUser(c, userParams)
}

func GetUserList(c *gin.Context) {
	service.GetUserList(c)
}
