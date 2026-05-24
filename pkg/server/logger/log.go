package logger

import (
	"github.com/Sephy314/cnps/pkg/server/status"
)

type LoggingLevel string

const (
	DEBUG LoggingLevel = "DEBUG"
	INFO  LoggingLevel = "INFO"
	WARN  LoggingLevel = "WARN"
	ERROR LoggingLevel = "ERROR"
	FATAL LoggingLevel = "FATAL"
)

type Log struct {
	Msg    string
	Level  LoggingLevel
	Fields map[string]interface{} `json:",omitempty"`
}

type ResponseLog struct {
	Log
	ReqID  string
	Status status.Status
}

type Logging struct {
	Msg   string
	Level LoggingLevel
}

type Logger interface {
	Print()
}
