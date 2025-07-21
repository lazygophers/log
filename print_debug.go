//go:build debug && !discard

// 当启用 debug 模式时（且未启用 discard），此文件提供调试日志功能。
// debug 模式下的日志输出特性：
//   - 会输出完整的调用栈信息。
//   - 包含源代码文件名和行号。
//   - 输出更详细的变量值信息。
//   - 启用高性能日志采集（避免影响主流程性能）。

package log

import "fmt"

// Debug 使用默认日志器记录调试日志。
// 参数 args: 要记录的日志内容，支持多个参数。
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf 使用默认日志器记录格式化调试日志
// 参数 format: 格式化字符串
// 参数 args: 格式化字符串的参数
func Debugf(format string, args ...interface{}) {
	std.Debug(fmt.Sprintf(format, args...))
}
