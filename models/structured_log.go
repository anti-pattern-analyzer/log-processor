package models

type StructuredLog struct {
	TraceID           string
	SpanID            string
	ParentSpanID      *string
	Source            string
	Destination       string
	Method            string
	Type              string
	Request           string
	HttpStatus        int
	Response          *string
	StartTimestampMs  uint64
	EndTimestampMs    *uint64
	DurationMs        uint64
	IsComplete        bool
	IsDurationUpdated bool
	IsProcessed       bool
}
