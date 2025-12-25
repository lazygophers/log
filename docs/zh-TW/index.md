---
titleSuffix: " | LazyGophers Log"
---

# LazyGophers 日誌庫

一個高性能的 Go 語言日誌庫

簡潔的 API、卓越的性能和靈活的配置

<div align="center" style="margin: 2rem 0">
  <img src="/log/public/logo.svg" alt="LazyGophers Log Logo" style="width: 200px; margin-bottom: 1rem">
  
  <div style="margin: 1rem 0; display: flex; align-items: center; gap: 10px">
    <a href="https://github.com/lazygophers/log/actions/workflows/ci.yml"><img src="https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg" alt="CI Status"></a>
    <a href="https://goreportcard.com/report/github.com/lazygophers/log"><img src="https://goreportcard.com/badge/github.com/lazygophers/log" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/lazygophers/log"><img src="https://pkg.go.dev/badge/github.com/lazygophers/log.svg" alt="Go Reference"></a>
    <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  </div>
  
  <div style={{ margin: '2rem 0' }}>
    <a href="#quick-start" style="padding: 10px 20px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px; margin: 0 10px">快速開始</a>
    <a href="/zh-TW/API" style="padding: 10px 20px; background-color: #6c757d; color: white; text-decoration: none; border-radius: 5px; margin: 0 10px">API 參考</a>
  </div>
</div>

## ✨ 核心特性

### 高性能
基於 zap 構建，採用物件池和條件欄位記錄技術，確保出色的性能表現

### 豐富的日誌級別
支援 Trace、Debug、Info、Warn、Error、Fatal、Panic 七個日誌級別

### 靈活的配置
支援日誌級別控制、呼叫者資訊記錄、跟蹤資訊、自定義前綴後綴等

### 文件輪換
內置日誌文件輪換功能，支援按小時自動輪換日誌文件

### Zap 兼容性
與 zap WriteSyncer 無縫集成，支援自定義輸出目標

### 簡潔的 API
API 設計類似於標準日誌庫，易於使用和遷移

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

| 文檔 | 描述 |
|------|------|
| [API 參考](/zh-TW/API) | 詳細的 API 文檔 |
| [版本歷史](/zh-TW/CHANGELOG) | 查看所有版本更新記錄 |
| [貢獻指南](/zh-TW/CONTRIBUTING) | 如何為項目貢獻代碼 |
| [行為準則](/zh-TW/CODE_OF_CONDUCT) | 社區行為規範 |
| [安全政策](/zh-TW/SECURITY) | 安全漏洞報告流程 |

## 🌍 多語言文檔

- [🇺🇸 English](/)
- [🇨🇳 簡體中文](/zh-CN/)
- [🇹🇼 繁體中文](/zh-TW/)

## 📄 許可證

本項目採用 MIT 許可證 - 查看 [LICENSE](/zh-TW/LICENSE) 文件了解詳情。

## 🤝 貢獻

我們歡迎貢獻！請查看 [貢獻指南](/zh-TW/CONTRIBUTING) 了解詳情。

---

<div align="center">
  <p><strong>LazyGophers Log</strong> 旨在成為 Go 開發者的首選日誌解決方案，既注重性能又注重易用性。無論您是構建小型工具還是大型分布式系統，這個庫都能提供恰到好處的功能和易用性平衡。</p>
  <a href="https://github.com/lazygophers/log" style="display: inline-block; padding: 12px 24px; background-color: #24292e; color: white; text-decoration: none; border-radius: 6px; margin: 1rem 0">
    <span style="margin-right: 8px">⭐</span>Star on GitHub
  </a>
</div>
