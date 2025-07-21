// Package log 提供核心日志接口
//
// 包含标准日志级别函数和格式化版本，所有函数都是线程安全的
package log

import "fmt"

// Trace 记录跟踪级别日志（最详细级别）。
// 参数: args - 可变参数，日志内容。
// 使用场景: 开发阶段的详细流程追踪。
func Trace(args ...interface{}) {
	std.Trace(args...)
}

// Log 记录指定级别的日志。
// 参数: level - 日志级别, args - 可变参数，日志内容。
func Log(level Level, args ...interface{}) {
	std.Log(level, args...)
}

// Info 记录信息级别日志（常规操作信息）。
// 参数: args - 可变参数，日志内容。
func Info(args ...interface{}) {
	std.Info(args...)
}

// Warn 记录警告级别日志（非错误异常）。
// 参数: args - 可变参数，日志内容。
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Error 记录错误级别日志（可恢复错误）。
// 参数: args - 可变参数，日志内容。
func Error(args ...interface{}) {
	std.Error(args...)
}

// Panic 记录致命错误日志并触发panic。
// 参数: args - 可变参数，日志内容。
// 注意: 记录日志后会立即触发panic。
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal 记录致命错误日志并退出程序。
// 参数: args - 可变参数，日志内容。
// 注意: 记录日志后会调用os.Exit(1)。
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Logf 记录指定级别的格式化日志。
// 参数: level - 日志级别, format - 格式字符串, args - 格式化参数。
func Logf(level Level, format string, args ...interface{}) {
	std.Logf(level, format, args...)
}

// Infof 记录信息级别格式化日志。
// 参数: format - 格式字符串, args - 格式化参数。
func Infof(format string, args ...interface{}) {
	std.Info(fmt.Sprintf(format, args...))
}

// Warnf 记录警告级别格式化日志。
// 参数: format - 格式字符串, args - 格式化参数。
func Warnf(format string, args ...interface{}) {
	std.Warn(fmt.Sprintf(format, args...))
}

// Errorf 记录错误级别格式化日志。
// 参数: format - 格式字符串, args - 格式化参数。
func Errorf(format string, args ...interface{}) {
	std.Error(fmt.Sprintf(format, args...))
}

// Panicf 记录致命错误格式化日志并触发panic。
// 参数: format - 格式字符串, args - 格式化参数。
func Panicf(format string, args ...interface{}) {
	std.Panic(fmt.Sprintf(format, args...))
}

// Fatalf 记录致命错误格式化日志并退出程序。
// 参数: format - 格式字符串, args - 格式化参数。
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// StartMsg 输出启动信息。
// 特殊: 通常在应用启动时调用，输出初始化完成信息。
func StartMsg() {
	std.StartMsg()
}
