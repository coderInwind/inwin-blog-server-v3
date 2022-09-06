package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitTagRouter(router *gin.RouterGroup) {
	tagRouter := router.Group("/tag")
	var tagApi api.TagApi
	{
		tagRouter.GET("/list", tagApi.GetTagList)
	}
}
