package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/interner/serializer"
	"inwind-blog-server-v3/interner/service"
	"inwind-blog-server-v3/utils"
)

func GetBlogList(c *gin.Context) {
	var blogListService service.BlogListSerivce
	//验证参数
	if err := c.ShouldBind(&blogListService); err != nil {
		c.JSON(400, serializer.BuildErrorResponse(utils.InvalidParams, utils.GetMsg(utils.InvalidParams)))
	} else {
		result := blogListService.GetBlogList()
		c.JSON(200, result)
	}

}

func GetBlogDetail(c *gin.Context) {
	var blogDetailService service.BlogDetailService
	if err := c.ShouldBind(&blogDetailService); err != nil {
		c.JSON(400, serializer.BuildErrorResponse(utils.InvalidParams, utils.GetMsg(utils.InvalidParams)))
	} else {
		result := blogDetailService.GetBlogDetail()
		c.JSON(200, result)
	}
}

func EditBlog(c *gin.Context) {
	var editBlogService service.EditBlogService
	var blog model.Blog
	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(400, serializer.BuildErrorResponse(utils.InvalidParams, utils.GetMsg(utils.InvalidParams)))
	}
	editBlogService.EditBlog()
}
