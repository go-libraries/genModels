package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Orm *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:sa@tcp(127.0.0.1:3306)/mgtj_pay?charset=utf8mb4")
	if err != nil {
		panic("连接数据库失败")
	}
	Orm = db
}
