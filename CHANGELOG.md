# 📋 Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2026-05-05

### Added
- **Hook 接口系统**: 支持在日志写入前进行修改、过滤或丰富日志条目
- **JSON 序列化优化**: Entry 实现 `json.Marshaler` 接口，支持条件字段序列化
- **测试覆盖率提升**: 从 93.0% 提升到 95.0%，新增 40+ 测试用例
- **文档重写**: 根据实际代码状态深度重写所有文档

### Changed
- **架构简化**: 移除 `logger_ctx` 和相关 zap 代码，简化日志处理逻辑
- **Entry 重构**: 将 `Entry` 移至 `constant` 包，优化内存布局（200字节，2%填充浪费）
- **日志轮转改进**: 完善按小时轮转功能，支持基于大小的分片
- **测试文件规范化**: 统一测试文件命名标准，移除独立的覆盖率测试文件

### Removed
- `logger_ctx.go` 和 `logger_ctx_test.go` — context 日志层
- 独立的覆盖率测试文件（`coverage_boost_test.go` 等）
- 过时的文档和 API 说明

### Fixed
- 修复格式化器的调用者信息处理
- 改进时间戳缓存机制
- 优化对象池复用逻辑

### Performance
- **缓存优化**: Entry 结构按访问频率和大小排列，最小化缓存行浪费
- **时间缓存**: 格式化的时间戳缓存为字符串，避免重复格式化
- **早期检查**: 级别检查在外层完成，避免不必要的 Entry 填充

### Testing
- 测试覆盖率: **95.0%** (486 个测试用例)
- 所有测试合并到标准测试文件
- 新增边界情况和分支覆盖测试

## [1.0.0] - 2024-01-01

### Added
- 核心日志功能，支持多个级别（Trace、Debug、Info、Warn、Error、Fatal、Panic）
- 线程安全的日志器实现，使用对象池优化
- 构建标签支持（default、debug、release、discard 模式）
- 自定义格式化器接口，提供默认文本格式化器
- 多 Writer 输出支持
- 异步写入能力，适用于高吞吐场景
- 按小时自动轮转日志文件
- 上下文感知日志，支持 Goroutine ID 和 Trace ID 追踪
- 可配置的调用者信息（文件名、行号、函数名）
- 全局包级便捷函数
- Zap 日志器集成支持

### Performance
- 使用 `sync.Pool` 对 Entry 对象和缓冲区进行对象池优化
- 早期级别检查，避免昂贵的操作
- 异步写入器实现非阻塞日志写入
- 针对不同环境的构建标签优化

### Build Tags
- **Default**: 完整功能，包含调试消息
- **Debug**: 增强调试信息和调用者详情
- **Release**: 生产优化，禁用调试消息
- **Discard**: 最高性能，空操作日志

### Core Features
- **Logger**: 主日志器结构，可配置级别、输出、格式化器
- **Entry**: 日志记录结构，包含全面的元数据
- **Levels**: 七个日志级别，从 Panic（最高）到 Trace（最低）
- **Formatters**: 可插拔格式化系统
- **Writers**: 文件轮转和异步写入支持
- **Context**: Goroutine ID 和分布式追踪支持

### API Highlights
- 流式配置 API，支持方法链式调用
- 简单和格式化日志方法（`.Info()` 和 `.Infof()`）
- 日志器克隆，支持隔离配置
- 前缀和后缀消息自定义
- 调用者信息开关

### Testing
- 全面的测试套件，93.5% 覆盖率
- 多构建标签测试支持
- 自动化测试工作流
- 性能基准测试

## [0.9.0] - 2023-12-15

### Added
- 初始项目结构
- 基本日志功能
- 级别过滤
- 文件输出支持

### Changed
- 使用对象池改进性能
- 增强错误处理

## [0.8.0] - 2023-12-01

### Added
- 多 Writer 支持
- 自定义格式化器接口
- 异步写入能力

### Fixed
- 高吞吐场景下的内存泄漏
- 并发访问的竞态条件

## [0.7.0] - 2023-11-15

### Added
- 构建标签支持，实现条件编译
- Trace 和 Debug 级别日志
- 调用者信息追踪

### Changed
- 优化内存分配模式
- 改进线程安全性

## [0.6.0] - 2023-11-01

### Added
- 日志轮转功能
- 上下文感知日志
- Goroutine ID 追踪

## [0.5.0] - 2023-10-15

### Added
- JSON 格式化器
- 多个输出目标
- 性能基准测试

### Changed
- 重构核心日志引擎
- 改进 API 一致性

## [0.4.0] - 2023-10-01

### Added
- Fatal 和 Panic 级别日志
- 全局包函数
- 配置验证

### Fixed
- 输出同步问题
- 内存使用优化

## [0.3.0] - 2023-09-15

### Added
- 自定义日志级别
- 格式化器接口
- 线程安全操作

### Changed
- 简化 API 设计
- 增强文档

## [0.2.0] - 2023-09-01

### Added
- 文件输出支持
- 基于级别的过滤
- 基本格式化选项

### Fixed
- 性能瓶颈
- 内存泄漏

## [0.1.0] - 2023-08-15

### Added
- 初始发布
- 基本控制台日志
- 简单级别支持（Info、Warn、Error）
- 核心日志器结构

## 版本历史摘要

| 版本 | 发布日期 | 主要特性 |
|------|----------|----------|
| 1.1.0 | 2026-05-05 | Hook 系统、Entry 优化、95% 测试覆盖率、文档重写 |
| 1.0.0 | 2024-01-01 | 完整日志系统、构建标签、异步写入、全面文档 |
| 0.9.0 | 2023-12-15 | 性能改进、对象池 |
| 0.8.0 | 2023-12-01 | 多 Writer、异步写入、自定义格式化器 |
| 0.7.0 | 2023-11-15 | 构建标签、Trace/Debug 级别、调用者信息 |
| 0.6.0 | 2023-11-01 | 日志轮转、上下文日志、Goroutine 追踪 |
| 0.5.0 | 2023-10-15 | JSON 格式化器、多输出、基准测试 |
| 0.4.0 | 2023-10-01 | Fatal/Panic 级别、全局函数 |
| 0.3.0 | 2023-09-15 | 自定义级别、格式化器接口 |
| 0.2.0 | 2023-09-01 | 文件输出、级别过滤 |
| 0.1.0 | 2023-08-15 | 初始发布、基本控制台日志 |

## 迁移指南

### 从 v1.0.x 迁移到 v1.1.0

#### 破坏性变更
- `Entry` 类型现在位于 `constant` 包中
- 主包通过类型别名导出：`type Entry = constant.Entry`
- 如果代码直接引用 `Entry`，需要更新导入路径

#### 新功能
- Hook 接口支持：
  ```go
  type Hook interface {
      OnWrite(entry interface{}) interface{}
  }
  ```
- JSON 序列化优化：
  ```go
  // Entry 现在实现 json.Marshaler
  data, _ := json.Marshal(entry)
  ```

#### 推荐更新
```go
// 旧代码
if entry.Level == log.InfoLevel {
    // ...
}

// 新代码（兼容）
if entry.Level == log.InfoLevel {
    // ...
}

// 或使用 constant 包
if entry.Level == constant.InfoLevel {
    // ...
}
```

### 从 v0.9.x 迁移到 v1.0.0

#### 破坏性变更
- 无破坏性变更 - v1.0.0 与 v0.9.x 向后兼容

#### 新功能
- 增强的构建标签支持
- 全面的文档
- 专业的项目模板
- 安全报告流程

### 从 v0.8.x 迁移到 v0.9.x

#### 破坏性变更
- 移除了弃用的配置方法
- 更改了内部缓冲区管理

#### 迁移步骤
1. 更新导入路径（如需要）
2. 替换弃用的方法：
   ```go
   // 旧（已弃用）
   logger.SetOutputFile("app.log")
   
   // 新
   file, _ := os.Create("app.log")
   logger.SetOutput(file)
   ```

## 贡献

欢迎贡献！请参阅[贡献指南](CONTRIBUTING.md)了解详情：
- 报告错误和请求功能
- 代码提交流程
- 开发环境设置
- 测试要求
- 文档标准

## 安全

关于安全漏洞，请参阅[安全策略](SECURITY.md)了解：
- 支持的版本
- 报告流程
- 响应时间线
- 安全最佳实践

## 支持

- 📖 [文档](docs/)
- 🐛 [问题追踪](https://github.com/lazygophers/log/issues)
- 💬 [讨论](https://github.com/lazygophers/log/discussions)
- 📧 Email: support@lazygophers.com

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 多语言文档

本变更日志提供多种语言版本：

- [🇺🇸 English](CHANGELOG.md) (当前)
- [🇨🇳 简体中文](docs/CHANGELOG_zh-CN.md)
- [🇹🇼 繁體中文](docs/CHANGELOG_zh-TW.md)
- [🇫🇷 Français](docs/CHANGELOG_fr.md)
- [🇷🇺 Русский](docs/CHANGELOG_ru.md)
- [🇪🇸 Español](docs/CHANGELOG_es.md)
- [🇸🇦 العربية](docs/CHANGELOG_ar.md)

---

**跟踪每一次改进，随时了解 LazygoPHers Log 的演进！🚀**

---

*本变更日志在每次发布时自动更新。获取最新信息，请查看 [GitHub Releases](https://github.com/lazygophers/log/releases) 页面。*
