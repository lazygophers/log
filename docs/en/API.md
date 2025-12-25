---
titleSuffix: " | LazyGophers Log"
---

# 📚 API Documentation

## Overview

LazyGophers Log provides a comprehensive logging API that supports multiple log levels, custom formatting, asynchronous writing, and build tag optimization. This document covers all public APIs, configuration options, and usage patterns.

## Table of Contents

-   [Core Types](#core-types)
-   [Logger API](#logger-api)
-   [Global Functions](#global-functions)
-   [Log Levels](#log-levels)
-   [Formatters](#formatters)
-   [Output Writers](#output-writers)
-   [Contextual Logging](#contextual-logging)
-   [Build Tags](#build-tags)
-   [Performance Optimization](#performance-optimization)
-   [Examples](#examples)

## Core Types

### Logger

The main logger struct that provides all logging functionality.

```go
type Logger struct {
    // Contains private fields for thread-safe operations
}
```

#### Constructor

```go
func New() *Logger
```

Creates a new logger instance with default configuration:

-   Level: `DebugLevel`
-   Output: `os.Stdout`
-   Formatter: Default text formatter
-   Caller tracking: Disabled

**Example:**

```go
logger := log.New()
logger.Info("New logger created")
```

### Entry

Represents a single log entry with all associated metadata.

```go
type Entry struct {
    Time       time.Time     // Timestamp when the entry was created
    Level      Level         // Log level
    Message    string        // Log message
    Pid        int          // Process ID
    Gid        uint64       // Goroutine ID
    TraceID    string       // Trace ID for distributed tracing
    CallerName string       // Caller function name
    CallerFile string       // Caller file path
    CallerLine int          // Caller line number
}
```

## Logger API

### Configuration Methods

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Sets the minimum log level. Messages below this level will be ignored.

**Parameters:**

-   `level`：The minimum log level to process

**Return:**

-   `*Logger`：Returns itself to support method chaining

**Example:**

```go
logger.SetLevel(log.InfoLevel)
logger.Debug("This won't be displayed")  // Ignored
logger.Info("This will be displayed")    // Processed
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Sets one or more output destinations for log messages.

**Parameters:**

-   `writers`：One or more `io.Writer` output destinations

**Return:**

-   `*Logger`：Returns itself to support method chaining

**Example:**

```go
// Single output
logger.SetOutput(os.Stdout)

// Multiple outputs
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

Sets a custom formatter for log output.

**Parameters:**

-   `formatter`：A formatter that implements the `Format` interface

**Return:**

-   `*Logger`：Returns itself to support method chaining

**Example:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enabled bool) *Logger
```

Enables or disables caller information in log entries.

**Parameters:**

-   `enabled`：Whether to include caller information

**Return:**

-   `*Logger`：Returns itself to support method chaining

**Example:**

```go
logger.EnableCaller(true)
logger.Info("This will include file:line information")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Sets the stack depth for caller information when wrapping loggers.

**Parameters:**

-   `depth`：Number of stack frames to skip

**Return:**

-   `*Logger`：Returns itself to support method chaining

**Example:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // Skip the wrapper function
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Sets prefix or suffix text for all log messages.

**Parameters:**

-   `prefix/suffix`：Text to prepend/append to messages

**Return:**

-   `*Logger`：Returns itself to support method chaining

**Example:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // Output: [APP] Hello [END]
```

### Logging Methods

All logging methods have two variants: simple and formatted versions.

#### Trace Level

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

Logs at trace level (most detailed).

**Example:**

```go
logger.Trace("Detailed execution trace")
logger.Tracef("Processing item %d of %d", i, total)
```

#### Debug Level

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

Logs development information at debug level.

**Example:**

```go
logger.Debug("Variable state:", variable)
logger.Debugf("User %s authenticated successfully", username)
```

#### Info Level

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

Logs informational messages.

**Example:**

```go
logger.Info("Application started")
logger.Infof("Server listening on port %d", port)
```

#### Warn Level

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

Logs warning messages for potential issue situations.

**Example:**

```go
logger.Warn("Deprecated function called")
logger.Warnf("High memory usage: %d%%", memoryPercent)
```

#### Error Level

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

Logs error messages.

**Example:**

```go
logger.Error("Database connection failed")
logger.Errorf("Request processing failed: %v", err)
```

#### Fatal Level

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

Logs a fatal error and calls `os.Exit(1)`.

**Example:**

```go
logger.Fatal("Critical system error")
logger.Fatalf("Failed to start server: %v", err)
```

#### Panic Level

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

Logs an error message and calls `panic()`.

**Example:**

```go
logger.Panic("Unrecoverable error occurred")
logger.Panicf("Invalid state: %v", state)
```

### Utility Methods

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Creates a copy of the logger with identical configuration.

**Return:**

-   `*Logger`：New logger instance with copied settings

**Example:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

Creates a context-aware logger that accepts `context.Context` as the first parameter.

**Return:**

-   `LoggerWithCtx`：Context-aware logger instance

**Example:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "Context-aware message")
```

## Global Functions

Package-level functions that use the default global logger.

```go
func SetLevel(level Level)
func SetOutput(writers ...io.Writer)
func SetFormatter(formatter Format)
func EnableCaller(enabled bool)

func Trace(v ...any)
func Tracef(format string, v ...any)
func Debug(v ...any)
func Debugf(format string, v ...any)
func Info(v ...any)
func Infof(format string, v ...any)
func Warn(v ...any)
func Warnf(format string, v ...any)
func Error(v ...any)
func Errorf(format string, v ...any)
func Fatal(v ...any)
func Fatalf(format string, v ...any)
func Panic(v ...any)
func Panicf(format string, v ...any)
```

**Example:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("Using global logger")
```

## Log Levels

### Level Type

```go
type Level int8
```

### Available Levels

```go
const (
    PanicLevel Level = iota  // 0 - Panic and exit
    FatalLevel              // 1 - Fatal error and exit
    ErrorLevel              // 2 - Error conditions
    WarnLevel               // 3 - Warning conditions
    InfoLevel               // 4 - Informational messages
    DebugLevel              // 5 - Debug messages
    TraceLevel              // 6 - Most detailed tracing
)
```

### Level Methods

```go
func (l Level) String() string
```

Returns the string representation of the level.

**Example:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## Formatters

### Format Interface

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

Custom formatters must implement this interface.

### Default Formatter

Built-in text formatter with customizable options.

```go
type Formatter struct {
    // Configuration options
}
```

### JSON Formatter Example

```go
type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "caller":    fmt.Sprintf("%s:%d", entry.CallerFile, entry.CallerLine),
    }
    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }

    jsonData, _ := json.Marshal(data)
    return append(jsonData, '\n')
}

// Usage
logger.SetFormatter(&JSONFormatter{})
```

## Output Writers

### File Output and Rotation

```go
func GetOutputWriterHourly(filename string) io.Writer
```

Creates a writer that rotates log files hourly.

**Parameters:**

-   `filename`：Base filename for log files

**Return:**

-   `io.Writer`：Rotating file writer

**Example:**

```go
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// Creates files like: app-2024010115.log, app-2024010116.log, etc.
```

### Async Writer

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

Creates an asynchronous writer for high-performance logging.

**Parameters:**

-   `writer`：Underlying writer
-   `bufferSize`：Internal buffer size

**Return:**

-   `*AsyncWriter`：Async writer instance

**Methods:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**Example:**

```go
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## Contextual Logging

### LoggerWithCtx Interface

```go
type LoggerWithCtx interface {
    Trace(ctx context.Context, v ...any)
    Tracef(ctx context.Context, format string, v ...any)
    Debug(ctx context.Context, v ...any)
    Debugf(ctx context.Context, format string, v ...any)
    Info(ctx context.Context, v ...any)
    Infof(ctx context.Context, format string, v ...any)
    Warn(ctx context.Context, v ...any)
    Warnf(ctx context.Context, format string, v ...any)
    Error(ctx context.Context, v ...any)
    Errorf(ctx context.Context, format string, v ...any)
    Fatal(ctx context.Context, v ...any)
    Fatalf(ctx context.Context, format string, v ...any)
    Panic(ctx context.Context, v ...any)
    Panicf(ctx context.Context, format string, v ...any)
}
```

### Context Functions

```go
func SetTrace(traceID string)
func GetTrace() string
```

Sets and gets the trace ID for the current goroutine.

**Example:**

```go
log.SetTrace("trace-123-456")
log.Info("This message will include trace ID")

traceID := log.GetTrace()
fmt.Println("Current trace ID:", traceID)
```

## Build Tags

The library supports conditional compilation using build tags:

### Default Mode

```bash
go build
```

-   Full functionality enabled
-   Debug messages included
-   Standard performance

### Debug Mode

```bash
go build -tags debug
```

-   Enhanced debugging information
-   Additional runtime checks
-   Detailed caller information

### Release Mode

```bash
go build -tags release
```

-   Optimized for production
-   Debug messages disabled
-   Automatic log rotation enabled

### Discard Mode

```bash
go build -tags discard
```

-   Maximum performance
-   All logging operations are no-ops
-   Zero overhead

### Combined Modes

```bash
go build -tags "debug,discard"    # Debug with discard
go build -tags "release,discard"  # Release with discard
```

## Performance Optimization

### Object Pooling

The library internally uses `sync.Pool` to manage:

-   Log entry objects
-   Byte buffers
-   Formatter buffers

This reduces garbage collection pressure in high-throughput scenarios.

### Level Checking

Log level checks occur before expensive operations:

```go
// Efficient - message formatting only happens if level is enabled
logger.Debugf("Expensive operation result: %+v", expensiveCall())

// Less efficient when debug is disabled in production
result := expensiveCall()
logger.Debug("Result:", result)
```

### Asynchronous Writing

For high-throughput applications:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // Large buffer
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### Build Tag Optimization

Use appropriate build tags based on environment:

-   Development: Default or debug tag
-   Production: Release tag
-   Performance-critical: Discard tag

## Examples

### Basic Usage

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("Application starting")
    log.Warn("This is a warning")
    log.Error("This is an error")
}
```

### Custom Logger

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // Configure the logger
    logger.SetLevel(log.DebugLevel)
    logger.EnableCaller(true)
    logger.SetPrefixMsg("[MyApp] ")

    // Set output to file
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("Custom logger configured")
    logger.Debug("Debug info with caller")
}
```

### High Performance Logging

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Create hourly rotating file writer
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // Wrap with async writer for performance
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // Skip debug in production

    // High-throughput logging
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### Context-aware Logging

```go
package main

import (
    "context"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()
    ctxLogger := logger.CloneToCtx()

    ctx := context.Background()
    log.SetTrace("trace-123-456")

    ctxLogger.Info(ctx, "Processing user request")
    ctxLogger.Debug(ctx, "Validation completed")
}
```

### Custom JSON Formatter

```go
package main

import (
    "encoding/json"
    "os"
    "time"
    "github.com/lazygophers/log"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *log.Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339Nano),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "pid":       entry.Pid,
        "gid":       entry.Gid,
    }

    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }

    if entry.CallerName != "" {
        data["caller"] = map[string]interface{}{
            "function": entry.CallerName,
            "file":     entry.CallerFile,
            "line":     entry.CallerLine,
        }
    }

    jsonData, _ := json.MarshalIndent(data, "", "  ")
    return append(jsonData, '\n')
}

func main() {
    logger := log.New()
    logger.SetFormatter(&JSONFormatter{})
    logger.EnableCaller(true)
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("JSON formatted message")
}
```

## Error Handling

For performance reasons, most logger methods do not return errors. If you need error handling for output operations, implement a custom writer:

```go
type ErrorCapturingWriter struct {
    writer io.Writer
    lastError error
}

func (w *ErrorCapturingWriter) Write(data []byte) (int, error) {
    n, err := w.writer.Write(data)
    if err != nil {
        w.lastError = err
    }
    return n, err
}

func (w *ErrorCapturingWriter) LastError() error {
    return w.lastError
}
```

## Thread Safety

All logger operations are thread-safe and can be used concurrently across multiple goroutines without additional synchronization mechanisms.

---

## 🌍 Multilingual Documentation

This document is available in multiple languages:

-   [🇺🇸 English](API.md) (current)
-   [🇨🇳 Simplified Chinese](API_zh-CN.md)
-   [🇹🇼 Traditional Chinese](API_zh-TW.md)

---

**LazyGophers Log Complete API Reference - Build Better Applications with Exceptional Logging! 🚀**
