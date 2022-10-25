package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"inwind-blog-server-v3/global"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		// 签发人
		SigningKey: []byte(global.Config.Jwt.Issuer),
	}
}

func (j *JWT) CreateClaims(id uint, username string) Claims {
	jwtCfg := global.Config.Jwt

	claim := Claims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: jwtCfg.SigningKey,
			// 设置和jwt一样的生效时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+jwtCfg.ExpiresTime, 0)),
		},
	}

	return claim
}

func (j *JWT) GenerateToken(claims Claims) (string, error) {

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaim.SignedString(j.SigningKey)

	return token, err
}

func (j *JWT) ParseToken(tokenString string) *Claims {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		fmt.Println("错误了")
	}

	return token.Claims.(*Claims)
}
