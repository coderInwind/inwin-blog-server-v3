package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"inwind-blog-server-v3/global"
	"inwind-blog-server-v3/interner/config"
)

func Viper(path ...string) *viper.Viper {

	//判断环境
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件错误")
	}
	var c config.Config
	v.Unmarshal(&c)
	// 指向解析出的结构体实例
	global.Config = &c

	return v
}
