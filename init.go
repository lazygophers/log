package log

import (
	"os"
)

var (
	std = newLogger()

	pid = os.Getpid()
)

// Pid 返回当前进程ID
func Pid() int {
	return pid
}

// New 创建新的日志记录器实例
func New() *Logger {
	return newLogger()
}

// SetLevel 设置日志记录级别
func SetLevel(level Level) *Logger {
	return std.SetLevel(level)
}

// GetLevel 获取当前日志记录级别
func GetLevel() Level {
	return std.Level()
}

// Sync 刷新日志缓冲区
func Sync() {
	std.Sync()
}

// Clone 克隆日志记录器
func Clone() *Logger {
	return std.Clone()
}

// CloneToCtx 克隆为带上下文的日志记录器
func CloneToCtx() *LoggerWithCtx {
	return std.CloneToCtx()
}

// SetCallerDepth 设置调用栈深度
func SetCallerDepth(callerDepth int) *Logger {
	return std.SetCallerDepth(callerDepth)
}

// SetPrefixMsg 设置日志前缀消息
func SetPrefixMsg(prefixMsg string) *Logger {
	return std.SetPrefixMsg(prefixMsg)
}

// AppendPrefixMsg 追加日志前缀消息
func AppendPrefixMsg(prefixMsg string) *Logger {
	return std.AppendPrefixMsg(prefixMsg)
}

// SetSuffixMsg 设置日志后缀消息
func SetSuffixMsg(suffixMsg string) *Logger {
	return std.SetSuffixMsg(suffixMsg)
}

// AppendSuffixMsg 追加日志后缀消息
func AppendSuffixMsg(suffixMsg string) *Logger {
	return std.AppendSuffixMsg(suffixMsg)
}

// ParsingAndEscaping 启用/禁用日志解析和转义
func ParsingAndEscaping(disable bool) *Logger {
	return std.ParsingAndEscaping(disable)
}

// Caller 启用/禁用调用者信息记录
func Caller(disable bool) *Logger {
	return std.Caller(disable)
}
