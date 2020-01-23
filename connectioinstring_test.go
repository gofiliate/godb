package godb

import (
	"fmt"
	"testing"
)

func TestGetMysqlReadConnectionString(t *testing.T) {

	config, _ := LoadConfig()
	want := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbrUser, config.DbrPass, config.DbrHost, config.DbrPort, config.DbrName)
	if got := GetMysqlReadConnectionString(config); got != want {
		t.Errorf("GetMysqlReadConnectionString() = %q, want %q", got, want)
	}

}

func TestGetMysqlWriteConnectionString(t *testing.T) {

	config, _ := LoadConfig()
	want := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbwUser, config.DbwPass, config.DbwHost, config.DbwPort, config.DbwName)
	if got := GetMysqlWriteConnectionString(config); got != want {
		t.Errorf("GetMysqlWriteConnectionString() = %q, want %q", got, want)
	}

}



