package db

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

var (
	db              *gorm.DB
	err             error
	tables          = make([]string, 0)
	dataSourceName  = ""
	driverNameMysql = "mysql"
)

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
