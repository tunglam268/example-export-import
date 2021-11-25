package api

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	database "account.testing.csv/db"
	"account.testing.csv/model"
)

const batchsize int = 1000

func ImportFromCSV(res http.ResponseWriter, req *http.Request) {
	db := database.OpenGormDB()
	start1 := time.Now()
	var accounts = LoadAccountCSV()
	end1 := time.Since(start1)

	start2 := time.Now()
	err := db.CreateInBatches(accounts, batchsize).Error
	if err != nil {
		log.Println(err)
	}
	end2 := time.Since(start2)
	fmt.Printf("\n Time to read data from CSV file is : %v \n Time to write to DB is : %v \n", end1, end2)
	res.WriteHeader(http.StatusOK)
}

func LoadAccountCSV() []model.Account {
	var acc []model.Account
	file, err := os.Open("accounts.csv")
	if err != nil {
		log.Println(err)
	}
	reader := csv.NewReader(bufio.NewReader(file))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.ParseInt(line[0], 0, 64)
		phonenumber, err := strconv.ParseInt(line[3], 0, 64)
		balance, err := strconv.ParseInt(line[4], 0, 64)

		if err != nil {
			log.Println(err)
			os.Exit(2)
		}

		acc = append(acc, model.Account{
			Id:          int64(id),
			Name:        line[1],
			Address:     line[2],
			Phonenumber: int64(phonenumber),
			Balance:     int64(balance),
		})

	}
	return acc
}

func ExportDBToCSV(res http.ResponseWriter, req *http.Request) {
	db := database.OpenDB()
	rows, err := db.Query(fmt.Sprintf("SELECT * from accounts"))

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	totalValues := make([][]string, 0)
	for rows.Next() {

		//Save the contents of each line
		var s []string

		//Add the contents of each line to scanArgs, and also to values
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for _, v := range values {
			s = append(s, string(v))
			// print(len(s))
		}
		totalValues = append(totalValues, s)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	WriteToCSV(columns, totalValues)
}

func WriteToCSV(columns []string, totalValues [][]string) {
	f, err := os.Create("accounts.csv")
	// fmt.Println(columns)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	//f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	for i, row := range totalValues {
		//First write column name + first row of data
		if i == 0 {
			w.Write(columns)
			w.Write(row)
		} else {
			w.Write(row)
		}
	}
	w.Flush()
	fmt.Println("Finished processing:")
}
