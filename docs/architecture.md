# 架构文档

## 概述

lazygophers/log 是一个高性能、灵活的 Go 日志库，专注于简洁的 API 和强大的扩展能力。本文档详细说明了系统的架构设计、核心组件和性能优化策略。

## 核心设计原则

1. **性能优先**: 在日志热路径中，性能是首要考虑因素
2. **简洁 API**: 提供直观、易用的 API，隐藏复杂性
3. **可扩展性**: 通过接口支持自定义格式化器、Hook 和输出
4. **类型安全**: 利用 Go 的类型系统确保编译时安全
5. **零依赖循环**: 通过 constant 包避免循环依赖

## 项目结构

```
log/
├── constant/              # 核心类型和接口定义
│   ├── entry.go          # 日志条目结构
│   ├── interface.go      # Hook、Format 等接口
│   └── level.go          # 日志级别定义
├── logger.go             # 主日志器实现
├── formatter.go          # 文本格式化器
├── formatter_json.go     # JSON 格式化器
├── rotator.go            # 日志轮转
├── pool.go               # 对象池
├── init.go               # 全局日志器
├── output.go             # 输出管理
└── *_test.go             # 测试文件
```

## 核心组件

### 1. Logger（日志器）

**位置**: `logger.go`

**职责**:
- 提供日志记录 API
- 管理配置（级别、输出、格式化器）
- 协调 Entry 创建和处理
- 执行 Hook 链

**关键方法**:
```go
type Logger struct {
    level       Level
    out         zapcore.WriteSyncer
    Format      constant.Format
    callerDepth int
    PrefixMsg   []byte
    SuffixMsg   []byte
    enableCaller bool
    enableTrace  bool
    hooks       []constant.Hook
}

// 日志记录方法
func (p *Logger) Log(level Level, args ...interface{})
func (p *Logger) Logf(level Level, format string, args ...interface{})
func (p *Logger) Infow(msg string, args ...interface{})

// 配置方法（支持链式调用）
func (p *Logger) SetLevel(level Level) *Logger
func (p *Logger) SetOutput(writes ...io.Writer) *Logger
func (p *Logger) EnableCaller(enable bool) *Logger
func (p *Logger) AddHook(hook constant.Hook) *Logger
```

**性能优化**:
- 早期级别检查：`levelEnabled()` 在外层完成
- 条件字段填充：只在需要时填充调用者和追踪信息
- 对象池：Entry 从池中获取和归还

### 2. Entry（日志条目）

**位置**: `constant/entry.go`

**职责**:
- 表示单条日志记录
- 存储所有日志元数据
- 支持对象池复用

**内存布局优化**:
```go
type Entry struct {
    // 热路径字段 - 每次日志调用都访问
    Level      Level     // 4 字节
    Pid        int       // 4 字节
    Gid        int64     // 8 字节
    CallerLine int       // 4 字节

    // 时间戳 - 访问频繁但低于核心字段
    Time       time.Time // 16 字节
    TimeStr    string    // 16 字节（缓存格式化时间）
    TimeStrSet bool      // 1 字节

    // 字符串字段（各 16 字节）- 按访问频率排序
    Message    string    // 16 字节
    File       string    // 16 字节
    CallerFunc string    // 16 字节
    CallerDir  string    // 16 字节
    CallerName string    // 16 字节
    TraceId    string    // 16 字节

    // 字节切片（各 24 字节）- 较低频率
    PrefixMsg []byte     // 24 字节
    SuffixMsg []byte     // 24 字节

    // 结构化字段
    Fields []KV          // 动态大小
}
```

**总计**: 200 字节，仅 4 字节填充浪费（2%）

**关键特性**:
- 实现 `json.Marshaler` 接口，支持条件字段序列化
- `Reset()` 方法支持对象池复用
- 时间戳缓存避免重复格式化

### 3. Level（日志级别）

**位置**: `constant/level.go`

**设计**:
```go
type Level uint32

const (
    PanicLevel Level = iota  // 最高
    FatalLevel
    ErrorLevel
    WarnLevel
    InfoLevel
    DebugLevel
    TraceLevel                // 最低
)
```

**性能优化**:
```go
// 使用数组查找而非 switch，提升性能
var levelStrings = []string{
    PanicLevel: "panic",
    FatalLevel: "fatal",
    ErrorLevel: "error",
    WarnLevel:  "warn",
    InfoLevel:  "info",
    DebugLevel: "debug",
    TraceLevel: "trace",
}

func (level Level) String() string {
    if level >= 0 && int(level) < len(levelStrings) {
        return levelStrings[level]
    }
    return "trace"
}
```

### 4. Formatter（格式化器）

**位置**: `formatter.go`, `formatter_json.go`

**接口定义**:
```go
// 基本格式化接口
type Format interface {
    Format(entry interface{}) []byte
}

// 扩展格式化接口
type FormatFull interface {
    Format
    ParsingAndEscaping(disable bool)
    Caller(disable bool)
    Clone() Format
}
```

**文本格式化器** (`Formatter`):
- 彩色输出（支持 ANSI 颜色）
- 结构化字段支持（key=value 格式）
- 可配置的消息解析和转义
- 可配置的调用者信息

**JSON 格式化器** (`JSONFormatter`):
- 紧凑/美化打印模式
- 条件字段序列化（使用 `MarshalJSON`）
- 特殊字符转义
- 错误回退机制

### 5. Hook（钩子系统）

**位置**: `constant/interface.go`

**接口定义**:
```go
type Hook interface {
    OnWrite(entry interface{}) interface{}
}
```

**使用场景**:
- **过滤**: 返回 `nil` 跳过日志
- **修改**: 修改 Entry 内容
- **丰富**: 添加额外字段
- **集成**: 发送到第三方服务

**实现示例**:
```go
// 环境信息 Hook
type EnvironmentHook struct {
    env string
}

func (h *EnvironmentHook) OnWrite(entry interface{}) interface{} {
    if e, ok := entry.(*log.Entry); ok {
        e.Fields = append(e.Fields, log.KV{
            Key:   "environment",
            Value: h.env,
        })
    }
    return entry
}

// 敏感信息过滤 Hook
type SensitiveDataHook struct{}

func (h *SensitiveDataHook) OnWrite(entry interface{}) interface{} {
    if e, ok := entry.(*log.Entry); ok {
        if strings.Contains(e.Message, "password") {
            return nil // 过滤掉
        }
    }
    return entry
}
```

### 6. Rotator（日志轮转）

**位置**: `rotator.go`

**功能**:
- 按小时自动轮转
- 基于大小的分片（单文件超过限制时创建分片）
- 自动清理过期文件

**文件命名**:
```
<目录>/<YYYYMMDDHH>.log          # 主文件
<目录>/<YYYYMMDDHH>.log.1        # 分片 1
<目录>/<YYYYMMDDHH>.log.2        # 分片 2
...
```

**使用示例**:
```go
// 创建轮转写入器
// 参数：目录、单文件最大大小（字节）、保留文件数量
writer := GetOutputWriterHourly("/var/log/app", 100*1024*1024, 168)

logger := New()
logger.SetOutput(writer)
```

## 性能优化策略

### 1. 对象池

**Entry 池**:
```go
var entryPool = sync.Pool{
    New: func() interface{} {
        return &Entry{}
    },
}

func GetEntry() *Entry {
    return entryPool.Get().(*Entry)
}

func PutEntry(entry *Entry) {
    entry.Reset()
    entryPool.Put(entry)
}
```

**Buffer 池**:
```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func GetBuffer() *bytes.Buffer {
    b := bufferPool.Get().(*bytes.Buffer)
    b.Reset()
    return b
}

func PutBuffer(b *bytes.Buffer) {
    bufferPool.Put(b)
}
```

### 2. 早期检查

```go
func (p *Logger) Log(level Level, args ...interface{}) {
    // 早期级别检查，避免不必要的计算
    if !p.levelEnabled(level) {
        return
    }
    p.log(level, fastSprint(args...))
}

func (p *Logger) levelEnabled(level Level) bool {
    return level >= p.level
}
```

### 3. 内联优化

```go
//go:inline
func (p *Logger) fillTraceInfo(entry *Entry) {
    if p.enableTrace {
        entry.Gid = goid.Get()
        entry.TraceId = getTrace(entry.Gid)
    }
}
```

### 4. 时间戳缓存

```go
// 在 populateEntry 中缓存格式化的时间戳
entry.Time = time.Now()
entry.TimeStr = entry.Time.Format(time.RFC3339Nano)
entry.TimeStrSet = true
```

## 接口设计

### constant 包

**目的**: 避免循环依赖，提供独立的接口定义

**核心接口**:
```go
// Hook - 日志处理扩展
type Hook interface {
    OnWrite(entry interface{}) interface{}
}

// Format - 基本格式化
type Format interface {
    Format(entry interface{}) []byte
}

// FormatFull - 扩展格式化
type FormatFull interface {
    Format
    ParsingAndEscaping(disable bool)
    Caller(disable bool)
    Clone() Format
}

// Writer - 输出目标
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

**使用 `interface{}` 的原因**:
- 避免循环依赖
- 支持第三方实现
- 类型断言在实现内部完成

### 依赖关系

```
┌─────────────────────────────────┐
│  constant (独立)                 │
│  - Hook, Format, FormatFull     │
│  - Entry, Level                  │
│  - 无依赖                        │
└──────────┬──────────────────────┘
           │
           ├─────────────────┐
           │                 │
           ▼                 ▼
┌──────────────────┐  ┌──────────────────┐
│  log             │  │  第三方扩展      │
│  - Logger        │  │  - Hook 实现     │
│  - Formatter     │  │  - Format 实现   │
│  - Rotator       │  └──────────────────┘
│  - 使用 constant │
└──────────────────┘
```

## 日志流程

### 基本日志流程

```
1. 用户调用: logger.Info("message")
            ↓
2. 级别检查: levelEnabled(InfoLevel)
            ↓
3. 获取 Entry: entryPool.Get()
            ↓
4. 填充基础字段: populateEntry()
            ↓
5. 条件填充: fillTraceInfo(), fillCallerInfo()
            ↓
6. 执行 Hook 链: hooks[i].OnWrite(entry)
            ↓
7. 格式化: formatter.Format(entry)
            ↓
8. 写入输出: out.Write(formatted)
            ↓
9. 归还 Entry: entryPool.Put(entry)
```

### 结构化日志流程

```
1. 用户调用: logger.Infow("message", "key", "value")
            ↓
2. 级别检查: levelEnabled(InfoLevel)
            ↓
3. 获取 Entry: entryPool.Get()
            ↓
4. 填充基础字段: populateEntry()
            ↓
5. 解析键值对: populateFields()
            ↓
6. 条件填充: fillTraceInfo(), fillCallerInfo()
            ↓
7. 执行 Hook 链: hooks[i].OnWrite(entry)
            ↓
8. 格式化: formatter.Format(entry)
            ↓
9. 写入输出: out.Write(formatted)
            ↓
10. 归还 Entry: entryPool.Put(entry)
```

## 并发模型

### Logger 线程安全

- Logger 实例设计为线程安全
- 内部使用 `sync.Mutex` 保护共享状态
- Hook 执行按顺序进行

### Entry 非线程安全

- Entry 对象不是线程安全的
- 设计为单次使用（从池中获取 → 使用 → 归还）
- 不应保留 Entry 引用

### Rotator 线程安全

- 使用 `sync.Mutex` 保护文件操作
- 支持并发写入

## 扩展点

### 1. 自定义格式化器

```go
type MyFormatter struct{}

func (f *MyFormatter) Format(entry interface{}) []byte {
    e, ok := entry.(*log.Entry)
    if !ok {
        return nil
    }
    // 自定义格式化逻辑
    return []byte(fmt.Sprintf("[%s] %s\n", e.Level, e.Message))
}

// 使用
logger := log.New()
logger.Format = &MyFormatter{}
```

### 2. 自定义 Hook

```go
type MyHook struct{}

func (h *MyHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }
    // 处理或修改 entry
    e.Message = "[HOOK] " + e.Message
    return e // 返回 nil 跳过
}

// 使用
logger.AddHook(&MyHook{})
```

### 3. 自定义输出

```go
type MyWriter struct{}

func (w *MyWriter) Write(p []byte) (n int, err error) {
    // 自定义写入逻辑
    return os.Stdout.Write(p)
}

// 使用
logger := log.New()
logger.SetOutput(&MyWriter{})
```

## 最佳实践

### 1. 使用对象池

始终为频繁分配的对象使用对象池：
- Entry 对象
- Buffer 对象
- 其他临时对象

### 2. 早期检查

在昂贵的操作前检查日志级别：
```go
if !logger.levelEnabled(level) {
    return // 早期返回
}
// 执行昂贵的操作
```

### 3. 避免 Hook 阻塞

Hook 中应避免阻塞操作：
```go
func (h *MyHook) OnWrite(entry interface{}) interface{} {
    // ❌ 避免阻塞
    // http.Post(url, entry)
    
    // ✅ 使用 goroutine
    go func() {
        http.Post(url, entry)
    }()
    
    return entry
}
```

### 4. 注意内存泄漏

不要保留 Entry 引用：
```go
// ❌ 错误：保留引用
var cachedEntry *log.Entry
func logSomething() {
    entry := getEntry()
    cachedEntry = entry // 错误！
}

// ✅ 正确：及时归还
func logSomething() {
    entry := getEntry()
    defer putEntry(entry)
    // 使用 entry
}
```

## 性能基准

典型性能指标（基于基准测试）：

```
BenchmarkLoggerInfo-8           5000000    280 ns/op    120 B/op    3 allocs/op
BenchmarkLoggerInfoDisabled-8 100000000    10.5 ns/op     0 B/op    0 allocs/op
BenchmarkJSONFormatter-8        3000000    450 ns/op    280 B/op    5 allocs/op
BenchmarkTextFormatter-8       5000000    320 ns/op    180 B/op    4 allocs/op
```

**关键指标**:
- 禁用级别的日志几乎零开销（10.5 ns）
- 启用级别时，每次日志约 280-450 ns
- 内存分配极小（120-280 B）

## 未来改进

### 计划中的特性

1. **结构化日志增强**: 支持更多结构化字段类型
2. **性能优化**: 进一步减少内存分配
3. **更多输出格式**: 支持 YAML、XML 等格式
4. **日志采样**: 高频场景下的日志采样
5. **插件系统**: 更强大的扩展机制

### 欢迎贡献

欢迎提交 PR 和 Issue！请参阅[贡献指南](../CONTRIBUTING.md)。

---

**记住**: 这是一个专注于性能的日志库。任何更改都应考虑对日志热路径的性能影响。
