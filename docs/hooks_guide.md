# Hook 使用指南

## 概述

Hook 是 lazygophers/log 提供的强大扩展机制，允许在日志写入前拦截、修改、过滤或丰富日志条目。

## Hook 接口

### 定义

```go
package constant

// Hook 定义日志处理钩子接口
type Hook interface {
    // OnWrite 在日志写入前处理日志条目
    // 返回修改后的条目或 nil 跳过该条日志
    OnWrite(entry interface{}) interface{}
}
```

### 便捷实现

```go
// HookFunc 是函数类型的 Hook 便捷实现
type HookFunc func(entry interface{}) interface{}

func (h HookFunc) OnWrite(entry interface{}) interface{} {
    return h(entry)
}
```

## 基本用法

### 添加 Hook

```go
import (
    "github.com/lazygophers/log"
    "github.com/lazygophers/log/constant"
)

logger := log.New()

// 添加 Hook
logger.AddHook(&MyHook{})

// 或使用 HookFunc
logger.AddHook(constant.HookFunc(func(entry interface{}) interface{} {
    // 处理逻辑
    return entry
}))
```

### Hook 返回值

- **返回 `entry`**: 正常记录日志
- **返回修改后的 `entry`**: 记录修改后的日志
- **返回 `nil`**: 跳过该条日志

## 使用场景

### 1. 添加全局字段

```go
// EnvironmentHook 添加环境信息
type EnvironmentHook struct {
    Environment string
    Version     string
}

func (h *EnvironmentHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    // 添加全局字段
    e.Fields = append(e.Fields, log.KV{
        Key:   "environment",
        Value: h.Environment,
    }, log.KV{
        Key:   "version",
        Value: h.Version,
    })

    return e
}

// 使用
logger := log.New()
logger.AddHook(&EnvironmentHook{
    Environment: "production",
    Version:     "1.0.0",
})

logger.Info("服务启动")
// 输出: ... service=start environment=production version=1.0.0
```

### 2. 过滤敏感信息

```go
// SensitiveDataHook 过滤敏感数据
type SensitiveDataHook struct {
    keywords []string
}

func NewSensitiveDataHook() *SensitiveDataHook {
    return &SensitiveDataHook{
        keywords: []string{"password", "token", "secret", "api_key"},
    }
}

func (h *SensitiveDataHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    // 检查消息中是否包含敏感关键词
    msg := strings.ToLower(e.Message)
    for _, keyword := range h.keywords {
        if strings.Contains(msg, keyword) {
            return nil // 过滤掉
        }
    }

    // 检查字段中是否包含敏感信息
    for i := len(e.Fields) - 1; i >= 0; i-- {
        key := strings.ToLower(e.Fields[i].Key)
        for _, keyword := range h.keywords {
            if strings.Contains(key, keyword) {
                // 移除敏感字段
                e.Fields = append(e.Fields[:i], e.Fields[i+1:]...)
                break
            }
        }
    }

    return e
}

// 使用
logger := log.New()
logger.AddHook(NewSensitiveDataHook())

logger.Infow("用户登录", "username", "admin", "password", "secret123")
// 输出: ... 用户登录 username=admin (password 字段被过滤)
```

### 3. 修改日志内容

```go
// ModifyHook 修改日志内容
type ModifyHook struct{}

func (h *ModifyHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    // 添加前缀
    e.Message = "[PROCESSED] " + e.Message

    // 添加额外字段
    e.Fields = append(e.Fields, log.KV{
        Key:   "processed_by",
        Value: "ModifyHook",
    })

    return e
}

// 使用
logger := log.New()
logger.AddHook(&ModifyHook{})

logger.Info("原始消息")
// 输出: ... [PROCESSED] 原始消息 processed_by=ModifyHook
```

### 4. 条件过滤

```go
// LevelFilterHook 按级别过滤
type LevelFilterHook struct {
    MinLevel log.Level
}

func (h *LevelFilterHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    // 只记录指定级别及以上的日志
    if e.Level < h.MinLevel {
        return nil
    }

    return e
}

// 使用
logger := log.New()
logger.SetLevel(log.TraceLevel) // 日志器设置为最低级别
logger.AddHook(&LevelFilterHook{MinLevel: log.WarnLevel}) // Hook 过滤

logger.Trace("这条不会记录")
logger.Debug("这条也不会记录")
logger.Info("这条还是不会记录")
logger.Warn("这条会被记录")  // 通过
logger.Error("这条也会被记录") // 通过
```

### 5. 异步处理

```go
// AsyncHook 异步发送日志到远程服务
type AsyncHook struct {
    url     string
    client  *http.Client
    buffer  chan interface{}
    wg      sync.WaitGroup
}

func NewAsyncHook(url string) *AsyncHook {
    h := &AsyncHook{
        url:    url,
        client: &http.Client{Timeout: 5 * time.Second},
        buffer: make(chan interface{}, 1000),
    }

    // 启动后台处理
    h.wg.Add(1)
    go h.process()

    return h
}

func (h *AsyncHook) OnWrite(entry interface{}) interface{} {
    select {
    case h.buffer <- entry:
        // 成功发送到缓冲区
    default:
        // 缓冲区满，丢弃
    }

    return entry // 不影响原始日志
}

func (h *AsyncHook) process() {
    defer h.wg.Done()

    for entry := range h.buffer {
        e := entry.(*log.Entry)

        // 序列化
        data, _ := json.Marshal(e)

        // 发送
        req, _ := http.NewRequest("POST", h.url, bytes.NewReader(data))
        req.Header.Set("Content-Type", "application/json")
        resp, err := h.client.Do(req)

        if err != nil {
            // 处理错误
            continue
        }
        resp.Body.Close()
    }
}

func (h *AsyncHook) Close() {
    close(h.buffer)
    h.wg.Wait()
}

// 使用
hook := NewAsyncHook("https://logs.example.com/api/logs")
defer hook.Close()

logger := log.New()
logger.AddHook(hook)
```

### 6. 上下文感知

```go
// ContextHook 添加请求上下文信息
type ContextHook struct {
    context map[string]interface{}
    mu      sync.RWMutex
}

func NewContextHook() *ContextHook {
    return &ContextHook{
        context: make(map[string]interface{}),
    }
}

func (h *ContextHook) Set(key string, value interface{}) {
    h.mu.Lock()
    defer h.mu.Unlock()
    h.context[key] = value
}

func (h *ContextHook) Clear() {
    h.mu.Lock()
    defer h.mu.Unlock()
    h.context = make(map[string]interface{})
}

func (h *ContextHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    h.mu.RLock()
    defer h.mu.RUnlock()

    // 添加上下文字段
    for key, value := range h.context {
        e.Fields = append(e.Fields, log.KV{
            Key:   key,
            Value: value,
        })
    }

    return e
}

// 使用（例如在 HTTP 中间件中）
ctxHook := NewContextHook()
logger.AddHook(ctxHook)

// 在请求处理中
ctxHook.Set("request_id", uuid.New().String())
ctxHook.Set("user_id", userID)

defer ctxHook.Clear()

logger.Info("处理请求")
// 输出包含 request_id 和 user_id
```

### 7. 性能监控

```go
// MetricsHook 收集日志指标
type MetricsHook struct {
    mu           sync.Mutex
    levelCounts  map[log.Level]int64
    totalLogs    int64
    errorLogs    int64
}

func NewMetricsHook() *MetricsHook {
    return &MetricsHook{
        levelCounts: make(map[log.Level]int64),
    }
}

func (h *MetricsHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    h.mu.Lock()
    defer h.mu.Unlock()

    // 统计
    h.levelCounts[e.Level]++
    h.totalLogs++

    if e.Level >= log.ErrorLevel {
        h.errorLogs++
    }

    return e
}

func (h *MetricsHook) GetMetrics() map[string]int64 {
    h.mu.Lock()
    defer h.mu.Unlock()

    metrics := make(map[string]int64)
    for level, count := range h.levelCounts {
        metrics[level.String()] = count
    }
    metrics["total"] = h.totalLogs
    metrics["errors"] = h.errorLogs

    return metrics
}

// 使用
metricsHook := NewMetricsHook()
logger.AddHook(metricsHook)

// ... 记录日志 ...

// 获取指标
metrics := metricsHook.GetMetrics()
fmt.Printf("总日志数: %d, 错误数: %d\n", metrics["total"], metrics["errors"])
```

## 最佳实践

### 1. 避免 Hook 中的阻塞操作

```go
// ❌ 错误：阻塞操作
func (h *BadHook) OnWrite(entry interface{}) interface{} {
    http.Post("http://example.com/log", "application/json", data)
    return entry
}

// ✅ 正确：使用 goroutine
func (h *GoodHook) OnWrite(entry interface{}) interface{} {
    go func() {
        http.Post("http://example.com/log", "application/json", data)
    }()
    return entry
}
```

### 2. 处理类型断言失败

```go
func (h *MyHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        // 优雅处理未知类型
        return entry
    }

    // 处理逻辑
    return e
}
```

### 3. 注意并发安全

```go
type SafeHook struct {
    mu   sync.RWMutex
    data map[string]string
}

func (h *SafeHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    // 读取时使用读锁
    h.mu.RLock()
    defer h.mu.RUnlock()

    // 访问共享数据
    for k, v := range h.data {
        e.Fields = append(e.Fields, log.KV{Key: k, Value: v})
    }

    return e
}
```

### 4. 考虑性能影响

```go
// 轻量级 Hook - 只在需要时处理
type LightweightHook struct{}

func (h *LightweightHook) OnWrite(entry interface{}) interface{} {
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry
    }

    // 快速检查
    if e.Level < log.ErrorLevel {
        return e // 快速返回
    }

    // 只处理错误级别的日志
    // ... 复杂处理逻辑

    return e
}
```

### 5. Hook 执行顺序

多个 Hook 按添加顺序执行：

```go
logger := log.New()
logger.AddHook(&Hook1{}) // 先执行
logger.AddHook(&Hook2{}) // 再执行
logger.AddHook(&Hook3{}) // 最后执行

logger.Info("test")
// 执行顺序: Hook1 → Hook2 → Hook3 → 格式化 → 写入
```

## 完整示例

### 综合使用多个 Hook

```go
package main

import (
    "github.com/lazygophers/log"
    "github.com/lazygophers/log/constant"
)

func main() {
    logger := log.New()

    // 1. 添加环境信息
    logger.AddHook(&EnvironmentHook{
        Environment: "production",
        Version:     "1.0.0",
    })

    // 2. 过滤敏感信息
    logger.AddHook(NewSensitiveDataHook())

    // 3. 添加性能监控
    metricsHook := NewMetricsHook()
    logger.AddHook(metricsHook)

    // 使用日志器
    logger.Info("服务启动")
    logger.Infow("用户登录", "user_id", 123, "username", "admin")
    logger.Errorw("数据库错误", "error", "connection failed")

    // 获取指标
    metrics := metricsHook.GetMetrics()
    log.Infof("日志指标: %+v", metrics)
}
```

## 常见问题

### Q: Hook 中返回 nil 会怎样？

A: 返回 `nil` 会跳过该条日志，不会写入输出。可以用于过滤日志。

### Q: 如何在 Hook 中访问原始 Logger？

A: Hook 接口设计为无状态的。如果需要访问 Logger，建议在创建 Hook 时传入必要的信息。

### Q: Hook 会影响性能吗？

A: 轻量级 Hook 影响很小。但应避免在 Hook 中进行阻塞操作或大量计算。

### Q: 如何移除已添加的 Hook？

A: 当前不支持移除单个 Hook。如果需要动态管理，建议自行实现 Hook 管理器。

## 总结

Hook 是 lazygophers/log 的强大扩展机制：

- ✅ 灵活：支持任意自定义逻辑
- ✅ 高性能：轻量级实现
- ✅ 易用：简单的接口设计
- ✅ 安全：类型断言和错误处理

通过合理使用 Hook，可以实现：
- 添加全局字段
- 过滤敏感信息
- 修改日志内容
- 集成第三方服务
- 收集性能指标
- 实现自定义逻辑

欢迎探索更多 Hook 使用场景！
