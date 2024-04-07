package schemalog

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const REQUEST_EVENT_SCHEMA_LOG_NAME = "nio_v1_nio_request_events"

type RequestEvent struct {
	TraceId   string                 `json:"trace_id"`
	Url       string                 `json:"url"`
	Method    string                 `json:"method"`
	Params    []string               `json:"params"`
	Headers   []string               `json:"headers"`
	UserId    string                 `json:"user_id"`
	Timestamp *timestamppb.Timestamp `json:"timestamp"`
}

func (event *RequestEvent) MarshalJSON() ([]byte, error) {
	type Alias RequestEvent
	return json.Marshal(struct {
		SchemaLogName string `json:"_schemaLogName"`
		*Alias
	}{
		SchemaLogName: REQUEST_EVENT_SCHEMA_LOG_NAME,
		Alias:         (*Alias)(event),
	})
}
