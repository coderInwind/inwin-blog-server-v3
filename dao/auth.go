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

func Login(c *gin.Context) {
	var user model.User

	result := db.DbEngine.Take(&user)

	common.Response(c, http.StatusOK, 200, result, "请求成功")

}

func CreateUser(c *gin.Context, params common.UserParams) {

	encryptedPassword, isExists := c.Get("password")
	if isExists {
		fmt.Printf("%v存在", encryptedPassword)
	} else {
		fmt.Printf("%v不存在", encryptedPassword)
	}

	user := model.User{
		Username:  params.Username,
		Password:  params.Password,
		RoleId:    params.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.DbEngine.Create(&user)
	//fmt.Printf("%v", result)
	common.ResponseOk(c)
}
