package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db              *gorm.DB
	err             error
	tables          = make([]string, 0)
	dataSourceName  = ""
	driverNameMysql = "mysql"
)

func Open() *gorm.DB {
	db = Connection()
	return db
}

func Connection() *gorm.DB {
	db, err = gorm.Open("mysql", "root:@(localhost)/account?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	// db.AutoMigrate(&model.Account{})
	fmt.Println("Connection success")
	return db
}

func OpenDB() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/account?charset=utf8")
	if err != nil {
		panic(err.Error())

	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
