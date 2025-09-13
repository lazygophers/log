// Package log 提供了一个功能强大的日志记录工具，
// 支持多种输出、日志级别、自动轮转和清理，旨在为开发者提供开箱即用的日志解决方案。
package log

import (
	"io"
	"os"
	"path/filepath"
	"sync"
)

// SetOutput 为标准日志记录器（std）配置一个或多个输出目标。
//
// 这是一个便捷的函数，允许开发者快速将日志流重定向到不同的 `io.Writer`，
// 例如：控制台（os.Stdout）、文件、或网络连接。
//
// 使用示例：
//
//	// 同时输出到控制台和文件
//	fileWriter := log.GetOutputWriter("app.log")
//	log.SetOutput(os.Stdout, fileWriter)
//
// 参数：
//   - writes: 一个或多个实现了 `io.Writer` 接口的实例。
//
// 返回：
//   - 返回标准日志记录器 `*Logger` 的实例，以支持链式调用。
func SetOutput(writes ...io.Writer) *Logger {
	return std.SetOutput(writes...)
}

// GetOutputWriter 创建一个基础的文件日志输出器。
//
// 它会检查并创建日志文件所在的目录，并直接打开文件进行写入。
// 这是一个简单的文件写入器，不包含自动轮转和清理功能。
// 若需要更完善的日志轮转策略，请使用 `GetOutputWriterHourly`。
//
// 参数：
//   - filename: 日志文件的完整路径。
//
// 返回：
//   - 一个实现了 `io.Writer` 接口的日志输出器。
func GetOutputWriter(filename string) io.Writer {
	// 确保日志文件所在的目录存在
	ensureDir(filepath.Dir(filename))

	// 直接打开文件进行写入
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// 如果创建日志写入器失败，则触发 panic，因为这是关键功能
		std.Panicf("创建日志文件写入器 %s 失败: %v", filename, err)
	}
	return file
}

// ensureDir 确保指定的目录存在。
// 如果目录不存在，它会尝试创建该目录（包括所有必需的父目录）。
// 这是一个内部辅助函数，用于简化日志写入器初始化时的目录检查和创建逻辑。
//
// 参数：
//   - dir: 需要确保其存在的目录路径。
func ensureDir(dir string) {
	// 当目录路径不为 "." (当前目录) 且路径本身不是一个目录时，执行创建逻辑。
	// 这样做可以避免在不必要的情况下调用 os.MkdirAll，并处理路径已存在的情况。
	if dir != "." && !isDir(dir) {
		// 使用 os.ModePerm 权限创建目录。如果创建失败，记录一条错误日志。
		// 这里选择记录错误而不是 panic，是因为目录创建失败不应总是导致整个应用程序崩溃。
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			std.Errorf("无法创建日志目录 %s: %v", dir, err)
		}
	}
}

// isDir 检查给定路径是否为一个有效的目录。
//
// 参数：
//   - path: 需要检查的文件系统路径。
//
// 返回：
//   - 如果路径存在且为目录，则返回 `true`，否则返回 `false`。
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Output 接口定义了日志输出器的基本行为，它继承了 `io.Writer` 接口。
// 这样做是为了在类型系统中明确标识出用于日志输出的写入器。
type Output interface {
	io.Writer
}

var (
	// rotatorInstances 存储已创建的轮转器实例，避免重复创建
	rotatorInstances = make(map[string]*HourlyRotator)
	
	// rotatorMutex 保护对 rotatorInstances 的并发访问
	rotatorMutex = &sync.Mutex{}
)

// GetOutputWriterHourly 创建一个按小时轮转和自动清理的日志输出器。
//
// 这是推荐的生产环境日志文件配置，提供了完善的日志管理功能：
//   - **按小时轮转**：每小时创建一个新的日志文件，文件名格式为 `filenameYYYYMMDDHH.log`。
//   - **自动链接**：始终维护一个指向最新日志文件的软链接，方便实时查看。
//   - **大小限制**：当日志文件达到一定大小时（当前为 800MB），也会触发轮转。
//   - **历史保留**：最多保留最近 12 个小时的日志文件。
//   - **自动清理**：启动一个后台 goroutine，每 10 分钟检查并删除过期的日志文件。
//
// 参数：
//   - filename: 日志文件的基础路径和名称，例如 `/var/log/app/access`。
//
// 返回：
//   - 一个实现了 `Writer` 接口（即 `io.Writer`）的日志输出器。
func GetOutputWriterHourly(filename string) Writer {
	rotatorMutex.Lock()
	defer rotatorMutex.Unlock()
	
	// 检查是否已存在该文件的轮转器实例
	if rotator, exists := rotatorInstances[filename]; exists {
		return rotator
	}
	
	// 创建新的轮转器实例
	rotator := NewHourlyRotator(
		filename,
		1024*1024*8*100, // 800MB 大小限制
		12,               // 保留最近12个文件
	)
	
	// 存储实例供后续使用
	rotatorInstances[filename] = rotator
	
	return rotator
}
