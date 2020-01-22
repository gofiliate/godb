package godb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadConfig() (config ConnectionDetails, err error) {

	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Println("Cannot load config.json file. Make sure the file is in the root of the module and is valid json.")
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		fmt.Println("Error unmarshalling json :", err)
	}

	return
}
