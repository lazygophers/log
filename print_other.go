//go:build !debug && !release && !discard

// 构建约束：仅在非debug、非release和非discard模式下生效

// 本文件提供标准日志级别下的调试日志功能
package log

import "fmt"

// Debug 条件性记录调试日志
// 功能：当日志级别≥DebugLevel时记录提供的日志内容
// 参数：
//
//	args: 可变参数，接收任意类型的日志内容
func Debug(args ...interface{}) {
	if std.level >= DebugLevel {
		std.Debug(args...)
	}
}

// Debugf 条件性记录格式化调试日志
// 功能：当日志级别≥DebugLevel时记录格式化日志
// 参数：
//
//	format: 格式字符串
//	args: 格式化参数
func Debugf(format string, args ...interface{}) {
	if std.level >= DebugLevel {
		std.Debug(fmt.Sprintf(format, args...))
	}
}
