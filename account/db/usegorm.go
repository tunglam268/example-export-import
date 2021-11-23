package db

import (
	"fmt"

	model "account.testing.csv/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Open() *gorm.DB {
	db = Connection()
	return db
}

func Connection() *gorm.DB {
	db, err = gorm.Open("mysql", "root:@(localhost)/account?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&model.Account{})
	fmt.Println("Connection success")
	return db
}
