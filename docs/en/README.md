# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A high-performance, flexible Go logging library built on zap, providing rich features and a simple API.

## 📖 Documentation Languages

-   [🇺🇸 English](README.md)
-   [🇨🇳 简体中文](docs/README_zh-CN.md)
-   [🇹🇼 繁體中文](docs/README_zh-TW.md)
-   [🇫🇷 Français](docs/README_fr.md)
-   [🇷🇺 Русский](docs/README_ru.md)
-   [🇪🇸 Español](docs/README_es.md)
-   [🇸🇦 العربية](docs/README_ar.md)

## 🚀 Online Documentation

Visit our [GitHub Pages documentation](https://lazygophers.github.io/log/) for a better reading experience.

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

```go
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

```go
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

## 🔧 Configuration Options

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

| Feature          | lazygophers/log | zap    | logrus | standard log |
| ---------------- | --------------- | ------ | ------ | ------------ |
| Performance      | High            | High   | Medium | Low          |
| API Simplicity   | High            | Medium | High   | High         |
| Feature Richness | Medium          | High   | High   | Low          |
| Flexibility      | Medium          | High   | High   | Low          |
| Learning Curve   | Low             | Medium | Medium | Low          |

## 🔗 Related Documentation

-   [📚 API Documentation](docs/API.md) - Complete API reference
-   [🤝 Contributing Guide](docs/CONTRIBUTING.md) - How to contribute
-   [📋 Changelog](CHANGELOG.md) - Version history
-   [🔒 Security Policy](docs/SECURITY.md) - Security guidelines
-   [📜 Code of Conduct](docs/CODE_OF_CONDUCT.md) - Community guidelines

## 🚀 Getting Help

-   **GitHub Issues**: [Report bugs or request features](https://github.com/lazygophers/log/issues)
-   **GoDoc**: [API Documentation](https://pkg.go.dev/github.com/lazygophers/log)
-   **Examples**: [Usage examples](https://github.com/lazygophers/log/tree/main/examples)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🌍 Multilingual Documentation

This document is available in multiple languages:

-   [🇺🇸 English](README.md) (Current)
-   [🇨🇳 简体中文](docs/README_zh-CN.md)
-   [🇹🇼 繁體中文](docs/README_zh-TW.md)
-   [🇫🇷 Français](docs/README_fr.md)
-   [🇷🇺 Русский](docs/README_ru.md)
-   [🇪🇸 Español](docs/README_es.md)
-   [🇸🇦 العربية](docs/README_ar.md)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](docs/CONTRIBUTING.md) for details.

---

**lazygophers/log** is designed to be the go-to logging solution for Go developers who value both performance and simplicity. Whether you're building a small utility or a large-scale distributed system, this library provides the right balance of features and ease of use.
