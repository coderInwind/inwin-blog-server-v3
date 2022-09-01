package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/service"
)

func GetRoleList(c *gin.Context) {
	result := service.RoleListService()
	c.JSON(200, result)
}

func GetRoleOptions(c *gin.Context) {
	result := service.RoleOptionsService()
	c.JSON(200, result)
}
