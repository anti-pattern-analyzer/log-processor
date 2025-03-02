package services

import (
	"log"
	"log-processor/repository"
)

func UpdateGraphForCompletedTraces() error {
	logs, err := repository.GetCompletedUnprocessedTraces()
	if err != nil {
		return err
	}

	for _, structuredLog := range logs {
		err := repository.UpdateGraphForTrace(structuredLog)
		if err != nil {
			log.Printf("Error updating graph in Neo4j for span_id %s: %v", structuredLog.SpanID, err)
			continue
		}

		err = repository.MarkTraceAsProcessed(structuredLog.SpanID)
		if err != nil {
			log.Printf("Failed to mark trace as processed: %s", structuredLog.SpanID)
		}
	}

	log.Println("Graph updated successfully for completed traces.")
	return nil
}
