package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	var userApi api.ApiGroup
	{
		userRouter.GET("/list", userApi.GetUserList)

		userRouter.POST("/create", userApi.CreateUser)

		userRouter.GET("/logout", userApi.Logout)

		userRouter.GET("/info", userApi.GetInfo)
	}
}
