package db

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	database "account.testing.csv/db"
	model "account.testing.csv/model"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var err error

func genRandomString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func stubAccounts() (accounts []*model.Account) {
	for i := 0; i < 16300; i++ {
		account := &model.Account{
			Name:        genRandomString(10),
			Address:     genRandomString(10),
			Phonenumber: rand.Int63(),
			Balance:     rand.Int63(),
		}
		accounts = append(accounts, account)
	}

	return accounts
}

func InsertDB() {
	db := database.OpenDB()

	log.Printf("Successfully connected to database")
	accounts := stubAccounts()

	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, account := range accounts {
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		valueArgs = append(valueArgs, account.Name)
		valueArgs = append(valueArgs, account.Address)
		valueArgs = append(valueArgs, account.Balance)
		valueArgs = append(valueArgs, account.Phonenumber)

	}

	_, err = db.Query("CREATE DATABASE IF NOT EXISTS account;")
	if err != nil {
		log.Println(err)
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS `account`.`accounts` (`id` INT NOT NULL AUTO_INCREMENT,`name` VARCHAR(255) NULL ,`address` VARCHAR(255) NULL,`phonenumber` VARCHAR(255) NOT NULL ,`balance` VARCHAR(255) NOT NULL,PRIMARY KEY (`id`));")
	if err != nil {
		log.Println(err)
	}

	stmt := fmt.Sprintf("INSERT INTO accounts (name, address,phonenumber,balance) VALUES %s", strings.Join(valueStrings, ","))

	_, err = db.Exec(stmt, valueArgs...)
	if err != nil {
		log.Println(err)
	}

}
