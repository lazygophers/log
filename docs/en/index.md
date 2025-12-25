---
titleSuffix: " | LazyGophers Log"
---

# LazyGophers Log

A High-Performance Logging Library for Go

Simple API, excellent performance, and flexible configuration

![LazyGophers Log Logo](/log/public/logo.svg)

[![CI Status](https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg)](https://github.com/lazygophers/log/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/log.svg)](https://pkg.go.dev/github.com/lazygophers/log)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[Getting Started](#quick-start) | [API Reference](/API)

## ✨ Core Features

### High Performance

Built on zap, using object pooling and conditional field recording technology to ensure excellent performance

### Rich Log Levels

Supports seven log levels: Trace, Debug, Info, Warn, Error, Fatal, Panic

### Flexible Configuration

Supports log level control, caller information recording, trace information, custom prefix and suffix, etc.

### File Rotation

Built-in log file rotation function, supporting automatic hourly log file rotation

### Zap Compatibility

Seamlessly integrates with zap WriteSyncer, supporting custom output targets

### Simple API

API designed similar to the standard log library, easy to use and migrate

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

    customLogger.Info("This is a message from a custom logger")
}
```

## 📚 Documentation Navigation

| Document                            | Description                              |
| ----------------------------------- | ---------------------------------------- |
| [API Reference](/API)               | Detailed API documentation               |
| [Changelog](/CHANGELOG)             | View all version update records          |
| [Contributing Guide](/CONTRIBUTING) | How to contribute code to the project    |
| [Code of Conduct](/CODE_OF_CONDUCT) | Community code of conduct                |
| [Security Policy](/SECURITY)        | Security vulnerability reporting process |

## 🌍 Multilingual Documentation

-   [🇺🇸 English](/)
-   [🇨🇳 Simplified Chinese](/zh-CN/)
-   [🇹🇼 Traditional Chinese](/zh-TW/)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.

## 🤝 Contributing

We welcome contributions! Please see the [Contributing Guide](/CONTRIBUTING) for details.

---

**LazyGophers Log** aims to be the preferred logging solution for Go developers, focusing on both performance and usability. Whether you're building small tools or large distributed systems, this library provides the perfect balance of functionality and ease of use.

[⭐ Star on GitHub](https://github.com/lazygophers/log)
