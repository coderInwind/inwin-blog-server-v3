package dao

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/db"
	"inwind-blog-server-v3/model"
	"net/http"
)

func Login(c *gin.Context) {
	var user model.User

	result := db.DbEngine.Take(&user)

	common.Response(c, http.StatusOK, 200, result, "请求成功")

}
