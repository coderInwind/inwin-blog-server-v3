package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func NewDbEngine() {
	//连接数据库
	db, _ = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/blogsdata")

	// 尝试与数据库建立连接（校验dsn是否正确）
	err := db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	} else {
		fmt.Println("数据库连接成功")
	}

}

func GetDB() *sql.DB {
	return db
}
