package repository

import (
	"log"
	"log-processor/database"
	"log-processor/models"
)

func SaveRowLog(rowLog models.RowLog) error {
	query := `
		INSERT INTO raw_logs (
			trace_id, span_id, parent_span_id, source, destination,
			method, request, response, timestamp_ms
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := database.DB.Exec(query,
		rowLog.TraceID,
		rowLog.SpanID,
		rowLog.ParentSpanID,
		rowLog.Source,
		rowLog.Destination,
		rowLog.Method,
		rowLog.Request,
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
