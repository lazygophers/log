//go:build release && !discard

// 该构建标签确保此文件仅在满足以下两个条件时才被编译：
// 1. 构建时指定了 "release" 标签。
// 2. 构建时未指定 "discard" 标签。
// 这是 Go 语言提供的条件编译机制，用于在不同构建环境中使用不同的代码实现。

// Package log 提供发布模式下的日志功能（非丢弃模式）
//
// 此文件配置日志输出为按小时分割的日志文件，存储在系统的临时目录下
package log

import (
	"os"
	"path/filepath"
)

// ReleaseLogPath specifies the file path for logs in release mode.
// Default uses hourly rotation in system temp directory like GetOutputWriterHourly.
// This can be overridden at compile time using build flags:
// go build -ldflags "-X github.com/lazygophers/log.ReleaseLogPath=/path/to/log/file" -tags release
var ReleaseLogPath string = filepath.Join(os.TempDir(), "lazygophers", "log")

// init 初始化发布模式下的日志输出。
//
// 功能：
//   - 设置日志输出为按小时分割的文件。
//   - 文件存储在临时目录的lazygophers/log子目录下。
//   - 示例路径: /tmp/lazygophers/log/2023071015.log，
//     每小时生成新文件。
func init() {
	SetOutput(GetOutputWriterHourly(ReleaseLogPath))
}

// Debug 记录调试级别的日志。
//
// 仅当全局日志级别设置为 DebugLevel 或更低时，才会实际输出日志。
func Debug(args ...interface{}) {
	// 检查当前日志级别，确保只有在需要时才记录调试信息，以提高性能。
	if std.level >= DebugLevel {
		std.Debug(args...)
	}
}

// Debugf 记录格式化的调试级别日志。
//
// 仅当全局日志级别设置为 DebugLevel 或更低时，才会实际输出日志。
func Debugf(format string, args ...interface{}) {
	// 检查当前日志级别，确保只有在需要时才格式化并记录调试信息。
	if std.level >= DebugLevel {
		std.Debug(fmt.Sprintf(format, args...))
	}
}
