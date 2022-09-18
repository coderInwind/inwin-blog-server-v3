package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitUploadRouter(router *gin.RouterGroup) {
	uploadRouter := router.Group("/upload")
	uploadApi := api.UploadApi{}
	{
		uploadRouter.GET("/token", uploadApi.GetUploadUrl)
	}
}
