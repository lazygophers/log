<div align="center">
  <img src="https://raw.githubusercontent.com/lazygophers/log/main/docs/public/logo.svg" alt="LazyGophers Log Logo" width="200" height="200">
</div>

# LazyGophers Log

<div align="center">
  <h2>🚀 高性能、易用的 Go 语言日志库</h2>
  <p>基于 zap 构建，提供丰富特性和简洁 API</p>
  <div>
    <a href="https://github.com/lazygophers/log/actions/workflows/ci.yml"><img src="https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg" alt="CI Status"></a>
    <a href="https://goreportcard.com/report/github.com/lazygophers/log"><img src="https://goreportcard.com/badge/github.com/lazygophers/log" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/lazygophers/log"><img src="https://pkg.go.dev/badge/github.com/lazygophers/log.svg" alt="Go Reference"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  </div>
</div>

## ✨ 核心特性

<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mt-12">
  <div class="card">
    <div class="card-content">
      <h3>🚀 高性能</h3>
      <p>基于 zap 构建，采用对象池和条件字段记录技术，确保出色的性能表现</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>📊 丰富的日志级别</h3>
      <p>支持 Trace、Debug、Info、Warn、Error、Fatal、Panic 七个日志级别</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>⚙️ 灵活的配置</h3>
      <p>支持日志级别控制、调用者信息记录、跟踪信息、自定义前缀后缀等</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🔄 文件轮换</h3>
      <p>内置日志文件轮换功能，支持按小时自动轮换日志文件</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🔌 Zap 兼容性</h3>
      <p>与 zap WriteSyncer 无缝集成，支持自定义输出目标</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🎯 简洁的 API</h3>
      <p>API 设计类似于标准日志库，易于使用和迁移</p>
    </div>
  </div>
</div>

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

<div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-12">
  <div class="card">
    <div class="card-content">
      <h3>📖 核心文档</h3>
      <ul>
        <li><a href="/API">API 参考</a></li>
        <li><a href="/CHANGELOG">版本历史</a></li>
      </ul>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🤝 社区与贡献</h3>
      <ul>
        <li><a href="/CONTRIBUTING">贡献指南</a></li>
        <li><a href="/CODE_OF_CONDUCT">行为准则</a></li>
        <li><a href="/SECURITY">安全政策</a></li>
      </ul>
    </div>
  </div>
</div>

## 🌍 多语言文档

<div class="flex flex-wrap gap-4 mt-12">
  <a href="/" class="btn btn-primary">🇺🇸 English</a>
  <a href="/index_zh-CN" class="btn btn-secondary">🇨🇳 简体中文</a>
  <a href="/index_zh-TW" class="btn btn-secondary">🇹🇼 繁體中文</a>
</div>

## 📄 许可证

本项目采用 MIT 许可证 - 查看 <a href="/LICENSE">LICENSE</a> 文件了解详情。

## 🤝 贡献

我们欢迎贡献！请查看 <a href="/CONTRIBUTING">贡献指南</a> 了解详情。

---

<div align="center">
  <p><strong>LazyGophers Log</strong> 旨在成为 Go 开发者的首选日志解决方案，既注重性能又注重易用性。无论您是构建小型工具还是大型分布式系统，这个库都能提供恰到好处的功能和易用性平衡。</p>
  <a href="https://github.com/lazygophers/log" class="btn btn-large btn-primary mt-4">
    <span class="icon-github mr-2"></span>Star on GitHub
  </a>
</div>