package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitGalleryRouter(router *gin.RouterGroup) {
	galleryRouter := router.Group("/gallery")
	galleryApi := api.GalleryApi{}

	galleryRouter.POST("/upload", galleryApi.Upload)
	galleryRouter.GET("/list", galleryApi.GetGalleryList)
	galleryRouter.POST("/delete", galleryApi.DeletePicture)
}
