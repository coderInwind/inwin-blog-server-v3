package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitBaseUser(r *gin.RouterGroup) {
	baseRouter := r.Group("")
	
	var userApi api.UserApi
	{
		baseRouter.POST("/login", userApi.Login)
	}
}
