# 🚀 LazyGophers Log

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.9%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A high-performance, feature-rich logging library for Go applications with multi-build tag support, async writing, and extensive customization options.

## 📖 Documentation Languages

- [🇺🇸 English](README.md) (Current)
- [🇨🇳 简体中文](docs/README_zh-CN.md)
- [🇹🇼 繁體中文](docs/README_zh-TW.md)
- [🇫🇷 Français](docs/README_fr.md)
- [🇷🇺 Русский](docs/README_ru.md)
- [🇪🇸 Español](docs/README_es.md)
- [🇸🇦 العربية](docs/README_ar.md)

## ✨ Features

- **🚀 High Performance**: Object pooling and async writing support
- **🏗️ Build Tag Support**: Different behaviors for debug, release, and discard modes
- **🔄 Log Rotation**: Automatic hourly log file rotation
- **🎨 Rich Formatting**: Customizable log formats with color support
- **🔍 Context Tracing**: Goroutine ID and trace ID tracking
- **🔌 Framework Integration**: Native Zap logger integration
- **⚙️ Highly Configurable**: Flexible levels, outputs, and formatting
- **🧪 Well Tested**: 93.9% test coverage with 284+ test cases across all build configurations

## 🚀 Quick Start

### Installation

```bash
go get github.com/lazygophers/log
```

### Basic Usage

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Simple logging
    log.Info("Hello, World!")
    log.Debug("This is a debug message")
    log.Warn("This is a warning")
    log.Error("This is an error")

    // Formatted logging
    log.Infof("User %s logged in with ID %d", "john", 123)
    
    // With custom logger
    logger := log.New()
    logger.SetLevel(log.InfoLevel)
    logger.Info("Custom logger message")
}
```

### Advanced Usage

```go
package main

import (
    "context"
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Create logger with file output
    logger := log.New()
    
    // Set output to file with hourly rotation
    writer := log.GetOutputWriterHourly("./logs/app.log")
    logger.SetOutput(writer)
    
    // Configure formatting
    logger.SetLevel(log.DebugLevel)
    logger.SetPrefixMsg("[APP] ")
    logger.Caller(true) // Enable caller info
    
    // Context logging
    ctxLogger := logger.CloneToCtx()
    ctxLogger.Info(context.Background(), "Context-aware logging")
    
    // Async logging for high-throughput scenarios
    asyncWriter := log.NewAsyncWriter(writer, 1000)
    logger.SetOutput(asyncWriter)
    defer asyncWriter.Close()
    
    logger.Info("High-performance async logging")
}
```

## 🏗️ Build Tags

The library supports different build modes through Go build tags:

### Default Mode (No Tags)
```bash
go build
```
- Full logging functionality
- Debug messages enabled
- Standard performance

### Debug Mode
```bash
go build -tags debug
```
- Enhanced debug information
- Detailed caller information
- Performance profiling support

### Release Mode
```bash
go build -tags release
```
- Optimized for production
- Debug messages disabled
- Automatic log file rotation

### Discard Mode
```bash
go build -tags discard
```
- Maximum performance
- All logs discarded
- Zero logging overhead

### Combined Modes
```bash
go build -tags "debug,discard"    # Debug with discard
go build -tags "release,discard"  # Release with discard
```

## 📊 Log Levels

The library supports 7 log levels (from highest to lowest priority):

| Level | Value | Description |
|-------|-------|-------------|
| `PanicLevel` | 0 | Logs and then calls panic |
| `FatalLevel` | 1 | Logs and then calls os.Exit(1) |
| `ErrorLevel` | 2 | Error conditions |
| `WarnLevel` | 3 | Warning conditions |
| `InfoLevel` | 4 | Informational messages |
| `DebugLevel` | 5 | Debug-level messages |
| `TraceLevel` | 6 | Most verbose logging |

## 🔌 Framework Integration

### Zap Integration

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/lazygophers/log"
)

// Create a zap logger that writes to our log system
logger := log.New()
hook := log.NewZapHook(logger)

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    hook,
    zapcore.InfoLevel,
)
zapLogger := zap.New(core)

zapLogger.Info("Message from Zap", zap.String("key", "value"))
```

## 🧪 Testing

The library comes with comprehensive testing support:

```bash
# Run all tests
make test

# Run tests with coverage for all build tags
make coverage-all

# Quick test across all build tags
make test-quick

# Generate HTML coverage reports
make coverage-html
```

### Coverage Results by Build Tag

| Build Tag | Coverage | Description |
|-----------|----------|-------------|
| **Default** | 93.3% | Standard build with full debugging |
| **Debug** | 93.5% | Enhanced debug information |
| **Release** | **93.9%** | Production-optimized build |
| **Discard** | 93.5% | Maximum performance, no-op logging |
| **Debug+Discard** | 93.5% | Debug with discard optimization |
| **Release+Discard** | 93.7% | Release with discard optimization |

### Test Statistics
- **Total Test Functions**: 284+
- **Functions with 100% Coverage**: 125/138 (90.6%)
- **Build Tag Combinations Tested**: 6/6
- **Excluded from Coverage**: Fatal/Panic functions (safety reasons)

## ⚙️ Configuration Options

### Logger Configuration

```go
logger := log.New()

// Set minimum log level
logger.SetLevel(log.InfoLevel)

// Configure output
logger.SetOutput(os.Stdout) // Single writer
logger.SetOutput(writer1, writer2, writer3) // Multiple writers

// Customize messages
logger.SetPrefixMsg("[MyApp] ")
logger.SetSuffixMsg(" [END]")
logger.AppendPrefixMsg("Extra: ")

// Configure formatting
logger.ParsingAndEscaping(false) // Disable escape sequences
logger.Caller(true) // Enable caller information
logger.SetCallerDepth(4) // Adjust caller stack depth
```

## 📁 Log Rotation

Automatic log rotation with configurable intervals:

```go
// Hourly rotation
writer := log.GetOutputWriterHourly("./logs/app.log")

// The library will create files like:
// - app-2024010115.log (2024-01-01 15:00)
// - app-2024010116.log (2024-01-01 16:00)
// - app-2024010117.log (2024-01-01 17:00)
```

## 🔍 Context and Tracing

Built-in support for context-aware logging and distributed tracing:

```go
// Set trace ID for current goroutine
log.SetTrace("trace-123-456")

// Get trace ID
traceID := log.GetTrace()

// Context-aware logging
ctx := context.Background()
ctxLogger := log.CloneToCtx()
ctxLogger.Info(ctx, "Request processed", "user_id", 123)

// Automatic goroutine ID tracking
log.Info("This log includes goroutine ID automatically")
```

## 📈 Performance

The library is designed for high-performance applications:

- **Object Pooling**: Reuses log entry objects to reduce GC pressure
- **Async Writing**: Non-blocking log writes for high-throughput scenarios
- **Level Filtering**: Early filtering prevents expensive operations
- **Build Tag Optimization**: Compile-time optimization for different environments

### Benchmarks

```bash
# Run performance benchmarks
make benchmark

# Benchmark different build modes
make benchmark-debug
make benchmark-release  
make benchmark-discard
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](docs/CONTRIBUTING.md) for details.

### Development Setup

1. **Fork and Clone**
   ```bash
   git clone https://github.com/your-username/log.git
   cd log
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Run Tests**
   ```bash
   make test-all
   ```

4. **Submit Pull Request**
   - Follow our [PR Template](.github/pull_request_template.md)
   - Ensure tests pass
   - Update documentation if needed

## 📋 Requirements

- **Go**: 1.21 or higher
- **Dependencies**: 
  - `go.uber.org/zap` (for Zap integration)
  - `github.com/petermattis/goid` (for goroutine ID)
  - `github.com/lestrrat-go/file-rotatelogs` (for log rotation)
  - `github.com/google/uuid` (for trace IDs)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Zap](https://github.com/uber-go/zap) for inspiration and integration support
- [Logrus](https://github.com/sirupsen/logrus) for level design patterns
- The Go community for continuous feedback and improvements

## 📞 Support

- 📖 [Documentation](docs/)
- 🐛 [Issue Tracker](https://github.com/lazygophers/log/issues)
- 💬 [Discussions](https://github.com/lazygophers/log/discussions)
- 📧 Email: support@lazygophers.com

---

**Made with ❤️ by the LazyGophers team**
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
