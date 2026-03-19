---
pageType: home

hero:
    name: LazyGophers Log
    text: 高性能、灵活的 Go 日志库
    tagline: 基于 zap 构建，提供丰富的功能和简洁的 API
    actions:
        - theme: brand
          text: 快速开始
          link: /API
        - theme: alt
          text: 查看 GitHub
          link: https://github.com/lazygophers/log

features:
    - title: "高性能"
      details: 基于 zap 构建，采用对象池复用和条件字段记录，实现最优性能
      icon: 🚀
    - title: "丰富的日志级别"
      details: 支持 Trace、Debug、Info、Warn、Error、Fatal、Panic 级别
      icon: 📊
    - title: "灵活的配置"
      details: 可自定义日志级别、调用者信息、追踪信息、前缀、后缀和输出目标
      icon: ⚙️
    - title: "文件轮转"
      details: 内置每小时日志文件轮转支持
      icon: 🔄
    - title: "Zap 兼容性"
      details: 与 zap WriteSyncer 无缝集成
      icon: 🔌
    - title: "简洁的 API"
      details: 类似标准日志库的清晰 API，易于使用和集成
      icon: 🎯
---

## 快速开始

### 安装

```bash
go get github.com/lazygophers/log
```

### 基本用法

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 使用默认全局 logger
    log.Debug("调试信息")
    log.Info("普通信息")
    log.Warn("警告信息")
    log.Error("错误信息")

    // 使用格式化输出
    log.Infof("用户 %s 登录成功", "admin")

    // 自定义配置
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("这是来自自定义 logger 的日志")
}
```

## 文档

-   [API 参考](API.md) - 完整的 API 文档
-   [更新日志](CHANGELOG.md) - 版本历史
-   [贡献指南](CONTRIBUTING.md) - 如何贡献
-   [安全策略](SECURITY.md) - 安全指南
-   [行为准则](CODE_OF_CONDUCT.md) - 社区准则

## 性能比较

| 特性       | lazygophers/log | zap | logrus | 标准日志 |
| ---------- | --------------- | --- | ------ | -------- |
| 性能       | 高              | 高  | 中     | 低       |
| API 简洁性 | 高              | 中  | 高     | 高       |
| 功能丰富度 | 中              | 高  | 高     | 低       |
| 灵活性     | 中              | 高  | 高     | 低       |
| 学习曲线   | 低              | 中  | 中     | 低       |

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。
