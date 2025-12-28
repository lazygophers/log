---
titleSuffix: " | LazyGophers Log"
pageType: home

hero:
  name: LazyGophers Log
  text: A High-Performance Logging Library for Go
  tagline: Simple API, excellent performance, and flexible configuration
  image:
    src: /log/public/logo.svg
    alt: LazyGophers Log Logo
  actions:
    - theme: brand
      text: Getting Started
      link: /README
    - theme: alt
      text: API Reference
      link: /API
    - theme: alt
      text: View on GitHub
      link: https://github.com/lazygophers/log

features:
  - title: High Performance
    details: Built on zap with object pooling and conditional field recording technology to ensure excellent performance
    icon: ⚡
  - title: Rich Log Levels
    details: Supports seven log levels: Trace, Debug, Info, Warn, Error, Fatal, Panic
    icon: 📊
  - title: Flexible Configuration
    details: Supports log level control, caller information recording, trace information, custom prefix and suffix, etc.
    icon: ⚙️
  - title: File Rotation
    details: Built-in log file rotation function, supporting automatic hourly log file rotation
    icon: 🔄
  - title: Zap Compatibility
    details: Seamlessly integrates with zap WriteSyncer, supporting custom output targets
    icon: 🔗
  - title: Simple API
    details: API designed similar to the standard log library, easy to use and migrate
    icon: 🚀
  - title: Thread-Safe
    details: Lock-free design for most operations, ensuring thread safety without performance overhead
    icon: 🔒
  - title: Production Ready
    details: Battle-tested in production environments with comprehensive test coverage
    icon: ✅
---
