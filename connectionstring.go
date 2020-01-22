package godb

import (
	"fmt"
	"strings"
	"errors"
)

func GetReadConnectionString(config ConnectionDetails) (connectionString string) {

	if len(config.SslMode) == 0 {

		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbrUser, config.DbrPass, config.DbrHost, config.DbrPort, config.DbrName)
	} else {
		validSSL, _ := validateSslMode(config.SslMode)

		if validSSL {
			connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?ssl-mode=%s", config.DbrUser, config.DbrPass, config.DbrHost, config.DbrPort, config.DbrName, config.SslMode)
		}
	}

	return

}

func GetWriteConnectionString(config ConnectionDetails) (connectionString string) {

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbwUser, config.DbwPass, config.DbwHost, config.DbwPort, config.DbwName)

	return
}


func validateSslMode(mode string) (valid bool, err error) {

	valid = false
	if strings.ToUpper(mode) == "REQUIRED" || strings.ToUpper(mode) == "REQUIRED" || strings.ToUpper(mode) == "REQUIRED" || strings.ToUpper(mode) == "REQUIRED" {
		valid = true
	} else {
		err = errors.New("ERROR: Supplied SslMode incompitable")
	}
	return
}
