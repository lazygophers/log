---
pageType: home

hero:
  name: LazyGophers Log
  text: A high-performance, flexible Go logging library
  tagline: Built on zap with rich features and a simple API
  actions:
    - theme: brand
      text: Get Started
      link: /API
    - theme: alt
      text: View on GitHub
      link: https://github.com/lazygophers/log

features:
  - title: 'High Performance'
    details: Built on zap with object pooling and conditional field recording for optimal performance
    icon: 🚀
  - title: 'Rich Log Levels'
    details: Supports Trace, Debug, Info, Warn, Error, Fatal, and Panic levels
    icon: 📊
  - title: 'Flexible Configuration'
    details: Customize log levels, caller info, trace info, prefixes, suffixes, and output targets
    icon: ⚙️
  - title: 'File Rotation'
    details: Built-in hourly log file rotation support
    icon: 🔄
  - title: 'Zap Compatibility'
    details: Seamless integration with zap WriteSyncer
    icon: 🔌
  - title: 'Simple API'
    details: Clean API similar to standard log library, easy to use and integrate
    icon: 🎯
---

## Quick Start

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

## Documentation

- [API Reference](API.md) - Complete API documentation
- [Changelog](CHANGELOG.md) - Version history
- [Contributing](CONTRIBUTING.md) - How to contribute
- [Security Policy](SECURITY.md) - Security guidelines
- [Code of Conduct](CODE_OF_CONDUCT.md) - Community guidelines

## Performance Comparison

| Feature          | lazygophers/log | zap    | logrus | standard log |
| ---------------- | --------------- | ------ | ------ | ------------ |
| Performance      | High            | High   | Medium | Low          |
| API Simplicity   | High            | Medium | High   | High         |
| Feature Richness | Medium          | High   | High   | Low          |
| Flexibility      | Medium          | High   | High   | Low          |
| Learning Curve   | Low             | Medium | Medium | Low          |

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
