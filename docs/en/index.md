---
pageType: custom
---
import { HomeHero } from '@theme';

<HomeHero
  name="hero.name"
  text="hero.text"
  tagline="hero.tagline"
  image={{
    src: 'https://raw.githubusercontent.com/lazygophers/log/main/docs/public/logo.svg',
    alt: 'LazyGophers Log Logo'
  }}
  actions={[
    {
      text: 'hero.gettingStarted',
      link: '#quick-start',
      theme: 'brand'
    },
    {
      text: 'hero.apiReference',
      link: '/API',
      theme: 'alt'
    }
  ]}
/>

<div align="center" style={{ margin: '2rem 0' }}>
  <a href="https://github.com/lazygophers/log/actions/workflows/ci.yml"><img src="https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg" alt="CI Status"></a>
  <a href="https://goreportcard.com/report/github.com/lazygophers/log"><img src="https://goreportcard.com/badge/github.com/lazygophers/log" alt="Go Report Card" style={{ margin: '0 10px' }}></a>
  <a href="https://pkg.go.dev/github.com/lazygophers/log"><img src="https://pkg.go.dev/badge/github.com/lazygophers/log.svg" alt="Go Reference" style={{ margin: '0 10px' }}></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
</div>

import { HomeFeature } from '@theme';

## ✨ 核心特性

<HomeFeature
  features={[
    {
      title: 'feature.highPerformance',
      details: 'feature.highPerformance.desc',
      icon: '🚀',
      span: 3
    },
    {
      title: 'feature.richLevels',
      details: 'feature.richLevels.desc',
      icon: '📊',
      span: 3
    },
    {
      title: 'feature.flexibleConfig',
      details: 'feature.flexibleConfig.desc',
      icon: '⚙️',
      span: 3
    },
    {
      title: 'feature.fileRotation',
      details: 'feature.fileRotation.desc',
      icon: '🔄',
      span: 3
    },
    {
      title: 'feature.zapCompatibility',
      details: 'feature.zapCompatibility.desc',
      icon: '🔌',
      span: 3
    },
    {
      title: 'feature.simpleAPI',
      details: 'feature.simpleAPI.desc',
      icon: '🎯',
      span: 3
    }
  ]}
/>

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

import { LinkCard } from '@theme';

## 📚 文档导航

<div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))', gap: '1rem', margin: '2rem 0' }}>
  <LinkCard
    href="/API"
    title="API 参考"
    description="详细的 API 文档"
  />
  <LinkCard
    href="/CHANGELOG"
    title="版本历史"
    description="查看所有版本更新记录"
  />
  <LinkCard
    href="/CONTRIBUTING"
    title="贡献指南"
    description="如何为项目贡献代码"
  />
  <LinkCard
    href="/CODE_OF_CONDUCT"
    title="行为准则"
    description="社区行为规范"
  />
  <LinkCard
    href="/SECURITY"
    title="安全政策"
    description="安全漏洞报告流程"
  />
</div>

## 🌍 多语言文档

<div class="flex flex-wrap gap-4 mt-12">
  <a href="/" class="btn btn-primary">🇺🇸 English</a>
  <a href="/zh-CN/" class="btn btn-secondary">🇨🇳 简体中文</a>
  <a href="/zh-TW/" class="btn btn-secondary">🇹🇼 繁體中文</a>
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