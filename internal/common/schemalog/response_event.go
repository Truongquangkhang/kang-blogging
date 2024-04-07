package schemalog

import (
	"encoding/json"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const RESPONSE_EVENT_SCHEMA_LOG_NAME = "nio_v1_nio_response_events"

type ResponseEvent struct {
	TraceId    string                 `json:"trace_id"`
	Url        string                 `json:"url"`
	StatusCode int32                  `json:"status_code"`
	Code       string                 `json:"code"` // Code should be int32 but string because of fault when creating generated code
	Message    string                 `json:"message"`
	Timestamp  *timestamppb.Timestamp `json:"timestamp"`
}

// Same as RequestEvent
func (event *ResponseEvent) MarshalJSON() ([]byte, error) {
	type Alias ResponseEvent
	return json.Marshal(struct {
		SchemaLogName string `json:"_schemaLogName"`
		*Alias
	}{
		SchemaLogName: RESPONSE_EVENT_SCHEMA_LOG_NAME,
		Alias:         (*Alias)(event),
	})
}
