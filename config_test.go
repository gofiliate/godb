package godb

import (
	"testing"
	"fmt"

)

func TestReadLoadConfigExternal(t *testing.T) {

	LoadConfigExternal("read", "goapi", "password", "127.0.0.1", "3306", "goapi")

	fmt.Printf("username: %v password: %v host: %v port %v database: %v\n", readUser, readPass, readHost, readPort, readDB)

}