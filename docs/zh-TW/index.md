---
pageType: custom
titleSuffix: " | LazyGophers Log"
---

import { HomeHero } from '@theme';
import { HomeFeature } from '@theme';
import { LinkCard } from '@theme';

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
      link: '/zh-TW/API',
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

## 🚀 快速開始

### 安裝

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
    // 使用默認全局日誌器
    log.Debug("Debug 消息")
    log.Info("Info 消息")
    log.Warn("Warning 消息")
    log.Error("Error 消息")

    // 使用格式化輸出
    log.Infof("用戶 %s 登錄成功", "admin")

    // 自定義配置
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("這是來自自定義日誌器的消息")
}
```

## 📚 文檔導航

<div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))', gap: '1rem', margin: '2rem 0' }}>
  <LinkCard
    href="/zh-TW/API"
    title="API 參考"
    description="詳細的 API 文檔"
  />
  <LinkCard
    href="/zh-TW/CHANGELOG"
    title="版本歷史"
    description="查看所有版本更新記錄"
  />
  <LinkCard
    href="/zh-TW/CONTRIBUTING"
    title="貢獻指南"
    description="如何為項目貢獻代碼"
  />
  <LinkCard
    href="/zh-TW/CODE_OF_CONDUCT"
    title="行為準則"
    description="社區行為規範"
  />
  <LinkCard
    href="/zh-TW/SECURITY"
    title="安全政策"
    description="安全漏洞報告流程"
  />
</div>

## 🌍 多語言文檔

<div class="flex flex-wrap gap-4 mt-12">
  <a href="/" class="btn btn-secondary">🇺🇸 English</a>
  <a href="/zh-CN/" class="btn btn-secondary">🇨🇳 簡體中文</a>
  <a href="/zh-TW/" class="btn btn-primary">🇹🇼 繁體中文</a>
</div>

## 📄 許可證

本項目採用 MIT 許可證 - 查看 <a href="/zh-TW/LICENSE">LICENSE</a> 文件了解詳情。

## 🤝 貢獻

我們歡迎貢獻！請查看 <a href="/zh-TW/CONTRIBUTING">貢獻指南</a> 了解詳情。

---

<div align="center">
  <p><strong>LazyGophers Log</strong> 旨在成為 Go 開發者的首選日誌解決方案，既注重性能又注重易用性。無論您是構建小型工具還是大型分布式系統，這個庫都能提供恰到好處的功能和易用性平衡。</p>
  <a href="https://github.com/lazygophers/log" class="btn btn-large btn-primary mt-4">
    <span class="icon-github mr-2"></span>Star on GitHub
  </a>
</div>
