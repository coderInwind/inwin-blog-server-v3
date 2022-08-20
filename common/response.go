package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response[T any](c *gin.Context, httpStatus int, code int, data T, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func ResponseError(c *gin.Context, httpStatus int, code int, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "msg": msg})
}

func ResponseOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "The request is successful"})
}
