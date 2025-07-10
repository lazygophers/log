//go:build debug && discard

// Package log 提供在调试环境下使用的日志丢弃器。
//
// 当在调试模式并启用discard时，所有调试日志将被静默丢弃，而不是打印输出。
// 这在日志输出会带来显著性能开销或干扰的性能关键调试场景中非常有用。
package log

import (
	"fmt"
	"io"
)

func init() {
	SetOutput(io.Discard)
}

// Debug 记录一条调试日志，该日志在此模式下会被丢弃。
//
// 参数:
//
//	args: 要记录的值（行为与fmt.Print相同）
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf 记录一条格式化的调试日志，该日志在此模式下会被丢弃。
//
// 参数:
//
//	format: 格式化字符串（与fmt.Sprintf相同）
//	args: 要插入到格式化字符串中的值
func Debugf(format string, args ...interface{}) {
	std.Debug(fmt.Sprintf(format, args...))
}
