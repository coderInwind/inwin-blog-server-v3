package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/db"
	"inwind-blog-server-v3/model"
	"net/http"
	"time"
)

func CreateUser(c *gin.Context, params common.UserParams) {

	encryptedPassword, isExists := c.Get("password")

	if !isExists {
		return
	}

	user := model.User{
		Username:  params.Username,
		Password:  encryptedPassword.(string), //密码做加密处理
		RoleId:    params.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.DbEngine.Create(&user)
	common.ResponseOk(c)
}

func GetUserList(c *gin.Context) {
	var user []model.User
	// preload嵌套预加载 - 结构体中嵌套有结构体使用
	result := db.DbEngine.Preload("Role").Find(&user)
	fmt.Printf("%v", user)
	common.Response(c, http.StatusOK, 200, result, "The request is successful")
}
