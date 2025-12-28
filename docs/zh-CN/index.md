---
titleSuffix: ' | LazyGophers Log'
pageType: home

hero:
  name: LazyGophers 日志库
  text: 一个高性能的 Go 语言日志库
  tagline: 简洁的 API、卓越的性能和灵活的配置
  image:
    src: /log/public/logo.svg
    alt: LazyGophers Log Logo
  actions:
    - theme: brand
      text: 快速开始
      link: /zh-CN/README
    - theme: alt
      text: API 参考
      link: /zh-CN/API
    - theme: alt
      text: 查看 GitHub
      link: https://github.com/lazygophers/log

features:
  - title: 高性能
    details: 基于 zap 构建，采用对象池和条件字段记录技术，确保出色的性能表现
    icon: ⚡
  - title: 丰富的日志级别
    details: 支持 Trace、Debug、Info、Warn、Error、Fatal、Panic 七个日志级别
    icon: 📊
  - title: 灵活的配置
    details: 支持日志级别控制、调用者信息记录、跟踪信息、自定义前缀后缀等
    icon: ⚙️
  - title: 文件轮换
    details: 内置日志文件轮换功能，支持按小时自动轮换日志文件
    icon: 🔄
  - title: Zap 兼容性
    details: 与 zap WriteSyncer 无缝集成，支持自定义输出目标
    icon: 🔗
  - title: 简洁的 API
    details: API 设计类似于标准日志库，易于使用和迁移
    icon: 🚀
  - title: 线程安全
    details: 大多数操作采用无锁设计，确保线程安全且无性能开销
    icon: 🔒
  - title: 生产就绪
    details: 在生产环境中经过实战检验，具有全面的测试覆盖率
    icon: ✅
---
