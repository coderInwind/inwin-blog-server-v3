package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"inwind-blog-server-v3/interner/config"
)

var (
	DB     *gorm.DB
	Viper  *viper.Viper
	Config *config.Config
	Redis  *redis.Client
)
