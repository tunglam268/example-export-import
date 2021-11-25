package db

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Connector *gorm.DB
	err       error
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

func OpenGormDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/account?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Connection was successful!!")
	return db
}
