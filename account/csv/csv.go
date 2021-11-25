package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"account.testing.csv/model"
)

var (
	tables = make([]string, 0)
)

func LoadAccountCSV() []model.Account {
	var acc []model.Account
	file, err := os.Open("accounts.csv")
	if err != nil {
		log.Panic(err)
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
			log.Panic(err)
			os.Exit(2)
		}

		acc = append(acc, model.Account{
			Id:          id,
			Name:        line[1],
			Address:     line[2],
			Phonenumber: phonenumber,
			Balance:     balance,
		})

	}
	return acc
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
		if i == -1 {
			w.Write(columns)
			w.Write(row)
		} else {
			w.Write(row)
		}
	}
	w.Flush()
	fmt.Println("Finished processing:")
}
