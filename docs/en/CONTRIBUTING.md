---
titleSuffix: " | LazyGophers Log"
---
# 🤝 Contributing to LazyGophers Log

We welcome your contributions! We want to make contributing to LazyGophers Log as simple and transparent as possible, whether it's:

-   🐛 Reporting a bug
-   💬 Discussing the current state of the code
-   ✨ Submitting a feature request
-   🔧 Proposing a fix
-   🚀 Implementing new features

## 📋 Table of Contents

-   [Code of Conduct](#-code-of-conduct)
-   [Development Process](#-development-process)
-   [Getting Started](#-getting-started)
-   [Pull Request Process](#-pull-request-process)
-   [Coding Standards](#-coding-standards)
-   [Testing Guidelines](#-testing-guidelines)
-   [Build Tag Requirements](#-build-tag-requirements)
-   [Documentation](#-documentation)
-   [Issue Guidelines](#-issue-guidelines)
-   [Performance Considerations](#-performance-considerations)
-   [Security Guidelines](#-security-guidelines)
-   [Community](#-community)

## 📜 Code of Conduct

This project and all participants are governed by our [Code of Conduct](/CODE_OF_CONDUCT). By participating, you agree to abide by this code.

## 🔄 Development Process

We use GitHub to host code, track issues and feature requests, and accept pull requests.

### Workflow

1. **Fork** the repository
2. **Clone** your fork locally
3. **Create** a feature branch from `master`
4. **Make** your changes
5. **Test** thoroughly under all build tags
6. **Submit** a pull request

## 🚀 Getting Started

### Prerequisites

-   **Go 1.21+** - [Install Go](https://golang.org/doc/install)
-   **Git** - [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
-   **Make** (optional but recommended)

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
# Set Go environment (if not already set)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Optional: Install useful tools
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 Pull Request Process

### Before Submitting

1. **Search** for existing PRs to avoid duplicates
2. **Test** your changes under all build configurations
3. **Document** any breaking changes
4. **Update** related documentation
5. **Add** tests for new features

### PR Checklist

-   [ ] **Code Quality**

    -   [ ] Code follows project style guide
    -   [ ] No new lint warnings
    -   [ ] Proper error handling
    -   [ ] Efficient algorithms and data structures

-   [ ] **Testing**

    -   [ ] All existing tests pass: `make test`
    -   [ ] New tests added for new functionality
    -   [ ] Tests cover edge cases
    -   [ ] Performance tests (if applicable)

-   [ ] **Documentation**

    -   [ ] Code comments updated
    -   [ ] API documentation updated
    -   [ ] README/guide updates (if applicable)

-   [ ] **Build**
    -   [ ] Builds under all supported Go versions
    -   [ ] Builds under all build tags
    -   [ ] No new dependencies added unnecessarily

### Submitting

1. Push your changes to your fork
2. Create a pull request to the `master` branch
3. Fill out the PR template completely
4. Link any related issues

## 💻 Coding Standards

### General Guidelines

-   Follow Go best practices: [Effective Go](https://go.dev/doc/effective_go)
-   Use meaningful variable and function names
-   Keep functions small and focused
-   Write self-documenting code

### Specific Standards

-   **Code Style**: Use `gofmt` and `goimports` for formatting
-   **Error Handling**: Use proper error wrapping and context
-   **Logging**: Use the project's logging package appropriately
-   **Concurrency**: Follow Go concurrency patterns safely

### Build Tags

Some features may be conditionally compiled using build tags:

-   `dev`: Development features
-   `test`: Testing utilities
-   `bench`: Benchmarking tools

## 🧪 Testing Guidelines

### Running Tests

```bash
# Run all tests
make test

# Run tests quickly (without race detection)
make test-quick

# Run tests with race detection
make test-race

# Run tests for a specific package
make test-pkg pkg=github.com/lazygophers/log
```

### Writing Tests

-   Write unit tests for all new functionality
-   Use table-driven tests for multiple test cases
-   Follow the existing test patterns
-   Test both success and failure cases

### Coverage

```bash
# Generate test coverage
make test-coverage

# View coverage report
make test-coverage-view
```

## 🏷️ Build Tag Requirements

When adding features that should be conditionally compiled:

1. Use descriptive build tag names
2. Document build tags in the README
3. Ensure backward compatibility
4. Test with and without the build tag

## 📚 Documentation

### API Documentation

Update GoDoc comments for any public API changes:

```go
// LogLevel represents the severity level of a log message
// Example:
//     logger.SetLevel(log.InfoLevel)
type LogLevel int
```

### User Documentation

Update the appropriate documentation files for:

-   New features
-   API changes
-   Configuration options
-   Usage examples

## ❓ Issue Guidelines

### Reporting Bugs

When reporting bugs, please include:

-   **Go version**: Output of `go version`
-   **Package version**: Which version of the package you're using
-   **Description**: A clear and concise description of the bug
-   **Steps to reproduce**: Minimum steps to reproduce the issue
-   **Expected behavior**: What you expected to happen
-   **Actual behavior**: What actually happened
-   **Logs**: Any relevant log output
-   **Code example**: Minimal, complete, and verifiable example

### Feature Requests

When requesting features, please include:

-   **Description**: A clear and concise description of the feature
-   **Use case**: Why this feature would be useful
-   **Proposal**: A suggested implementation (optional)
-   **Alternatives**: Any alternative solutions you've considered

## ⚡ Performance Considerations

-   **Benchmark**: Add benchmarks for performance-sensitive code
-   **Profile**: Use Go's profiling tools to identify bottlenecks
-   **Optimize**: Focus on hot paths and critical sections
-   **Document**: Note any performance considerations in the code

## 🔒 Security Guidelines

If you discover a security vulnerability, please follow our [Security Policy](/SECURITY) to report it responsibly.

## 👥 Community

-   **GitHub Discussions**: For questions and discussions
-   **Issue Tracker**: For bug reports and feature requests
-   **Slack**: Join our community Slack channel

## 📄 License

By contributing to LazyGophers Log, you agree that your contributions will be licensed under the [MIT License](/LICENSE).

## 🌍 Multilingual Documentation

This document is available in multiple languages:

-   🇺🇸 [English](/CONTRIBUTING) (current)
-   🇨🇳 [简体中文](docs/CONTRIBUTING_zh-CN.md)
-   🇹🇼 [繁體中文](docs/CONTRIBUTING_zh-TW.md)
-   🇫🇷 [Français](docs/CONTRIBUTING_fr.md)
-   🇷🇺 [Русский](docs/CONTRIBUTING_ru.md)
-   🇪🇸 [Español](docs/CONTRIBUTING_es.md)
-   🇸🇦 [العربية](docs/CONTRIBUTING_ar.md)
