package initialize

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/middleware"
	"inwind-blog-server-v3/interner/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors)

	//公共的路由，不需要携带token
	PublicGroup := r.Group("/api")
	{
		router.InitBaseUser(PublicGroup)
	}

	//需要鉴权的路由
	PrivateGroup := r.Group("/api")
	PrivateGroup.Use(middleware.JWTAuth())

	{
		router.InitBlogRouter(PrivateGroup)
		router.InitUserRouter(PrivateGroup)
		router.InitTagRouter(PrivateGroup)
		router.InitGalleryRouter(PrivateGroup)
		router.InitUploadRouter(PrivateGroup)
		router.InitCoverRouter(PrivateGroup)
		//router.InitSubmitRecord(PrivateGroup)
	}

	return r
}
