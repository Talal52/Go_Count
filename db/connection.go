package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "host=localhost user=postgres password=12345 dbname=go_count sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database!")
}
