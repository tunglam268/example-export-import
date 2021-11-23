package api

import (
	exportdb "account.testing.csv/exportdata"
	importdb "account.testing.csv/importdata"
	insert "account.testing.csv/insert"
)

func Create() {
	insert.InsertDB()
}

func ImportDB() {
	importdb.ImportData()
}

func ExportDB() {
	exportdb.ExportData()
}
