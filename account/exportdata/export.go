package export

// Export data from Mysql to a CSV file.

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	tables         = make([]string, 0)
	dataSourceName = ""
)

const (
	driverNameMysql = "mysql"

	helpInfo = `Usage of mysqldataexport:
  -port int
     	the port for mysql,default:3306
  -addr string
    	the address for mysql,default:127.0.0.1
  -user string
    	the username for login mysql,default:root

  -pwd 	string
    	the password for login mysql by the username,default:
  -db 	string
    	the port for me to listen on,default:account
  -tables string
    	the tables will export data, multi tables separator by comma, default:op_log,sc_log,sys_log
	`
)

func init() {

	port := flag.Int("port", 3306, "the port for mysql,default:3306")
	addr := flag.String("addr", "127.0.0.1", "the address for mysql,default:127.0.0.1")
	user := flag.String("user", "root", "the username for login mysql,default:root")
	pwd := flag.String("pwd", "", "the password for login mysql by the username,default:")
	db := flag.String("db", "account", "the port for me to listen on,default:account")
	tabs := flag.String("tables", "accounts", "the tables will export data, multi tables separator by comma, default:op_log,sc_log,sys_log")

	flag.Usage = usage

	flag.Parse()

	tables = append(tables, strings.Split(*tabs, ",")...)

	dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", *user, *pwd, *addr, *port, *db)
}

func usage() {
	fmt.Fprint(os.Stderr, helpInfo)
	flag.PrintDefaults()
}

func querySQL(db *sql.DB, table string, ch chan bool) {
	fmt.Println("Start processing:", table)
	rows, err := db.Query(fmt.Sprintf("SELECT * from %s", table))

	if err != nil {
		panic(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	//Values: all values of a row. Put each field of each row into values. Values length = = number of columns
	values := make([]sql.RawBytes, len(columns))
	// print(len(values))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//Save contents of all lines totalValues
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
	writeToCSV(table+".csv", columns, totalValues)
	ch <- true
}

//writeToCSV
func writeToCSV(file string, columns []string, totalValues [][]string) {
	f, err := os.Create(file)
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
	fmt.Println("Finished processing:", file)
}

func ExportData() {

	count := len(tables)
	ch := make(chan bool, count)

	db, err := sql.Open(driverNameMysql, dataSourceName)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	// Open doesn't open a connection. Validate DSN data:

	for _, table := range tables {
		go querySQL(db, table, ch)
	}

	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Println("Done!")
}
