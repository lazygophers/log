//go:build debug && !discard

// 当启用 debug 模式时（且未启用 discard），此文件提供调试日志功能。
// debug 模式下的日志输出特性：
//   - 会输出完整的调用栈信息。
//   - 包含源代码文件名和行号。
//   - 输出更详细的变量值信息。
//   - 启用高性能日志采集（避免影响主流程性能）。

package log

import "fmt"

// Debug 使用默认日志器（std）记录调试级别的日志。
// 日志内容由多个参数 `args` 拼接而成。
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf 使用默认日志器（std）记录格式化的调试级别日志。
// `format` 是格式化模板，`args` 是对应的参数。
func Debugf(format string, args ...interface{}) {
	std.Debug(fmt.Sprintf(format, args...))
}
