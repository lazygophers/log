<div align="center">
  <img src="https://raw.githubusercontent.com/lazygophers/log/main/docs/public/logo.svg" alt="LazyGophers Log Logo" width="200" height="200">
</div>

# LazyGophers Log

<div align="center">
  <h2>🚀 高性能、易用的 Go 語言日誌庫</h2>
  <p>基於 zap 構建，提供豐富特性和簡潔 API</p>
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
      <p>基於 zap 構建，採用對象池和條件字段記錄技術，確保出色的性能表現</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>📊 豐富的日誌級別</h3>
      <p>支持 Trace、Debug、Info、Warn、Error、Fatal、Panic 七個日誌級別</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>⚙️ 靈活的配置</h3>
      <p>支持日誌級別控制、調用者信息記錄、跟蹤信息、自定義前綴後綴等</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🔄 文件輪換</h3>
      <p>內置日誌文件輪換功能，支持按小時自動輪換日誌文件</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🔌 Zap 兼容性</h3>
      <p>與 zap WriteSyncer 無縫集成，支持自定義輸出目標</p>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🎯 簡潔的 API</h3>
      <p>API 設計類似於標準日誌庫，易於使用和遷移</p>
    </div>
  </div>
</div>

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

<div class="grid grid-cols-1 md:grid-cols-2 gap-4 mt-12">
  <div class="card">
    <div class="card-content">
      <h3>📖 核心文檔</h3>
      <ul>
        <li><a href="/API_zh-TW">API 參考</a></li>
        <li><a href="/CHANGELOG">版本歷史</a></li>
      </ul>
    </div>
  </div>

  <div class="card">
    <div class="card-content">
      <h3>🤝 社區與貢獻</h3>
      <ul>
        <li><a href="/CONTRIBUTING">貢獻指南</a></li>
        <li><a href="/CODE_OF_CONDUCT">行為準則</a></li>
        <li><a href="/SECURITY">安全政策</a></li>
      </ul>
    </div>
  </div>
</div>

## 🌍 多語言文檔

<div class="flex flex-wrap gap-4 mt-12">
  <a href="/" class="btn btn-secondary">🇺🇸 English</a>
  <a href="/index_zh-CN" class="btn btn-secondary">🇨🇳 簡體中文</a>
  <a href="/index_zh-TW" class="btn btn-primary">🇹🇼 繁體中文</a>
</div>

## 📄 許可證

本項目採用 MIT 許可證 - 查看 <a href="/LICENSE">LICENSE</a> 文件了解詳情。

## 🤝 貢獻

我們歡迎貢獻！請查看 <a href="/CONTRIBUTING">貢獻指南</a> 了解詳情。

---

<div align="center">
  <p><strong>LazyGophers Log</strong> 旨在成為 Go 開發者的首選日誌解決方案，既注重性能又注重易用性。無論您是構建小型工具還是大型分布式系統，這個庫都能提供恰到好處的功能和易用性平衡。</p>
  <a href="https://github.com/lazygophers/log" class="btn btn-large btn-primary mt-4">
    <span class="icon-github mr-2"></span>Star on GitHub
  </a>
</div>