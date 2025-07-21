# lazygophers/log

[![AGPL-3.0 License](https://img.shields.io/badge/license-AGPL--3.0-blue)](LICENSE)

灵活可配置的Go日志库，支持多级别日志输出、自定义格式和输出目标。

## 功能特性

- 📊 多日志级别：Trace/Debug/Info/Warn/Error/Fatal/Panic
- 🎨 自定义日志格式（支持完整格式配置）
- 📁 灵活输出目标（支持多个io.Writer）
- 🧵 协程ID追踪
- ⏱️ 时间戳记录
- 🔍 调用栈追踪（文件、行号、函数名）
- ⚡ 高性能（使用sync.Pool减少内存分配）

## 安装

```bash
go get github.com/lazygophers/log
```

## 快速开始

```go
package main

import "github.com/lazygophers/log"

func main() {
    // 设置日志级别
    log.SetLevel(log.InfoLevel)
    
    // 记录日志
    log.Info("Application started")
    log.Debug("This is debug message") // 不会被输出
    log.Warn("Something might be wrong")
    
    // 格式化日志
    log.Infof("User %s logged in", "Alice")
    
    // 记录错误
    log.Error("Failed to connect database")
}
```

## API参考

### 核心结构

- `log`: 日志记录器主体
  - `SetLevel(level Level)`: 设置日志级别
  - `SetOutput(writers ...io.Writer)`: 设置输出目标
  - `Clone()`: 创建副本
  - `SetCallerDepth(depth int)`: 设置调用栈深度

### 日志方法

- `Trace/Tracef`: TRACE级别
- `Debug/Debugf`: DEBUG级别
- `Info/Infof`: INFO级别
- `Warn/Warnf`: WARN级别
- `Error/Errorf`: ERROR级别
- `Fatal/Fatalf`: FATAL级别（触发os.Exit）
- `Panic/Panicf`: PANIC级别（触发panic）

### 配置方法

- `SetPrefixMsg()`: 设置消息前缀
- `SetSuffixMsg()`: 设置消息后缀
- `ParsingAndEscaping()`: 启用/禁用转义
- `Caller()`: 启用/禁用调用者信息

## 贡献指南

欢迎贡献！请遵循以下流程：

1. Fork仓库
2. 创建特性分支 (`git checkout -b feature/your-feature`)
3. 提交更改 (`git commit -am 'Add some feature'`)
4. 推送到分支 (`git push origin feature/your-feature`)
5. 创建Pull Request

## 许可证

本项目采用 [GNU Affero General Public License v3.0](LICENSE) 授权。
