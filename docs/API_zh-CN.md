# 📚 API 文档

## 概述

LazyGophers Log 提供了一个全面的日志 API，支持多个日志级别、自定义格式化、异步写入和构建标签优化。本文档涵盖所有公共 API、配置选项和使用模式。

## 目录

- [核心类型](#核心类型)
- [Logger API](#logger-api)
- [全局函数](#全局函数)
- [日志级别](#日志级别)
- [格式化器](#格式化器)
- [输出写入器](#输出写入器)
- [上下文日志](#上下文日志)
- [构建标签](#构建标签)
- [性能优化](#性能优化)
- [示例](#示例)

## 核心类型

### Logger

提供所有日志功能的主要日志结构体。

```go
type Logger struct {
    // 包含用于线程安全操作的私有字段
}
```

#### 构造函数

```go
func New() *Logger
```

创建具有默认配置的新日志实例：
- 级别: `DebugLevel`
- 输出: `os.Stdout`
- 格式化器: 默认文本格式化器
- 调用者跟踪: 禁用

**示例:**
```go
logger := log.New()
logger.Info("新日志器已创建")
```

### Entry

表示单个日志条目及其所有关联元数据。

```go
type Entry struct {
    Time       time.Time     // 条目创建时的时间戳
    Level      Level         // 日志级别
    Message    string        // 日志消息
    Pid        int          // 进程 ID
    Gid        uint64       // Goroutine ID
    TraceID    string       // 分布式跟踪的跟踪 ID
    CallerName string       // 调用者函数名
    CallerFile string       // 调用者文件路径
    CallerLine int          // 调用者行号
}
```

## Logger API

### 配置方法

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

设置最小日志级别。低于此级别的消息将被忽略。

**参数:**
- `level`: 要处理的最小日志级别

**返回:**
- `*Logger`: 返回自身用于方法链接

**示例:**
```go
logger.SetLevel(log.InfoLevel)
logger.Debug("这不会被显示")  // 被忽略
logger.Info("这将被显示")     // 被处理
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

为日志消息设置一个或多个输出目标。

**参数:**
- `writers`: 一个或多个 `io.Writer` 目标

**示例:**
```go
// 单个输出
logger.SetOutput(os.Stdout)

// 多个输出
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

为日志输出设置自定义格式化器。

**示例:**
```go
logger.SetFormatter(&JSONFormatter{})
```

### 日志方法

所有日志方法都有两种变体：简单和格式化。

#### 级别方法

```go
// Trace 级别 - 最详细
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)

// Debug 级别 - 调试信息
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)

// Info 级别 - 信息消息
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)

// Warn 级别 - 警告消息
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)

// Error 级别 - 错误消息
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)

// Fatal 级别 - 致命错误，调用 os.Exit(1)
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)

// Panic 级别 - 记录错误并调用 panic()
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

**示例:**
```go
logger.Info("应用程序已启动")
logger.Errorf("处理请求失败: %v", err)
```

## 日志级别

### 可用级别

```go
const (
    PanicLevel Level = iota  // 0 - Panic 并退出
    FatalLevel              // 1 - 致命错误并退出  
    ErrorLevel              // 2 - 错误条件
    WarnLevel               // 3 - 警告条件
    InfoLevel               // 4 - 信息消息
    DebugLevel              // 5 - 调试消息
    TraceLevel              // 6 - 最详细的跟踪
)
```

## 格式化器

### Format 接口

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

自定义格式化器必须实现此接口。

### JSON 格式化器示例

```go
type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "caller":    fmt.Sprintf("%s:%d", entry.CallerFile, entry.CallerLine),
    }
    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }
    
    jsonData, _ := json.Marshal(data)
    return append(jsonData, '\n')
}

// 使用
logger.SetFormatter(&JSONFormatter{})
```

## 输出写入器

### 带轮转的文件输出

```go
func GetOutputWriterHourly(filename string) io.Writer
```

创建一个按小时轮转日志文件的写入器。

**示例:**
```go
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// 创建文件如: app-2024010115.log, app-2024010116.log, 等等
```

### 异步写入器

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

创建用于高性能日志记录的异步写入器。

**示例:**
```go
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## 上下文日志

### 上下文函数

```go
func SetTrace(traceID string)
func GetTrace() string
```

为当前 goroutine 设置和获取跟踪 ID。

**示例:**
```go
log.SetTrace("trace-123-456")
log.Info("此消息将包含跟踪 ID")

traceID := log.GetTrace()
fmt.Println("当前跟踪 ID:", traceID)
```

## 构建标签

库支持使用构建标签进行条件编译：

### 默认模式
```bash
go build
```
- 启用完整功能
- 包含调试消息
- 标准性能

### 调试模式
```bash
go build -tags debug
```
- 增强的调试信息
- 详细的调用者信息

### 发布模式
```bash
go build -tags release
```
- 为生产环境优化
- 禁用调试消息
- 启用自动日志轮转

### 丢弃模式
```bash
go build -tags discard
```
- 最大性能
- 所有日志都被丢弃
- 零日志开销

## 性能优化

### 对象池化

库内部使用 `sync.Pool` 来池化：
- 日志条目对象
- 字节缓冲区
- 格式化器缓冲区

这在高吞吐量场景中减少了垃圾收集压力。

### 级别检查

日志级别检查在昂贵操作之前进行：

```go
// 高效 - 仅在级别启用时才进行消息格式化
logger.Debugf("昂贵操作结果: %+v", expensiveCall())
```

### 异步写入

对于高吞吐量应用程序：

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // 大缓冲区
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

## 示例

### 基本使用

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("应用程序启动中")
    log.Warn("这是一个警告")
    log.Error("这是一个错误")
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
    logger := log.New()
    
    // 配置日志器
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)
    logger.SetPrefixMsg("[我的应用] ")
    
    // 设置输出到文件
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    logger.SetOutput(file)
    
    logger.Info("自定义日志器已配置")
    logger.Debug("带调用者的调试信息")
}
```

### 高性能日志

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 创建轮转文件写入器
    writer := log.GetOutputWriterHourly("./logs/app.log")
    
    // 用异步写入器包装以提高性能
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()
    
    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // 在生产环境中跳过调试
    
    // 高吞吐量日志记录
    for i := 0; i < 10000; i++ {
        logger.Infof("处理请求 %d", i)
    }
}
```

### 上下文感知日志

```go
package main

import (
    "context"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()
    ctxLogger := logger.CloneToCtx()
    
    ctx := context.Background()
    log.SetTrace("trace-123-456")
    
    ctxLogger.Info(ctx, "处理用户请求")
    ctxLogger.Debug(ctx, "验证完成")
}
```

## 错误处理

出于性能原因，大多数日志器方法不返回错误。如果您需要输出操作的错误处理，请实现自定义写入器。

## 线程安全

所有日志器操作都是线程安全的，可以从多个 goroutine 并发使用，无需额外同步。

---

## 🌍 多语言文档

本文档提供多种语言版本：

- [🇺🇸 English](API.md)
- [🇨🇳 简体中文](API_zh-CN.md)（当前）
- [🇹🇼 繁體中文](API_zh-TW.md)
- [🇫🇷 Français](API_fr.md)
- [🇷🇺 Русский](API_ru.md)
- [🇪🇸 Español](API_es.md)
- [🇸🇦 العربية](API_ar.md)

---

**LazyGophers Log 的完整 API 参考 - 用卓越的日志构建更好的应用程序！🚀**