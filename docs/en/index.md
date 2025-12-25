---
titleSuffix: " | LazyGophers Log"
---

# LazyGophers Log

A High-Performance Logging Library for Go

Simple API, excellent performance, and flexible configuration

<div align="center" style="margin: 2rem 0">
  <img src="/log/public/logo.svg" alt="LazyGophers Log Logo" style="width: 200px; margin-bottom: 1rem">
  
  <div style="margin: 1rem 0; display: flex; align-items: center; gap: 10px">
    <a href="https://github.com/lazygophers/log/actions/workflows/ci.yml"><img src="https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg" alt="CI Status"></a>
    <a href="https://goreportcard.com/report/github.com/lazygophers/log"><img src="https://goreportcard.com/badge/github.com/lazygophers/log" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/lazygophers/log"><img src="https://pkg.go.dev/badge/github.com/lazygophers/log.svg" alt="Go Reference"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  </div>
  
  <div style="margin: 2rem 0">
    <a href="#quick-start" style="padding: 10px 20px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px; margin: 0 10px">Getting Started</a>
    <a href="/API" style="padding: 10px 20px; background-color: #6c757d; color: white; text-decoration: none; border-radius: 5px; margin: 0 10px">API Reference</a>
  </div>
</div>

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

<div align="center">
  <p><strong>LazyGophers Log</strong> aims to be the preferred logging solution for Go developers, focusing on both performance and usability. Whether you're building small tools or large distributed systems, this library provides the perfect balance of functionality and ease of use.</p>
  <a href="https://github.com/lazygophers/log" style="display: inline-block; padding: 12px 24px; background-color: #24292e; color: white; text-decoration: none; border-radius: 6px; margin: 1rem 0">
    <span style="margin-right: 8px">⭐</span>Star on GitHub
  </a>
</div>
