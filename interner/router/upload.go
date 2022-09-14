package router

import "github.com/gin-gonic/gin"

func InitUploadRouter(router *gin.RouterGroup) {
	uploadRouter := router.Group("/upload")

	uploadRouter.GET("/token")
}
