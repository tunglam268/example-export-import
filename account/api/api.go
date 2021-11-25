package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	sp "account.testing.csv/csv"
	database "account.testing.csv/db"
)

const batchsize int = 1000

func ImportFromCSV(res http.ResponseWriter, req *http.Request) {
	db := database.OpenGormDB()
	start1 := time.Now()
	var accounts = sp.LoadAccountCSV()
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
	sp.WriteToCSV(columns, totalValues)
}
