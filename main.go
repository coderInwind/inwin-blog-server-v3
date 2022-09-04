package main

import (
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/initialize"
)

func main() {
	global.DB = initialize.Gorm()

	r := initialize.Routers()

	r.Run(":4000")
}
