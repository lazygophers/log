// Package log 提供了一个功能强大的日志记录工具，
// 支持多种输出、日志级别、自动轮转和清理，旨在为开发者提供开箱即用的日志解决方案。
package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"sync"
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
	ensureDir(filepath.Dir(filename))

	// 初始化 rotatelogs，即使不轮转，也用它来管理文件写入
	hook, err := rotatelogs.New(filename)
	if err != nil {
		// 如果创建日志写入器失败，则触发 panic，因为这是关键功能
		std.Panicf("创建日志文件写入器 %s 失败: %v", filename, err)
	}
	return hook
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
			Errorf("无法创建日志目录 %s: %v", dir, err)
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
	// cleanRotatelogOnce 用于确保对同名日志文件的清理任务（goroutine）只启动一次。
	// 这是一个简单的并发控制机制，防止因重复调用 `GetOutputWriterHourly` 而创建多个清理协程。
	// key: 日志文件的基础名称
	// value: bool，标记是否已启动清理协程
	cleanRotatelogOnce = make(map[string]bool)

	// cleanMutex 是一个互斥锁，用于保护对 `cleanRotatelogOnce` 的并发访问。
	// 在高并发场景下，多个 goroutine 可能会同时调用 `GetOutputWriterHourly`，
	// 如果没有锁，可能会导致对同一个日志文件启动多个清理协程，造成资源浪费和潜在的竞争问题。
	// 使用此锁可确保 "只启动一次" 的逻辑在并发环境下是安全的。
	cleanMutex = &sync.Mutex{}
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
	ensureDir(filepath.Dir(filename))

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

	// 通过加锁来确保对 `cleanRotatelogOnce` 的检查和写入是原子操作，防止并发场景下的 race condition。
	cleanMutex.Lock()
	// 检查是否已经为该日志文件启动了清理协程
	if _, ok := cleanRotatelogOnce[filename]; !ok {
		// 如果尚未启动，则标记为已启动，并释放锁，然后启动一个新的 goroutine 执行清理任务。
		// 必须在启动 goroutine 之前释放锁，以避免潜在的死锁或长时间的锁持有。
		cleanRotatelogOnce[filename] = true
		cleanMutex.Unlock()

		go func() {
			// 后台清理协程：这是一个无限循环，旨在定期清理旧的日志文件，以防磁盘空间被占满。
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
	} else {
		// 如果清理协程已经启动，则直接释放锁。
		cleanMutex.Unlock()
	}

	return hook
}
