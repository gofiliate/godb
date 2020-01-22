package godb

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	var blankStruct ConnectionDetails
	want := reflect.TypeOf(blankStruct)

	if got, _ := LoadConfig(); reflect.TypeOf(got) != want {
		t.Errorf("LoadConfig() = %q, want %q", got, want)
	}

}