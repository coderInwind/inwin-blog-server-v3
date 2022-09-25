package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitCoverRouter(router *gin.RouterGroup) {
	coverRouter := router.Group("/cover")
	coverApi := api.ApiGroupApp
	{
		coverRouter.GET("/list", coverApi.GetCoverList)
	}
}
