package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitRoleRouter(router *gin.RouterGroup) {
	roleRouter := router.Group("/role")
	roleApi := api.ApiGroupApp.RoleApi
	{
		roleRouter.GET("/list", roleApi.GetRoleList)
	}
}
