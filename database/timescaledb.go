package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectTimescaleDB() {
	var err error
	connStr := "host=localhost port=5433 user=user password=password dbname=logs sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to TimescaleDB: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("TimescaleDB connection failed: %v", err)
	}

	fmt.Println("Connected to TimescaleDB!")
}
