package log

import (
	"fmt"
)

// Level represents log level, compatible with logrus
type Level uint32

const (
	// PanicLevel represents panic level (highest)
	PanicLevel Level = iota

	// FatalLevel represents fatal error level
	FatalLevel

	// ErrorLevel represents error level
	ErrorLevel

	// WarnLevel represents warning level
	WarnLevel

	// InfoLevel represents info level
	InfoLevel

	// DebugLevel represents debug level
	DebugLevel

	// TraceLevel represents trace level (lowest)
	TraceLevel
)

// levelStrings provides fast array lookup for level string conversion
var levelStrings = []string{
	PanicLevel: "panic",
	FatalLevel: "fatal",
	ErrorLevel: "error",
	WarnLevel:  "warn",
	InfoLevel:  "info",
	DebugLevel: "debug",
	TraceLevel: "trace",
}

// String implements fmt.Stringer interface
// Optimized to use array lookup instead of switch for better performance
func (level Level) String() string {
	if level >= 0 && int(level) < len(levelStrings) {
		return levelStrings[level]
	}
	// Default fallback for unknown levels
	return "trace"
}

// MarshalText implements encoding.TextMarshaler interface for logrus compatibility
func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("trace"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		return []byte("warn"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}
	return nil, fmt.Errorf("invalid logrus level %d", level)
}
