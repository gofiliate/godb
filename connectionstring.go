package godb

import (
	"fmt"
)

func GetMysqlReadConnectionString() (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", readUser, readPass, readHost, readPort, readDB)

	return

}

func GetMysqlWriteConnectionString() (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", writeUser, writePass, writeHost, writePort, writeDB)

	return
}
