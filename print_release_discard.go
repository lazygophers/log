//go:build release && discard

// Package log 在 release 和 discard 构建标记同时启用时，提供一个日志丢弃的实现。
// 在此模式下，所有日志输出都将被重定向到 io.Discard，以实现零日志开销，从而最大化运行时性能。
package log

import "io"

// init 函数在包初始化时执行。
// 它通过调用 SetOutput(io.Discard) 将全局日志输出重定向到空设备，
// 从而有效地禁用所有级别的日志记录。
func init() {
	SetOutput(io.Discard)
}

// Debug 在丢弃模式下是一个空操作。
// 调用此函数不会产生任何日志输出，也几乎没有性能开销。
// 这样设计是为了在不修改业务逻辑代码的情况下，通过构建标记来移除Debug日志。
func Debug(args ...interface{}) {
}

// Debugf 在丢弃模式下是一个空操作。
// 调用此函数不会产生任何日志输出，也不会进行格式化操作，性能开销极低。
// 这样设计是为了在不修改业务逻辑代码的情况下，通过构建标记来移除Debugf日志。
func Debugf(format string, args ...interface{}) {
}
