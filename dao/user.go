package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/db"
	"inwind-blog-server-v3/model"
)

func CreateUser(c *gin.Context, params common.UserParams) {

	encryptedPassword, isExists := c.Get("password")

	if !isExists {
		return
	}

	user := model.User{
		Username: params.Username,
		Password: encryptedPassword.(string), //密码做加密处理
		RoleId:   params.Role,
	}

	db.DbEngine.Create(&user)
	common.ResponseOk(c)
}

func SelectUserList(s int, i int) ([]model.User, int64) {
	// preload嵌套预加载 - 结构体中嵌套有结构体使用
	var users []model.User
	var user model.User
	var total int64

	fmt.Println(s, i)

	db.DbEngine.Preload("Role").Limit(s).Offset((i - 1) * s).Find(&users)
	db.DbEngine.Model(user).Count(&total)
	return users, total
}
