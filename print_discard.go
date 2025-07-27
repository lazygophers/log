//go:build !release && !debug && discard

/*
Package log 提供 discard 模式实现：
该模式会静默丢弃所有日志输出，
适用于性能测试和基准测试场景，
在这些场景中需要最小化日志系统对性能的影响。

此文件通过编译标签 (build tag) "!release && !debug && discard" 控制，
确保只在独立的 "discard" 模式下生效。
*/
package log

import (
	"fmt"
	"io"
)

// init 函数在包初始化时，将标准日志记录器的输出设置为 io.Discard。
// io.Discard 是一个特殊的 io.Writer，它会消费并丢弃所有写入的数据，
// 从而实现零成本的日志操作。
func init() {
	SetOutput(io.Discard)
}

// Debug 在 discard 模式下是一个空操作 (no-op)。
// 该函数仅为保持 API 的完整性，实际不会产生任何日志输出。
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf 在 discard 模式下是一个空操作 (no-op)。
// 该函数仅为保持 API 的完整性，格式化操作会执行，但结果会被丢弃。
func Debugf(format string, args ...interface{}) {
	std.Debug(fmt.Sprintf(format, args...))
}
