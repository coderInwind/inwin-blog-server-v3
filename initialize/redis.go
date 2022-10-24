package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"inwind-blog-server-v3/global"
)

func Redis() *redis.Client {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(pong, err)
	}
	return client
}
