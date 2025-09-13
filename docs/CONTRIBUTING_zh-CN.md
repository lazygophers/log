# 🤝 为 LazyGophers Log 贡献代码

我们非常欢迎您的参与！我们希望让为 LazyGophers Log 贡献代码变得尽可能简单和透明，无论是：

- 🐛 报告错误
- 💬 讨论代码现状
- ✨ 提交功能请求
- 🔧 提出修复方案
- 🚀 实现新功能

## 📋 目录

- [行为准则](#-行为准则)
- [开发流程](#-开发流程)
- [快速开始](#-快速开始)
- [拉取请求流程](#-拉取请求流程)
- [编码标准](#-编码标准)
- [测试指南](#-测试指南)
- [构建标签要求](#️-构建标签要求)
- [文档编写](#-文档编写)
- [问题提交指南](#-问题提交指南)
- [性能考虑](#-性能考虑)
- [安全指南](#-安全指南)
- [社区](#-社区)

## 📜 行为准则

本项目和所有参与者均受我们的[行为准则](CODE_OF_CONDUCT_zh-CN.md)约束。参与即表示您同意遵守此准则。

## 🔄 开发流程

我们使用 GitHub 来托管代码、跟踪问题和功能请求，以及接受拉取请求。

### 工作流程

1. **Fork** 仓库
2. **克隆** 您的 fork 到本地
3. **从 `master` 创建**功能分支
4. **进行**更改
5. **全面测试**所有构建标签
6. **提交**拉取请求

## 🚀 快速开始

### 前提条件

- **Go 1.21+** - [安装 Go](https://golang.org/doc/install)
- **Git** - [安装 Git](https://git-scm.com/book/zh/v2/起步-安装-Git)
- **Make**（可选但推荐）

### 本地开发环境设置

```bash
# 1. 在 GitHub 上 fork 仓库
# 2. 克隆您的 fork
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. 添加上游远程仓库
git remote add upstream https://github.com/lazygophers/log.git

# 4. 安装依赖
go mod tidy

# 5. 验证安装
make test-quick
```

### 环境配置

```bash
# 设置您的 Go 环境（如果尚未设置）
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# 可选：安装有用的工具
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 拉取请求流程

### 提交前检查

1. **搜索**现有 PR 以避免重复
2. **测试**您的更改在所有构建配置下的表现
3. **记录**任何破坏性更改
4. **更新**相关文档
5. **为新功能添加**测试

### PR 检查清单

- [ ] **代码质量**
  - [ ] 代码遵循项目风格指南
  - [ ] 无新的代码检查警告
  - [ ] 适当的错误处理
  - [ ] 高效的算法和数据结构

- [ ] **测试**
  - [ ] 所有现有测试通过：`make test`
  - [ ] 为新功能添加了新测试
  - [ ] 测试覆盖率保持或提升
  - [ ] 所有构建标签测试：`make test-all`

- [ ] **文档**
  - [ ] 代码有适当注释
  - [ ] API 文档已更新（如需要）
  - [ ] README 已更新（如需要）
  - [ ] 多语言文档已更新（如面向用户）

- [ ] **构建兼容性**
  - [ ] 默认模式：`go build`
  - [ ] 调试模式：`go build -tags debug`
  - [ ] 发布模式：`go build -tags release`
  - [ ] 丢弃模式：`go build -tags discard`
  - [ ] 组合模式已测试

### PR 模板

提交拉取请求时请使用我们的 [PR 模板](.github/pull_request_template.md)。

## 📏 编码标准

### Go 风格指南

我们遵循标准的 Go 风格指南，并有一些补充：

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

### 命名约定

- **包名**：简短、小写、尽可能单词
- **函数**：CamelCase，描述性
- **变量**：本地变量 camelCase，导出变量 CamelCase
- **常量**：导出 CamelCase，未导出 camelCase
- **接口**：通常以 "er" 结尾（如 `Writer`、`Formatter`）

### 代码组织

```
project/
├── docs/           # 多语言文档
├── .github/        # GitHub 模板和工作流
├── logger.go       # 主日志实现
├── entry.go        # 日志条目结构
├── level.go        # 日志级别
├── formatter.go    # 日志格式化
├── output.go       # 输出管理
└── *_test.go      # 测试文件与源码同目录
```

### 错误处理

```go
// ✅ 推荐：返回错误，不要 panic
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ 避免：在库代码中 panic
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // 不要这样做
    }
    return &Logger{...}
}
```

## 🧪 测试指南

### 测试结构

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
            // 测试实现
        })
    }
}
```

### 覆盖率要求

- **最低要求**：新代码 90% 覆盖率
- **目标**：整体覆盖率 95%+
- **所有构建标签**必须保持覆盖率
- 使用 `make coverage-all` 验证

### 测试命令

```bash
# 快速测试所有构建标签
make test-quick

# 带覆盖率的完整测试套件
make test-all

# 覆盖率报告
make coverage-html

# 基准测试
make benchmark
```

## 🏗️ 构建标签要求

所有更改必须与我们的构建标签系统兼容：

### 支持的构建标签

- **默认** (`go build`)：完整功能
- **调试** (`go build -tags debug`)：增强调试
- **发布** (`go build -tags release`)：生产优化
- **丢弃** (`go build -tags discard`)：最大性能

### 构建标签测试

```bash
# 测试每个构建配置
make test-default
make test-debug  
make test-release
make test-discard

# 测试组合
make test-debug-discard
make test-release-discard

# 全部测试
make test-all
```

### 构建标签指南

```go
//go:build debug
// +build debug

package log

// 调试特定实现
```

## 📚 文档编写

### 代码文档

- **所有导出函数**必须有清晰注释
- **复杂算法**需要解释
- **非平凡用法**需要示例
- **线程安全性**说明（如适用）

```go
// SetLevel 设置最小日志级别。
// 低于此级别的日志将被忽略。
// 此方法是线程安全的。
//
// 示例:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // 不会输出
//   logger.Info("visible")   // 会输出
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### README 更新

添加功能时，请更新：
- 主 README.md
- `docs/` 中所有语言特定的 README
- 代码示例
- 功能列表

## 🐛 问题提交指南

### 错误报告

使用[错误报告模板](.github/ISSUE_TEMPLATE/bug_report.md)并包含：

- **清晰描述**问题
- **重现步骤**
- **期望与实际行为**
- **环境详情**（操作系统、Go 版本、构建标签）
- **最小代码示例**

### 功能请求

使用[功能请求模板](.github/ISSUE_TEMPLATE/feature_request.md)并包含：

- **清晰的功能动机**
- **建议的 API** 设计
- **实现考虑**
- **破坏性更改分析**

### 问题咨询

使用[问题模板](.github/ISSUE_TEMPLATE/question.md)用于：

- 使用问题
- 配置帮助
- 最佳实践
- 集成指导

## 🚀 性能考虑

### 基准测试

始终对性能敏感的更改进行基准测试：

```bash
# 运行基准测试
go test -bench=. -benchmem

# 前后对比
go test -bench=. -benchmem > before.txt
# 进行更改
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### 性能指南

- **最小化热路径**中的分配
- **使用对象池**处理频繁创建的对象
- **提前返回**对禁用的日志级别
- **避免反射**在性能关键代码中
- **先分析再优化**

### 内存管理

```go
// ✅ 良好：使用对象池
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

### 敏感数据

- **永不记录**密码、令牌或敏感数据
- **清理**日志消息中的用户输入
- **避免**记录完整的请求/响应主体
- **使用**结构化日志获得更好的控制

```go
// ✅ 良好
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ 不好
logger.Infof("User login: %+v", userRequest) // 可能包含密码
```

### 依赖管理

- 保持依赖**最新**
- **仔细审查**新依赖
- **最小化**外部依赖
- **使用** `go mod verify` 检查完整性

## 👥 社区

### 获取帮助

- 📖 [文档](../README_zh-CN.md)
- 💬 [GitHub 讨论](https://github.com/lazygophers/log/discussions)
- 🐛 [问题跟踪器](https://github.com/lazygophers/log/issues)
- 📧 邮箱：support@lazygophers.com

### 交流指南

- **尊重**和包容
- **提问前先搜索**
- **请求帮助时提供上下文**
- **力所能及地帮助他人**
- **遵循**[行为准则](CODE_OF_CONDUCT_zh-CN.md)

## 🎯 贡献认可

贡献者通过以下方式获得认可：

- **README 贡献者**部分
- **发布说明**提及
- **GitHub 贡献者**图表
- **社区感谢**帖子

## 📝 许可证

通过贡献，您同意您的贡献将在 MIT 许可证下获得许可。

---

## 🌍 多语言文档

本文档提供多种语言版本：

- [🇺🇸 English](CONTRIBUTING.md)
- [🇨🇳 简体中文](CONTRIBUTING_zh-CN.md)（当前）
- [🇹🇼 繁體中文](CONTRIBUTING_zh-TW.md)
- [🇫🇷 Français](CONTRIBUTING_fr.md)
- [🇷🇺 Русский](CONTRIBUTING_ru.md)
- [🇪🇸 Español](CONTRIBUTING_es.md)
- [🇸🇦 العربية](CONTRIBUTING_ar.md)

---

**感谢您为 LazyGophers Log 贡献代码！🚀**