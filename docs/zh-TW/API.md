---
titleSuffix: ' | LazyGophers Log'
---
# 📚 API Documentation

## 概述

LazyGophers Log 提供了一個全面的日誌記錄 API，支援多日誌級別、自定義格式化、非同步写入和構建標籤優化。本檔案涵蓋了所有公共 API、配置選項和使用模式。

## 目录

-   [核心類型](#核心類型)
-   [Logger API](#logger-api)
-   [全局函式](#全局函式)
-   [日誌級別](#日誌級別)
-   [格式化器](#格式化器)
-   [输出写入器](#输出写入器)
-   [上下文日誌](#上下文日誌)
-   [構建標籤](#構建標籤)
-   [性能優化](#性能優化)
-   [示例](#示例)

## 核心類型

### Logger

提供所有日誌記錄功能的主要日誌記錄器結構體。

```go
type Logger struct {
    // 包含用於執行緒安全操作的私有欄位
}
```

#### 建構函式

```go
func New() *Logger
```

建立一個具有預設配置的新日誌記錄器實例：

-   級別：`DebugLevel`
-   輸出：`os.Stdout`
-   格式化器：預設文本格式化器
-   呼叫者追蹤：停用

**範例：**

```go title="建立日誌器"
logger := log.New()
logger.Info("新日誌記錄器已建立")
```

### Entry

表示具有所有關聯元資料的单个日誌條目。

```go
type Entry struct {
    Time       time.Time     // 條目建立時的時間戳
    Level      Level         // 日誌級別
    Message    string        // 日誌訊息
    Pid        int          // 進程 ID
    Gid        uint64       // Goroutine ID
    TraceID    string       // 分佈式追蹤的追蹤 ID
    CallerName string       // 呼叫者函式名
    CallerFile string       // 呼叫者檔案路徑
    CallerLine int          // 呼叫者行號
}
```

## Logger API

### 配置方法

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

设置最低日誌級別。低于此級別的訊息将被忽略。

**参数：**

-   `level`：要处理的最低日誌級別

**返回值：**

-   `*Logger`：返回自身以支援方法链式调用

**範例：**

```go title="设置日誌級別"
logger.SetLevel(log.InfoLevel)
logger.Debug("这不会被显示")  // 被忽略
logger.Info("这会被显示")    // 被处理
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

设置日誌訊息的一個或多个输出目标。

**参数：**

-   `writers`：一個或多个 `io.Writer` 输出目标

**返回值：**

-   `*Logger`：返回自身以支援方法链式调用

**範例：**

```go title="设置输出目标"
// 单一输出
logger.SetOutput(os.Stdout)

// 多个输出
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

设置日誌输出的自定義格式化器。

**参数：**

-   `formatter`：实现 `Format` 接口的格式化器

**返回值：**

-   `*Logger`：返回自身以支援方法链式调用

**範例：**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

启用或停用日誌條目中的呼叫者信息記錄。

**参数：**

-   `enable`：是否启用呼叫者信息（传入 `true` 表示启用）

**返回值：**

-   `*Logger`：返回自身以支援方法链式调用

**範例：**

```go
logger.EnableCaller(true)
logger.Info("这将包含檔案:行號信息")

logger.EnableCaller(false)
logger.Info("这不会包含檔案:行號信息")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

控制格式化器中的呼叫者信息。

**参数：**

-   `disable`：是否停用呼叫者信息（传入 `true` 表示停用）

**返回值：**

-   `*Logger`：返回自身以支援方法链式调用

**範例：**

```go
logger.Caller(false)  // 不停用，显示呼叫者信息
logger.Info("这将包含檔案:行號信息")

logger.Caller(true)   // 停用呼叫者信息
logger.Info("这不会包含檔案:行號信息")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

设置包装日誌記錄器時呼叫者信息的堆栈深度。

**参数：**

-   `depth`：要跳过的堆栈帧数

**返回值：**

-   `*Logger`：返回自身以支援方法链式调用

**範例：**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // 跳过包装函式
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

为所有日誌訊息设置前缀或后缀文本。

**参数：**

-   `prefix/suffix`：要前置/后置到訊息的文本

**返回值：**

-   `*Logger`：返回自身以支援方法链式调用

**範例：**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // 输出: [APP] Hello [END]
```

### 日誌記錄方法

所有日誌記錄方法都有两种变体：简单版本和格式化版本。

#### Trace 級別

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

在 trace 級別記錄日誌（最详细）。

**範例：**

```go
logger.Trace("详细执行追蹤")
logger.Tracef("处理第 %d 项，共 %d 项", i, total)
```

#### Debug 級別

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

在 debug 級別記錄开发信息。

**範例：**

```go
logger.Debug("变量状态:", variable)
logger.Debugf("用户 %s 认证成功", username)
```

#### Info 級別

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

記錄信息性訊息。

**範例：**

```go
logger.Info("应用程序已启动")
logger.Infof("服务器监听端口 %d", port)
```

#### Warn 級別

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

記錄警告訊息，用於潜在问题情况。

**範例：**

```go
logger.Warn("已调用弃用函式")
logger.Warnf("内存使用率高: %d%%", memoryPercent)
```

#### Error 級別

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

記錄错误訊息。

**範例：**

```go
logger.Error("数据库连接失败")
logger.Errorf("处理请求失败: %v", err)
```

#### Fatal 級別

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

記錄致命错误并调用 `os.Exit(1)`。

:::danger 破坏性操作
`Fatal` 和 `Fatalf` 会在日誌記錄后立即调用 `os.Exit(1)` 终止進程。请确保仅在不可恢复的错误情况下使用。`defer` 语句**不会**被执行。
:::

**範例：**

```go
logger.Fatal("关键系统错误")
logger.Fatalf("无法启动服务器: %v", err)
```

#### Panic 級別

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

記錄错误訊息并调用 `panic()`。

:::danger 破坏性操作
`Panic` 和 `Panicf` 会在日誌記錄后调用 `panic()`。与 `Fatal` 不同，`panic` 可以被 `recover()` 捕获，但如果未捕获则会终止程序。
:::

**範例：**

```go
logger.Panic("发生不可恢复错误")
logger.Panicf("无效状态: %v", state)
```

### 实用方法

#### Clone

```go
func (l *Logger) Clone() *Logger
```

建立具有相同配置的日誌記錄器副本。

**返回值：**

-   `*Logger`：具有复制设置的新日誌記錄器實例

**範例：**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

建立一個上下文感知的日誌記錄器，接受 `context.Context` 作为第一個参数。

**返回值：**

-   `LoggerWithCtx`：上下文感知的日誌記錄器實例

**範例：**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "上下文感知訊息")
```

## 全局函式

使用預設全局日誌記錄器的包级函式。

```go
func SetLevel(level Level)
func SetOutput(writers ...io.Writer)
func SetFormatter(formatter Format)
func Caller(disable bool)

func Trace(v ...any)
func Tracef(format string, v ...any)
func Debug(v ...any)
func Debugf(format string, v ...any)
func Info(v ...any)
func Infof(format string, v ...any)
func Warn(v ...any)
func Warnf(format string, v ...any)
func Error(v ...any)
func Errorf(format string, v ...any)
func Fatal(v ...any)
func Fatalf(format string, v ...any)
func Panic(v ...any)
func Panicf(format string, v ...any)
```

**範例：**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("使用全局日誌記錄器")
```

## 日誌級別

### Level 類型

```go
type Level int8
```

### 可用級別

```go
const (
    PanicLevel Level = iota  // 0 - Panic 并退出
    FatalLevel              // 1 - 致命错误并退出
    ErrorLevel              // 2 - 错误条件
    WarnLevel               // 3 - 警告条件
    InfoLevel               // 4 - 信息性訊息
    DebugLevel              // 5 - 调试訊息
    TraceLevel              // 6 - 最详细的追蹤
)
```

### Level 方法

```go
func (l Level) String() string
```

返回級別的字符串表示。

**範例：**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## 格式化器

### Format 接口

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

自定義格式化器必须实现此接口。

### 預設格式化器

具有可自定義選項的内置文本格式化器。

```go
type Formatter struct {
    // 配置選項
}
```

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

### 檔案输出与轮转

```go
func GetOutputWriterHourly(filename string) io.Writer
```

建立一個每小時轮转日誌檔案的写入器。

**参数：**

-   `filename`：日誌檔案的基础檔案名

**返回值：**

-   `io.Writer`：轮转檔案写入器

**範例：**

```go title="按小時轮转日誌"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// 建立类似的檔案：app-2024010115.log, app-2024010116.log 等
```

### 非同步写入器

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

为高性能日誌記錄建立非同步写入器。

**参数：**

-   `writer`：底层写入器
-   `bufferSize`：内部缓冲区大小

**返回值：**

-   `*AsyncWriter`：非同步写入器實例

**方法：**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**範例：**

```go title="非同步写入器"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## 上下文日誌

### LoggerWithCtx 接口

```go
type LoggerWithCtx interface {
    Trace(ctx context.Context, v ...any)
    Tracef(ctx context.Context, format string, v ...any)
    Debug(ctx context.Context, v ...any)
    Debugf(ctx context.Context, format string, v ...any)
    Info(ctx context.Context, v ...any)
    Infof(ctx context.Context, format string, v ...any)
    Warn(ctx context.Context, v ...any)
    Warnf(ctx context.Context, format string, v ...any)
    Error(ctx context.Context, v ...any)
    Errorf(ctx context.Context, format string, v ...any)
    Fatal(ctx context.Context, v ...any)
    Fatalf(ctx context.Context, format string, v ...any)
    Panic(ctx context.Context, v ...any)
    Panicf(ctx context.Context, format string, v ...any)
}
```

### 上下文函式

```go
func SetTrace(traceID string)
func GetTrace() string
```

设置和获取当前 goroutine 的追蹤 ID。

**範例：**

```go
log.SetTrace("trace-123-456")
log.Info("此訊息将包含追蹤 ID")

traceID := log.GetTrace()
fmt.Println("当前追蹤 ID:", traceID)
```

## 構建標籤

该库支援使用構建標籤进行条件编译：

:::info 構建標籤说明
構建標籤通过 `go build -tags` 参数指定，不同標籤会改变日誌库的编译行为和运行時特性。选择合适的標籤可以在开发便利性和生产性能之间取得平衡。
:::

### 預設模式

```bash
go build
```

-   启用完整功能
-   包含调试訊息
-   标准性能

### 调试模式

```bash
go build -tags debug
```

-   增强的调试信息
-   额外的运行時检查
-   详细的呼叫者信息

### 发布模式

```bash
go build -tags release
```

-   为生产环境優化
-   调试訊息被停用
-   启用自动日誌轮转

### 丢弃模式

```bash
go build -tags discard
```

-   最大性能
-   所有日誌操作都是空操作
-   零开销

### 组合模式

```bash
go build -tags "debug,discard"    # 调试与丢弃
go build -tags "release,discard"  # 发布与丢弃
```

## 性能優化

:::tip 性能最佳实践
该库通过对象池、級別检查前置和非同步写入等机制进行了深度性能優化。在高吞吐量场景下，建议组合使用非同步写入器和适当的構建標籤以获得最佳性能。
:::

### 对象池

该库在内部使用 `sync.Pool` 来管理：

-   日誌條目对象
-   字节缓冲区
-   格式化器缓冲区

这减少了高吞吐量场景下的垃圾收集压力。

### 級別检查

日誌級別检查发生在昂贵操作之前：

```go
// 高效 - 仅当級別启用時才进行訊息格式化
logger.Debugf("昂贵操作结果: %+v", expensiveCall())

// 在生产环境中调试被停用時效率较低
result := expensiveCall()
logger.Debug("结果:", result)
```

### 非同步写入

对于高吞吐量应用程序：

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // 大缓冲区
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### 構建標籤優化

根据环境使用适当的構建標籤：

-   开发：預設或调试標籤
-   生产：发布標籤
-   性能关键：丢弃標籤

## 示例

### 基本用法

```go title="基本用法"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("应用程序启动中")
    log.Warn("这是一個警告")
    log.Error("这是一個错误")
}
```

### 自定義日誌記錄器

```go title="自定義配置"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // 配置日誌記錄器
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  // 停用呼叫者信息
    logger.SetPrefixMsg("[MyApp] ")

    // 设置输出到檔案
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("自定義日誌記錄器已配置")
    logger.Debug("带呼叫者的调试信息")
}
```

### 高性能日誌記錄

```go title="高性能日誌記錄"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 建立轮转檔案写入器
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // 使用非同步写入器提升性能
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // 生产环境跳过 debug 日誌

    // 高吞吐量日誌記錄
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### 上下文感知日誌記錄

```go title="上下文感知日誌記錄"
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

### 自定義 JSON 格式化器

```go title="自定義 JSON 格式化器"
package main

import (
    "encoding/json"
    "os"
    "time"
    "github.com/lazygophers/log"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *log.Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339Nano),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "pid":       entry.Pid,
        "gid":       entry.Gid,
    }

    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }

    if entry.CallerName != "" {
        data["caller"] = map[string]interface{}{
            "function": entry.CallerName,
            "file":     entry.CallerFile,
            "line":     entry.CallerLine,
        }
    }

    jsonData, _ := json.MarshalIndent(data, "", "  ")
    return append(jsonData, '\n')
}

func main() {
    logger := log.New()
    logger.SetFormatter(&JSONFormatter{})
    logger.Caller(true)  // 停用呼叫者信息
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("JSON格式化訊息")
}
```

## 错误处理

:::warning 注意
出于性能考虑，大多数日誌記錄器方法不返回错误。如果写入失败，日誌将被静默丢弃。如果需要错误感知能力，请使用自定義写入器。
:::

如果您需要对输出操作进行错误处理，请实现自定義写入器：

```go title="错误捕获写入器"
type ErrorCapturingWriter struct {
    writer io.Writer
    lastError error
}

func (w *ErrorCapturingWriter) Write(data []byte) (int, error) {
    n, err := w.writer.Write(data)
    if err != nil {
        w.lastError = err
    }
    return n, err
}

func (w *ErrorCapturingWriter) LastError() error {
    return w.lastError
}
```

## 執行緒安全

:::tip 并发安全
所有 `Logger` 實例的方法都是執行緒安全的，可以在多个 goroutine 中并发使用，无需额外的同步机制。但请注意，单个 `Entry` 对象**不是**執行緒安全的，属于一次性使用。
:::

---

## 🌍 多语言文档

本檔案提供多种语言版本：

-   [🇺🇸 English](../en/API.md)
-   [🇨🇳 简体中文](API.md)（当前）
-   [🇹🇼 繁體中文](../zh-TW/API.md)
-   [🇫🇷 Français](../fr/API.md)
-   [🇷🇺 Русский](../ru/API.md)
-   [🇪🇸 Español](../es/API.md)
-   [🇸🇦 العربية](../ar/API.md)

---

**LazyGophers Log 完整 API 参考 - 使用卓越的日誌記錄構建更好的应用程序！🚀**
