package godb

import (
	"fmt"
)

func GetMysqlReadConnectionString(config ConnectionDetails) (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbrUser, config.DbrPass, config.DbrHost, config.DbrPort, config.DbrName)

	return

}

func GetMysqlWriteConnectionString(config ConnectionDetails) (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbwUser, config.DbwPass, config.DbwHost, config.DbwPort, config.DbwName)

	return
}
