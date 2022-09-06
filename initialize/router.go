package initialize

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/router"
)

//gin.Default()里面调用了gin.New();在调用完gin.New()得到Engine 实例后,还调用了engine.Use(Logger(), Recovery());gin.Default()获取到的Engine 实例集成了Logger 和 Recovery 中间件
//New 返回一个新的空白 Engine 实例，没有附加任何中间件。

func Routers() *gin.Engine {
	r := gin.Default()

	publicGroup := r.Group("/api")
	{
		router.InitBlogRouter(publicGroup)
		router.InitUserRouter(publicGroup)
		router.InitTagRouter(publicGroup)
	}

	return r
}
