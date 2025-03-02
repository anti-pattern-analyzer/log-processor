package mappers

import (
	"log-processor/models"
	"log-processor/payload/request"
)

// MapRowLogRequestToRowLog converts RowLogRequestDTO → RowLog Model
func MapRowLogRequestToRowLog(dto request.RowLogRequestDTO) models.RowLog {
	return models.RowLog{
		TraceID:      dto.TraceID,
		SpanID:       dto.SpanID,
		ParentSpanID: dto.ParentSpanID,
		Source:       dto.Source,
		Destination:  dto.Destination,
		Method:       dto.Method,
		Type:         dto.Type,
		Request:      dto.Request,
		Response:     dto.Response,
		TimestampMs:  dto.Timestamp,
	}
}

// MapRowLogToStructuredLog converts RowLog → StructuredLog Model
func MapRowLogToStructuredLog(rowLog models.RowLog) models.StructuredLog {
	return models.StructuredLog{
		TraceID:           rowLog.TraceID,
		SpanID:            rowLog.SpanID,
		ParentSpanID:      rowLog.ParentSpanID,
		Source:            rowLog.Source,
		Destination:       rowLog.Destination,
		Method:            rowLog.Method,
		Type:              rowLog.Type,
		Request:           rowLog.Request,
		Response:          rowLog.Response,
		StartTimestampMs:  rowLog.TimestampMs,
		EndTimestampMs:    nil,
		DurationMs:        0,
		IsComplete:        false,
		IsDurationUpdated: false,
		IsProcessed:       false,
	}
}
