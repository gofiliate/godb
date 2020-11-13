package godb

import (
	"testing"
	"fmt"
)

type Brand struct {
	brand_ident string
	brand_name string
	status_id int
}

func TestWorkFlow(t *testing.T) {
	fmt.Println("Test Workflow Method")
	ConnectMysqlRead()

	query, err := DBR.Prepare("SELECT brand_ident as brand_ident, brand_name as brand_name, status_id as status_id FROM brands")

	if err != nil {
		fmt.Printf("Cannot Prepare Query: %v", err.Error())
		return
	}

	rows, err := query.Query()

	var brands []Brand
	for rows.Next() {

		var brd Brand
		rows.Scan(&brd.brand_ident, &brd.brand_name, &brd.status_id)
		brands = append(brands, brd)
	}

	fmt.Printf("Rows are: %v\n", brands)

}
