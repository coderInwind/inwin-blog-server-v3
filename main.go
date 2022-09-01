package main

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/router"
)

func main() {
	global.NewDbEngine()

	//测试

	r := router.NewRouter()
	r.Run(":4000")

}
