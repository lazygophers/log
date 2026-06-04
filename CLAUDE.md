# CLAUDE.md

本文件为 Claude Code (claude.ai/code) 在本仓库中工作时提供指导。

## 项目概述

**lazygophers/log** - 高性能 Go 日志库。7 级日志、sync.Pool 对象池、Hook 扩展、按小时轮转。
主包覆盖率 ~92%，hooks 包 ~99%，总计 546 测试用例。

## 快速参考

### 开发命令

```bash
go test ./...                    # 运行所有测试
go test -v ./...                 # 详细输出
go test -bench=. -benchmem       # 基准测试
go test -cover ./...             # 覆盖率
go build ./...                   # 构建
go vet ./...                     # 静态分析
```

### 项目结构

```
log/
├── constant/              # 核心类型（Entry, Level, Hook, Format 接口）
├── hooks/                 # 预置 Hook 实现（独立 module）
├── logctx/                # context-aware Logger（独立 module）
├── zap/                   # zap WriteSyncer 适配（独立 module）
├── logger.go              # Logger 主结构 + 日志方法
├── formatter.go           # 文本格式化器
├── formatter_json.go      # JSON 格化器
├── entry.go / pool.go     # Entry 对象池
├── rotator.go             # 按小时 + 按大小轮转
├── writer_async.go        # 异步 Writer
├── output.go              # 输出管理 + GetOutputWriterHourly
├── trace.go               # Trace ID 管理（sync.Map）
├── utils.go               # fastSprint/fastSprintf 快速路径
├── init.go                # 全局 std logger
├── env.go                 # APP_ENV 自动级别设置
├── print.go               # 全局函数代理（Info/Warn/Error 等）
├── print_funcs.go         # build tag: debug 输出到 stdout
├── print_funcs_release.go # build tag: release 输出到文件
├── print_funcs_discard.go # build tag: discard 丢弃
├── config.go              # 默认配置常量
└── doc.go                 # 包文档
```

### 模块关系

`constant`（独立 pkg，无外部依赖）← `log`（主包）← `hooks`/`logctx`/`zap`（独立 module）

### 深入文档

| 文档 | 内容 |
|------|------|
| [docs/architecture.md](docs/architecture.md) | 完整架构设计、性能优化策略、日志流程图、并发模型 |
| [docs/hooks_guide.md](docs/hooks_guide.md) | Hook 使用教程、7 个使用场景、最佳实践 |

## 代码规范

### 测试规范（严格遵守）

- 测试文件名 = 源文件名 + `_test.go`
- **禁止**独立覆盖率文件（`coverage_boost_test.go`）
- **禁止**功能特定测试文件（`formatter_coverage_test.go`）
- 测试覆盖率保持 90% 以上
- 使用表驱动测试

### 代码风格

- `go fmt` 格式化
- 热点路径 `//go:inline`，栈深度敏感 `//go:noinline`
- 注释只写 WHY 不写 WHAT

### 踩坑 / 红线

- **Hook 断言必须用 `*constant.Entry`**：`hooks/` 包的 Hook 通过 `OnWrite(entry interface{})` 接收 `*constant.Entry`，断言必须写 `entry.(*constant.Entry)`，禁止用本地 struct 做 entryLike 断言
- **Entry.Reset() 必须清理所有字段**：池化复用时，`Fields`/`PrefixMsg`/`SuffixMsg` 必须 `[:0]`，否则残留数据泄露到后续日志行
- **Hook 过滤后必须归池**：`applyHooks` 返回 nil 时，原始 entry 仍需 `putEntry`，否则池泄漏
- **Logger 不用 Mutex**：主 Logger 热路径无锁，靠 early level check + pool 保证安全；子包 `logctx` 用 Mutex 持锁做 I/O（不同设计）
- **Level 排序**：`PanicLevel=0`(最高) → `TraceLevel=6`(最低)；`levelEnabled` = `p.level >= level`

## 常见任务

### 添加新日志级别

1. `constant/level.go` 添加常量 + 更新 `levelStrings`
2. `logger.go` 添加级别方法（Xxx/Xxxf/Xxxw 三件套）
3. `print.go` 添加全局函数代理

### 实现自定义 Hook

```go
type MyHook struct{}
func (h *MyHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*constant.Entry)  // 必须用这个断言
    if !ok { return entry }
    // 处理...
    return e // nil = 跳过
}
```

### 实现自定义格式化器

```go
type MyFormatter struct{}
func (f *MyFormatter) Format(entry interface{}) []byte {
    e := entry.(*constant.Entry)
    return []byte(fmt.Sprintf("[%s] %s\n", e.Level, e.Message))
}
```

## 重要注意事项

- Entry 从池获取，用完自动归还——不要保留引用
- Logger 线程安全，Entry 不线程安全（单次使用）
- 扩展点：`Format` 接口、`Hook` 接口、`io.Writer`
- `hooks/`/`logctx/`/`zap/` 是独立 Go module，有各自 go.mod，测试需单独 cd 进去跑

---

**记住**: 性能日志库。所有改动考虑热路径影响。
