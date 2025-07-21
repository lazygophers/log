package log

import (
	"fmt"
)

// Level 表示日志级别
type Level uint32

const (
	// PanicLevel 表示 panic 级别（最高级别）
	PanicLevel Level = iota

	// FatalLevel 表示致命错误级别
	FatalLevel

	// ErrorLevel 表示错误级别
	ErrorLevel

	// WarnLevel 表示警告级别
	WarnLevel

	// InfoLevel 表示信息级别
	InfoLevel

	// DebugLevel 表示调试级别
	DebugLevel

	// TraceLevel 表示跟踪级别（最低级别）
	TraceLevel
)

// String 返回日志级别的字符串表示。
//
// 参数:
//
//	level: 要转换的日志级别
//
// 返回值:
//
//	日志级别的字符串名称
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
		return "trace"
	}
}

// MarshalText 将日志级别编码为文本形式。
//
// 参数:
//
//	level: 要编码的日志级别
//
// 返回值:
//
//	[]byte: 编码后的字节切片
//	error: 编码错误（如果级别无效）
func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("trace"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		return []byte("warning"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}

	return nil, fmt.Errorf("not a valid logrus level %d", level)
}
