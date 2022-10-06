package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	//判断环境
	fmt.Println(gin.Mode())

	v := viper.New()
	return v
}
