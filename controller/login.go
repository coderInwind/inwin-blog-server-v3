package controller

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/db"
)

func Login(c *gin.Context) {

	db := db.GetDB()

	//多行查询
	row := db.QueryRow("select * from test") //获取所有数据
	var id string

	row.Scan(&id)

	c.JSON(200, gin.H{"id": id})
}
