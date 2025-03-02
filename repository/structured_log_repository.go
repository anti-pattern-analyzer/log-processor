package repository

import (
	"log"
	"log-processor/database"
	"log-processor/models"
)

func SaveStructuredLog(structuredLog models.StructuredLog) error {
	query := `
		INSERT INTO structured_row_logs (
			trace_id, span_id, parent_span_id, source, destination,
			method, request, response, type, start_timestamp_ms, end_timestamp_ms, duration_ms, is_duration_updated
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NULL, 0, FALSE)
		ON CONFLICT (span_id) 
		DO UPDATE SET 
			response = EXCLUDED.response,
			end_timestamp_ms = CASE 
				WHEN EXCLUDED.response IS NOT NULL THEN EXCLUDED.start_timestamp_ms 
				ELSE structured_row_logs.end_timestamp_ms 
			END,
			duration_ms = CASE 
				WHEN EXCLUDED.response IS NOT NULL THEN EXCLUDED.start_timestamp_ms - structured_row_logs.start_timestamp_ms 
				ELSE structured_row_logs.duration_ms 
			END,
			is_duration_updated = CASE 
				WHEN EXCLUDED.response IS NOT NULL THEN TRUE 
				ELSE structured_row_logs.is_duration_updated 
			END;
	`

	_, err := database.DB.Exec(query,
		structuredLog.TraceID,
		structuredLog.SpanID,
		structuredLog.ParentSpanID,
		structuredLog.Source,
		structuredLog.Destination,
		structuredLog.Method,
		structuredLog.Type,
		structuredLog.Request,
		structuredLog.Response,
		structuredLog.StartTimestampMs,
	)

	if err != nil {
		log.Printf("Error inserting/updating structured log: %v", err)
		return err
	}

	log.Println("Structured log inserted/updated:", structuredLog.TraceID)
	return nil
}

func MarkCompletedTraces() error {
	query := `
		UPDATE structured_row_logs
		SET is_complete = TRUE
		WHERE trace_id IN (
			SELECT trace_id FROM structured_row_logs
			GROUP BY trace_id
			HAVING (EXTRACT(EPOCH FROM NOW()) * 1000 - MAX(end_timestamp_ms)) > 10000
		) AND is_complete = FALSE;
	`

	_, err := database.DB.Exec(query)
	if err != nil {
		log.Printf("Error updating completed traces: %v", err)
		return err
	}

	log.Println("Traces marked as complete!")
	return nil
}
