package main

import "account.testing.csv/api"

func main() {
	for i := 0; i < 1; i++ {
		api.Create()
	}
	api.ExportDB()
	api.ImportDB()
}
