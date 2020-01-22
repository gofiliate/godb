package godb

import (
	"fmt"
	"testing"
)

func TestGetReadConnectionString(t *testing.T) {

	config, _ := LoadConfig()

	want := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbrUser, config.DbrPass, config.DbrHost, config.DbrPort, config.DbrName)
	fmt.Printf("Want R: %s\n", want)

	if got := GetReadConnectionString(config); got != want {
		t.Errorf("GetReadConnectionString() = %q, want %q", got, want)
	}

}

func TestGetWriteConnectionString(t *testing.T) {

	config, _ := LoadConfig()

	want := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbwUser, config.DbwPass, config.DbwHost, config.DbwPort, config.DbwName)
	fmt.Printf("Want W: %s\n", want)

	if got := GetWriteConnectionString(config); got != want {
		t.Errorf("GetWriteConnectionString() = %q, want %q", got, want)
	}

}



