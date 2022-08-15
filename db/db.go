package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DbEngine *gorm.DB

func NewDbEngine() {
	//连接数据库
	//db, _ = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/blogsdata")
	//
	//// 尝试与数据库建立连接（校验dsn是否正确）
	//err := db.Ping()
	//if err != nil {
	//	fmt.Println("数据库连接失败")
	//	return
	//} else {
	//	fmt.Println("数据库连接成功")
	//}
	var err error
	dsn := "root:mysql@tcp(127.0.0.1:3306)/blogsdata?charset=utf8mb4&parseTime=True&loc=Local"
	DbEngine, err = gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("数据库连接失败")
		return
	} else {
		fmt.Println("数据库连接成功")
	}

}

//func GetDB() *gorm.DB {
//	return db
//}
