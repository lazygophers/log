//go:build release && discard

// 此文件仅在 release 和 discard 模式同时启用时生效。

// 本文件实现发布版本中的日志丢弃模式，
// 旨在通过完全禁止日志输出提升运行时性能。
// 在 release 和 discard 模式同时启用时，
// 所有日志调用将被重定向到空设备，
// 避免任何日志处理开销。
package log

import "io"

// init 初始化日志配置
// 功能：全局重定向日志输出到 io.Discard
// 效果：完全抑制日志输出以最大化运行时性能
func init() {
	SetOutput(io.Discard)
}

// Debug 空实现函数（丢弃模式）
// 设计意图：避免不必要的日志参数处理开销
func Debug(args ...interface{}) {
}

// Debugf 空实现函数（丢弃模式）
// 设计意图：避免不必要的日志格式化处理开销
func Debugf(format string, args ...interface{}) {
}
