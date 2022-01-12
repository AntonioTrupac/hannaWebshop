package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connecting")

	db, err := sql.Open("mysql", "ql6e013y0w78:pscale_pw_m0h2YiWrWT9HNiCx9y6816aqe-Yl0J64k_QEtODqF38@tcp(kqmwg2ezxey8.eu-west-3.psdb.cloud)/hannawebshop?tls=true")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	insert, err := db.Query("INSERT INTO user VALUES ( 3, 'Karlo' )")

	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
