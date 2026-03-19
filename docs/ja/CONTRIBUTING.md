---
titleSuffix: ' | LazyGophers Log'
---
# 🤝 LazyGophers Log への貢献

貢献を歓迎します！貢献を可能な限り簡単かつ透明にするために、以下を含むあらゆる種類の貢献を受け付けます：

-   🐛 バグの報告
-   💬 コードの現在の状態に関する議論
-   ✨ 機能リクエスト
-   🔧 修正の提案
-   🚀 新機能の実装

## 📋 目次

-   [行動規範](#-行動規範)
-   [開発フロー](#-開発フロー)
-   [始め方](#-始め方)
-   [プルリクエストのプロセス](#-プルリクエストのプロセス)
-   [コーディング標準](#-コーディング標準)
-   [テストガイド](#-テストガイド)
-   [ビルドタグの要件](#️-ビルドタグの要件)
-   [ドキュメント](#-ドキュメント)
-   [問題ガイド](#-問題ガイド)
-   [パフォーマンスの考慮](#-パフォーマンスの考慮)
-   [セキュリティガイド](#-セキュリティガイド)
-   [コミュニティ](#-コミュニティ)

## 📜 行動規範

このプロジェクトおよびそのすべての参加者は、私たちの[行動規範](/ja/CODE_OF_CONDUCT)に従います。参加することで、この規範に従うことに同意したとみなされます。

## 🔄 開発フロー

GitHub を使用してコードをホストし、問題と機能リクエストを追跡し、プルリクエストを受け付けています。

### ワークフロー

:::note 開発フローの概要
1. **Fork** リポジトリ
2. **Clone** fork をローカルに
3. **作成** `master` ブランチから機能ブランチを作成
4. **実行** 変更を行う
5. **テスト** すべてのビルドタグで徹底的にテスト
6. **提出** プルリクエスト
:::

## 🚀 始め方

### 前提条件

-   **Go 1.21+** - [Go をインストール](https://golang.org/doc/install)
-   **Git** - [Git をインストール](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
-   **Make**（推奨だが任意）

### ローカル開発のセットアップ

```bash title="リポジトリをクローンして開発環境を設定"
# 1. GitHub でリポジトリを Fork
# 2. fork を Clone
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. アップストリームのリモートリポジトリを追加
git remote add upstream https://github.com/lazygophers/log.git

# 4. 依存関係をインストール
go mod tidy

# 5. インストールを確認
make test-quick
```

### 環境設定

:::info 環境設定
Go 環境変数が正しく設定されていることを確認し、最適な開発体験を得るために推奨される開発ツールをインストールしてください。
:::

```bash title="環境設定"
# Go 環境を設定（まだ設定されていない場合）
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# 積極的なツールのインストール（任意）
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 プルリクエストのプロセス

### プルリクエスト提出前

1. **検索** 既存の PR をして重複を避ける
2. **テスト** 変更がすべてのビルド構成で動作することを確認
3. **記録** すべての破壊的変更
4. **更新** 関連するドキュメント
5. **追加** 新機能のテスト

### PR チェックリスト

:::warning PR を提出する前に以下のすべてを確認してください
チェックリストの要件を満たさない PR はマージされません。
:::

-   [ ] **コード品質**

    -   [ ] コードがプロジェクトのスタイルガイドに従う
    -   [ ] 新しい lint 警告がない
    -   [ ] 適切なエラーハンドリング
    -   [ ] 効率的なアルゴリズムとデータ構造

-   [ ] **テスト**

    -   [ ] すべての既存のテストが通る：`make test`
    -   [ ] 新機能のために新しいテストを追加
    -   [ ] テストカバレッジが維持または向上
    -   [ ] すべてのビルドタグがテストされる：`make test-all`

-   [ ] **ドキュメント**

    -   [ ] コードに適切なコメント
    -   [ ] API ドキュメントが更新（必要な場合）
    -   [ ] README が更新（必要な場合）
    -   [ ] 多言語ドキュメントが更新（ユーザー向けの場合）

-   [ ] **ビルド互換性**
    -   [ ] デフォルトモード：`go build`
    -   [ ] デバッグモード：`go build -tags debug`
    -   [ ] リリースモード：`go build -tags release`
    -   [ ] ドロップモード：`go build -tags discard`
    -   [ ] 組み合わせモードがテストされる

### PR テンプレート

プルリクエストを提出する際は、[PR テンプレート](https://github.com/lazygophers/log/blob/main/.github/pull_request_template.md)を使用してください。

## 📏 コーディング標準

### Go スタイルガイド

:::tip Go コード規範
標準の Go スタイルガイドに従い、いくつかの補足事項があります。コードが `go fmt` と `goimports` でフォーマットされていることを確認してください。
:::

```go
// ✅ Good
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ Bad
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### 命名規約

-   **パッケージ**: 短く、小文字、可能であれば単語1つ
-   **関数**: キャメルケース、記述的
-   **変数**: ローカル変数は camelCase、エクスポートされた変数は CamelCase
-   **定数**: エクスポートされた定数は CamelCase、非エクスポートされた定数は camelCase
-   **インターフェース**: 通常 "er" で終わる（例：`Writer`、`Formatter`）

### コードの整理

```
project/
├── docs/           # 多言語ドキュメント
├── .github/        # GitHub テンプレートとワークフロー
├── logger.go       # メインロギング実装
├── entry.go        # ログエントリ構造
├── level.go        # ログレベル
├── formatter.go    # ログフォーマット
├── output.go       # 出力管理
└── *_test.go      # ソースコードと共存するテスト
```

### エラーハンドリング

:::tip エラーハンドリングのベストプラクティス
ライブラリコードは panic ではなくエラーを返し、呼び出し元が例外処理を決定できるようにします。
:::

```go title="エラーハンドリングの例"
// ✅ 推荐：エラーを返し、panic は避ける
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ 避ける：ライブラリコードで panic を使用
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // しないでください
    }
    return &Logger{...}
}
```

## 🧪 テストガイド

### テスト構造

```go title="テーブル駆動テストの例"
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
            // テスト実装
        })
    }
}
```

### カバレッジ要件

:::warning カバレッジの厳格要件
新規コードのカバレッジが 90% 未満の PR は CI チェックに合格しません。
:::

-   **最低要件**: 新規コードのカバレッジ 90%
-   **目標**: 全体のカバレッジ 95%+
-   **すべてのビルドタグ**でカバレッジを維持
-   `make coverage-all` で検証

### テストコマンド

```bash title="テストを実行"
# すべてのビルドタグでクイックテスト
make test-quick

# カバレッジを含む完全なテストスイート
make test-all

# カバレッジレポート
make coverage-html

# ベンチマークテスト
make benchmark
```

## 🏗️ ビルドタグの要件

:::warning ビルド互換性
すべての変更はビルドタグシステムと互換である必要があり、すべてのビルドタグテストに合格しないコードはマージされません。
:::

### サポートされるビルドタグ

-   **デフォルト** (`go build`): 完全な機能
-   **デバッグ** (`go build -tags debug`): 強化されたデバッグ機能
-   **リリース** (`go build -tags release`): 本番環境向けの最適化
-   **ドロップ** (`go build -tags discard`): 最大パフォーマンス

### ビルドタグのテスト

:::info ビルドタグの説明
プロジェクトはビルドタグを使用して条件付きコンパイルを実装しており、異なるタグは異なる実行モードを対応します。提出前にすべてのタグでテストされていることを確認してください。
:::

```bash title="ビルドタグのテスト"
# 各ビルド設定でテスト
make test-default
make test-debug
make test-release
make test-discard

# 組み合わせをテスト
make test-debug-discard
make test-release-discard

# すべてを一度にテスト
make test-all
```

### ビルドタグガイド

```go
//go:build debug
// +build debug

package log

// デバッグ固有の実装
```

## 📚 ドキュメント

### コードドキュメント

-   **すべてのエクスポートされた関数**には明確なコメントが必要
-   **複雑なアルゴリズム**には説明が必要
-   **例**は非自明な使用に役立つ
-   **スレッドセーフ**の説明（該当する場合）

```go
// SetLevel は最小のログレベルを設定します。
// このレベルより下のログは無視されます。
// このメソッドはスレッドセーフです。
//
// Example:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // 出力されない
//   logger.Info("visible")   // 出力される
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### README の更新

機能を追加する際は、以下を更新してください：

-   メインの README.md
-   `docs/` ですべての言語固有の README
-   コードの例
-   機能リスト

## 🐛 問題ガイド

### バグの報告

[バグレポートテンプレート](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/bug_report.md)を使用し、以下を含めます：

-   **明確な問題の説明**
-   **再現手順**
-   **期待される動作と実際の動作**
-   **環境の詳細**（OS、Go バージョン、ビルドタグ）
-   **最小のコード例**

### 機能リクエスト

[機能リクエストテンプレート](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/feature_request.md)を使用し、以下を含めます：

-   **明確な機能の動機**
-   **提案された API** 設計
-   **実装の考慮事項**
-   **破壊的変更の分析**

### 質問

[質問テンプレート](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/question.md)を使用して以下を扱います：

-   使用に関する問題
-   設定に関するヘルプ
-   ベストプラクティス
-   統合ガイダンス

## 🚀 パフォーマンスの考慮

### ベンチマーク

常にパフォーマンスに敏感な変更に対してベンチマークを実行します：

```bash title="ベンチマークを実行"
# ベンチマークを実行
go test -bench=. -benchmem

# 変更前後のパフォーマンスを比較
go test -bench=. -benchmem > before.txt
# 変更を行う
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### パフォーマンスガイド

:::tip パフォーマンス最適化のポイント
これはパフォーマンスに敏感なログライブラリであり、あらゆる変更はホットパスへの影響を考慮する必要があります。
:::

-   **最小化** ホットパス内のメモリ割り当て
-   **オブジェクトプールを使用** 頻繁に作成されるオブジェクト用
-   **早期リターン** 無効なログレベル用
-   **反射を避ける** パフォーマンス重要なコードで
-   **最適化前にプロファイリングを行う**

### メモリ管理

```go
// ✅ 推荐：オブジェクトプールを使用
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

## 🔒 セキュリティガイド

### 敏感なデータ

:::warning セキュリティに関する注意
ログからの漏洩により深刻なセキュリティ事故が発生する可能性があるため、以下の規範を遵守してください。
:::

-   **決して記録しない** パスワード、トークン、または機密データ
-   **クリーンアップ** ログメッセージ内のユーザー入力
-   **避ける** 完全なリクエスト/レスポンス本文を記録
-   **使用する** 結果として制御できるログ記録

```go
// ✅ 推荐
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ 避ける
logger.Infof("User login: %+v", userRequest) // パスワードを含む可能性
```

### 依存関係

-   依存関係を**最新に保つ**
-   **慎重に審査** 新しい依存関係
-   **最小化** 外部依存関係
-   **使用する** `go mod verify` で整合性をチェック

## 👥 コミュニティ

### ヘルプを得る

-   📖 [ドキュメント](README.md)
-   💬 [GitHub ディスカッション](https://github.com/lazygophers/log/discussions)
-   🐛 [問題トラッカー](https://github.com/lazygophers/log/issues)
-   📧 Email: support@lazygophers.com

### コミュニケーションガイド

-   **尊重し、包摂的でいる**
-   **質問する前に検索する**
-   **ヘルプを求める際にコンテキストを提供する**
-   **力を尽くして他者を助ける**
-   **従う** [行動規範](/ja/CODE_OF_CONDUCT)

## 🎯 認定

貢献者は以下の方法で認定されます：

-   **README 貢献者**セクション
-   **リリースノート**での言及
-   **GitHub 貢献者**グラフ
-   **コミュニティへの感謝**投稿

## 📝 ライセンス

貢献することで、貢献が MIT ライセンスの下でライセンスされることに同意します。

---

## 🌍 多言語ドキュメント

このドキュメントは複数の言語バージョンを提供しています：

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CONTRIBUTING.md)
-   [🇨🇳 簡体字中国語](/zh-CN/CONTRIBUTING)（現在）
-   [🇹🇼 繁体字中国語](https://lazygophers.github.io/log/zh-TW/CONTRIBUTING.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CONTRIBUTING.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CONTRIBUTING.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CONTRIBUTING.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CONTRIBUTING.md)

---

**LazyGophers Log への貢献に感謝します！🚀**
