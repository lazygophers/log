package log

import (
	"fmt"
)

// Level 定义日志级别类型
type Level uint32

const (
	PanicLevel Level = iota // 最高级别，触发 panic
	FatalLevel              // 致命错误，程序退出
	ErrorLevel              // 错误级别
	WarnLevel               // 警告级别
	InfoLevel               // 信息级别
	DebugLevel              // 调试级别
	TraceLevel              // 跟踪级别，最详细
)

// String 返回日志级别的字符串表示
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

// MarshalText 实现日志级别的文本序列化
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
