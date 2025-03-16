package response

type StructuredLogResponseDTO struct {
	TraceID        string  `json:"trace_id"`
	SpanID         string  `json:"span_id"`
	ParentSpanID   *string `json:"parent_span_id,omitempty"`
	Source         string  `json:"source"`
	Destination    string  `json:"destination"`
	Method         string  `json:"method"`
	Type           string  `json:"type"`
	Request        string  `json:"request"`
	HttpStatus     int     `json:"http_status_code"`
	Response       *string `json:"response"`
	StartTimestamp uint64  `json:"start_timestamp_ms"`
	EndTimestamp   *uint64 `json:"end_timestamp_ms"`
	DurationMs     uint64  `json:"duration_ms"`
	IsComplete     bool    `json:"is_complete"`
}
