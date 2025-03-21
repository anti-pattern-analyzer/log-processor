package util

import (
	"errors"
	"log-processor/payload/request"
	"strconv"
	"strings"
)

func ParseLogLine(logLine string) (request.RowLogRequestDTO, error) {
	parts := strings.SplitN(logLine, "|", 2)
	if len(parts) < 2 {
		return request.RowLogRequestDTO{}, errors.New("invalid log format")
	}

	timestampStr := strings.TrimSpace(parts[0])
	data := strings.TrimSpace(parts[1])

	fields := make(map[string]string)
	for _, field := range strings.Split(data, ",") {
		kv := strings.SplitN(strings.TrimSpace(field), "=", 2)
		if len(kv) != 2 {
			continue
		}
		fields[kv[0]] = kv[1]
	}

	timestampMs, err := strconv.ParseUint(timestampStr, 10, 64)
	if err != nil {
		return request.RowLogRequestDTO{}, errors.New("invalid timestamp format, expected milliseconds as uint64")
	}

	var parentSpanID *string
	if val, exists := fields["parent_span_id"]; exists {
		val = strings.TrimSpace(val)
		if val == "" || val == "null" {
			parentSpanID = nil
		} else {
			parentSpanID = &val
		}
	}

	var httpStatus int
	var response *string
	if val, exists := fields["http_status"]; exists {
		parts := strings.SplitN(strings.TrimSpace(val), " response=", 2)
		statusPart := strings.TrimSpace(parts[0])

		httpStatus, err = strconv.Atoi(statusPart)
		if err != nil {
			httpStatus = 500
		}

		if len(parts) > 1 && strings.TrimSpace(parts[1]) != "" {
			resp := strings.TrimSpace(parts[1])
			response = &resp
		}
	}

	logDTO := request.RowLogRequestDTO{
		Timestamp:    timestampMs,
		TraceID:      fields["trace_id"],
		SpanID:       fields["span_id"],
		ParentSpanID: parentSpanID,
		Source:       fields["source"],
		Destination:  fields["destination"],
		Method:       fields["method"],
		Type:         fields["type"],
		Request:      fields["request"],
		HttpStatus:   httpStatus,
		Response:     response,
	}

	return logDTO, nil
}
