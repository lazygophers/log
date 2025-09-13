# 🚀 LazyGophers Log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![DeepWiki](https://img.shields.io/badge/DeepWiki-documented-blue?logo=bookstack&logoColor=white)](https://deepwiki.ai/docs/lazygophers/log)
[![Go.Dev Downloads](https://pkg.go.dev/badge/github.com/lazygophers/log.svg)](https://pkg.go.dev/github.com/lazygophers/log)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/log)
[![Goproxy.io](https://goproxy.io/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.io/stats/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一個高效能、功能豐富的 Go 日誌庫，支援多建置標籤、非同步寫入和廣泛的自訂選項。

## 📖 文檔語言

- [🇺🇸 English](../README.md)
- [🇨🇳 简体中文](README.zh-CN.md)
- [🇹🇼 繁體中文](README.zh-TW.md) (目前)
- [🇫🇷 Français](README.fr.md)
- [🇷🇺 Русский](README.ru.md)
- [🇪🇸 Español](README.es.md)
- [🇸🇦 العربية](README.ar.md)

## ✨ 特色

- **🚀 高效能**: 物件池和非同步寫入支援
- **🏗️ 建置標籤支援**: 為除錯、發佈和捨棄模式提供不同行為
- **🔄 日誌輪轉**: 自動按小時輪轉日誌檔案
- **🎨 豐富格式化**: 可客製化的日誌格式和顏色支援
- **🔍 上下文追蹤**: Goroutine ID 和追蹤 ID 跟踪
- **🔌 框架整合**: 原生 Zap 日誌框架整合
- **⚙️ 高度可配置**: 靈活的級別、輸出和格式化配置
- **🧪 充分測試**: 在所有建置配置下達到 93.0% 測試涵蓋率

## 🚀 快速開始

### 安裝

```bash
go get github.com/lazygophers/log
```

### 基本用法

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 簡單日誌記錄
    log.Info("你好，世界！")
    log.Debug("這是一條除錯訊息")
    log.Warn("這是一條警告")
    log.Error("這是一條錯誤")

    // 格式化日誌記錄
    log.Infof("使用者 %s 登入，ID 為 %d", "張三", 123)
    
    // 使用自訂日誌器
    logger := log.New()
    logger.SetLevel(log.InfoLevel)
    logger.Info("自訂日誌器訊息")
}
```

### 進階用法

```go
package main

import (
    "context"
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 建立帶檔案輸出的日誌器
    logger := log.New()
    
    // 設定輸出到按小時輪轉的檔案
    writer := log.GetOutputWriterHourly("./logs/app.log")
    logger.SetOutput(writer)
    
    // 配置格式化
    logger.SetLevel(log.DebugLevel)
    logger.SetPrefixMsg("[APP] ")
    logger.Caller(true) // 啟用呼叫者資訊
    
    // 上下文日誌記錄
    ctxLogger := logger.CloneToCtx()
    ctxLogger.Info(context.Background(), "上下文感知日誌記錄")
    
    // 高吞吐量場景的非同步日誌記錄
    asyncWriter := log.NewAsyncWriter(writer, 1000)
    logger.SetOutput(asyncWriter)
    defer asyncWriter.Close()
    
    logger.Info("高效能非同步日誌記錄")
}
```

## 🏗️ 建置標籤

該庫透過 Go 建置標籤支援不同的建置模式：

### 預設模式（無標籤）
```bash
go build
```
- 完整日誌功能
- 除錯訊息啟用
- 標準效能

### 除錯模式
```bash
go build -tags debug
```
- 增強除錯資訊
- 詳細呼叫者資訊
- 效能分析支援

### 發佈模式
```bash
go build -tags release
```
- 針對正式環境最佳化
- 除錯訊息停用
- 自動日誌檔案輪轉

### 捨棄模式
```bash
go build -tags discard
```
- 最大效能
- 所有日誌被捨棄
- 零日誌開銷

### 組合模式
```bash
go build -tags "debug,discard"    # 除錯 + 捨棄
go build -tags "release,discard"  # 發佈 + 捨棄
```

## 📊 日誌級別

該庫支援 7 個日誌級別（從最高到最低優先級）：

| 級別 | 值 | 描述 |
|------|---|------|
| `PanicLevel` | 0 | 記錄日誌然後呼叫 panic |
| `FatalLevel` | 1 | 記錄日誌然後呼叫 os.Exit(1) |
| `ErrorLevel` | 2 | 錯誤條件 |
| `WarnLevel` | 3 | 警告條件 |
| `InfoLevel` | 4 | 資訊訊息 |
| `DebugLevel` | 5 | 除錯級別訊息 |
| `TraceLevel` | 6 | 最詳細的日誌記錄 |

## 🔌 框架整合

### Zap 整合

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/lazygophers/log"
)

// 建立一個寫入到我們日誌系統的 zap 日誌器
logger := log.New()
hook := log.NewZapHook(logger)

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    hook,
    zapcore.InfoLevel,
)
zapLogger := zap.New(core)

zapLogger.Info("來自 Zap 的訊息", zap.String("key", "value"))
```

## 🧪 測試

該庫提供全面的測試支援：

```bash
# 執行所有測試
make test

# 執行所有建置標籤的涵蓋率測試
make coverage-all

# 快速測試所有建置標籤
make test-quick

# 產生 HTML 涵蓋率報告
make coverage-html
```

### 按建置標籤的涵蓋率結果

| 建置標籤 | 涵蓋率 |
|----------|--------|
| 預設 | 92.9% |
| 除錯 | 93.1% |
| 發佈 | 93.5% |
| 捨棄 | 93.1% |
| 除錯+捨棄 | 93.1% |
| 發佈+捨棄 | 93.3% |

## ⚙️ 配置選項

### 日誌器配置

```go
logger := log.New()

// 設定最小日誌級別
logger.SetLevel(log.InfoLevel)

// 配置輸出
logger.SetOutput(os.Stdout) // 單個寫入器
logger.SetOutput(writer1, writer2, writer3) // 多個寫入器

// 客製化訊息
logger.SetPrefixMsg("[MyApp] ")
logger.SetSuffixMsg(" [END]")
logger.AppendPrefixMsg("額外: ")

// 配置格式化
logger.ParsingAndEscaping(false) // 停用跳脫序列
logger.Caller(true) // 啟用呼叫者資訊
logger.SetCallerDepth(4) // 調整呼叫者堆疊深度
```

## 📁 日誌輪轉

可配置間隔的自動日誌輪轉：

```go
// 按小時輪轉
writer := log.GetOutputWriterHourly("./logs/app.log")

// 該庫將建立如下檔案：
// - app-2024010115.log (2024-01-01 15:00)
// - app-2024010116.log (2024-01-01 16:00)
// - app-2024010117.log (2024-01-01 17:00)
```

## 🔍 上下文和追蹤

內建支援上下文感知日誌記錄和分散式追蹤：

```go
// 為目前 goroutine 設定追蹤 ID
log.SetTrace("trace-123-456")

// 取得追蹤 ID
traceID := log.GetTrace()

// 上下文感知日誌記錄
ctx := context.Background()
ctxLogger := log.CloneToCtx()
ctxLogger.Info(ctx, "請求已處理", "user_id", 123)

// 自動 goroutine ID 跟踪
log.Info("此日誌自動包含 goroutine ID")
```

## 📈 效能

該庫為高效能應用程式而設計：

- **物件池**: 重用日誌條目物件以減少 GC 壓力
- **非同步寫入**: 高吞吐量場景的非阻塞日誌寫入
- **級別過濾**: 早期過濾防止昂貴操作
- **建置標籤最佳化**: 不同環境的編譯時最佳化

### 基準測試

```bash
# 執行效能基準測試
make benchmark

# 不同建置模式的基準測試
make benchmark-debug
make benchmark-release  
make benchmark-discard
```

## 🤝 貢獻

我們歡迎貢獻！請查看我們的[貢獻指南](CONTRIBUTING.md)了解詳情。

### 開發環境設定

1. **Fork 並複製**
   ```bash
   git clone https://github.com/your-username/log.git
   cd log
   ```

2. **安裝相依性**
   ```bash
   go mod tidy
   ```

3. **執行測試**
   ```bash
   make test-all
   ```

4. **提交 Pull Request**
   - 遵循我們的 [PR 範本](../.github/pull_request_template.md)
   - 確保測試通過
   - 如需要請更新文件

## 📋 需求

- **Go**: 1.19 或更高版本
- **相依性**: 
  - `go.uber.org/zap` (用於 Zap 整合)
  - `github.com/petermattis/goid` (用於 goroutine ID)
  - `github.com/lestrrat-go/file-rotatelogs` (用於日誌輪轉)
  - `github.com/google/uuid` (用於追蹤 ID)

## 📄 授權

本專案採用 MIT 授權 - 請查看 [LICENSE](../LICENSE) 檔案了解詳情。

## 🙏 致謝

- [Zap](https://github.com/uber-go/zap) 提供靈感和整合支援
- [Logrus](https://github.com/sirupsen/logrus) 提供級別設計模式
- Go 社群持續的回饋和改進

## 📞 支援

- 📖 [文件](../docs/)
- 🐛 [問題追蹤](https://github.com/lazygophers/log/issues)
- 💬 [討論](https://github.com/lazygophers/log/discussions)
- 📧 信箱: support@lazygophers.com

---

**由 LazyGophers 團隊用 ❤️ 製作**