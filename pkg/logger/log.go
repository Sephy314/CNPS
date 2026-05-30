package logger

import (
	"github.com/Sephy314/cnps/pkg/types/status"
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
	Msg   any
	Level LoggingLevel
}

type ResponseLog struct {
	Log
	ReqID  string
	Status status.Status
}
