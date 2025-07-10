// Package log 提供日志初始化和管理功能
package log

import (
	"os"
)

var (
	// std 是标准日志记录器实例
	std = newLogger()

	// pid 存储当前进程ID
	pid = os.Getpid()
)

// Pid 返回当前进程ID
func Pid() int {
	return pid
}

// New 创建并返回一个新的Logger实例
func New() *Logger {
	return newLogger()
}

// SetLevel 设置标准日志记录器的日志级别
// 参数:
//
//	level: 要设置的日志级别
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetLevel(level Level) *Logger {
	return std.SetLevel(level)
}

// GetLevel 返回标准日志记录器的当前日志级别
// 返回值:
//
//	Level: 当前日志级别
func GetLevel() Level {
	return std.Level()
}

// Sync 刷新标准日志记录器的缓冲区
func Sync() {
	std.Sync()
}

// Clone 创建标准日志记录器的副本
// 返回值:
//
//	*Logger: 新的日志记录器实例
func Clone() *Logger {
	return std.Clone()
}

// CloneToCtx 创建标准日志记录器的副本并返回LoggerWithCtx实例
// 返回值:
//
//	*LoggerWithCtx: 新的上下文日志记录器实例
func CloneToCtx() *LoggerWithCtx {
	return std.CloneToCtx()
}

// SetCallerDepth 设置日志调用堆栈深度
// 参数:
//
//	callerDepth: 调用堆栈深度
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetCallerDepth(callerDepth int) *Logger {
	return std.SetCallerDepth(callerDepth)
}

// SetPrefixMsg 设置日志消息前缀
// 参数:
//
//	prefixMsg: 前缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetPrefixMsg(prefixMsg string) *Logger {
	return std.SetPrefixMsg(prefixMsg)
}

// AppendPrefixMsg 追加日志消息前缀
// 参数:
//
//	prefixMsg: 要追加的前缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func AppendPrefixMsg(prefixMsg string) *Logger {
	return std.AppendPrefixMsg(prefixMsg)
}

// SetSuffixMsg 设置日志消息后缀
// 参数:
//
//	suffixMsg: 后缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetSuffixMsg(suffixMsg string) *Logger {
	return std.SetSuffixMsg(suffixMsg)
}

// AppendSuffixMsg 追加日志消息后缀
// 参数:
//
//	suffixMsg: 要追加的后缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func AppendSuffixMsg(suffixMsg string) *Logger {
	return std.AppendSuffixMsg(suffixMsg)
}

// ParsingAndEscaping 启用/禁用日志消息的解析和转义
// 参数:
//
//	disable: true表示禁用，false表示启用
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func ParsingAndEscaping(disable bool) *Logger {
	return std.ParsingAndEscaping(disable)
}

// Caller 启用/禁用调用者信息（文件名和行号）的显示
// 参数:
//
//	disable: true表示禁用，false表示启用
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func Caller(disable bool) *Logger {
	return std.Caller(disable)
}
