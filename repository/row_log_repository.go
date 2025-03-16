package repository

import (
	"fmt"
	"log"
	"log-processor/database"
	"log-processor/models"
)

func SaveRowLog(rowLog models.RowLog) error {
	query := `
		INSERT INTO raw_logs (
			trace_id, span_id, parent_span_id, source, destination,
			method, type, request, http_status, response, timestamp_ms
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
	`

	_, err := database.DB.Exec(query,
		rowLog.TraceID,
		rowLog.SpanID,
		rowLog.ParentSpanID,
		rowLog.Source,
		rowLog.Destination,
		rowLog.Method,
		rowLog.Type,
		rowLog.Request,
		rowLog.HttpStatus,
		rowLog.Response,
		rowLog.TimestampMs,
	)

	if err != nil {
		log.Printf("Error inserting raw log: %v", err)
		return err
	}

	log.Println("Raw log saved:", rowLog.TraceID)
	return nil
}

// GetAllRowLogs fetches all logs from the database with sorting (asc/desc)
func GetAllRowLogs(sortOrder string) ([]models.RowLog, error) {
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	query := fmt.Sprintf(`
		SELECT trace_id, span_id, parent_span_id, source, destination,
			   method, type, request, http_status, response, timestamp_ms
		FROM raw_logs 
		ORDER BY trace_id %s, parent_span_id IS NULL DESC, timestamp_ms ASC;
	`, sortOrder)

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Error fetching row logs: %v", err)
		return nil, err
	}

	var logs []models.RowLog
	for rows.Next() {
		var logEntry models.RowLog
		err := rows.Scan(
			&logEntry.TraceID, &logEntry.SpanID, &logEntry.ParentSpanID,
			&logEntry.Source, &logEntry.Destination, &logEntry.Method,
			&logEntry.Type, &logEntry.Request, &logEntry.HttpStatus,
			&logEntry.Response, &logEntry.TimestampMs,
		)
		if err != nil {
			log.Printf("Error scanning row log: %v", err)
			continue
		}
		logs = append(logs, logEntry)
	}
	return logs, nil
}
