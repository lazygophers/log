//go:build debug && discard

// 该文件仅在构建时同时满足 "debug" 和 "discard" 构建标签时才会被编译。
// 这是 Go 语言条件编译的一种方式，用于在特定场景下启用或禁用某些代码。

// Package log 提供在调试环境下使用的日志丢弃器。
//
// 当在调试模式并启用discard时，所有调试日志将被静默丢弃，而不是打印输出。
// 这在日志输出会带来显著性能开销或干扰的性能关键调试场景中非常有用。
package log

import (
	"fmt"
	"io"
)

// init 函数在包初始化时自动执行。
// 它将日志输出设置为 io.Discard，从而有效地丢弃所有日志消息。
func init() {
	SetOutput(io.Discard)
}

// Debug 记录一条调试日志，该日志在此模式下会被丢弃。
// 它通过调用标准记录器的 Debug 方法来实现，但由于输出被设置为 io.Discard，
// 因此不会产生任何实际输出。
//
// 参数:
//
//	args: 要记录的值（行为与fmt.Print相同）
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf 记录一条格式化的调试日志，该日志在此模式下会被丢弃。
// 它首先使用 fmt.Sprintf 格式化日志消息，然后调用标准记录器的 Debug 方法。
// 同样地，由于输出被设置为 io.Discard，日志消息不会被打印。
//
// 参数:
//
//	format: 格式化字符串（与fmt.Sprintf相同）
//	args: 要插入到格式化字符串中的值
func Debugf(format string, args ...interface{}) {
	std.Debug(fmt.Sprintf(format, args...))
}
