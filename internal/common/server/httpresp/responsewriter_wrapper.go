package httpresp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// We can't log response body by ResponseWriter after handling
// so we need a wrapper to wrap the response
// ResponseWriterWrapper struct is used to log the response
type ResponseWriterWrapper struct {
	w          *http.ResponseWriter
	body       *bytes.Buffer
	statusCode *int
}

// NewResponseWriterWrapper function creates a wrapper for the http.ResponseWriter
func NewResponseWriterWrapper(w http.ResponseWriter) ResponseWriterWrapper {
	var buf bytes.Buffer
	var statusCode int = 200
	return ResponseWriterWrapper{
		w:          &w,
		body:       &buf,
		statusCode: &statusCode,
	}
}

func (rww ResponseWriterWrapper) GetBodyData() (int32, string) {
	var jsonStr BaseResponse
	if json.Unmarshal(rww.body.Bytes(), &jsonStr) != nil {
		return -1, "Error parsing response to json"
	}
	return jsonStr.Code, jsonStr.Message
}

func (rww ResponseWriterWrapper) GetStatusCode() *int {
	return rww.statusCode
}

func (rww ResponseWriterWrapper) Write(buf []byte) (int, error) {
	rww.body.Write(buf)
	return (*rww.w).Write(buf)
}

// Header function overwrites the http.ResponseWriter Header() function
func (rww ResponseWriterWrapper) Header() http.Header {
	return (*rww.w).Header()
}

// WriteHeader function overwrites the http.ResponseWriter WriteHeader() function
func (rww ResponseWriterWrapper) WriteHeader(statusCode int) {
	(*rww.statusCode) = statusCode
	(*rww.w).WriteHeader(statusCode)
}
