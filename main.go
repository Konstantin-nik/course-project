package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
	//hello string
	//ans   string
	//num   int
)

func main() {
	go RunServer()
	connStr := "user=postgres password=1122334455qq dbname=productdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer db.Close()
	fmt.Print("All is ok!")
	result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Print("Wow!")
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}
