package utils

import (
	"context"
	"inwind-blog-server-v3/global"
	"time"
)

func SetRedisJwt(jwt string, username string) (err error) {
	// 过期时间等于jwt过期时间
	time := time.Duration(global.Config.Jwt.ExpiresTime) * time.Second
	err = global.Redis.Set(context.Background(), username, jwt, time).Err()
	return err

}
