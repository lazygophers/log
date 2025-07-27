package log

import (
	"fmt"
)

// Level 表示日志级别，与 logrus 库兼容。
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

// String 实现了 fmt.Stringer 接口，返回日志级别的小写字符串表示。
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
		// 默认返回 "trace"，确保在遇到未知或未定义的级别时有一个安全的回退值。
		return "trace"
	}
}

// MarshalText 实现了 encoding.TextMarshaler 接口，
// 用于将日志级别序列化为文本格式（例如，在 JSON 编码中），以兼容 logrus。
func (level Level) MarshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("trace"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		// 注意：在文本序列化时，WarnLevel 使用 "warning" 而非 "warn"，以保持与 logrus 的兼容性。
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
