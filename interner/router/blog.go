package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

func InitBlogRouter(router *gin.RouterGroup) {
	blogRouter := router.Group("/blog")
	blogApi := api.BlogApi{}
	{
		blogRouter.GET("/list", blogApi.GetBlogList)
		blogRouter.POST("/detail", blogApi.GetBlogDetail)
		blogRouter.POST("/edit", blogApi.EditBlog)
		blogRouter.POST("/create", blogApi.CreateBlog)
		blogRouter.POST("/delete", blogApi.DeleteBlog)
		blogRouter.POST("/detail", blogApi.GetBlogDetail)
	}
}
