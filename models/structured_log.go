package models

type StructuredLog struct {
	TraceID           string
	SpanID            string
	ParentSpanID      *string
	Source            string
	Destination       string
	Method            string
	Request           string
	Response          *string
	StartTimestampMs  uint64
	EndTimestampMs    *uint64
	DurationMs        uint64
	IsComplete        bool
	IsDurationUpdated bool
}
