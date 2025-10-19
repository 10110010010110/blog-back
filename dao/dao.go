package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var DB *gorm.DB

func Initmysql() {
	dsn := "root:root@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic(err)
	}

}
