# 🤝 Contributing to LazyGophers Log

We love your input! We want to make contributing to LazyGophers Log as easy and transparent as possible, whether it's:

- 🐛 Reporting a bug
- 💬 Discussing the current state of the code
- ✨ Submitting a feature request
- 🔧 Proposing a fix
- 🚀 Implementing a new feature

## 📋 Table of Contents

- [Code of Conduct](#-code-of-conduct)
- [Development Process](#-development-process)
- [Getting Started](#-getting-started)
- [Pull Request Process](#-pull-request-process)
- [Coding Standards](#-coding-standards)
- [Testing Guidelines](#-testing-guidelines)
- [Build Tag Requirements](#️-build-tag-requirements)
- [Documentation](#-documentation)
- [Issue Guidelines](#-issue-guidelines)
- [Performance Considerations](#-performance-considerations)
- [Security Guidelines](#-security-guidelines)
- [Community](#-community)

## 📜 Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## 🔄 Development Process

We use GitHub to host code, track issues and feature requests, as well as accept pull requests.

### Workflow

1. **Fork** the repository
2. **Clone** your fork locally
3. **Create** a feature branch from `master`
4. **Make** your changes
5. **Test** thoroughly across all build tags
6. **Submit** a pull request

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+** - [Install Go](https://golang.org/doc/install)
- **Git** - [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- **Make** (optional but recommended)

### Local Development Setup

```bash
# 1. Fork the repository on GitHub
# 2. Clone your fork
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. Add upstream remote
git remote add upstream https://github.com/lazygophers/log.git

# 4. Install dependencies
go mod tidy

# 5. Verify installation
make test-quick
```

### Environment Setup

```bash
# Set up your Go environment (if not already done)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Optional: Install useful tools
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 Pull Request Process

### Before Submitting

1. **Search** existing PRs to avoid duplicates
2. **Test** your changes across all build configurations
3. **Document** any breaking changes
4. **Update** relevant documentation
5. **Add** tests for new functionality

### PR Checklist

- [ ] **Code Quality**
  - [ ] Code follows project style guidelines
  - [ ] No new linting warnings
  - [ ] Proper error handling
  - [ ] Efficient algorithms and data structures

- [ ] **Testing**
  - [ ] All existing tests pass: `make test`
  - [ ] New tests added for new functionality
  - [ ] Test coverage maintained or improved
  - [ ] All build tags tested: `make test-all`

- [ ] **Documentation**
  - [ ] Code is properly commented
  - [ ] API documentation updated (if needed)
  - [ ] README updated (if needed)
  - [ ] Multilingual docs updated (if user-facing)

- [ ] **Build Compatibility**
  - [ ] Default mode: `go build`
  - [ ] Debug mode: `go build -tags debug`
  - [ ] Release mode: `go build -tags release`
  - [ ] Discard mode: `go build -tags discard`
  - [ ] Combined modes tested

### PR Template

Please use our [PR template](.github/pull_request_template.md) when submitting pull requests.

## 📏 Coding Standards

### Go Style Guide

We follow the standard Go style guide with some additions:

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

### Naming Conventions

- **Packages**: Short, lowercase, single word when possible
- **Functions**: CamelCase, descriptive
- **Variables**: camelCase for local, CamelCase for exported
- **Constants**: CamelCase for exported, camelCase for unexported
- **Interfaces**: Usually end with "er" (e.g., `Writer`, `Formatter`)

### Code Organization

```
project/
├── docs/           # Documentation in multiple languages
├── .github/        # GitHub templates and workflows
├── logger.go       # Main logger implementation
├── entry.go        # Log entry structure
├── level.go        # Log levels
├── formatter.go    # Log formatting
├── output.go       # Output management
└── *_test.go      # Tests co-located with source
```

### Error Handling

```go
// ✅ Preferred: Return errors, don't panic
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ Avoid: Panicking in library code
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // Don't do this
    }
    return &Logger{...}
}
```

## 🧪 Testing Guidelines

### Test Structure

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
            // Test implementation
        })
    }
}
```

### Coverage Requirements

- **Minimum**: 90% coverage for new code
- **Target**: 95%+ coverage overall
- **All build tags** must maintain coverage
- Use `make coverage-all` to verify

### Testing Commands

```bash
# Quick test across all build tags
make test-quick

# Full test suite with coverage
make test-all

# Coverage reports
make coverage-html

# Benchmarks
make benchmark
```

## 🏗️ Build Tag Requirements

All changes must be compatible with our build tag system:

### Supported Build Tags

- **Default** (`go build`): Full functionality
- **Debug** (`go build -tags debug`): Enhanced debugging
- **Release** (`go build -tags release`): Production optimized
- **Discard** (`go build -tags discard`): Maximum performance

### Build Tag Testing

```bash
# Test each build configuration
make test-default
make test-debug  
make test-release
make test-discard

# Test combinations
make test-debug-discard
make test-release-discard

# All in one
make test-all
```

### Build Tag Guidelines

```go
//go:build debug
// +build debug

package log

// Debug-specific implementations
```

## 📚 Documentation

### Code Documentation

- **All exported functions** must have clear comments
- **Complex algorithms** need explanation
- **Examples** for non-trivial usage
- **Thread safety** notes where applicable

```go
// SetLevel sets the minimum logging level.
// Logs below this level will be ignored.
// This method is thread-safe.
//
// Example:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // Won't output
//   logger.Info("visible")   // Will output
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### README Updates

When adding features, update:
- Main README.md
- All language-specific READMEs in `docs/`
- Code examples
- Feature lists

## 🐛 Issue Guidelines

### Bug Reports

Use the [bug report template](.github/ISSUE_TEMPLATE/bug_report.md) and include:

- **Clear description** of the issue
- **Steps to reproduce** 
- **Expected vs actual behavior**
- **Environment details** (OS, Go version, build tags)
- **Minimal code sample**

### Feature Requests

Use the [feature request template](.github/ISSUE_TEMPLATE/feature_request.md) and include:

- **Clear motivation** for the feature
- **Proposed API** design
- **Implementation considerations**
- **Breaking change analysis**

### Questions

Use the [question template](.github/ISSUE_TEMPLATE/question.md) for:

- Usage questions
- Configuration help
- Best practices
- Integration guidance

## 🚀 Performance Considerations

### Benchmarking

Always benchmark performance-sensitive changes:

```bash
# Run benchmarks
go test -bench=. -benchmem

# Compare before/after
go test -bench=. -benchmem > before.txt
# Make changes
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### Performance Guidelines

- **Minimize allocations** in hot paths
- **Use object pooling** for frequently created objects
- **Early return** for disabled log levels
- **Avoid reflection** in performance-critical code
- **Profile before optimizing**

### Memory Management

```go
// ✅ Good: Use object pools
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

## 🔒 Security Guidelines

### Sensitive Data

- **Never log** passwords, tokens, or sensitive data
- **Sanitize** user input in log messages
- **Avoid** logging entire request/response bodies
- **Use** structured logging for better control

```go
// ✅ Good
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ Bad
logger.Infof("User login: %+v", userRequest) // May contain password
```

### Dependencies

- Keep dependencies **up to date**
- **Review** new dependencies carefully
- **Minimize** external dependencies
- **Use** `go mod verify` to check integrity

## 👥 Community

### Getting Help

- 📖 [Documentation](../README.md)
- 💬 [GitHub Discussions](https://github.com/lazygophers/log/discussions)
- 🐛 [Issue Tracker](https://github.com/lazygophers/log/issues)
- 📧 Email: support@lazygophers.com

### Communication Guidelines

- **Be respectful** and inclusive
- **Search** before asking questions
- **Provide context** when asking for help
- **Help others** when you can
- **Follow** the [Code of Conduct](CODE_OF_CONDUCT.md)

## 🎯 Recognition

Contributors are recognized in several ways:

- **README contributors** section
- **Release notes** mentions
- **GitHub contributor** graphs
- **Community appreciation** posts

## 📝 License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

## 🌍 Multilingual Documentation

This document is available in multiple languages:

- [🇺🇸 English](CONTRIBUTING.md) (Current)
- [🇨🇳 简体中文](CONTRIBUTING_zh-CN.md)
- [🇹🇼 繁體中文](CONTRIBUTING_zh-TW.md)
- [🇫🇷 Français](CONTRIBUTING_fr.md)
- [🇷🇺 Русский](CONTRIBUTING_ru.md)
- [🇪🇸 Español](CONTRIBUTING_es.md)
- [🇸🇦 العربية](CONTRIBUTING_ar.md)

---

**Thank you for contributing to LazyGophers Log! 🚀**