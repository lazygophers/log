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

// String implements fmt.Stringer interface
func (level Level) String() string {
	switch level {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	default:
		// Default fallback for unknown levels
		return "trace"
	}
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
		// Use "warning" instead of "warn" for logrus compatibility
		return []byte("warning"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}
	return nil, fmt.Errorf("invalid logrus level %d", level)
}
