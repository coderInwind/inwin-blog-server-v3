package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/api"
	"inwind-blog-server-v3/middleware"
)

func NewRouter() *gin.Engine {

	r := gin.New()

	r.Use(middleware.Cors())

	auth := r.Group("/auth")
	{
		auth.POST("/login", api.Login)
	}
	// 路由
	user := r.Group("/user")
	{
		user.GET("/list", api.GetUserList)
		user.POST("/create", api.CreateUser)
	}

	role := r.Group("/role")
	{
		role.GET("/list", api.GetRoleList)
		role.GET("/options", api.GetRoleOptions)
	}

	blog := r.Group("/blog")
	{
		blog.GET("/list", api.GetBlogList)
	}

	return r
}
