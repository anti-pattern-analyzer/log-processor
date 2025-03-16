package database

import (
	"log"
)

func RunMigrations() {
	if DB == nil {
		log.Fatal("Database connection not initialized. Call ConnectTimescaleDB() first.")
	}

	log.Println("Running TimescaleDB Migrations...")

	migrations := []string{
		`CREATE TABLE IF NOT EXISTS raw_logs (
			trace_id UUID NOT NULL,
			span_id UUID NOT NULL,
			parent_span_id UUID NULL,
			source TEXT NOT NULL,
			destination TEXT NOT NULL,
			method TEXT NOT NULL,
			type TEXT NOT NULL,
			request TEXT NOT NULL,
			http_status INT NOT NULL,
			response TEXT NULL,
			timestamp_ms BIGINT NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS structured_row_logs (
			trace_id UUID NOT NULL,
			span_id UUID PRIMARY KEY,
			parent_span_id UUID NULL,
			source TEXT NOT NULL,
			destination TEXT NOT NULL,
			method TEXT NOT NULL,
			type TEXT NOT NULL,
			request TEXT NOT NULL,
			http_status INT NOT NULL,
			response TEXT NULL,
			start_timestamp_ms BIGINT NOT NULL,
			end_timestamp_ms BIGINT NULL,
			duration_ms BIGINT DEFAULT 0,
			is_complete BOOLEAN DEFAULT FALSE,
			is_duration_updated BOOLEAN DEFAULT FALSE,
			is_processed BOOLEAN DEFAULT FALSE
		);`,
	}

	for _, query := range migrations {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	log.Println("Migrations applied successfully!")
}
