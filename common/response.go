package common

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code int, data *gorm.DB, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data.Value,
		"msg":  msg,
	})
}

func ResponseError(c *gin.Context, httpStatus int, code int, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "msg": msg})
}

func ResponseOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "The request is successful"})
}
