# 🚀 LazyGophers Log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![DeepWiki](https://img.shields.io/badge/DeepWiki-documented-blue?logo=bookstack&logoColor=white)](https://deepwiki.ai/docs/lazygophers/log)
[![Go.Dev Downloads](https://pkg.go.dev/badge/github.com/lazygophers/log.svg)](https://pkg.go.dev/github.com/lazygophers/log)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/log)
[![Goproxy.io](https://goproxy.io/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.io/stats/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个高性能、功能丰富的 Go 日志库，支持多构建标签、异步写入和广泛的自定义选项。

## 📖 文档语言

- [🇺🇸 English](../README.md)
- [🇨🇳 简体中文](README.zh-CN.md) (当前)
- [🇹🇼 繁體中文](README.zh-TW.md)
- [🇫🇷 Français](README.fr.md)
- [🇷🇺 Русский](README.ru.md)
- [🇪🇸 Español](README.es.md)
- [🇸🇦 العربية](README.ar.md)

## ✨ 特性

- **🚀 高性能**: 对象池和异步写入支持
- **🏗️ 构建标签支持**: 为调试、发布和丢弃模式提供不同行为
- **🔄 日志轮转**: 自动按小时轮转日志文件
- **🎨 丰富格式化**: 可定制的日志格式和颜色支持
- **🔍 上下文追踪**: Goroutine ID 和追踪 ID 跟踪
- **🔌 框架集成**: 原生 Zap 日志框架集成
- **⚙️ 高度可配置**: 灵活的级别、输出和格式化配置
- **🧪 充分测试**: 在所有构建配置下达到 93.0% 测试覆盖率

## 🚀 快速开始

### 安装

```bash
go get github.com/lazygophers/log
```

### 基本用法

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 简单日志记录
    log.Info("你好，世界！")
    log.Debug("这是一条调试消息")
    log.Warn("这是一条警告")
    log.Error("这是一条错误")

    // 格式化日志记录
    log.Infof("用户 %s 登录，ID 为 %d", "张三", 123)
    
    // 使用自定义日志器
    logger := log.New()
    logger.SetLevel(log.InfoLevel)
    logger.Info("自定义日志器消息")
}
```

### 高级用法

```go
package main

import (
    "context"
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 创建带文件输出的日志器
    logger := log.New()
    
    // 设置输出到按小时轮转的文件
    writer := log.GetOutputWriterHourly("./logs/app.log")
    logger.SetOutput(writer)
    
    // 配置格式化
    logger.SetLevel(log.DebugLevel)
    logger.SetPrefixMsg("[APP] ")
    logger.Caller(true) // 启用调用者信息
    
    // 上下文日志记录
    ctxLogger := logger.CloneToCtx()
    ctxLogger.Info(context.Background(), "上下文感知日志记录")
    
    // 高吞吐量场景的异步日志记录
    asyncWriter := log.NewAsyncWriter(writer, 1000)
    logger.SetOutput(asyncWriter)
    defer asyncWriter.Close()
    
    logger.Info("高性能异步日志记录")
}
```

## 🏗️ 构建标签

该库通过 Go 构建标签支持不同的构建模式：

### 默认模式（无标签）
```bash
go build
```
- 完整日志功能
- 调试消息启用
- 标准性能

### 调试模式
```bash
go build -tags debug
```
- 增强调试信息
- 详细调用者信息
- 性能分析支持

### 发布模式
```bash
go build -tags release
```
- 针对生产环境优化
- 调试消息禁用
- 自动日志文件轮转

### 丢弃模式
```bash
go build -tags discard
```
- 最大性能
- 所有日志被丢弃
- 零日志开销

### 组合模式
```bash
go build -tags "debug,discard"    # 调试 + 丢弃
go build -tags "release,discard"  # 发布 + 丢弃
```

## 📊 日志级别

该库支持 7 个日志级别（从最高到最低优先级）：

| 级别 | 值 | 描述 |
|------|---|------|
| `PanicLevel` | 0 | 记录日志然后调用 panic |
| `FatalLevel` | 1 | 记录日志然后调用 os.Exit(1) |
| `ErrorLevel` | 2 | 错误条件 |
| `WarnLevel` | 3 | 警告条件 |
| `InfoLevel` | 4 | 信息消息 |
| `DebugLevel` | 5 | 调试级别消息 |
| `TraceLevel` | 6 | 最详细的日志记录 |

## 🔌 框架集成

### Zap 集成

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/lazygophers/log"
)

// 创建一个写入到我们日志系统的 zap 日志器
logger := log.New()
hook := log.NewZapHook(logger)

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    hook,
    zapcore.InfoLevel,
)
zapLogger := zap.New(core)

zapLogger.Info("来自 Zap 的消息", zap.String("key", "value"))
```

## 🧪 测试

该库提供全面的测试支持：

```bash
# 运行所有测试
make test

# 运行所有构建标签的覆盖率测试
make coverage-all

# 快速测试所有构建标签
make test-quick

# 生成 HTML 覆盖率报告
make coverage-html
```

### 按构建标签的覆盖率结果

| 构建标签 | 覆盖率 |
|----------|--------|
| 默认 | 92.9% |
| 调试 | 93.1% |
| 发布 | 93.5% |
| 丢弃 | 93.1% |
| 调试+丢弃 | 93.1% |
| 发布+丢弃 | 93.3% |

## ⚙️ 配置选项

### 日志器配置

```go
logger := log.New()

// 设置最小日志级别
logger.SetLevel(log.InfoLevel)

// 配置输出
logger.SetOutput(os.Stdout) // 单个写入器
logger.SetOutput(writer1, writer2, writer3) // 多个写入器

// 自定义消息
logger.SetPrefixMsg("[MyApp] ")
logger.SetSuffixMsg(" [END]")
logger.AppendPrefixMsg("额外: ")

// 配置格式化
logger.ParsingAndEscaping(false) // 禁用转义序列
logger.Caller(true) // 启用调用者信息
logger.SetCallerDepth(4) // 调整调用者栈深度
```

## 📁 日志轮转

可配置间隔的自动日志轮转：

```go
// 按小时轮转
writer := log.GetOutputWriterHourly("./logs/app.log")

// 该库将创建如下文件：
// - app-2024010115.log (2024-01-01 15:00)
// - app-2024010116.log (2024-01-01 16:00)
// - app-2024010117.log (2024-01-01 17:00)
```

## 🔍 上下文和追踪

内置支持上下文感知日志记录和分布式追踪：

```go
// 为当前 goroutine 设置追踪 ID
log.SetTrace("trace-123-456")

// 获取追踪 ID
traceID := log.GetTrace()

// 上下文感知日志记录
ctx := context.Background()
ctxLogger := log.CloneToCtx()
ctxLogger.Info(ctx, "请求已处理", "user_id", 123)

// 自动 goroutine ID 跟踪
log.Info("此日志自动包含 goroutine ID")
```

## 📈 性能

该库为高性能应用程序而设计：

- **对象池**: 重用日志条目对象以减少 GC 压力
- **异步写入**: 高吞吐量场景的非阻塞日志写入
- **级别过滤**: 早期过滤防止昂贵操作
- **构建标签优化**: 不同环境的编译时优化

### 基准测试

```bash
# 运行性能基准测试
make benchmark

# 不同构建模式的基准测试
make benchmark-debug
make benchmark-release  
make benchmark-discard
```

## 🤝 贡献

我们欢迎贡献！请查看我们的[贡献指南](CONTRIBUTING.md)了解详情。

### 开发环境设置

1. **Fork 并克隆**
   ```bash
   git clone https://github.com/your-username/log.git
   cd log
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **运行测试**
   ```bash
   make test-all
   ```

4. **提交 Pull Request**
   - 遵循我们的 [PR 模板](../.github/pull_request_template.md)
   - 确保测试通过
   - 如需要请更新文档

## 📋 要求

- **Go**: 1.19 或更高版本
- **依赖项**: 
  - `go.uber.org/zap` (用于 Zap 集成)
  - `github.com/petermattis/goid` (用于 goroutine ID)
  - `github.com/lestrrat-go/file-rotatelogs` (用于日志轮转)
  - `github.com/google/uuid` (用于追踪 ID)

## 📄 许可证

本项目采用 MIT 许可证 - 请查看 [LICENSE](../LICENSE) 文件了解详情。

## 🙏 致谢

- [Zap](https://github.com/uber-go/zap) 提供灵感和集成支持
- [Logrus](https://github.com/sirupsen/logrus) 提供级别设计模式
- Go 社区持续的反馈和改进

## 📞 支持

- 📖 [文档](../docs/)
- 🐛 [问题跟踪](https://github.com/lazygophers/log/issues)
- 💬 [讨论](https://github.com/lazygophers/log/discussions)
- 📧 邮箱: support@lazygophers.com

---

**由 LazyGophers 团队用 ❤️ 制作**