package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Authority int    `json:"authority"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte("测试")

func GenerateToken(username string, password string) (string, error) {

	claim := Claims{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "inwind-blog",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * time.Duration(1))), //生效时间24小时
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := tokenClaim.SignedString(jwtSecret)

	return token, err
}
