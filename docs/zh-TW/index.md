---
titleSuffix: " | LazyGophers Log"
---

# LazyGophers 日誌庫

一個高性能的 Go 語言日誌庫

簡潔的 API、卓越的性能和靈活的配置

![LazyGophers Log Logo](/log/public/logo.svg)

[![CI Status](https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg)](https://github.com/lazygophers/log/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/log.svg)](https://pkg.go.dev/github.com/lazygophers/log)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[快速開始](#quick-start) | [API 參考](/zh-TW/API)

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

-   [🇺🇸 English](/)
-   [🇨🇳 簡體中文](/zh-CN/)
-   [🇹🇼 繁體中文](/zh-TW/)

## 📄 許可證

本項目採用 MIT 許可證 - 查看 [LICENSE](/zh-TW/LICENSE) 文件了解詳情。

## 🤝 貢獻

我們歡迎貢獻！請查看 [貢獻指南](/zh-TW/CONTRIBUTING) 了解詳情。

---

**LazyGophers Log** 旨在成為 Go 開發者的首選日誌解決方案，既注重性能又注重易用性。無論您是構建小型工具還是大型分布式系統，這個庫都能提供恰到好處的功能和易用性平衡。

[⭐ Star on GitHub](https://github.com/lazygophers/log)