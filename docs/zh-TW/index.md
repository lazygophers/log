---
titleSuffix: " | LazyGophers Log"
pageType: home

hero:
  name: LazyGophers 日誌庫
  text: 一個高性能的 Go 語言日誌庫
  tagline: 簡潔的 API、卓越的性能和靈活的配置
  image:
    src: /log/public/logo.svg
    alt: LazyGophers Log Logo
  actions:
    - theme: brand
      text: 快速開始
      link: /zh-TW/README
    - theme: alt
      text: API 參考
      link: /zh-TW/API
    - theme: alt
      text: 查看 GitHub
      link: https://github.com/lazygophers/log

features:
  - title: 高性能
    details: 基於 zap 構建，採用物件池和條件欄位記錄技術，確保出色的性能表現
    icon: ⚡
  - title: 豐富的日誌級別
    details: 支援 Trace、Debug、Info、Warn、Error、Fatal、Panic 七個日誌級別
    icon: 📊
  - title: 靈活的配置
    details: 支援日誌級別控制、呼叫者資訊記錄、跟蹤資訊、自定義前綴後綴等
    icon: ⚙️
  - title: 文件輪換
    details: 內置日誌文件輪換功能，支援按小時自動輪換日誌文件
    icon: 🔄
  - title: Zap 兼容性
    details: 與 zap WriteSyncer 無縫集成，支援自定義輸出目標
    icon: 🔗
  - title: 簡潔的 API
    details: API 設計類似於標準日誌庫，易於使用和遷移
    icon: 🚀
  - title: 線程安全
    details: 大多數操作採用無鎖設計，確保線程安全且無性能開銷
    icon: 🔒
  - title: 生產就緒
    details: 在生產環境中經過實戰檢驗，具有全面的測試覆蓋率
    icon: ✅
---
