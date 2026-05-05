# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-95.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

高性能、灵活的 Go 日志库，专注于简洁的 API 和强大的扩展能力。

## 特性

- **高性能设计**
  - 基于 `sync.Pool` 的对象池优化，减少 GC 压力
  - 早期级别检查，避免不必要的计算
  - 内联函数优化热点路径
  - 缓存友好的 Entry 结构布局（200字节，仅2%填充浪费）

- **丰富的日志级别**
  - 七个日志级别：Trace、Debug、Info、Warn、Error、Fatal、Panic
  - 兼容 logrus 级别约定

- **灵活的配置选项**
  - 日志级别控制
  - 调用者信息记录（文件名、行号、函数名）
  - 追踪信息（Goroutine ID、Trace ID）
  - 自定义日志前缀和后缀
  - 多输出目标支持（控制台、文件、自定义 Writer）
  - 可插拔的格式化器（文本/JSON）

- **强大的扩展机制**
  - Hook 接口：在日志写入前进行修改、过滤或丰富
  - 自定义格式化器：实现 `Format` 或 `FormatFull` 接口
  - 多 Writer 支持：同时输出到多个目标

- **日志轮转**
  - 按小时自动轮转
  - 基于大小的分片支持
  - 自动清理过期日志文件

- **结构化日志**
  - 支持键值对字段（`Infow`、`Debugw` 等方法）
  - 自定义 JSON 序列化（Entry 实现 `MarshalJSON` 接口）

## 安装

```bash
go get github.com/lazygophers/log
```

## 快速开始

### 基础使用

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 使用默认全局日志器
    log.Debug("调试信息")
    log.Info("普通信息")
    log.Warn("警告信息")
    log.Error("错误信息")

    // 格式化输出
    log.Infof("用户 %s 登录成功", "admin")

    // 自定义配置
    logger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    logger.Info("来自自定义日志器的消息")
}
```

### 结构化日志

```go
package main

import "github.com/lazygophers/log"

func main() {
    // 使用键值对记录结构化数据
    log.Infow("用户登录",
        "user_id", 12345,
        "username", "admin",
        "ip", "192.168.1.100",
        "success", true,
    )

    // 输出示例：
    // (12345) 2026-05-05 12:34:56.789+08:00 [INFO] 用户登录 user_id=12345 username=admin ip=192.168.1.100 success=true
}
```

### 自定义日志器

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 创建带文件输出的日志器
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(
            os.Stdout,
            log.GetOutputWriterHourly("/var/log/myapp.log"),
        )

    logger.Debug("带调用者信息的调试消息")
    logger.Info("带追踪信息的普通消息")
}
```

### JSON 格式输出

```go
package main

import "github.com/lazygophers/log"

func main() {
    // 切换到 JSON 格式
    logger := log.New()
    logger.Format = &log.JSONFormatter{}

    logger.Infow("服务启动",
        "port", 8080,
        "env", "production",
    )

    // 输出示例：
    // {"level":"info","message":"服务启动","pid":12345,"time":"2026-05-05T12:34:56.789+08:00","fields":{"port":8080,"env":"production"}}
}
```

### 使用 Hook

```go
package main

import (
    "github.com/lazygophers/log"
    "github.com/lazygophers/log/constant"
)

// EnvironmentHook 添加环境信息到每条日志
type EnvironmentHook struct {
    env string
}

func (h *EnvironmentHook) OnWrite(entry interface{}) interface{} {
    if e, ok := entry.(*log.Entry); ok {
        e.Fields = append(e.Fields, log.KV{
            Key:   "environment",
            Value: h.env,
        })
    }
    return entry
}

func main() {
    logger := log.New()
    logger.AddHook(&EnvironmentHook{env: "production"})

    logger.Info("这条日志会包含 environment 字段")
}
```

### 日志级别控制

```go
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // 只有 Warn 及以上级别会被记录
    logger.Debug("这条不会记录")  // 忽略
    logger.Info("这条也不会记录")   // 忽略
    logger.Warn("这条会被记录")    // 记录
    logger.Error("这条也会被记录")  // 记录
}
```

## 配置选项

### 日志器配置方法

| 方法 | 描述 | 默认值 |
|------|------|--------|
| `SetLevel(level)` | 设置最低日志级别 | `DebugLevel` |
| `EnableCaller(enable)` | 启用/禁用调用者信息 | `true` |
| `EnableTrace(enable)` | 启用/禁用追踪信息 | `true` |
| `SetCallerDepth(depth)` | 设置调用者栈深度 | `4` |
| `SetPrefixMsg(prefix)` | 设置日志前缀 | `""` |
| `SetSuffixMsg(suffix)` | 设置日志后缀 | `""` |
| `SetOutput(writers...)` | 设置输出目标 | `os.Stdout` |
| `AddHook(hook)` | 添加处理 Hook | `nil` |
| `Clone()` | 创建日志器副本 | - |

### 日志级别

| 级别 | 描述 |
|------|------|
| `TraceLevel` | 最详细，用于详细追踪 |
| `DebugLevel` | 调试信息 |
| `InfoLevel` | 普通信息 |
| `WarnLevel` | 警告信息 |
| `ErrorLevel` | 错误信息 |
| `FatalLevel` | 致命错误（调用 `os.Exit(1)`） |
| `PanicLevel` | 恐慌错误（调用 `panic()`） |

## 架构设计

### 核心组件

- **Logger** (`logger.go`): 主日志结构，提供链式配置 API
- **Entry** (`constant/entry.go`): 日志条目结构，优化的内存布局
- **Level** (`constant/level.go`): 日志级别定义和实用函数
- **Formatter** (`formatter.go`, `formatter_json.go`): 日志格式化实现
- **Rotator** (`rotator.go`): 基于时间和大小的日志轮转
- **Hook** (`constant/interface.go`): 日志处理扩展接口

### 性能优化

- **对象池**: Entry 和 Buffer 对象复用，减少内存分配
- **早期级别检查**: 在外层检查日志级别，避免不必要的操作
- **内联函数**: 热点路径函数使用 `//go:inline` 标记
- **缓存优化**: Entry 结构按访问频率和大小排列，最小化缓存行浪费
- **时间戳缓存**: 格式化的时间戳缓存为字符串，避免重复格式化

### 接口设计

```go
// Hook 接口 - 日志处理扩展
type Hook interface {
    OnWrite(entry interface{}) interface{}
}

// Format 接口 - 日志格式化
type Format interface {
    Format(entry interface{}) []byte
}

// FormatFull 接口 - 扩展格式化
type FormatFull interface {
    Format
    ParsingAndEscaping(disable bool)
    Caller(disable bool)
    Clone() Format
}
```

## 日志轮转

```go
// 创建按小时轮转的日志写入器
// 参数：日志目录、单文件最大大小（字节）、保留文件数量
writer := log.GetOutputWriterHourly("/var/log/app", 100*1024*1024, 168)

logger := log.New()
logger.SetOutput(writer)

// 轮转特性：
// - 按小时自动创建新文件
// - 单文件超过大小限制时创建分片
// - 自动清理超过保留数量的旧文件
```

## 测试

```bash
# 运行所有测试
go test ./...

# 运行测试并显示覆盖率
go test -cover ./...

# 运行基准测试
go test -bench=. -benchmem

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**当前测试覆盖率**: 95.0% (486 个测试用例)

## 性能对比

| 特性 | lazygophers/log | zap | logrus | 标准库 |
|------|-----------------|-----|--------|--------|
| 性能 | 高 | 高 | 中 | 低 |
| API 简洁性 | 高 | 中 | 高 | 高 |
| 功能丰富度 | 中 | 高 | 高 | 低 |
| 扩展性 | 中 | 高 | 高 | 低 |
| 学习曲线 | 低 | 中 | 中 | 低 |

## 文档

- [API 文档](docs/en/API.md) - 完整 API 参考
- [架构文档](docs/architecture.md) - 设计决策和架构细节
- [Hook 指南](docs/hooks_guide.md) - Hook 使用教程
- [贡献指南](CONTRIBUTING.md) - 如何贡献代码
- [变更日志](CHANGELOG.md) - 版本历史
- [安全策略](SECURITY.md) - 漏洞报告流程

## 多语言文档

- [🇺🇸 English](README.md)
- [🇨🇳 简体中文](docs/zh-CN/README.md)
- [🇹🇼 繁體中文](docs/zh-TW/README.md)
- [🇫🇷 Français](docs/fr/README.md)
- [🇪🇸 Español](docs/es/README.md)
- [🇷🇺 Русский](docs/ru/README.md)
- [🇸🇦 العربية](docs/ar/README.md)
- [🇯🇵 日本語](docs/ja/README.md)
- [🇩🇪 Deutsch](docs/de/README.md)
- [🇰🇷 한국어](docs/ko/README.md)
- [🇵🇹 Português](docs/pt/README.md)
- [🇳🇱 Nederlands](docs/nl/README.md)
- [🇵🇱 Polski](docs/pl/README.md)
- [🇮🇹 Italiano](docs/it/README.md)
- [🇹🇷 Türkçe](docs/tr/README.md)

## 获取帮助

- **GitHub Issues**: [报告问题或请求功能](https://github.com/lazygophers/log/issues)
- **GoDoc**: [API 文档](https://pkg.go.dev/github.com/lazygophers/log)
- **讨论**: [GitHub Discussions](https://github.com/lazygophers/log/discussions)

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 贡献

欢迎贡献！请参阅[贡献指南](CONTRIBUTING.md)了解详情。

---

**lazygophers/log** 旨在成为重视性能和简洁性的 Go 开发者的首选日志解决方案。无论是构建小型工具还是大规模分布式系统，本库都能在功能和易用性之间提供恰当的平衡。
