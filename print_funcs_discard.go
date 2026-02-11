//go:build discard

package log

import "io"

// init 函数在包初始化时，将标准日志记录器的输出设置为 io.Discard。
// io.Discard 是一个特殊的 io.Writer，它会消费并丢弃所有写入的数据，
// 从而实现零成本的日志操作。
func init() {
	SetOutput(io.Discard)
}

// Debug 记录调试级别的日志。
//
// 仅当全局日志级别设置为 DebugLevel 或更低时，才会实际输出日志。
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf 记录格式化的调试级别日志。
//
// 仅当全局日志级别设置为 DebugLevel 或更低时，才会实际输出日志。
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
