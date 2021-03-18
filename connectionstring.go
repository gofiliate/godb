package github.com/gofiliate/godb

import (
	"fmt"
)

func GetMysqlReadConnectionString() (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", readUser, readPass, readHost, readPort, readDB)

	return

}

func GetMysqlWriteConnectionString() (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", writeUser, writePass, writeHost, writePort, writeDB)

	return
}
