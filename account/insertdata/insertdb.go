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
	stmt := fmt.Sprintf("INSERT INTO accounts (name, address,phonenumber,balance) VALUES %s", strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	if err != nil {
		log.Println(err)
	}

}
