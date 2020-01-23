package godb

import (
	"testing"
	"fmt"
)

type Organization struct {
	organization_id string
	organization_name string
	status_id int
}

func TestWorkFlow(t *testing.T) {
	fmt.Println("Test Workflow Method")
	ConnectMysqlRead()

	query, err := DBR.Prepare("SELECT organization_id as organization_id, organization_name as organization_name, status_id as status_id FROM organizations")

	if err != nil {
		fmt.Println("Cannot Prepare Query")
		return
	}

	rows, err := query.Query()

	var organizations []Organization
	for rows.Next() {

		var org Organization
		rows.Scan(&org.organization_id, &org.organization_name, &org.status_id)
		organizations = append(organizations, org)
	}

	fmt.Printf("Rows are: %v\n", organizations)

}
