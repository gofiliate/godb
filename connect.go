package godb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DBR *sql.DB
var DBW *sql.DB

func ConnectMysqlRead()(err error) {


	readConnStr := GetMysqlReadConnectionString()
	DBR, err = sql.Open("mysql", readConnStr)

	if err != nil {
		fmt.Printf("Error attempting to establish READ database connection: %s\n", err.Error())
	}

	err = DBR.Ping()
	if err != nil {
		fmt.Printf("Error pinging DBR: %v\n", err.Error())
	}
	return
}

func ConnectMysqlWrite()(err error) {


	writeConnStr := GetMysqlWriteConnectionString()
	DBW, err = sql.Open("mysql", writeConnStr)

	if err != nil {
		fmt.Printf("Error attempting to establish WRITE database connection: %s\n", err.Error())
	}

	err = DBW.Ping()
	if err != nil {
		fmt.Printf("Error pinging DBR: %v\n", err.Error())
	}

	return
}
