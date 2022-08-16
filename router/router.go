package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/controller"
	"inwind-blog-server-v3/middleware"
)

func NewRouter() *gin.Engine {
	// 认证
	r := gin.New()
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
	}
	// 路由
	user := r.Group("/user")
	{
		user.GET("/list", middleware.Encrypt(), controller.GetUserList)
		user.POST("/create", middleware.Encrypt(), controller.CreateUser)
	}

	return r
}
