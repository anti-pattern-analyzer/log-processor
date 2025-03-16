package services

import (
	"log-processor/mappers"
	"log-processor/payload/response"
	"log-processor/repository"
)

// GetGroupedRowLogs retrieves and structures row logs hierarchically with sorting
func GetGroupedRowLogs(sortOrder string) (map[string][]response.RowLogResponseDTO, error) {
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	rowLogs, err := repository.GetAllRowLogs(sortOrder)
	if err != nil {
		return nil, err
	}

	traceMap := make(map[string][]response.RowLogResponseDTO)
	var responseDTOs []response.RowLogResponseDTO

	for _, log := range rowLogs {
		responseDTOs = append(responseDTOs, mappers.MapRowLogToRowLogResponseDto(log))
	}

	for _, log := range responseDTOs {
		traceMap[log.TraceID] = append(traceMap[log.TraceID], log)
	}

	return traceMap, nil
}
