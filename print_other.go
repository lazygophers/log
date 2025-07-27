//go:build !debug && !release && !discard

// 构建约束：此文件仅在没有 "debug"、"release" 或 "discard" 构建标签时编译。
// 这意味着它在标准的开发和生产环境中生效，但在特定的调试或发布模式下被排除。
package log

import "fmt"

// Debug 记录调试级别的日志。
//
// 仅当全局日志级别设置为 DebugLevel 或更低时，才会实际输出日志。
func Debug(args ...interface{}) {
	// 检查当前日志级别，确保只有在需要时才记录调试信息，以提高性能。
	if std.level >= DebugLevel {
		std.Debug(args...)
	}
}

// Debugf 记录格式化的调试级别日志。
//
// 仅当全局日志级别设置为 DebugLevel 或更低时，才会实际输出日志。
func Debugf(format string, args ...interface{}) {
	// 检查当前日志级别，确保只有在需要时才格式化并记录调试信息。
	if std.level >= DebugLevel {
		std.Debug(fmt.Sprintf(format, args...))
	}
}
