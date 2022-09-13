package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
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
		SigningKey: []byte("测试"),
	}
}

func (j *JWT) CreateClaims(id uint, username string) Claims {
	claim := Claims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "inwind-blog",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * time.Duration(1))), //生效时间24小时
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
