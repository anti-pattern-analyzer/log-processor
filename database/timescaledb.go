package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func ConnectTimescaleDB() {
	var err error

	host := getEnv("TIMESCALEDB_HOST", "localhost")
	port := getEnv("TIMESCALEDB_PORT", "5433")
	user := getEnv("TIMESCALEDB_USER", "user")
	password := getEnv("TIMESCALEDB_PASSWORD", "password")
	dbname := getEnv("TIMESCALEDB_NAME", "logs")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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
