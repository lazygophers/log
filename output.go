// Package log 提供了一个功能强大的日志记录工具，
// 支持多种输出、日志级别、自动轮转和清理，旨在为开发者提供开箱即用的日志解决方案。
package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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
// 它会检查并创建日志文件所在的目录，并使用 `file-rotatelogs`
// 创建一个简单的日志写入器，但**不**包含自动轮转和清理功能。
// 若需要更完善的日志轮转策略，请使用 `GetOutputWriterHourly`。
//
// 参数：
//   - filename: 日志文件的完整路径。
//
// 返回：
//   - 一个实现了 `io.Writer` 接口的日志输出器。
func GetOutputWriter(filename string) io.Writer {
	// 确保日志文件所在的目录存在
	dir := filepath.Dir(filename)
	if dir != "." && !isDir(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			// 在创建目录失败时记录错误，但不中断程序
			Errorf("无法创建日志目录 %s: %v", dir, err)
		}
	}

	// 初始化 rotatelogs，即使不轮转，也用它来管理文件写入
	hook, err := rotatelogs.New(filename)
	if err != nil {
		// 如果创建日志写入器失败，则触发 panic，因为这是关键功能
		std.Panicf("创建日志文件写入器 %s 失败: %v", filename, err)
	}
	return hook
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
	// cleanRotatelogOnce 用于确保对同名日志文件的清理任务（goroutine）只启动一次。
	// 这是一个简单的并发控制机制，防止因重复调用 `GetOutputWriterHourly` 而创建多个清理协程。
	// key: 日志文件的基础名称
	// value: bool，标记是否已启动清理协程
	cleanRotatelogOnce = make(map[string]bool)
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
	// 确保日志文件所在的目录存在
	dir := filepath.Dir(filename)
	if dir != "." && !isDir(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			Errorf("无法创建日志目录 %s: %v", dir, err)
		}
	}

	// 配置 rotatelogs 实现日志轮转
	hook, err := rotatelogs.New(
		filename+"%Y%m%d%H.log",                      // 日志文件命名模式
		rotatelogs.WithLinkName(filename+".log"),     // 创建一个指向最新日志的软链接
		rotatelogs.WithRotationSize(1024*1024*8*100), // 按大小轮转（800MB）
		rotatelogs.WithRotationTime(time.Hour),       // 按小时轮转
		rotatelogs.WithRotationCount(12),             // 保留最近12个文件
	)
	if err != nil {
		std.Panicf("创建按小时轮转的日志写入器 %s 失败: %v", filename, err)
	}

	// 确保对每个文件只启动一个清理协程
	if _, ok := cleanRotatelogOnce[filename]; !ok {
		go func() {
			// 无限循环，定期执行清理任务
			for {
				// 读取日志目录下的所有文件
				files, err := os.ReadDir(filepath.Dir(filename))
				if err != nil {
					// 使用 fmt.Printf 以避免日志系统自身的循环依赖问题
					fmt.Printf("读取日志目录 %s 用于清理时出错: %v\n", filepath.Dir(filename), err)
					continue
				}

				// 筛选出文件名
				var filesOnly []string
				for _, file := range files {
					if !file.IsDir() {
						filesOnly = append(filesOnly, file.Name())
					}
				}

				// 按文件名降序排序，最新的文件排在前面
				sort.Slice(filesOnly, func(i, j int) bool {
					return filesOnly[i] > filesOnly[j]
				})

				// 清理超过保留数量的旧日志
				for i, s := range filesOnly {
					// 跳过需要保留的最新文件和软链接
					if i < 12 || s == filepath.Base(filename)+".log" {
						continue
					}

					// 构造待删除文件的完整路径并执行删除
					fullPath := filepath.Join(filepath.Dir(filename), s)
					fmt.Printf("正在删除旧的日志文件: %s\n", fullPath)
					if err = os.Remove(fullPath); err != nil {
						Errorf("删除旧日志文件 %s 失败: %v", fullPath, err)
					}
				}

				// 等待10分钟后再次执行清理
				time.Sleep(time.Minute * 10)
			}
		}()
		// 标记该文件的清理协程已启动
		cleanRotatelogOnce[filename] = true
	}

	return hook
}
