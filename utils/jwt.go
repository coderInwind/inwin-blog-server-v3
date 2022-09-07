package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	Id       uint   `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte("测试"),
	}
}

func (j *JWT) GenerateToken(id uint, username string) (string, error) {
	claim := Claims{
		Id:       id,
		Password: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "inwind-blog",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * time.Duration(1))), //生效时间24小时
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := tokenClaim.SignedString(j.SigningKey)

	return token, err
}
