package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "anaet:password/book-store?charset=utf8&parseTime=True&loc=Local")

	if err {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
