package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/controller"
)

func NewRouter() *gin.Engine {

	r := gin.New()

	r.GET("/login", controller.Login)

	return r
}
