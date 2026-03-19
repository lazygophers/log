---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个高性能、灵活的 Go 日志库，基于 zap 构建，提供丰富的功能和简洁的 API。

## 📖 文档语言

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 简体中文](README.md)（当前）
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)

## ✨ 特性

-   **🚀 高性能**：基于 zap 构建，采用对象池复用 Entry 对象，减少内存分配
-   **📊 丰富的日志级别**：支持 Trace、Debug、Info、Warn、Error、Fatal、Panic 级别
-   **⚙️ 灵活的配置**：
    -   日志级别控制
    -   调用者信息记录
    -   追踪信息（包括 goroutine ID）
    -   自定义日志前缀和后缀
    -   自定义输出目标（控制台、文件等）
    -   日志格式化选项
-   **🔄 文件轮转**：支持每小时日志文件轮转
-   **🔌 Zap 兼容性**：与 zap WriteSyncer 无缝集成
-   **🎯 简洁的 API**：类似标准日志库的清晰 API，易于使用

## 🚀 快速开始

### 安装

:::tip 安装
```bash
go get github.com/lazygophers/log
```
:::

### 基本用法

```go title="快速开始"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 使用默认全局 logger
    log.Debug("调试信息")
    log.Info("普通信息")
    log.Warn("警告信息")
    log.Error("错误信息")

    // 使用格式化输出
    log.Infof("用户 %s 登录成功", "admin")

    // 自定义配置
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("这是来自自定义 logger 的日志")
}
```

## 📚 高级用法

### 带文件输出的自定义 Logger

```go title="文件输出配置"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 创建带文件输出的 logger
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("带调用者信息的调试日志")
    logger.Info("带追踪信息的普通日志")
}
```

### 日志级别控制

```go title="日志级别控制"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // 只有 warn 及以上级别会被记录
    logger.Debug("这不会被记录")  // 忽略
    logger.Info("这不会被记录")   // 忽略
    logger.Warn("这会被记录")    // 记录
    logger.Error("这会被记录")   // 记录
}
```

## 🎯 使用场景

### 适用场景

-   **Web 服务和 API 后端**：请求追踪、结构化日志、性能监控
-   **微服务架构**：分布式追踪（TraceID）、统一日志格式、高吞吐量
-   **命令行工具**：级别控制、简洁输出、错误报告
-   **批处理任务**：文件轮转、长时间运行、资源优化

### 特别优势

-   **对象池优化**：Entry 和 Buffer 对象复用，减少 GC 压力
-   **异步写入**：高吞吐量场景（10000+ 条/秒）无阻塞
-   **TraceID 支持**：分布式系统请求追踪，与 OpenTelemetry 集成
-   **零配置启动**：开箱即用，渐进式配置

## 🔧 配置选项

:::note 配置选项
以下方法均支持链式调用，可组合使用构建自定义 Logger。
:::

### Logger 配置

| 方法                  | 描述                | 默认值       |
| --------------------- | ------------------- | ----------- |
| `SetLevel(level)`       | 设置最低日志级别     | `DebugLevel` |
| `EnableCaller(enable)`  | 启用/禁用调用者信息  | `false`      |
| `EnableTrace(enable)`   | 启用/禁用追踪信息    | `false`      |
| `SetCallerDepth(depth)` | 设置调用者深度       | `2`          |
| `SetPrefixMsg(prefix)`  | 设置日志前缀         | `""`         |
| `SetSuffixMsg(suffix)`  | 设置日志后缀         | `""`         |
| `SetOutput(writers...)` | 设置输出目标         | `os.Stdout`  |

### 日志级别

| 级别        | 描述                        |
| ----------- | --------------------------- |
| `TraceLevel` | 最详细，用于详细跟踪        |
| `DebugLevel` | 调试信息                    |
| `InfoLevel`  | 普通信息                    |
| `WarnLevel`  | 警告信息                    |
| `ErrorLevel` | 错误信息                    |
| `FatalLevel` | 致命错误（调用 os.Exit(1)） |
| `PanicLevel` | 恐慌错误（调用 panic()）    |

## 🏗️ 架构

### 核心组件

-   **Logger**：主日志结构，具有可配置的级别、输出、格式化器和调用者深度
-   **Entry**：单个日志记录，包含全面的元数据支持
-   **Level**：日志级别定义和工具函数
-   **Format**：日志格式化接口和实现

### 性能优化

-   **对象池**：重用 Entry 对象以减少内存分配
-   **条件记录**：仅在需要时记录昂贵字段
-   **快速级别检查**：在最外层检查日志级别
-   **无锁设计**：大多数操作不需要锁

## 📊 性能比较

:::info 性能对比
以下数据基于基准测试，实际性能可能因环境和配置不同而有所差异。
:::

| 特性          | lazygophers/log | zap    | logrus | 标准日志 |
| ------------- | --------------- | ------ | ------ | -------- |
| 性能          | 高              | 高     | 中     | 低       |
| API 简洁性    | 高              | 中     | 高     | 高       |
| 功能丰富度    | 中              | 高     | 高     | 低       |
| 灵活性        | 中              | 高     | 高     | 低       |
| 学习曲线      | 低              | 中     | 中     | 低       |

## ❓ 常见问题

### 如何选择合适的日志级别？

-   **开发环境**：使用 `DebugLevel` 或 `TraceLevel` 获取详细信息
-   **生产环境**：使用 `InfoLevel` 或 `WarnLevel` 减少开销
-   **性能测试**：使用 `PanicLevel` 禁用所有日志

### 如何在生产环境优化性能？

:::warning 注意
在高吞吐量场景下，建议结合异步写入和合理的日志级别来优化性能。
:::

1. 使用 `AsyncWriter` 异步写入：

```go title="异步写入配置"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. 调整日志级别，避免不必要的日志：

```go title="级别优化"
logger.SetLevel(log.InfoLevel)  // 跳过 Debug 和 Trace
```

3. 使用条件日志减少开销：

```go title="条件日志"
if logger.Level >= log.DebugLevel {
    logger.Debug("详细调试信息")
}
```

### `Caller` 和 `EnableCaller` 有什么区别？

-   **`EnableCaller(enable bool)`**：控制 Logger 是否收集调用者信息
    -   `EnableCaller(true)` 启用调用者信息收集
-   **`Caller(disable bool)`**：控制 Formatter 是否输出调用者信息
    -   `Caller(true)` 禁用调用者信息输出

推荐使用 `EnableCaller` 进行全局控制。

### 如何实现自定义格式化器？

实现 `Format` 接口：

```go title="自定义格式化器"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 相关文档

-   [📚 API 文档](API.md) - 完整的 API 参考
-   [🤝 贡献指南](/zh-CN/CONTRIBUTING) - 如何贡献
-   [📋 变更日志](/zh-CN/CHANGELOG) - 版本历史
-   [🔒 安全政策](/zh-CN/SECURITY) - 安全指南
-   [📜 行为准则](/zh-CN/CODE_OF_CONDUCT) - 社区准则

## 🚀 获取帮助

-   **GitHub Issues**：[报告 bug 或请求功能](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[API 文档](https://pkg.go.dev/github.com/lazygophers/log)
-   **示例**：[使用示例](https://github.com/lazygophers/log/tree/main/examples)

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](/zh-CN/LICENSE) 文件。

## 🤝 贡献

我们欢迎贡献！请查看我们的 [贡献指南](/zh-CN/CONTRIBUTING) 了解详情。

---

**lazygophers/log** 旨在成为重视性能和简洁性的 Go 开发者的首选日志解决方案。无论您是构建小型工具还是大规模分布式系统，该库都能提供功能和易用性之间的良好平衡。
