# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

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

## Project Architecture

This is a high-performance Go logging library (`github.com/lazygophers/log`) designed for simplicity and extensibility.

### Core Components

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

### Key Design Patterns

- **Object Pooling**: Extensive use of `sync.Pool` for Entry objects and byte buffers to reduce GC pressure
- **Interface-based Design**: Pluggable formatters and writers via interfaces
- **Chain Configuration**: Logger methods return `*Logger` for fluent configuration
- **Conditional Compilation**: Build tags for different environments (debug/release)
- **Performance Optimization**: Level checking before expensive operations, buffer reuse

### Dependencies

- `go.uber.org/zap` - Used for WriteSyncer and multi-writer functionality
- `github.com/petermattis/goid` - Goroutine ID extraction for tracing
- `github.com/lestrrat-go/file-rotatelogs` - Time-based log file rotation
- `github.com/google/uuid` - UUID generation (likely for trace IDs)

### Testing Structure

Tests are co-located with source files (`*_test.go`). Key test files:
- `logger_test.go` - Core logger functionality
- `benchmark_test.go` - Performance benchmarks  
- `formatter_test.go` - Format testing
- `pool_test.go` - Object pool behavior