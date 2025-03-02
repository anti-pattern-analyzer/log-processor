package request

type RowLogRequestDTO struct {
	TraceID      string  `json:"trace_id"`
	SpanID       string  `json:"span_id"`
	ParentSpanID *string `json:"parent_span_id,omitempty"`
	Source       string  `json:"source"`
	Destination  string  `json:"destination"`
	Method       string  `json:"method"`
	Type         string  `json:"type"`
	Request      string  `json:"request"`
	Response     *string `json:"response"`
	Timestamp    uint64  `json:"timestamp"`
}
