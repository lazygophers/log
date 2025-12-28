---
pageType: home

hero:
  name: LazyGophers Log
  text: 高效能、靈活的 Go 日誌庫
  tagline: 基於 zap 構建，提供豐富的功能和簡潔的 API
  actions:
    - theme: brand
      text: 快速開始
      link: /zh-TW/API
    - theme: alt
      text: 查看 GitHub
      link: https://github.com/lazygophers/log

features:
  - title: '高效能'
    details: 基於 zap 構建，採用物件池複用和條件欄位記錄，實現最佳效能
    icon: 🚀
  - title: '豐富的日誌級別'
    details: 支援 Trace、Debug、Info、Warn、Error、Fatal、Panic 級別
    icon: 📊
  - title: '靈活的配置'
    details: 可自訂日誌級別、呼叫者資訊、追蹤資訊、前綴、後綴和輸出目標
    icon: ⚙️
  - title: '檔案輪轉'
    details: 內建每小時日誌檔案輪轉支援
    icon: 🔄
  - title: 'Zap 相容性'
    details: 與 zap WriteSyncer 無縫整合
    icon: 🔌
  - title: '簡潔的 API'
    details: 類似標準日誌庫的清晰 API，易於使用和整合
    icon: 🎯
---

## 快速開始

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
    // 使用預設全域 logger
    log.Debug("除錯資訊")
    log.Info("一般資訊")
    log.Warn("警告資訊")
    log.Error("錯誤資訊")

    // 使用格式化輸出
    log.Infof("使用者 %s 登入成功", "admin")

    // 自訂配置
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("這是來自自訂 logger 的日誌")
}
```

## 文件

- [API 參考](API.md) - 完整的 API 文件
- [更新日誌](CHANGELOG.md) - 版本歷史
- [貢獻指南](CONTRIBUTING.md) - 如何貢獻
- [安全策略](SECURITY.md) - 安全指南
- [行為準則](CODE_OF_CONDUCT.md) - 社群準則

## 效能比較

| 特性          | lazygophers/log | zap    | logrus | 標準日誌 |
| ------------- | --------------- | ------ | ------ | -------- |
| 效能          | 高              | 高     | 中     | 低       |
| API 簡潔性    | 高              | 中     | 高     | 高       |
| 功能豐富度    | 中              | 高     | 高     | 低       |
| 靈活性        | 中              | 高     | 高     | 低       |
| 學習曲線      | 低              | 中     | 中     | 低       |

## 授權

本專案採用 MIT 授權 - 詳見 [LICENSE](LICENSE) 檔案。
