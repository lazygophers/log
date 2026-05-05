package log

import "github.com/lazygophers/log/constant"

// Level re-exports constant.Level for convenience
type Level = constant.Level

// Level constants re-exported from constant package
const (
	PanicLevel = constant.PanicLevel
	FatalLevel = constant.FatalLevel
	ErrorLevel = constant.ErrorLevel
	WarnLevel  = constant.WarnLevel
	InfoLevel  = constant.InfoLevel
	DebugLevel = constant.DebugLevel
	TraceLevel = constant.TraceLevel
)
