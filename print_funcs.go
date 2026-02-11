//go:build !discard && (debug || canary || (!debug && !release && !canary))

package log

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
