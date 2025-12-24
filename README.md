# lazygophers/log

高性能、灵活的 Go 日志库，基于 zap 构建，提供丰富的功能和简洁的 API。

## 特性

-   **高性能**：基于 zap 构建，采用对象池复用 Entry 对象，减少内存分配
-   **丰富的日志级别**：支持 Trace、Debug、Info、Warn、Error、Fatal、Panic 级别
-   **灵活的配置**：
    -   日志级别控制
    -   调用者信息记录
    -   追踪信息记录（包含 goroutine ID）
    -   自定义日志前缀和后缀
    -   自定义输出目标（控制台、文件等）
    -   日志格式化选项
-   **文件轮转**：支持按小时轮转日志文件
-   **与 zap 兼容**：可以无缝集成 zap 的 WriteSyncer
-   **简洁的 API**：提供类似标准库 log 的简洁 API，易于使用

## 安装

```bash
go get github.com/lazygophers/log
```

## 快速开始

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 使用默认配置的全局 logger
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

    customLogger.Info("这是自定义 logger 的日志")
}
```

## 配置

### 日志级别

```go
// 设置全局日志级别
log.SetLevel(log.InfoLevel)

// 检查日志级别是否启用
if log.Enabled(log.DebugLevel) {
    log.Debug("调试信息")
}
```

支持的日志级别（从低到高）：

-   TraceLevel
-   DebugLevel
-   InfoLevel
-   WarnLevel
-   ErrorLevel
-   FatalLevel
-   PanicLevel

### 调用者信息

```go
// 启用调用者信息（默认启用）
log.EnableCaller(true)

// 设置调用者深度
log.SetCallerDepth(2)
```

### 追踪信息

```go
// 启用追踪信息（默认启用）
log.EnableTrace(true)
// 追踪信息包含 goroutine ID 和 Trace ID
```

### 日志前缀和后缀

```go
// 设置日志前缀
log.SetPrefixMsg("[MyApp]")

// 设置日志后缀
log.SetSuffixMsg("[end]")
```

### 输出目标

```go
// 设置输出到控制台
log.SetOutput(os.Stdout)

// 设置输出到多个目标
log.SetOutput(os.Stdout, os.Stderr)

// 设置输出到文件（带轮转）
log.SetOutput(log.GetOutputWriterHourly("/var/log/myapp.log"))
```

### 日志格式化

```go
// 禁用日志内容解析和转义（默认启用）
log.ParsingAndEscaping(true)
```

## 高级用法

### 克隆 logger

```go
// 克隆 logger 并修改配置
childLogger := log.New().Clone().SetPrefixMsg("[Child]")
```

### 上下文 logger

```go
// 创建带上下文的 logger
ctxLogger := log.Ctx(ctx)
ctxLogger.Info("带上下文的日志")
```

### 与 zap 集成

```go
// 使用 zap 的 WriteSyncer
zapWriter := zapcore.AddSync(os.Stdout)
log.SetOutput(zapWriter)
```

## 性能优化

-   使用对象池复用 Entry 对象，减少内存分配
-   条件性记录昂贵的字段（如调用者信息、追踪信息）
-   支持日志级别过滤，避免不必要的日志处理

## 测试

运行测试：

```bash
make test
```

运行性能测试：

```bash
make bench
```

## 贡献

欢迎提交 issues 和 pull requests！

## 许可证

MIT License

## 版本历史

查看 [CHANGELOG.md](CHANGELOG.md) 获取版本历史。
