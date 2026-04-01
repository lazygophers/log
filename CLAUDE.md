# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 🚀 Quick Start

### Development Commands

**Testing:**
- `go test ./...` - Run all tests in the project
- `go test -v ./...` - Run tests with verbose output  
- `go test -bench=.` - Run all benchmarks
- `go test -cover ./...` - Run tests with coverage report

**Building and Linting:**
- `go build ./...` - Build all packages
- `go vet ./...` - Run Go vet static analysis
- `go fmt ./...` - Format all Go source files
- `go mod tidy` - Clean up go.mod dependencies

## 🏗️ Project Architecture

This is a high-performance Go logging library (`github.com/lazygophers/log`) designed for simplicity and extensibility.

### 📋 Core Components

**Logger (`logger.go`):**
- Main `Logger` struct with configurable level, output, formatter, and caller depth
- Supports multi-writer output via zap's WriteSyncer
- Object pooling with `sync.Pool` for performance optimization
- Levels: Trace, Debug, Info, Warn, Error, Fatal, Panic
- Provides both simple (`.Info()`) and formatted (`.Infof()`) logging methods

**Entry System (`entry.go`):**
- `Entry` struct represents a single log record with metadata:
  - Process ID (Pid), Goroutine ID (Gid), Trace ID
  - Timestamp, level, message, caller information (file, line, func)
  - Prefix/suffix message support
- Uses `sync.Pool` (`entryPool`) for object reuse to minimize allocations

**Level System (`level.go`):**
- 7 log levels from PanicLevel (highest) to TraceLevel (lowest)
- Compatible with logrus library conventions
- Implements `fmt.Stringer` and `encoding.TextMarshaler` interfaces

**Formatting (`formatter.go`):**
- `Format` interface for pluggable log formatting
- `FormatFull` interface extends Format with parsing/escaping and caller controls  
- Default `Formatter` struct provides standard text formatting
- Uses buffer pooling (`pool.go`) for efficient string building

**Output Management (`output.go`, `writer.go`, `writer_async.go`):**
- File rotation support via `GetOutputWriterHourly()` 
- Async writing capabilities for high-throughput scenarios
- Integration with `file-rotatelogs` for time-based log rotation

**Global Logger (`init.go`):**
- Package-level functions (Info, Debug, etc.) that delegate to a global logger instance
- Default logger configured with DebugLevel and stdout output

## 🔧 Key Design Patterns

- **Object Pooling**: Extensive use of `sync.Pool` for Entry objects and byte buffers to reduce GC pressure
- **Interface-based Design**: Pluggable formatters and writers via interfaces
- **Chain Configuration**: Logger methods return `*Logger` for fluent configuration
- **Conditional Compilation**: Build tags for different environments (debug/release)
- **Performance Optimization**: Level checking before expensive operations, buffer reuse

## 📦 Dependencies

- `go.uber.org/zap` - Used for WriteSyncer and multi-writer functionality
- `github.com/petermattis/goid` - Goroutine ID extraction for tracing
- `github.com/lestrrat-go/file-rotatelogs` - Time-based log file rotation
- `github.com/google/uuid` - UUID generation (likely for trace IDs)

## 🧪 Testing Structure

Tests are co-located with source files (`*_test.go`). Key test files:
- `logger_test.go` - Core logger functionality
- `benchmark_test.go` - Performance benchmarks  
- `formatter_test.go` - Format testing
- `pool_test.go` - Object pool behavior

## 📚 Documentation Structure

### Main Documentation
- `README.md` - Primary project documentation (English)
- `docs/README_zh-CN.md` - Chinese documentation
- `docs/API.md` - Complete API reference
- `CHANGELOG.md` - Version history

### Contributing & Community
- `docs/CONTRIBUTING.md` - Contribution guidelines
- `docs/CODE_OF_CONDUCT.md` - Community standards
- `docs/SECURITY.md` - Security policy

### GitHub Templates
- `.github/ISSUE_TEMPLATE/` - Bug reports, feature requests, questions
- `.github/pull_request_template.md` - PR template

## 🎯 Development Guidelines

### Code Style
- Follow Go standard formatting (`go fmt`)
- Use descriptive variable and function names
- Add comments for complex logic
- Maintain consistent error handling patterns

### Performance Considerations
- Always use object pooling for frequently allocated objects
- Check log levels before expensive operations
- Use conditional compilation for environment-specific optimizations
- Benchmark performance-critical code paths

### Testing Requirements
- Write tests for all public APIs
- Include benchmarks for performance-critical code
- Test across different build tag configurations
- Ensure test coverage remains above 90%

## 🔄 Build Tags

This project uses build tags for conditional compilation:

- `debug` - Enables additional debugging features
- `release` - Optimizes for production performance  
- `discard` - Discards all log output (for testing)

## 🚨 Important Notes

### Memory Management
- The library extensively uses `sync.Pool` for performance
- Always return objects to pools after use
- Be aware of pool contention in highly concurrent scenarios

### Concurrency Safety
- Logger instances are designed to be thread-safe
- Individual Entry objects are not thread-safe (single-use)
- Use appropriate synchronization when extending functionality

### Extension Points
- The `Format` interface allows custom formatting implementations
- Output can be customized via `io.Writer` implementations
- Log levels can be extended by implementing the `Level` interface

## 📞 Getting Help

- Check the [API Documentation](docs/API.md) for detailed usage
- Review [examples](examples/) for common patterns
- Open an [issue](https://github.com/lazygophers/log/issues) for bugs or questions
- Refer to [contributing guidelines](docs/CONTRIBUTING.md) for development help

---

**Remember**: This is a performance-focused logging library. Always consider the performance implications of any changes, especially in the hot path of logging operations.

<!-- gitnexus:start -->
# GitNexus — Code Intelligence

This project is indexed by GitNexus as **log** (0 symbols, 0 relationships, 0 execution flows). Use the GitNexus MCP tools to understand code, assess impact, and navigate safely.

> If any GitNexus tool warns the index is stale, run `npx gitnexus analyze` in terminal first.

## Always Do

- **MUST run impact analysis before editing any symbol.** Before modifying a function, class, or method, run `gitnexus_impact({target: "symbolName", direction: "upstream"})` and report the blast radius (direct callers, affected processes, risk level) to the user.
- **MUST run `gitnexus_detect_changes()` before committing** to verify your changes only affect expected symbols and execution flows.
- **MUST warn the user** if impact analysis returns HIGH or CRITICAL risk before proceeding with edits.
- When exploring unfamiliar code, use `gitnexus_query({query: "concept"})` to find execution flows instead of grepping. It returns process-grouped results ranked by relevance.
- When you need full context on a specific symbol — callers, callees, which execution flows it participates in — use `gitnexus_context({name: "symbolName"})`.

## When Debugging

1. `gitnexus_query({query: "<error or symptom>"})` — find execution flows related to the issue
2. `gitnexus_context({name: "<suspect function>"})` — see all callers, callees, and process participation
3. `READ gitnexus://repo/log/process/{processName}` — trace the full execution flow step by step
4. For regressions: `gitnexus_detect_changes({scope: "compare", base_ref: "main"})` — see what your branch changed

## When Refactoring

- **Renaming**: MUST use `gitnexus_rename({symbol_name: "old", new_name: "new", dry_run: true})` first. Review the preview — graph edits are safe, text_search edits need manual review. Then run with `dry_run: false`.
- **Extracting/Splitting**: MUST run `gitnexus_context({name: "target"})` to see all incoming/outgoing refs, then `gitnexus_impact({target: "target", direction: "upstream"})` to find all external callers before moving code.
- After any refactor: run `gitnexus_detect_changes({scope: "all"})` to verify only expected files changed.

## Never Do

- NEVER edit a function, class, or method without first running `gitnexus_impact` on it.
- NEVER ignore HIGH or CRITICAL risk warnings from impact analysis.
- NEVER rename symbols with find-and-replace — use `gitnexus_rename` which understands the call graph.
- NEVER commit changes without running `gitnexus_detect_changes()` to check affected scope.

## Tools Quick Reference

| Tool | When to use | Command |
|------|-------------|---------|
| `query` | Find code by concept | `gitnexus_query({query: "auth validation"})` |
| `context` | 360-degree view of one symbol | `gitnexus_context({name: "validateUser"})` |
| `impact` | Blast radius before editing | `gitnexus_impact({target: "X", direction: "upstream"})` |
| `detect_changes` | Pre-commit scope check | `gitnexus_detect_changes({scope: "staged"})` |
| `rename` | Safe multi-file rename | `gitnexus_rename({symbol_name: "old", new_name: "new", dry_run: true})` |
| `cypher` | Custom graph queries | `gitnexus_cypher({query: "MATCH ..."})` |

## Impact Risk Levels

| Depth | Meaning | Action |
|-------|---------|--------|
| d=1 | WILL BREAK — direct callers/importers | MUST update these |
| d=2 | LIKELY AFFECTED — indirect deps | Should test |
| d=3 | MAY NEED TESTING — transitive | Test if critical path |

## Resources

| Resource | Use for |
|----------|---------|
| `gitnexus://repo/log/context` | Codebase overview, check index freshness |
| `gitnexus://repo/log/clusters` | All functional areas |
| `gitnexus://repo/log/processes` | All execution flows |
| `gitnexus://repo/log/process/{name}` | Step-by-step execution trace |

## Self-Check Before Finishing

Before completing any code modification task, verify:
1. `gitnexus_impact` was run for all modified symbols
2. No HIGH/CRITICAL risk warnings were ignored
3. `gitnexus_detect_changes()` confirms changes match expected scope
4. All d=1 (WILL BREAK) dependents were updated

## Keeping the Index Fresh

After committing code changes, the GitNexus index becomes stale. Re-run analyze to update it:

```bash
npx gitnexus analyze
```

If the index previously included embeddings, preserve them by adding `--embeddings`:

```bash
npx gitnexus analyze --embeddings
```

To check whether embeddings exist, inspect `.gitnexus/meta.json` — the `stats.embeddings` field shows the count (0 means no embeddings). **Running analyze without `--embeddings` will delete any previously generated embeddings.**

> Claude Code users: A PostToolUse hook handles this automatically after `git commit` and `git merge`.

## CLI

| Task | Read this skill file |
|------|---------------------|
| Understand architecture / "How does X work?" | `.claude/skills/gitnexus/gitnexus-exploring/SKILL.md` |
| Blast radius / "What breaks if I change X?" | `.claude/skills/gitnexus/gitnexus-impact-analysis/SKILL.md` |
| Trace bugs / "Why is X failing?" | `.claude/skills/gitnexus/gitnexus-debugging/SKILL.md` |
| Rename / extract / split / refactor | `.claude/skills/gitnexus/gitnexus-refactoring/SKILL.md` |
| Tools, resources, schema reference | `.claude/skills/gitnexus/gitnexus-guide/SKILL.md` |
| Index, status, clean, wiki CLI commands | `.claude/skills/gitnexus/gitnexus-cli/SKILL.md` |

<!-- gitnexus:end -->
