---
titleSuffix: " | LazyGophers Log"
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A high-performance, flexible Go logging library built on zap, providing rich features and a simple API.

## 📖 Documentation Languages

-   [🇺🇸 English](README.md) (Current)
-   [🇨🇳 简体中文](../zh-CN/README.md)
-   [🇹🇼 繁體中文](../zh-TW/README.md)
-   [🇫🇷 Français](../fr/README.md)
-   [🇪🇸 Español](../es/README.md)
-   [🇷🇺 Русский](../ru/README.md)
-   [🇸🇦 العربية](../ar/README.md)
-   [🇯🇵 日本語](../ja/README.md)
-   [🇩🇪 Deutsch](../de/README.md)
-   [🇰🇷 한국어](../ko/README.md)
-   [🇵🇹 Português](../pt/README.md)
-   [🇳🇱 Nederlands](../nl/README.md)
-   [🇵🇱 Polski](../pl/README.md)
-   [🇮🇹 Italiano](../it/README.md)
-   [🇹🇷 Türkçe](../tr/README.md)

## ✨ Features

-   **🚀 High Performance**: Built on zap with object pooling and conditional field recording
-   **📊 Rich Log Levels**: Trace, Debug, Info, Warn, Error, Fatal, Panic levels
-   **⚙️ Flexible Configuration**:
    -   Log level control
    -   Caller information recording
    -   Trace information (including goroutine ID)
    -   Custom log prefixes and suffixes
    -   Custom output targets (console, files, etc.)
    -   Log formatting options
-   **🔄 File Rotation**: Hourly log file rotation support
-   **🔌 Zap Compatibility**: Seamless integration with zap WriteSyncer
-   **🎯 Simple API**: Clean API similar to standard log library, easy to use

## 🚀 Quick Start

### Installation

:::tip Installation
```bash
go get github.com/lazygophers/log
```
:::

### Basic Usage

```go title="Quick Start"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Use default global logger
    log.Debug("Debug message")
    log.Info("Info message")
    log.Warn("Warning message")
    log.Error("Error message")

    // Use formatted output
    log.Infof("User %s logged in successfully", "admin")

    // Custom configuration
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("This is a log from custom logger")
}
```

## 📚 Advanced Usage

### Custom Logger with File Output

```go title="File Output Configuration"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Create logger with file output
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Debug message with caller info")
    logger.Info("Info message with trace info")
}
```

### Log Level Control

```go title="Log Level Control"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Only warn and above will be logged
    logger.Debug("This won't be logged")  // Ignored
    logger.Info("This won't be logged")   // Ignored
    logger.Warn("This will be logged")    // Logged
    logger.Error("This will be logged")   // Logged
}
```

## 🎯 Use Cases

### Applicable Scenarios

-   **Web Services and API Backends**: Request tracing, structured logging, performance monitoring
-   **Microservices Architecture**: Distributed tracing (TraceID), unified log format, high throughput
-   **Command-line Tools**: Level control, clean output, error reporting
-   **Batch Tasks**: File rotation, long-running, resource optimization

### Special Advantages

-   **Object Pool Optimization**: Entry and Buffer object reuse, reducing GC pressure
-   **Async Writing**: High throughput scenarios (10000+ logs/second) without blocking
-   **TraceID Support**: Request tracing in distributed systems, integration with OpenTelemetry
-   **Zero Configuration**: Ready to use, progressive configuration

## 🔧 Configuration Options

:::note Configuration Options
All the following methods support chaining and can be combined to build a custom Logger.
:::

### Logger Configuration

| Method                  | Description                | Default      |
| ----------------------- | -------------------------- | ------------ |
| `SetLevel(level)`       | Set minimum log level      | `DebugLevel` |
| `EnableCaller(enable)`  | Enable/disable caller info | `false`      |
| `EnableTrace(enable)`   | Enable/disable trace info  | `false`      |
| `SetCallerDepth(depth)` | Set caller depth           | `2`          |
| `SetPrefixMsg(prefix)`  | Set log prefix             | `""`         |
| `SetSuffixMsg(suffix)`  | Set log suffix             | `""`         |
| `SetOutput(writers...)` | Set output targets         | `os.Stdout`  |

### Log Levels

| Level        | Description                        |
| ------------ | ---------------------------------- |
| `TraceLevel` | Most verbose, for detailed tracing |
| `DebugLevel` | Debug information                  |
| `InfoLevel`  | General information                |
| `WarnLevel`  | Warning messages                   |
| `ErrorLevel` | Error messages                     |
| `FatalLevel` | Fatal errors (calls os.Exit(1))    |
| `PanicLevel` | Panic errors (calls panic())       |

## 🏗️ Architecture

### Core Components

-   **Logger**: Main logging structure with configurable options
-   **Entry**: Individual log record with comprehensive field support
-   **Level**: Log level definitions and utility functions
-   **Format**: Log formatting interface and implementations

### Performance Optimization

-   **Object Pooling**: Reuses Entry objects to reduce memory allocation
-   **Conditional Recording**: Only records expensive fields when needed
-   **Fast Level Checking**: Checks log level at the outermost layer
-   **Lock-Free Design**: Most operations don't require locks

## 📊 Performance Comparison

:::info Performance Comparison
The following data is based on benchmarks; actual performance may vary depending on environment and configuration.
:::

| Feature          | lazygophers/log | zap    | logrus | standard log |
| ---------------- | --------------- | ------ | ------ | ------------ |
| Performance      | High            | High   | Medium | Low          |
| API Simplicity   | High            | Medium | High   | High         |
| Feature Richness | Medium          | High   | High   | Low          |
| Flexibility      | Medium          | High   | High   | Low          |
| Learning Curve   | Low             | Medium | Medium | Low          |

## ❓ Frequently Asked Questions

### How to choose the right log level?

-   **Development Environment**: Use `DebugLevel` or `TraceLevel` for detailed information
-   **Production Environment**: Use `InfoLevel` or `WarnLevel` to reduce overhead
-   **Performance Testing**: Use `PanicLevel` to disable all logs

### How to optimize performance in production?

:::warning Note
In high throughput scenarios, it's recommended to combine async writing with reasonable log levels to optimize performance.
:::

1. Use `AsyncWriter` for async writing：

```go title="Async Writer Configuration"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Adjust log levels to avoid unnecessary logs：

```go title="Level Optimization"
logger.SetLevel(log.InfoLevel)  // Skip Debug and Trace
```

3. Use conditional logging to reduce overhead：

```go title="Conditional Logging"
if logger.Level >= log.DebugLevel {
    logger.Debug("Detailed debug info")
}
```

### What's the difference between `Caller` and `EnableCaller`?

-   **`EnableCaller(enable bool)`**：Controls whether the Logger collects caller information
    -   `EnableCaller(true)` enables caller information collection
-   **`Caller(disable bool)`**：Controls whether the Formatter outputs caller information
    -   `Caller(true)` disables caller information output

It's recommended to use `EnableCaller` for global control.

### How to implement a custom formatter?

Implement the `Format` interface：

```go title="Custom Formatter"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Related Documentation

-   [📚 API Documentation](API.md) - Complete API reference
-   [🤝 Contributing Guide](/en/CONTRIBUTING) - How to contribute
-   [📋 Changelog](/en/CHANGELOG) - Version history
-   [🔒 Security Policy](/en/SECURITY) - Security guidelines
-   [📜 Code of Conduct](/en/CODE_OF_CONDUCT) - Community guidelines

## 🚀 Getting Help

-   **GitHub Issues**：[Report bugs or request features](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[API Documentation](https://pkg.go.dev/github.com/lazygophers/log)
-   **Examples**：[Usage examples](https://github.com/lazygophers/log/tree/main/examples)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](/en/LICENSE) file for details.

## 🌟 Multilingual Documentation

This document is available in multiple languages：

-   [🇺🇸 English](README.md) (Current)
-   [🇨🇳 简体中文](README.md)
-   [🇹🇼 繁體中文](README.md)
-   [🇫🇷 Français](README.md)
-   [🇷🇺 Русский](README.md)
-   [🇪🇸 Español](README.md)
-   [🇸🇦 العربية](README.md)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](/en/CONTRIBUTING) for details.

---

**lazygophers/log** is designed to be the go-to logging solution for Go developers who value both performance and simplicity. Whether you're building a small utility or a large-scale distributed system, this library provides the right balance of features and ease of use.
