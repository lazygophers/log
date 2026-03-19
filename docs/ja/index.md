---
pageType: home

hero:
    name: LazyGophers Log
    text: 高性能で柔軟なGoロギングライブラリ
    tagline: zapベースで構築され、豊富な機能とシンプルなAPIを提供
    actions:
        - theme: brand
          text: クイックスタート
          link: /API
        - theme: alt
          text: GitHubで見る
          link: https://github.com/lazygophers/log

features:
    - title: "高性能"
      details: zapベースで構築され、オブジェクトプールの再利用と条件付きフィールド記録により、最適なパフォーマンスを実現
      icon: 🚀
    - title: "豊富なログレベル"
      details: Trace、Debug、Info、Warn、Error、Fatal、Panicレベルをサポート
      icon: 📊
    - title: "柔軟な設定"
      details: ログレベル、呼び出し元情報、トレース情報、プレフィックス、サフィックス、出力先をカスタマイズ可能
      icon: ⚙️
    - title: "ログローテーション"
      details: 毎時のログファイルローテーションサポートを内蔵
      icon: 🔄
    - title: "Zap互換性"
      details: zap WriteSyncerとのシームレスな統合
      icon: 🔌
    - title: "シンプルなAPI"
      details: 標準ログライブラリに似た明確なAPIで、使用と統合が簡単
      icon: 🎯
---

## クイックスタート

### インストール

```bash
go get github.com/lazygophers/log
```

### 基本的な使用方法

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // デフォルトのグローバルloggerを使用
    log.Debug("デバッグ情報")
    log.Info("通常情報")
    log.Warn("警告情報")
    log.Error("エラー情報")

    // フォーマット付き出力を使用
    log.Infof("ユーザー %s がログインしました", "admin")

    // カスタム設定
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("これはカスタムloggerからのログです")
}
```

## ドキュメント

-   [APIリファレンス](API.md) - 完全なAPIドキュメント
-   [更新履歴](/ja/CHANGELOG) - バージョン履歴
-   [貢献ガイド](/ja/CONTRIBUTING) - 貢献する方法
-   [セキュリティポリシー](/ja/SECURITY) - セキュリティガイド
-   [行動規範](/ja/CODE_OF_CONDUCT) - コミュニティガイドライン

## パフォーマンス比較

| 特性       | lazygophers/log | zap | logrus | 標準ログ |
| ---------- | --------------- | --- | ------ | -------- |
| パフォーマンス | 高              | 高  | 中     | 低       |
| APIの簡潔さ | 高              | 中  | 高     | 高       |
| 機能の豊富さ | 中              | 高  | 高     | 低       |
| 柔軟性     | 中              | 高  | 高     | 低       |
| 学習曲線   | 低              | 中  | 中     | 低       |

## ライセンス

このプロジェクトはMITライセンスの下で提供されています - 詳細は[LICENSE](/ja/LICENSE)ファイルをご覧ください。
