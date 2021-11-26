package api

import (
	"log"
	"net/http"
	"time"

	sp "account.testing.csv/csv"
	database "account.testing.csv/db"
	export "account.testing.csv/exportdata"
)

const batchsize int = 1000

func ImportFromCSV(res http.ResponseWriter, req *http.Request) {
	db := database.OpenGormDB()
	startDB := time.Now()
	var accounts = sp.LoadAccountCSV()
	endDB := time.Since(startDB)

	startImport := time.Now()
	err := db.CreateInBatches(accounts, batchsize).Error
	if err != nil {
		log.Panic(err)
	}
	endImport := time.Since(startImport)
	log.Printf("\n Time to read data from CSV file is : %v \n Time to write to DB is : %v \n", endDB, endImport)
	res.WriteHeader(http.StatusOK)
}

func ExportDBToCSV(res http.ResponseWriter, req *http.Request) {
	startEx := time.Now()
	export.ExportData()
	endEx := time.Since(startEx)
	log.Printf("Time to export db to csv : %v", endEx)
	// db := database.OpenDB()

	// rows, err := db.Query(fmt.Sprintf("SELECT * from accounts"))

	// columns, err := rows.Columns()
	// if err != nil {
	// 	log.Panic(err.Error())
	// }
	// values := make([]sql.RawBytes, len(columns))
	// scanArgs := make([]interface{}, len(values))
	// for i := range values {
	// 	scanArgs[i] = &values[i]
	// }
	// totalValues := make([][]string, 0)
	// for rows.Next() {

	// 	//Save the contents of each line
	// 	var s []string

	// 	//Add the contents of each line to scanArgs, and also to values
	// 	err = rows.Scan(scanArgs...)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	for _, v := range values {
	// 		s = append(s, string(v))
	// 		// print(len(s))
	// 	}
	// 	totalValues = append(totalValues, s)
	// }

	// if err = rows.Err(); err != nil {
	// 	panic(err.Error())
	// }
	// sp.WriteToCSV(columns, totalValues)
	res.WriteHeader(http.StatusOK)
}
