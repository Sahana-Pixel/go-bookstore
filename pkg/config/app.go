package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// open connection with db
// if error it will go inside err othersiws d
func Connect() {
	d, err := gorm.Open("mysql", "username:password)/table_name?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)

	}
	db = d //whterv in d will be transferd to db
}

// return db
// getdb fun 2ill be used inother files
// uswed in book.go callled
func GetDB() *gorm.DB {
	return db
}
