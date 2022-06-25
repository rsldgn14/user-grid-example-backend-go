package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	dns = "root:root@tcp(localhost:3306)/usersdb?parseTime=True&charset=utf8"
)

func Connect() {
	d, err := gorm.Open("mysql", dns)
	if err != nil {
		panic(err)

	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
