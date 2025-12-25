---
titleSuffix: ' | LazyGophers Log'
---

# LazyGophers 日志库

一个高性能的 Go 语言日志库

简洁的 API、卓越的性能和灵活的配置

<div align="center" style="margin: 2rem 0">
  <img src="/log/public/logo.svg" alt="LazyGophers Log Logo" style="width: 200px; margin-bottom: 1rem">
  
  <div style="margin: 1rem 0; display: flex; align-items: center; gap: 10px">
    <a href="https://github.com/lazygophers/log/actions/workflows/ci.yml"><img src="https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg" alt="CI Status"></a>
    <a href="https://goreportcard.com/report/github.com/lazygophers/log"><img src="https://goreportcard.com/badge/github.com/lazygophers/log" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/lazygophers/log"><img src="https://pkg.go.dev/badge/github.com/lazygophers/log.svg" alt="Go Reference"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  </div>
  
  <div style={{ margin: '2rem 0' }}>
    <a href="#quick-start" style="padding: 10px 20px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px; margin: 0 10px">快速开始</a>
    <a href="/zh-CN/API" style="padding: 10px 20px; background-color: #6c757d; color: white; text-decoration: none; border-radius: 5px; margin: 0 10px">API 参考</a>
  </div>
</div>

## ✨ 核心特性

### 高性能
基于 zap 构建，采用对象池和条件字段记录技术，确保出色的性能表现

### 丰富的日志级别
支持 Trace、Debug、Info、Warn、Error、Fatal、Panic 七个日志级别

### 灵活的配置
支持日志级别控制、调用者信息记录、跟踪信息、自定义前缀后缀等

### 文件轮换
内置日志文件轮换功能，支持按小时自动轮换日志文件

### Zap 兼容性
与 zap WriteSyncer 无缝集成，支持自定义输出目标

### 简洁的 API
API 设计类似于标准日志库，易于使用和迁移

## 🚀 快速开始

### 安装

```bash
go get github.com/lazygophers/log
```

### 基本使用

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 使用默认全局日志器
    log.Debug("Debug 消息")
    log.Info("Info 消息")
    log.Warn("Warning 消息")
    log.Error("Error 消息")

    // 使用格式化输出
    log.Infof("用户 %s 登录成功", "admin")

    // 自定义配置
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("这是来自自定义日志器的消息")
}
```

## 📚 文档导航

| 文档 | 描述 |
|------|------|
| [API 参考](/zh-CN/API) | 详细的 API 文档 |
| [版本历史](/zh-CN/CHANGELOG) | 查看所有版本更新记录 |
| [贡献指南](/zh-CN/CONTRIBUTING) | 如何为项目贡献代码 |
| [行为准则](/zh-CN/CODE_OF_CONDUCT) | 社区行为规范 |
| [安全政策](/zh-CN/SECURITY) | 安全漏洞报告流程 |

## 🌍 多语言文档

- [🇺🇸 English](/)
- [🇨🇳 简体中文](/zh-CN/)
- [🇹🇼 繁體中文](/zh-TW/)

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](/zh-CN/LICENSE) 文件了解详情。

## 🤝 贡献

我们欢迎贡献！请查看 [贡献指南](/zh-CN/CONTRIBUTING) 了解详情。

---

<div align="center">
  <p><strong>LazyGophers Log</strong> 旨在成为 Go 开发者的首选日志解决方案，既注重性能又注重易用性。无论您是构建小型工具还是大型分布式系统，这个库都能提供恰到好处的功能和易用性平衡。</p>
  <a href="https://github.com/lazygophers/log" style="display: inline-block; padding: 12px 24px; background-color: #24292e; color: white; text-decoration: none; border-radius: 6px; margin: 1rem 0">
    <span style="margin-right: 8px">⭐</span>Star on GitHub
  </a>
</div>