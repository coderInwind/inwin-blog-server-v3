package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/response"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/model"
	"inwind-blog-server-v3/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		res := response.NewResponse(c)
		if token == "" {
			res.FailWithMsg(errcode.UnauthorizedTokenError)
			c.Abort() //跳出
			return
		}

		//解析校验
		jwt := utils.NewJWT()
		claims := jwt.ParseToken(token)

		// 将用户信息传入上下文
		c.Set("claims", claims)

		//查询用户的权限等级
		user := model.User{}
		global.DB.Preload("Role").Find(&user)
		fmt.Println(user)
		c.Next()
	}

}
