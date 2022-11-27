package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitMenuRouter(router *gin.RouterGroup) {
	routeRouter := router.Group("/route")
	MenuApi := api.ApiGroupApp.MenuApi
	{
		routeRouter.GET("get_routes", MenuApi.GetMenuMap)
	}
}
