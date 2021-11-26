package main

import (
	"log"
	"net/http"

	"account.testing.csv/api"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the HTTP server on port 5000")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/import", api.ImportFromCSV).Methods("POST")
	router.HandleFunc("/api/export", api.ExportDBToCSV).Methods("GET")
	router.HandleFunc("/api/insert", api.InsertToDB).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))

}
