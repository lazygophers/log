---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

高性能で柔軟な Go ログライブラリ。zap をベースに構築されており、豊富な機能とシンプルな API を提供します。

## 📖 言語

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 簡体字中国語](README.md) (現在)
-   [🇹🇼 繁体字中国語](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)

## ✨ 機能

-   **🚀 高性能**：zap をベースにしたオブジェクトプールと条件付きフィールド記録
-   **📊 様々なログレベル**：Trace、Debug、Info、Warn、Error、Fatal、Panic
-   **⚙️ 柔軟な設定**：
    -   ログレベルの制御
    -   コールバック情報の記録
    -   トレース情報（goroutine ID を含む）
    -   カスタムプレフィックスとサフィックス
    -   カスタム出力先（コンソール、ファイルなど）
    -   ログフォーマットオプション
-   **🔄 ファイルローテーション**：1時間ごとのログファイルローテーション対応
-   **🔌 Zap との互換性**：zap WriteSyncer とのシームレスな統合
-   **🎯 シンプルな API**：標準のログライブラリと同じくクリアな API、使いやすい

## 🚀 クイックスタート

### インストール

:::tip インストール
```bash
go get github.com/lazygophers/log
```
:::

### 基本的な使用方法

```go title="クイックスタート"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // デフォルトのグローバル logger を使用
    log.Debug("デバッグメッセージ")
    log.Info("情報メッセージ")
    log.Warn("警告メッセージ")
    log.Error("エラーメッセージ")

    // フォーマット出力を使用
    log.Infof("ユーザー %s が正常にログインしました", "admin")

    // カスタム設定
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("これはカスタム logger からのログです")
}
```

## 📚 詳細な使用方法

### ファイル出力付きカスタム Logger

```go title="ファイル出力設定"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // ファイル出力付きの logger を作成
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("コールバック情報付きのデバッグログ")
    logger.Info("トレース情報付きの情報ログ")
}
```

### ログレベルの制御

```go title="ログレベルの制御"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // warn とそれ以上のレベルのみが記録されます
    logger.Debug("これは記録されません")  // 無視
    logger.Info("これは記録されません")   // 無視
    logger.Warn("これは記録されます")    // 記録
    logger.Error("これは記録されます")   // 記録
}
```

## 🎯 使用シーン

### 適用シーン

-   **Web サービスと API バックエンド**：リクエスト追跡、構造化ログ、パフォーマンス監視
-   **マイクロサービスアーキテクチャ**：分散トレース（TraceID）、統一されたログ形式、高スループット
-   **コマンドラインツール**：レベル制御、クリアな出力、エラーレポート
-   **バッチ処理タスク**：ファイルローテーション、長時間実行、リソース最適化

### 特別な利点

-   **オブジェクトプール最適化**：Entry と Buffer オブジェクトの再利用、GC 圧力の削減
-   **非同期書き込み**：高スループットシナリオ（10000+ ログ/秒）でのブロッキングなし
-   **TraceID サポート**：分散システムのリクエスト追跡、OpenTelemetry との統合
-   **ゼロ設定で起動**：すぐに使用可能、段階的な設定

## 🔧 設定オプション

:::note 設定オプション
以下のすべてのメソッドはチェーン呼び出しをサポートしており、カスタム Logger を構築するために組み合わせることができます。
:::

### Logger 設定

| メソッド                  | 説明                | デフォルト      |
| --------------------- | ------------------- | ------------ |
| `SetLevel(level)`       | 最小ログレベルを設定     | `DebugLevel` |
| `EnableCaller(enable)`  | コールバック情報を有効化/無効化 | `false`      |
| `EnableTrace(enable)`   | トレース情報を有効化/無効化  | `false`      |
| `SetCallerDepth(depth)` | コールバック深度を設定   | `2`          |
| `SetPrefixMsg(prefix)`  | ログプレフィックスを設定  | `""`         |
| `SetSuffixMsg(suffix)`  | ログサフィックスを設定  | `""`         |
| `SetOutput(writers...)` | 出力先を設定         | `os.Stdout`  |

### ログレベル

| レベル        | 説明                        |
| ----------- | --------------------------- |
| `TraceLevel` | 最も詳細、詳細な追跡用        |
| `DebugLevel` | デバッグ情報                  |
| `InfoLevel`  | 一般情報                      |
| `WarnLevel`  | 警告メッセージ                  |
| `ErrorLevel` | エラーメッセージ                  |
| `FatalLevel` | 致命的なエラー（os.Exit(1) を呼び出し）|
| `PanicLevel` | パニックエラー（panic() を呼び出し）|

## 🏗️ アーキテクチャ

### コアコンポーネント

-   **Logger**：設定可能なオプション付きの主なログ構造
-   **Entry**：包括的なフィールドサポート付きの個別ログレコード
-   **Level**：ログレベルの定義とユーティリティ関数
-   **Format**：ログフォーマットインターフェースと実装

### パフォーマンス最適化

-   **オブジェクトプール**：メモリ割り当てを削減するために Entry オブジェクトを再利用
-   **条件付き記録**：必要な場合のみ高価なフィールドを記録
-   **高速レベルチェック**：最外層でログレベルをチェック
-   **ロックフリー設計**：大多数の操作はロックを必要としない

## 📊 パフォーマンス比較

:::info パフォーマンス比較
以下のデータはベンチマークに基づいています。実際のパフォーマンスは環境と設定によって異なる場合があります。
:::

| 特性          | lazygophers/log | zap    | logrus | 標準ログ |
| ------------- | --------------- | ------ | ------ | -------- |
| パフォーマンス      | 高              | 高     | 中     | 低       |
| API のシンプルさ    | 高              | 中     | 高     | 高       |
| 機能の豊富さ      | 中              | 高     | 高     | 低       |
| 柔軟性          | 中              | 高     | 高     | 低       |
| 学習曲線         | 低              | 中     | 中     | 低       |

## ❓ よくある質問

### 適切なログレベルを選択するには？

-   **開発環境**：詳細情報を得るために `DebugLevel` または `TraceLevel` を使用
-   **本番環境**：オーバーヘッドを削減するために `InfoLevel` または `WarnLevel` を使用
-   **パフォーマンステスト**：すべてのログを無効にするために `PanicLevel` を使用

### 本番環境でパフォーマンスを最適化するには？

:::warning 注意
高スループットシナリオでは、パフォーマンスを最適化するために非同期書き込みと適切なログレベルを組み合わせることをお勧めします。
:::

1. 非同期書き込みに `AsyncWriter` を使用：

```go title="非同期書き込み設定"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. 不要なログを避けるためにログレベルを調整：

```go title="レベル最適化"
logger.SetLevel(log.InfoLevel)  // Debug と Trace をスキップ
```

3. オーバーヘッドを削減するために条件付きログを使用：

```go title="条件付きログ"
if logger.Level >= log.DebugLevel {
    logger.Debug("詳細なデバッグ情報")
}
```

### `Caller` と `EnableCaller` の違いは？

-   **`EnableCaller(enable bool)`**：Logger がコールバック情報を収集するかどうかを制御
    -   `EnableCaller(true)` はコールバック情報収集を有効化
-   **`Caller(disable bool)`**：Formatter がコールバック情報を出力するかどうかを制御
    -   `Caller(true)` はコールバック情報出力を無効化

グローバル制御には `EnableCaller` を使用することをお勧めします。

### カスタムフォーマッターを実装するには？

`Format` インターフェースを実装します：

```go title="カスタムフォーマッター"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 関連ドキュメント

-   [📚 API ドキュメント](API.md) - 完全な API リファレンス
-   [🤝 貢献ガイド](/ja/CONTRIBUTING) - 貢献の方法
-   [📋 変更履歴](/ja/CHANGELOG) - バージョン履歴
-   [🔒 セキュリティポリシー](/ja/SECURITY) - セキュリティガイドライン
-   [📜 コードオブコンダクト](/ja/CODE_OF_CONDUCT) - コミュニティガイドライン

## 🚀 ヘルプを得る

-   **GitHub Issues**：[バグを報告または機能をリクエスト](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[API ドキュメント](https://pkg.go.dev/github.com/lazygophers/log)
-   [✓ サンプル](https://github.com/lazygophers/log/tree/main/examples)

## 📄 ライセンス

このプロジェクトは MIT ライセンスの下でライセンスされています - 詳細については [LICENSE](/ja/LICENSE) ファイルをご覧ください。

## 🤝 貢献

貢献を歓迎します！[貢献ガイド](/ja/CONTRIBUTING)をご覧ください。

---

**lazygophers/log** は、パフォーマンスとシンプルさを重視する Go 開発者のために設計された主要なログソリューションです。小さなユーティリティを構築している場合でも、大規模な分散システムを構築している場合でも、このライブラリは機能と使いやすさの間の良いバランスを提供します。
