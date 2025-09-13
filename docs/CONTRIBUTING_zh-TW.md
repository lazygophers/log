# 🤝 為 LazyGophers Log 貢獻程式碼

我們非常歡迎您的參與！我們希望讓為 LazyGophers Log 貢獻程式碼變得盡可能簡單和透明，無論是：

- 🐛 回報錯誤
- 💬 討論程式碼現狀
- ✨ 提交功能請求
- 🔧 提出修復方案
- 🚀 實作新功能

## 📋 目錄

- [行為準則](#-行為準則)
- [開發流程](#-開發流程)
- [快速開始](#-快速開始)
- [拉取請求流程](#-拉取請求流程)
- [編碼標準](#-編碼標準)
- [測試指南](#-測試指南)
- [建置標籤要求](#️-建置標籤要求)
- [文件撰寫](#-文件撰寫)
- [問題提交指南](#-問題提交指南)
- [效能考量](#-效能考量)
- [安全指南](#-安全指南)
- [社群](#-社群)

## 📜 行為準則

本專案和所有參與者均受我們的[行為準則](CODE_OF_CONDUCT_zh-TW.md)約束。參與即表示您同意遵守此準則。

## 🔄 開發流程

我們使用 GitHub 來託管程式碼、追蹤問題和功能請求，以及接受拉取請求。

### 工作流程

1. **Fork** 儲存庫
2. **複製** 您的 fork 到本機
3. **從 `master` 建立**功能分支
4. **進行**變更
5. **全面測試**所有建置標籤
6. **提交**拉取請求

## 🚀 快速開始

### 前置條件

- **Go 1.21+** - [安裝 Go](https://golang.org/doc/install)
- **Git** - [安裝 Git](https://git-scm.com/book/zh-tw/v2/開始-Git-安裝教學)
- **Make**（可選但建議）

### 本機開發環境設定

```bash
# 1. 在 GitHub 上 fork 儲存庫
# 2. 複製您的 fork
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. 新增上游遠端儲存庫
git remote add upstream https://github.com/lazygophers/log.git

# 4. 安裝相依套件
go mod tidy

# 5. 驗證安裝
make test-quick
```

### 環境設定

```bash
# 設定您的 Go 環境（如果尚未設定）
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# 可選：安裝有用的工具
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 拉取請求流程

### 提交前檢查

1. **搜尋**現有 PR 以避免重複
2. **測試**您的變更在所有建置設定下的表現
3. **記錄**任何破壞性變更
4. **更新**相關文件
5. **為新功能新增**測試

### PR 檢查清單

- [ ] **程式碼品質**
  - [ ] 程式碼遵循專案風格指南
  - [ ] 無新的程式碼檢查警告
  - [ ] 適當的錯誤處理
  - [ ] 高效的演算法和資料結構

- [ ] **測試**
  - [ ] 所有現有測試通過：`make test`
  - [ ] 為新功能新增了新測試
  - [ ] 測試涵蓋率保持或提升
  - [ ] 所有建置標籤測試：`make test-all`

- [ ] **文件**
  - [ ] 程式碼有適當註解
  - [ ] API 文件已更新（如需要）
  - [ ] README 已更新（如需要）
  - [ ] 多語言文件已更新（如面向使用者）

- [ ] **建置相容性**
  - [ ] 預設模式：`go build`
  - [ ] 除錯模式：`go build -tags debug`
  - [ ] 發布模式：`go build -tags release`
  - [ ] 捨棄模式：`go build -tags discard`
  - [ ] 組合模式已測試

### PR 模板

提交拉取請求時請使用我們的 [PR 模板](.github/pull_request_template.md)。

## 📏 編碼標準

### Go 風格指南

我們遵循標準的 Go 風格指南，並有一些補充：

```go
// ✅ 良好
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ 不好
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### 命名慣例

- **套件名**：簡短、小寫、儘可能單字
- **函式**：CamelCase，描述性
- **變數**：本機變數 camelCase，匯出變數 CamelCase
- **常數**：匯出 CamelCase，未匯出 camelCase
- **介面**：通常以 "er" 結尾（如 `Writer`、`Formatter`）

### 程式碼組織

```
project/
├── docs/           # 多語言文件
├── .github/        # GitHub 範本和工作流程
├── logger.go       # 主要日誌實作
├── entry.go        # 日誌條目結構
├── level.go        # 日誌級別
├── formatter.go    # 日誌格式化
├── output.go       # 輸出管理
└── *_test.go      # 測試檔案與原始碼同目錄
```

### 錯誤處理

```go
// ✅ 建議：回傳錯誤，不要 panic
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ 避免：在函式庫程式碼中 panic
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // 不要這樣做
    }
    return &Logger{...}
}
```

## 🧪 測試指南

### 測試結構

```go
func TestLogger_Info(t *testing.T) {
    tests := []struct {
        name     string
        level    Level
        message  string
        expected bool
    }{
        {"info level allows info", InfoLevel, "test", true},
        {"warn level blocks info", WarnLevel, "test", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // 測試實作
        })
    }
}
```

### 涵蓋率要求

- **最低要求**：新程式碼 90% 涵蓋率
- **目標**：整體涵蓋率 95%+
- **所有建置標籤**必須保持涵蓋率
- 使用 `make coverage-all` 驗證

### 測試指令

```bash
# 快速測試所有建置標籤
make test-quick

# 帶涵蓋率的完整測試套件
make test-all

# 涵蓋率報告
make coverage-html

# 效能基準測試
make benchmark
```

## 🏗️ 建置標籤要求

所有變更必須與我們的建置標籤系統相容：

### 支援的建置標籤

- **預設** (`go build`)：完整功能
- **除錯** (`go build -tags debug`)：增強除錯
- **發布** (`go build -tags release`)：生產環境優化
- **捨棄** (`go build -tags discard`)：最大效能

### 建置標籤測試

```bash
# 測試每個建置設定
make test-default
make test-debug  
make test-release
make test-discard

# 測試組合
make test-debug-discard
make test-release-discard

# 全部測試
make test-all
```

### 建置標籤指南

```go
//go:build debug
// +build debug

package log

// 除錯特定實作
```

## 📚 文件撰寫

### 程式碼文件

- **所有匯出函式**必須有清楚註解
- **複雜演算法**需要解釋
- **非平凡用法**需要範例
- **執行緒安全性**說明（如適用）

```go
// SetLevel 設定最小日誌級別。
// 低於此級別的日誌將被忽略。
// 此方法是執行緒安全的。
//
// 範例:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // 不會輸出
//   logger.Info("visible")   // 會輸出
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### README 更新

新增功能時，請更新：
- 主要 README.md
- `docs/` 中所有語言特定的 README
- 程式碼範例
- 功能清單

## 🐛 問題提交指南

### 錯誤報告

使用[錯誤報告範本](.github/ISSUE_TEMPLATE/bug_report.md)並包含：

- **清楚描述**問題
- **重現步驟**
- **期望與實際行為**
- **環境詳情**（作業系統、Go 版本、建置標籤）
- **最小程式碼範例**

### 功能請求

使用[功能請求範本](.github/ISSUE_TEMPLATE/feature_request.md)並包含：

- **清楚的功能動機**
- **建議的 API** 設計
- **實作考量**
- **破壞性變更分析**

### 問題諮詢

使用[問題範本](.github/ISSUE_TEMPLATE/question.md)用於：

- 使用問題
- 設定說明
- 最佳實務
- 整合指導

## 🚀 效能考量

### 基準測試

始終對效能敏感的變更進行基準測試：

```bash
# 執行基準測試
go test -bench=. -benchmem

# 前後對比
go test -bench=. -benchmem > before.txt
# 進行變更
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### 效能指南

- **最小化熱路徑**中的配置
- **使用物件池**處理頻繁建立的物件
- **提前回傳**對停用的日誌級別
- **避免反射**在效能關鍵程式碼中
- **先分析再優化**

### 記憶體管理

```go
// ✅ 良好：使用物件池
var entryPool = sync.Pool{
    New: func() interface{} {
        return &Entry{}
    },
}

func getEntry() *Entry {
    return entryPool.Get().(*Entry)
}

func putEntry(e *Entry) {
    e.Reset()
    entryPool.Put(e)
}
```

## 🔒 安全指南

### 敏感資料

- **永不記錄**密碼、令牌或敏感資料
- **清理**日誌訊息中的使用者輸入
- **避免**記錄完整的請求/回應主體
- **使用**結構化日誌獲得更好的控制

```go
// ✅ 良好
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ 不好
logger.Infof("User login: %+v", userRequest) // 可能包含密碼
```

### 相依套件管理

- 保持相依套件**最新**
- **仔細審查**新的相依套件
- **最小化**外部相依套件
- **使用** `go mod verify` 檢查完整性

## 👥 社群

### 獲取協助

- 📖 [文件](../README_zh-TW.md)
- 💬 [GitHub 討論](https://github.com/lazygophers/log/discussions)
- 🐛 [問題追蹤器](https://github.com/lazygophers/log/issues)
- 📧 電子郵件：support@lazygophers.com

### 交流指南

- **尊重**和包容
- **提問前先搜尋**
- **請求協助時提供上下文**
- **力所能及地協助他人**
- **遵循**[行為準則](CODE_OF_CONDUCT_zh-TW.md)

## 🎯 貢獻認可

貢獻者透過以下方式獲得認可：

- **README 貢獻者**部分
- **發布說明**提及
- **GitHub 貢獻者**圖表
- **社群感謝**貼文

## 📝 授權

透過貢獻，您同意您的貢獻將在 MIT 授權條款下獲得授權。

---

## 🌍 多語言文件

本文件提供多種語言版本：

- [🇺🇸 English](CONTRIBUTING.md)
- [🇨🇳 简体中文](CONTRIBUTING_zh-CN.md)
- [🇹🇼 繁體中文](CONTRIBUTING_zh-TW.md)（目前）
- [🇫🇷 Français](CONTRIBUTING_fr.md)
- [🇷🇺 Русский](CONTRIBUTING_ru.md)
- [🇪🇸 Español](CONTRIBUTING_es.md)
- [🇸🇦 العربية](CONTRIBUTING_ar.md)

---

**感謝您為 LazyGophers Log 貢獻程式碼！🚀**