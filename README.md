# Gofiliate\godb
A simple to use Go Module to handle seperate Read/Write connections to a MySql database.

## Installation
To add to your project
```go
# sudo go get github.com/gofiliate/godb
```

## Configuration
Connection details for the database should be placed in a file called config.json. Simply copy the config.dist.json file to config.json

```json
# sudo cp config.dist.json config.json 
# nano config.json 
{
  "DbrUser" : "read_username",
  "DbrPass" : "read_password",
  "DbrHost" : "read_host",
  "DbrPort" : "read_port",
  "DbrName" : "read_database",
  "DbwUser" : "write_username",
  "DbwPass" : "write_password",
  "DbwHost" : "write_host",
  "DbwPort" : "write_port",
  "DbwName" : "write_database"
}
```

## Usage 

```go

package main


import (
    "github.com/gofiliate/godb"
    "fmt"
)

type Organization struct {
	organization_id string
	organization_name string
	status_id int
}


func main() {
    
        //Create a Read Only Connection
	godb.ConnectMysqlRead()

	query, err := godb.DBR.Prepare("SELECT organization_id as organization_id, organization_name as organization_name, status_id as status_id FROM organizations")

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






```
