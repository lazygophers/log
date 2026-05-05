# CLAUDE.md

本文件为 Claude Code (claude.ai/code) 在本仓库中工作时提供指导。

## 项目概述

**lazygophers/log** - 高性能 Go 日志库，专注于简洁 API 和强大扩展能力。

### 核心特性

- **测试覆盖率**: 95.0% (486 个测试用例)
- **性能优化**: 基于 `sync.Pool` 的对象池、缓存友好结构布局、早期级别检查
- **扩展机制**: Hook 接口、可插拔格式化器、多 Writer 支持
- **日志轮转**: 按小时轮转 + 基于大小分片

## 快速参考

### 开发命令

**测试:**
```bash
go test ./...                    # 运行所有测试
go test -v ./...                 # 详细输出
go test -bench=. -benchmem       # 基准测试
go test -cover ./...             # 覆盖率报告
```

**构建和检查:**
```bash
go build ./...                   # 构建所有包
go vet ./...                     # 静态分析
go fmt ./...                     # 格式化代码
go mod tidy                      # 整理依赖
```

## 架构设计

### 核心组件

**Logger** (`logger.go`):
- 主日志结构，支持链式配置
- 使用 `zapcore.WriteSyncer` 作为输出接口
- Hook 系统支持日志处理扩展
- 级别检查在外层进行，避免不必要的计算

**Entry** (`constant/entry.go`):
- 日志条目结构，优化的内存布局（200 字节，仅 2% 填充浪费）
- 字段按访问频率和大小排列，最小化缓存行浪费
- 实现 `json.Marshaler` 接口，支持自定义 JSON 序列化
- 使用 `Reset()` 方法支持对象池复用

**Level** (`constant/level.go`):
- 七个日志级别：Panic、Fatal、Error、Warn、Info、Debug、Trace
- 兼容 logrus 级别约定
- 使用数组查找优化 `String()` 方法性能

**Formatter** (`formatter.go`, `formatter_json.go`):
- `Format` 接口：基本格式化
- `FormatFull` 接口：扩展控制（解析/转义、调用者信息、克隆）
- 默认文本格式化器：彩色输出、结构化字段支持
- JSON 格式化器：紧凑/美化打印、条件字段序列化

**Rotator** (`rotator.go`):
- 按小时自动轮转日志文件
- 基于大小的分片支持（单文件超过限制时创建分片）
- 自动清理过期文件

**Hook** (`constant/interface.go`):
- `OnWrite(entry) interface{}` 方法：在日志写入前处理
- 可以修改、过滤或丰富日志条目
- 返回 `nil` 跳过该条日志

### 接口设计

```go
// constant/interface.go

type Hook interface {
    OnWrite(entry interface{}) interface{}
}

type Format interface {
    Format(entry interface{}) []byte
}

type FormatFull interface {
    Format
    ParsingAndEscaping(disable bool)
    Caller(disable bool)
    Clone() Format
}
```

### 性能优化模式

- **对象池**: Entry 和 Buffer 使用 `sync.Pool` 复用
- **早期检查**: 级别检查在 `levelEnabled()` 中完成，避免昂贵的 `populateEntry` 等操作
- **内联函数**: 热点路径使用 `//go:inline` 标记
- **时间缓存**: `TimeStr` 和 `TimeStrSet` 缓存格式化时间戳
- **条件字段**: 调用者和追踪信息按需填充

### 依赖管理

**主要依赖:**
- `go.uber.org/zap` - WriteSyncer 接口和多 Writer 支持
- `github.com/petermattis/goid` - Goroutine ID 提取
- `github.com/lestrrat-go/file-rotatelogs` - 时间日志轮转

**内部包:**
- `github.com/lazygophers/log/constant` - 核心类型和接口定义

## 代码规范

### 测试规范

**测试文件命名（严格限制）:**
- 测试文件名必须是源文件名 + `_test.go`
- **禁止**创建独立的覆盖率测试文件（如 `coverage_boost_test.go`）
- **禁止**创建功能特定的测试文件（如 `formatter_coverage_test.go`）
- 所有测试必须合并到对应的标准测试文件中

**测试文件组织:**
```
logger.go           → logger_test.go
formatter.go        → formatter_test.go
rotator.go          → rotator_test.go
constant/entry.go   → entry_test.go
```

**测试要求:**
- 所有公共 API 必须有测试
- 性能关键代码需要基准测试
- 测试覆盖率保持在 95% 以上
- 使用表驱动测试（table-driven tests）

### 代码风格

- 遵循 Go 标准格式化（`go fmt`）
- 使用描述性的变量和函数名
- 复杂逻辑添加注释
- 保持一致的错误处理模式
- 使用 `//go:inline` 标记热点路径的函数

### 性能考虑

- 始终为频繁分配的对象使用对象池
- 在昂贵的操作前检查日志级别
- 为性能关键代码路径编写基准测试
- 注意 Entry 结构的缓存行对齐

## 重要设计决策

### Entry 移至 constant 包

**原因:**
- Entry 是日志系统的核心数据结构
- 移至 constant 包避免循环依赖
- 便于其他包引用而不依赖主包

**影响:**
- 主包通过类型别名 `type Entry = constant.Entry` 导出
- 所有 Entry 操作保持不变
- 测试文件需要导入 `constant` 包

### Hook 系统

**设计:**
```go
type Hook interface {
    OnWrite(entry interface{}) interface{}
}
```

**使用场景:**
- 添加全局字段（环境、版本）
- 过滤敏感信息
- 修改日志内容
- 集成第三方服务

**注意:**
- Hook 返回 `nil` 跳过日志
- 多 Hook 按添加顺序执行
- Hook 中应避免阻塞操作

### Logger Clone

**深度复制:**
```go
func (p *Logger) Clone() *Logger {
    l := Logger{
        level:        p.level,
        out:          p.out,
        callerDepth:  p.callerDepth,
        PrefixMsg:    p.PrefixMsg,
        SuffixMsg:    p.SuffixMsg,
        enableCaller: p.enableCaller,
        enableTrace:  p.enableTrace,
    }

    // Clone formatter if it supports cloning
    switch f := p.Format.(type) {
    case constant.FormatFull:
        l.Format = f.Clone()
    default:
        l.Format = f
    }

    // Copy hooks
    if len(p.hooks) > 0 {
        l.hooks = make([]constant.Hook, len(p.hooks))
        copy(l.hooks, p.hooks)
    }

    return &l
}
```

**使用场景:**
- 创建独立配置的日志器
- 线程安全的日志器副本
- 隔离的测试环境

## 常见任务

### 添加新的日志级别

1. 在 `constant/level.go` 添加常量
2. 更新 `levelStrings` 数组
3. 在所有相关文件中添加级别方法

### 实现自定义格式化器

```go
type MyFormatter struct{}

func (f *MyFormatter) Format(entry interface{}) []byte {
    e := entry.(*constant.Entry)
    // 自定义格式化逻辑
    return []byte(fmt.Sprintf("[%s] %s\n", e.Level, e.Message))
}

// 使用
logger := log.New()
logger.Format = &MyFormatter{}
```

### 实现自定义 Hook

```go
type MyHook struct{}

func (h *MyHook) OnWrite(entry interface{}) interface{} {
    e := entry.(*constant.Entry)
    // 处理或修改 entry
    return e // 返回 nil 跳过
}

// 使用
logger.AddHook(&MyHook{})
```

## 重要注意事项

### 内存管理

- Entry 对象从池中获取，使用后自动归还
- 不要保留 Entry 对象的引用
- 注意 Hook 中的内存泄漏

### 并发安全

- Logger 实例设计为线程安全
- Entry 对象不是线程安全的（单次使用）
- 扩展功能时使用适当的同步

### 扩展点

- `Format` 接口：自定义格式化
- `Hook` 接口：日志处理扩展
- `io.Writer` 实现：自定义输出

## 故障排查

### 性能问题

1. 检查是否启用了不必要的调用者/追踪信息
2. 验证日志级别设置是否合理
3. 运行基准测试识别瓶颈
4. 检查 Hook 中是否有阻塞操作

### 内存问题

1. 确认没有保留 Entry 对象引用
2. 检查 Hook 中的内存泄漏
3. 验证对象池配置
4. 使用 pprof 分析内存分配

### 测试问题

1. 确认测试文件命名正确
2. 检查是否正确导入 `constant` 包
3. 验证接口实现是否完整
4. 使用 `-v` 标志查看详细输出

## 获取帮助

- 查看 [API 文档](docs/en/API.md)
- 参考 [架构文档](docs/architecture.md)
- 阅读 [Hook 指南](docs/hooks_guide.md)
- 在 [Issues](https://github.com/lazygophers/log/issues) 报告问题

---

**记住**: 这是一个专注于性能的日志库。在任何更改中都要考虑对日志操作热路径的性能影响。
