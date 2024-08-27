package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	connstr := "name:pass/tablename?charset=utf8&parseTime=True&loc=Local"
	// connect to mysql db
	d, err := gorm.Open("mysql", connstr)
	if err != nil {
		log.Printf("Error connecting to db: %v", err)
		panic(err)
	}
	// assign the mysql db to gorm db
	db = d
}

func GetDB() *gorm.DB {
	return db
}
