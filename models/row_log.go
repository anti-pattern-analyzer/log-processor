package models

type RowLog struct {
	TraceID      string
	SpanID       string
	ParentSpanID *string
	Source       string
	Destination  string
	Method       string
	Type         string
	Request      string
	Response     *string
	TimestampMs  uint64
}
