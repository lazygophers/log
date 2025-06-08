### 模块说明
1. **AsyncWriter**（`writer_async.go`）
   - 实现 Writer 接口适配 🔄 (`Write([]byte)`)
   - 基于 sync.Pool 构建零值优化 ∞ (`entryPool`, `bufPool`)
   - 通过 channel 实现内存安全写入 🔐

2. **Formatter**（`formatter.go`）
   - 支持 ZapHook 集成 🔄 (`NewZapHook(Logger)`)
   - 实现颜色编码与包名解析 🎨
   - 通过 sync.Pool 缓冲 byte 数组 📦 `msgBufPool` 

3. **池化架构**
   - entryPool: `var entryPool = sync.Pool{...}`
   - bufPool: `var bufPool = sync.Pool{...}`
   - 零值初始化优化性能 💡

4. **ZapHook 机制**
   - 接口类型: `func(entry zapcore.Entry) error`
   - 通过 `NewZapHook(Logger)` 初始化
   - 与 Write([]byte) 实现解耦设计
