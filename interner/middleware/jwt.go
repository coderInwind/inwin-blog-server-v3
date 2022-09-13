package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common/errcode"
	"inwind-blog-server-v3/common/response"
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
		fmt.Println(claims.RegisteredClaims.ExpiresAt)

		c.Next()
	}

}
