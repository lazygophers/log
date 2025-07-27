# 🐰 lazygophers/log

<p align="center">
  <strong>一个为追求极致简洁与扩展性而生的 Go 日志库。</strong>
</p>

<p align="center">
  <a href="https://github.com/lazygophers/log/actions/workflows/go.yml"><img src="https://github.com/lazygophers/log/actions/workflows/go.yml/badge.svg" alt="Build Status"></a>
  <a href="https://codecov.io/gh/lazygophers/log"><img src="https://codecov.io/gh/lazygophers/log/branch/main/graph/badge.svg" alt="codecov"></a>
  <a href="https://goreportcard.com/report/github.com/lazygophers/log"><img src="https://goreportcard.com/badge/github.com/lazygophers/log" alt="Go Report Card"></a>
  <a href="https://godoc.org/github.com/lazygophers/log"><img src="https://godoc.org/github.com/lazygophers/log?status.svg" alt="GoDoc"></a>
  <a href="https://github.com/lazygophers/log/releases"><img src="https://img.shields.io/github/release/lazygophers/log.svg" alt="GitHub release"></a>
  <a href="https://opensource.org/licenses/Apache-2.0"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="License"></a>
</p>

---

**`lazygophers/log`** 提供了一套优雅、直观的 API，摒弃了繁杂的配置，让您专注于业务逻辑本身。通过实现 `io.Writer` 和 `Format` 接口，您可以随心所欲地定制日志的输出目标与展现形式，无论是写入文件、发送到远程服务，还是集成到您自己的监控系统，都游刃有余。

## ✨ 功能特性

- **多日志级别**: `Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`, `Panic`
- **灵活输出目标**: 支持同时向多个 `io.Writer` 输出。
- **自定义格式**: 通过实现 `Format` 接口，轻松定制 JSON、Logfmt 等任意格式。
- **调用栈追踪**: 精准定位日志来源（文件、行号、函数名）。
- **协程安全**: 所有方法均为协程安全。
- **性能导向**: 清晰的性能优化路线图，致力于零内存分配。

## 🚀 性能与路线图

我们始终致力于将 `lazygophers/log` 打造成一款高性能、功能丰富的日志库。通过最近的基准测试，我们对当前版本的性能有了清晰的认识，并制定了明确的优化路线图。

### 基准测试

以下是当前版本与业界领先的 `zap.SugaredLogger` 在相同条件下的性能对比。

- **CPU**: Apple M3
- **Go**: 1.23.0

| 测试场景 (Benchmark)          | `lazygophers/log` (耗时/操作) | `zap.SugaredLogger` (耗时/操作) | 内存分配 (allocs/op) |
| :---------------------------- | :------------------------------ | :--------------------------------- | :--------------------- |
| **简单日志** (Simple)         | 944.6 ns/op                     | 207.8 ns/op                        | 9 allocs/op            |
| **带5个字段** (With 5 Fields) | 1343 ns/op                      | 746.3 ns/op                        | 13 allocs/op           |
| **格式化日志** (Infof)        | 1002 ns/op                      | 301.4 ns/op                        | 9 allocs/op            |

### 性能分析与未来计划

当前版本在性能上与 `zap` 存在一定差距，主要瓶颈在于**内存分配**。性能优化将是下一阶段的核心任务。

- [ ] **v0.2.0 - 内存优化**: 全面引入 `sync.Pool` 对象池技术，复用日志对象，显著减少高频日志场景下的内存分配和GC压力。
- [ ] **v0.3.0 - API 优化**: 引入 `With` 方法，支持结构化上下文日志，进一步提升字段处理效率。
- [ ] **长期 - 零分配探索**: 探索在特定场景下实现零内存分配的可能性。

我们相信，通过持续的迭代和优化，`lazygophers/log` 将在性能上达到一流水准。欢迎您关注我们的进展，也欢迎随时提出宝贵的建议！

## 📦 安装

```bash
go get github.com/lazygophers/log
```

## 快速开始

```go
package main

import "github.com/lazygophers/log"

func main() {
    // 默认级别为 Info
    log.Info("Application started")
    log.Debug("This is a debug message") // 这条日志不会被输出

    log.SetLevel(log.DebugLevel)
    log.Debug("Now, this debug message will be visible.")

    log.Warnf("User %s might not have permission", "Alice")
    log.Error("Failed to connect to the database")
}
```

## 🔧 高级用法

### 1. 多目标输出

您可以轻松地将日志同时输出到多个目标，例如，同时在控制台显示 `INFO` 级别以上的日志，并将所有日志（包括 `DEBUG`）保存到文件。

```go
package main

import (
	"os"
	"github.com/lazygophers/log"
)

func main() {
	logFile, _ := os.Create("app.log")
	defer logFile.Close()

	// 同时输出到控制台和文件
	log.SetOutput(os.Stdout, logFile)
	log.SetLevel(log.DebugLevel)

	log.Info("This message appears on both stdout and in app.log.")
	log.Debug("This message only appears in app.log.")
}
```

### 2. 自定义日志格式

通过实现 `Format` 接口，您可以创建自己的日志格式，比如输出结构化的 JSON。

```go
package main

import (
	"encoding/json"
	"os"
	"github.com/lazygophers/log"
)

// JSONFormatter 实现了 log.Format 接口
type JSONFormatter struct{}

// Format 将日志条目格式化为 JSON
func (f *JSONFormatter) Format(entry *log.Entry) []byte {
	data := map[string]interface{}{
		"level":   entry.Level.String(),
		"time":    entry.Time,
		"message": entry.Message,
		"caller":  entry.CallerName,
	}
	serialized, _ := json.Marshal(data)
	return append(serialized, '\n')
}

func main() {
	jsonLogger := log.New()
	jsonLogger.SetFormatter(&JSONFormatter{})
	jsonLogger.SetOutput(os.Stdout)

	jsonLogger.Info("This is a JSON formatted log.")
	jsonLogger.Warnf("User %s login failed.", "admin")
}
```

### 3. 按时间/大小轮转日志

使用内置的 `GetOutputWriterHourly` 可以方便地实现日志文件的按小时轮转和自动清理。

```go
package main

import (
	"time"
	"github.com/lazygophers/log"
)

func main() {
    // 日志将写入 logs/access-YYYY-MM-DD-HH.log，并按小时分割
	fileWriter := log.GetOutputWriterHourly("logs/access")
	log.SetOutput(fileWriter)

	log.Info("Service started, access log recording.")
}
```

### 4. 独立的日志实例

通过 `Clone()` 或 `New()`，您可以为应用的不同模块创建独立的 `Logger` 实例，它们可以拥有完全不同的配置（级别、输出、格式等），互不干扰。

```go
package main

import (
	"os"
	"github.com/lazygophers/log"
)

func main() {
	// 全局 logger
	log.Info("Message from the global logger.")

	// 为数据库模块创建一个独立的 logger
	dbLogger := log.New()
	dbLogger.SetPrefix("[Database]")

	// 为 HTTP 模块克隆并配置一个新的 logger
	httpLogger := dbLogger.Clone()
	httpLogger.SetPrefix("[HTTP]")

	dbLogger.Debug("Connecting to the database...")
	httpLogger.Info("New request received: /api/users")
}
```

## 📖 API 参考

### 主要方法

| 方法                     | 描述                                     |
| ------------------------ | ---------------------------------------- |
| `SetLevel(level Level)`    | 设置日志级别。                           |
| `SetOutput(w ...io.Writer)` | 设置一个或多个输出目标。                 |
| `SetFormatter(f Format)` | 设置自定义的日志格式化器。             |
| `SetCallerDepth(d int)`  | 调整调用栈深度，用于封装场景。       |
| `SetPrefix(p string)`      | 为每条日志添加统一前缀。                 |
| `New() *Logger`            | 创建一个全新的、独立的 Logger 实例。     |
| `Clone() *Logger`          | 创建一个继承当前配置的 Logger 副本。     |

### 日志级别方法

- `Trace(v ...any)` / `Tracef(format string, v ...any)`
- `Debug(v ...any)` / `Debugf(format string, v ...any)`
- `Info(v ...any)` / `Infof(format string, v ...any)`
- `Warn(v ...any)` / `Warnf(format string, v ...any)`
- `Error(v ...any)` / `Errorf(format string, v ...any)`
- `Fatal(v ...any)` / `Fatalf(format string, v ...any)` (执行后调用 `os.Exit(1)`)
- `Panic(v ...any)` / `Panicf(format string, v ...any)` (执行后调用 `panic()`)

## ❤️ 贡献指南

欢迎各种形式的贡献！无论是提交 Issue、发起 Pull Request，还是改进文档，我们都非常欢迎。

1.  **Fork** 本仓库
2.  创建您的特性分支 (`git checkout -b feature/your-amazing-feature`)
3.  提交您的更改 (`git commit -am 'Add some amazing feature'`)
4.  推送到分支 (`git push origin feature/your-amazing-feature`)
5.  创建一个 **Pull Request**

## 📜 许可证

本项目采用 [Apache 2.0 License](LICENSE) 授权。
