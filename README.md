# log

[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/log)

## 功能特性

- **高性能日志系统**：基于sync.Pool优化内存分配
- **多输出支持**：同时写入多个输出目标（文件/终端/网络）
- **灵活格式化**：支持JSON、文本、Zap集成等多种格式
- **分布式追踪**：内置trace ID上下文传递能力
- **多级日志控制**：Debug/Info/Warn/Error/Fatal级别管理

## 快速开始

```go
// 基础使用示例
logger.Info("系统启动", "组件", "main")
```

## 核心组件

### Logger

- `NewLogger()` - 创建日志记录器
- `SetLevel()` - 动态调整日志级别
- `With()` - 创建带上下文的子记录器

### 输出管理

- `AddOutput()` - 添加新输出目标

## 性能优化

- 使用sync.Pool管理Entry和缓冲区对象
- 异步写入模式(writer_async.go)降低延迟
- 零内存分配的格式化实现

## 贡献指南

1. 环境要求: Go 1.21+
2. 开发流程:
   - `go mod tidy` 更新依赖
   - 提交规范: `feat(module): description` / `fix(module): description`
