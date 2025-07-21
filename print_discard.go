//go:build !release && !debug && discard

/*
Package log 提供 discard 模式实现：
该模式会静默丢弃所有日志输出，
适用于性能测试和基准测试场景，
在这些场景中需要最小化日志系统对性能的影响。
*/
package log

import (
	"fmt"
	"io"
)

// init 初始化日志输出为 io.Discard。
// io.Discard 是特殊的 io.Writer 实现，
// 会丢弃所有写入的数据。
// 此配置在包初始化时自动执行，
// 确保 discard 模式生效。
func init() {
	SetOutput(io.Discard)
}

// Debug 在 discard 模式下不会产生实际输出
// 该函数仅用于兼容日志接口，调用时不会执行任何操作
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf 在 discard 模式下不会产生实际输出
// 该函数仅用于兼容日志接口，调用时不会执行任何操作
func Debugf(format string, args ...interface{}) {
	std.Debug(fmt.Sprintf(format, args...))
}
