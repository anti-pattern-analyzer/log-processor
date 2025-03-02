package services

import (
	"log"
	"log-processor/mappers"
	"log-processor/payload/request"
	"log-processor/repository"
)

func ProcessRowLog(logDTO request.RowLogRequestDTO) error {
	rowLog := mappers.MapRowLogRequestToRowLog(logDTO)

	err := repository.SaveRowLog(rowLog)
	if err != nil {
		log.Printf("Error saving raw log: %v", err)
		return err
	}

	structuredLog := mappers.MapRowLogToStructuredLog(rowLog)
	err = repository.SaveStructuredLog(structuredLog)
	if err != nil {
		log.Printf("Error saving structured log: %v", err)
		return err
	}

	return nil
}
