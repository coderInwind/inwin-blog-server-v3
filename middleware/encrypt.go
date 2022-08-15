package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt() gin.HandlerFunc {
	return func(c *gin.Context) {

		pswStr := c.PostForm("password")

		psw := []byte(pswStr)
		hash, err := bcrypt.GenerateFromPassword(psw, bcrypt.MinCost)

		if err != nil {
			return
		}
		c.Set("password", string(hash))
		c.Next()
	}

}
