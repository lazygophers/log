# 📚 API 文件

## 概述

LazyGophers Log 提供了一個全面的日誌 API，支援多個日誌級別、自定義格式化、非同步寫入和建置標籤優化。本文件涵蓋所有公共 API、設定選項和使用模式。

## 目錄

- [核心類型](#核心類型)
- [Logger API](#logger-api)
- [全域函式](#全域函式)
- [日誌級別](#日誌級別)
- [格式化器](#格式化器)
- [輸出寫入器](#輸出寫入器)
- [上下文日誌](#上下文日誌)
- [建置標籤](#建置標籤)
- [效能優化](#效能優化)
- [範例](#範例)

## 核心類型

### Logger

提供所有日誌功能的主要日誌結構體。

```go
type Logger struct {
    // 包含用於執行緒安全操作的私有欄位
}
```

#### 建構函式

```go
func New() *Logger
```

建立具有預設設定的新日誌實例：
- 級別: `DebugLevel`
- 輸出: `os.Stdout`
- 格式化器: 預設文字格式化器
- 呼叫者追蹤: 停用

**範例:**
```go
logger := log.New()
logger.Info("新日誌器已建立")
```

### Entry

表示單個日誌條目及其所有關聯元數據。

```go
type Entry struct {
    Time       time.Time     // 條目建立時的時間戳
    Level      Level         // 日誌級別
    Message    string        // 日誌訊息
    Pid        int          // 程序 ID
    Gid        uint64       // Goroutine ID
    TraceID    string       // 分散式追蹤的追蹤 ID
    CallerName string       // 呼叫者函式名
    CallerFile string       // 呼叫者檔案路徑
    CallerLine int          // 呼叫者行號
}
```

## Logger API

### 設定方法

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

設定最小日誌級別。低於此級別的訊息將被忽略。

**參數:**
- `level`: 要處理的最小日誌級別

**傳回:**
- `*Logger`: 傳回自身用於方法鏈接

**範例:**
```go
logger.SetLevel(log.InfoLevel)
logger.Debug("這不會被顯示")  // 被忽略
logger.Info("這將被顯示")     // 被處理
```

## 日誌級別

### 可用級別

```go
const (
    PanicLevel Level = iota  // 0 - Panic 並退出
    FatalLevel              // 1 - 致命錯誤並退出  
    ErrorLevel              // 2 - 錯誤條件
    WarnLevel               // 3 - 警告條件
    InfoLevel               // 4 - 資訊訊息
    DebugLevel              // 5 - 除錯訊息
    TraceLevel              // 6 - 最詳細的追蹤
)
```

## 格式化器

### Format 介面

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

自定義格式化器必須實作此介面。

### JSON 格式化器範例

```go
type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "caller":    fmt.Sprintf("%s:%d", entry.CallerFile, entry.CallerLine),
    }
    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }
    
    jsonData, _ := json.Marshal(data)
    return append(jsonData, '\n')
}

// 使用
logger.SetFormatter(&JSONFormatter{})
```

## 輸出寫入器

### 帶輪轉的檔案輸出

```go
func GetOutputWriterHourly(filename string) io.Writer
```

建立一個按小時輪轉日誌檔案的寫入器。

**範例:**
```go
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// 建立檔案如: app-2024010115.log, app-2024010116.log, 等等
```

### 非同步寫入器

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

建立用於高效能日誌記錄的非同步寫入器。

**範例:**
```go
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## 建置標籤

程式庫支援使用建置標籤進行條件編譯：

### 預設模式
```bash
go build
```
- 啟用完整功能
- 包含除錯訊息
- 標準效能

### 除錯模式
```bash
go build -tags debug
```
- 增強的除錯資訊
- 詳細的呼叫者資訊

### 發布模式
```bash
go build -tags release
```
- 為生產環境優化
- 停用除錯訊息
- 啟用自動日誌輪轉

### 丟棄模式
```bash
go build -tags discard
```
- 最大效能
- 所有日誌都被丟棄
- 零日誌開銷

## 效能優化

### 物件池化

程式庫內部使用 `sync.Pool` 來池化：
- 日誌條目物件
- 位元組緩衝區
- 格式化器緩衝區

這在高吞吐量情境中減少了垃圾收集壓力。

### 級別檢查

日誌級別檢查在昂貴操作之前進行：

```go
// 高效 - 僅在級別啟用時才進行訊息格式化
logger.Debugf("昂貴操作結果: %+v", expensiveCall())
```

## 範例

### 基本使用

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("應用程式啟動中")
    log.Warn("這是一個警告")
    log.Error("這是一個錯誤")
}
```

### 自定義日誌器

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()
    
    // 設定日誌器
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)
    logger.SetPrefixMsg("[我的應用] ")
    
    // 設定輸出到檔案
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    logger.SetOutput(file)
    
    logger.Info("自定義日誌器已設定")
    logger.Debug("帶呼叫者的除錯資訊")
}
```

### 高效能日誌

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 建立輪轉檔案寫入器
    writer := log.GetOutputWriterHourly("./logs/app.log")
    
    // 用非同步寫入器包裝以提高效能
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()
    
    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // 在生產環境中跳過除錯
    
    // 高吞吐量日誌記錄
    for i := 0; i < 10000; i++ {
        logger.Infof("處理請求 %d", i)
    }
}
```

---

## 🌍 多語言文件

本文件提供多種語言版本：

- [🇺🇸 English](API.md)
- [🇨🇳 简体中文](API_zh-CN.md)
- [🇹🇼 繁體中文](API_zh-TW.md)（目前）
- [🇫🇷 Français](API_fr.md)
- [🇷🇺 Русский](API_ru.md)
- [🇪🇸 Español](API_es.md)
- [🇸🇦 العربية](API_ar.md)

---

**LazyGophers Log 的完整 API 參考 - 用卓越的日誌建置更好的應用程式！🚀**