package main

import (
	"couchdb"
	"fmt"
)

func main() {
	cdbcli, err := couchdb.Init("./config.json")
	if err != nil {
		fmt.Println(err)
	}

	cdbcli.GetAllHeaders()

	// ab, _ := json.Marshal(cdbcli)
	// fmt.Println(string(ab))
}
