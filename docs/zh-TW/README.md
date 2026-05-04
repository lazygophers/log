---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一個高性能、靈活的 Go 日誌庫，基於 zap 構建，提供豐富的功能和簡潔的 API。

## 📖 文檔語言

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/)
-   [🇹🇼 繁體中文](README.md)（當前）
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/)
-   [🇵🇹 Português](https://lazygophers.github.io/log/pt/)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/)
-   [🇵🇱 Polski](https://lazygophers.github.io/log/pl/)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/)
-   [🇹🇷 Türkçe](https://lazygophers.github.io/log/tr/)

## ✨ 特性

-   **🚀 高性能**：基於 zap 構建，採用物件池重用 Entry 物件，減少記憶體分配
-   **📊 豐富的日誌級別**：支援 Trace、Debug、Info、Warn、Error、Fatal、Panic 級別
-   **⚙️ 靈活的配置**：
    -   日誌級別控制
    -   調用者資訊記錄
    -   追蹤資訊（包括 goroutine ID）
    -   自定義日誌前綴和後綴
    -   自定義輸出目標（控制台、文件等）
    -   日誌格式化選項
-   **🔄 文件輪轉**：支援每小時日誌文件輪轉
-   **🔌 Zap 相容性**：與 zap WriteSyncer 無縫集成
-   **🎯 簡潔的 API**：類似標準日誌庫的清晰 API，易於使用

## 🚀 快速開始

### 安裝

:::tip 安裝
```bash
go get github.com/lazygophers/log
```
:::

### 基本用法

```go title="快速開始"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 使用預設全域 logger
    log.Debug("調試資訊")
    log.Info("普通資訊")
    log.Warn("警告資訊")
    log.Error("錯誤資訊")

    // 使用格式化輸出
    log.Infof("用戶 %s 登錄成功", "admin")

    // 自定義配置
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("這是來自自定義 logger 的日誌")
}
```

## 📚 高級用法

### 帶檔案輸出的自定義 Logger

```go title="檔案輸出配置"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 創建帶檔案輸出的 logger
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("帶調用者資訊的調試日誌")
    logger.Info("帶追蹤資訊的普通日誌")
}
```

### 日誌級別控制

```go title="日誌級別控制"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // 只有 warn 及以上級別會被記錄
    logger.Debug("這不會被記錄")  // 忽略
    logger.Info("這不會被記錄")   // 忽略
    logger.Warn("這會被記錄")    // 記錄
    logger.Error("這會被記錄")   // 記錄
}
```

## 🎯 使用場景

### 適用場景

-   **Web 服務和 API 後端**：請求追蹤、結構化日誌、性能監控
-   **微服務架構**：分佈式追蹤（TraceID）、統一日誌格式、高吞吐量
-   **命令列工具**：級別控制、簡潔輸出、錯誤報告
-   **批處理任務**：檔案輪轉、長時間運行、資源優化

### 特別優勢

-   **物件池優化**：Entry 和 Buffer 物件重用，減少 GC 壓力
-   **非同步寫入**：高吞吐量場景（10000+ 條/秒）無阻塞
-   **TraceID 支援**：分佈式系統請求追蹤，與 OpenTelemetry 集成
-   **零配置啟動**：開箱即用，漸進式配置

## 🔧 配置選項

:::note 配置選項
以下方法均支援鏈式呼叫，可組合使用構建自定義 Logger。
:::

### Logger 配置

| 方法                  | 描述                | 預設值       |
| --------------------- | ------------------- | ----------- |
| `SetLevel(level)`       | 設定最低日誌級別     | `DebugLevel` |
| `EnableCaller(enable)`  | 啟用/禁用調用者資訊  | `false`      |
| `EnableTrace(enable)`   | 啟用/禁用追蹤資訊    | `false`      |
| `SetCallerDepth(depth)` | 設定調用者深度       | `2`          |
| `SetPrefixMsg(prefix)`  | 設定日誌前綴         | `""`         |
| `SetSuffixMsg(suffix)`  | 設定日誌後綴         | `""`         |
| `SetOutput(writers...)` | 設定輸出目標         | `os.Stdout`  |

### 日誌級別

| 級別        | 描述                        |
| ----------- | --------------------------- |
| `TraceLevel` | 最詳細，用於詳細追蹤        |
| `DebugLevel` | 調試資訊                    |
| `InfoLevel`  | 普通資訊                    |
| `WarnLevel`  | 警告資訊                    |
| `ErrorLevel` | 錯誤資訊                    |
| `FatalLevel` | 致命錯誤（呼叫 os.Exit(1)） |
| `PanicLevel` | 恐慌錯誤（呼叫 panic()）    |

## 🏗️ 架構

### 核心組件

-   **Logger**：主日誌結構，具有可配置的級別、輸出、格式化器和調用者深度
-   **Entry**：單個日誌記錄，包含全面的元資料支援
-   **Level**：日誌級別定義和工具函數
-   **Format**：日誌格式化介面和實現

### 性能優化

-   **物件池**：重用 Entry 物件以減少記憶體分配
-   **條件記錄**：僅在需要時記錄昂貴字段
-   **快速級別檢查**：在最外層檢查日誌級別
-   **無鎖設計**：大多數操作不需要鎖

## 📊 性能比較

:::info 性能對比
以下資料基於基準測試，實際性能可能因環境和配置不同而有所差異。
:::

| 特性          | lazygophers/log | zap    | logrus | 標準日誌 |
| ------------- | --------------- | ------ | ------ | -------- |
| 性能          | 高              | 高     | 中     | 低       |
| API 簡潔性    | 高              | 中     | 高     | 高       |
| 功能豐富度    | 中              | 高     | 高     | 低       |
| 靈活性        | 中              | 高     | 高     | 低       |
| 學習曲線      | 低              | 中     | 中     | 低       |

## ❓ 常見問題

### 如何選擇合適的日誌級別？

-   **開發環境**：使用 `DebugLevel` 或 `TraceLevel` 獲取詳細資訊
-   **生產環境**：使用 `InfoLevel` 或 `WarnLevel` 減少開銷
-   **性能測試**：使用 `PanicLevel` 禁用所有日誌

### 如何在生產環境優化性能？

:::warning 注意
在高吞吐量場景下，建議結合非同步寫入和合理的日誌級別來優化性能。
:::

1. 使用 `AsyncWriter` 非同步寫入：

```go title="非同步寫入配置"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. 調整日誌級別，避免不必要的日誌：

```go title="級別優化"
logger.SetLevel(log.InfoLevel)  // 跳過 Debug 和 Trace
```

3. 使用條件日誌減少開銷：

```go title="條件日誌"
if logger.Level >= log.DebugLevel {
    logger.Debug("詳細調試資訊")
}
```

### `Caller` 和 `EnableCaller` 有什麼區別？

-   **`EnableCaller(enable bool)`**：控制 Logger 是否收集調用者資訊
    -   `EnableCaller(true)` 啟用調用者資訊收集
-   **`Caller(disable bool)`**：控制 Formatter 是否輸出調用者資訊
    -   `Caller(true)` 禁用調用者資訊輸出

推薦使用 `EnableCaller` 進行全域控制。

### 如何實現自定義格式化器？

實現 `Format` 介面：

```go title="自定義格式化器"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 相關文檔

-   [📚 API 文檔](API.md) - 完整的 API 參考
-   [🤝 貢獻指南](/zh-TW/CONTRIBUTING) - 如何貢獻
-   [📋 變更日誌](/zh-TW/CHANGELOG) - 版本歷史
-   [🔒 安全政策](/zh-TW/SECURITY) - 安全指南
-   [📜 行為準則](/zh-TW/CODE_OF_CONDUCT) - 社區準則

## 🚀 獲取幫助

-   **GitHub Issues**：[報告 bug 或請求功能](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[API 文檔](https://pkg.go.dev/github.com/lazygophers/log)
-   **示例**：[使用示例](https://github.com/lazygophers/log/tree/main/examples)

## 📄 許可證

本專案採用 MIT 許可證 - 詳見 [LICENSE](/zh-TW/LICENSE) 檔案。

## 🤝 貢獻

我們歡迎貢獻！請查看我們的 [貢獻指南](/zh-TW/CONTRIBUTING) 了解詳情。

---

**lazygophers/log** 旨在成為重視性能和簡潔性的 Go 開發者的首選日誌解決方案。無論您是構建小型工具還是大規模分佈式系統，該庫都能提供功能和易用性之間的良好平衡。