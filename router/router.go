package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/controller"
	"inwind-blog-server-v3/middleware"
)

func NewRouter() *gin.Engine {

	r := gin.New()

	r.POST("/login", controller.Login)

	r.POST("/create", middleware.Encrypt(), controller.CreateUser)

	return r
}
