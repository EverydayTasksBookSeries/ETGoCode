package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=hello sslmode=disable password=123456"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("INSERT INTO student VALUES(4, '赵云')")
	if err != nil {
		panic(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rows affected = %d\n", rowCount)
}
