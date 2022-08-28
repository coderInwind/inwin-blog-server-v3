package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/serializer"
	"inwind-blog-server-v3/service"
	"inwind-blog-server-v3/utils"
)

func GetBlogList(c *gin.Context) {
	var blogListService service.BlogListSerivce
	//验证参数
	if err := c.ShouldBind(&blogListService); err != nil {
		c.JSON(utils.ERROR, serializer.BuildErrorResponse(utils.InvalidParams, utils.GetMsg(utils.InvalidParams)))
	} else {
		result := blogListService.GetBlogList()
		c.JSON(utils.SUCCESS, result)
	}

}

func GetBlogDetail(c *gin.Context) {
	var blogDetailService service.BlogDetail
	if err := c.ShouldBind(&blogDetailService); err != nil {
		c.JSON(utils.ERROR, serializer.BuildErrorResponse(utils.InvalidParams, utils.GetMsg(utils.InvalidParams)))
	} else {
		result := blogDetailService.BlogDetailService()
	}

}
