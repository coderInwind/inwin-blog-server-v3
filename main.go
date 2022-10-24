package main

import (
	"fmt"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/initialize"
)

func main() {
	global.Viper = initialize.Viper()
	global.DB = initialize.Gorm()
	global.Redis = initialize.Redis()

	fmt.Println(global.Config.Jwt.SigningKey)
	r := initialize.Routers()

	r.Run(":4500")
}
