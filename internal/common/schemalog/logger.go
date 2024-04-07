package schemalog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var ignoredURL = []string{
	"/metrics",
	"/healthz",
	"/grpc.health.v1.Health/Check",
}

var logger = NewSchemaLogger()

func NewSchemaLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&RawLogFormatter{})
	return logger
}

type RawLogFormatter struct {
}

func (f *RawLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%v\n", entry.Message)), nil
}

func logRequestEvent(requestEvent RequestEvent) {
	jsonBytes, _ := requestEvent.MarshalJSON()

	logger.Info(string(jsonBytes))
}

func logResponseEvent(responseEvent ResponseEvent) {
	jsonBytes, _ := responseEvent.MarshalJSON()

	logger.Info(string(jsonBytes))
}
