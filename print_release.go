//go:build release && !discard

package log

import (
	"os"
	"path/filepath"
)

var ReleaseLogPath string = filepath.Join(os.TempDir(), "lazygophers", "log")

// init 初始化发布模式下的日志输出。
func init() {
	SetOutput(GetOutputWriterHourly(ReleaseLogPath))
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
