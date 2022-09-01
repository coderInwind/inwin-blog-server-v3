package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbEngine *gorm.DB

func NewDbEngine() {
	var err error
	dsn := "root:mysql@tcp(127.0.0.1:3306)/inwind_blog?charset=utf8mb4&parseTime=True&loc=Local"

	DbEngine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		fmt.Println("数据库连接失败")
		return
	} else {
		fmt.Println("数据库连接成功")
	}
}
