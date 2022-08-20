package main

import (
	"inwind-blog-server-v3/db"
	"inwind-blog-server-v3/router"
)

func main() {
	db.NewDbEngine()

	r := router.NewRouter()
	r.Run(":4000")

}
