---
titleSuffix: " | LazyGophers Log"
---

# 📚 API Documentation

## Overview

LazyGophers Log provides a comprehensive logging API with multiple log levels, custom formatting, structured logging, and a powerful Hook system. This document covers all public APIs, configuration options, and usage patterns.

## Table of Contents

- [Core Types](#core-types)
- [Logger API](#logger-api)
- [Global Functions](#global-functions)
- [Log Levels](#log-levels)
- [Structured Logging](#structured-logging)
- [Formatters](#formatters)
- [Hooks](#hooks)
- [Output Writers](#output-writers)
- [Build Tags](#build-tags)
- [Examples](#examples)

## Core Types

### Logger

The main logger struct providing all logging functionality.

```go
type Logger struct {
    // Contains private fields for thread-safe operations
}
```

**Constructor:**

```go
func New() *Logger
```

Creates a new logger instance with default configuration:
- Level: `DebugLevel`
- Output: `os.Stdout`
- Formatter: Default text formatter with parsing/escaping disabled
- Caller tracking: `enabled`
- Trace tracking: `enabled`

**Example:**

```go
logger := log.New()
logger.Info("New logger created")
```

### Entry

Represents a single log entry with comprehensive metadata.

> **Note:** Entry is defined in the `constant` package to avoid circular dependencies.

```go
type Entry struct {
    // Hot path fields - accessed on every log call
    Level      Level     // Log level
    Pid        int       // Process ID
    Gid        int64     // Goroutine ID
    CallerLine int       // Caller line number

    // Timestamp
    Time       time.Time // Timestamp
    TimeStr    string    // Cached formatted timestamp
    TimeStrSet bool      // Internal flag

    // String fields (16 bytes each) - ordered by access frequency
    Message    string    // Log message
    File       string    // Caller file path
    CallerFunc string    // Caller function name
    CallerDir  string    // Caller directory
    CallerName string    // Caller short name
    TraceId    string    // Trace ID for distributed tracing

    // Byte slices (24 bytes each)
    PrefixMsg []byte     // Log prefix message
    SuffixMsg []byte     // Log suffix message

    // Structured fields
    Fields []KV          // Key-value pairs
}
```

**Key Methods:**

```go
// Reset resets Entry to initial values for pool reuse
func (e *Entry) Reset()

// MarshalJSON implements json.Marshaler for custom serialization
func (e *Entry) MarshalJSON() ([]byte, error)
```

### KV

Represents a key-value pair for structured logging.

```go
type KV struct {
    Key   string
    Value interface{}
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
- `level`: The minimum log level to process

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.SetLevel(log.InfoLevel)
logger.Debug("This won't be displayed")  // Ignored
logger.Info("This will be displayed")    // Processed
```

#### Level

```go
func (l *Logger) Level() Level
```

Returns the current log level.

**Returns:**
- `Level`: The current minimum log level

**Example:**

```go
currentLevel := logger.Level()
fmt.Printf("Current level: %v\n", currentLevel)
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Sets one or more output destinations for log messages.

**Parameters:**
- `writers`: One or more `io.Writer` output destinations

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
// Single output
logger.SetOutput(os.Stdout)

// Multiple outputs
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)

// With rotating writer
logger.SetOutput(
    os.Stdout,
    log.GetOutputWriterHourly("/var/log/app.log", 100*1024*1024, 168),
)
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Sets the caller stack depth for determining the calling location.

**Parameters:**
- `depth`: Stack depth (default: 4)

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.SetCallerDepth(5)
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Enables or disables caller information in log output.

**Parameters:**
- `enable`: `true` to include caller information

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.EnableCaller(true)   // Include caller info
logger.EnableCaller(false)  // Exclude caller info
```

#### EnableTrace

```go
func (l *Logger) EnableTrace(enable bool) *Logger
```

Enables or disables trace information (Goroutine ID, Trace ID).

**Parameters:**
- `enable`: `true` to include trace information

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.EnableTrace(true)   // Include trace info
logger.EnableTrace(false)  // Exclude trace info
```

#### SetPrefixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
```

Sets a prefix message for all log entries.

**Parameters:**
- `prefix`: The prefix string

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.SetPrefixMsg("[MyApp] ")
logger.Info("message")  // Output: [MyApp] message
```

#### AppendPrefixMsg

```go
func (l *Logger) AppendPrefixMsg(prefix string) *Logger
```

Appends to the existing prefix message.

**Parameters:**
- `prefix`: The string to append

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.SetPrefixMsg("[App] ")
logger.AppendPrefixMsg("[v1.0] ")
logger.Info("message")  // Output: [App] [v1.0] message
```

#### SetSuffixMsg

```go
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Sets a suffix message for all log entries.

**Parameters:**
- `suffix`: The suffix string

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.SetSuffixMsg(" [END]")
logger.Info("message")  // Output: message [END]
```

#### AppendSuffixMsg

```go
func (l *Logger) AppendSuffixMsg(suffix string) *Logger
```

Appends to the existing suffix message.

**Parameters:**
- `suffix`: The string to append

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.SetSuffixMsg(" [1]")
logger.AppendSuffixMsg(" [2]")
logger.Info("message")  // Output: message [1] [2]
```

#### ParsingAndEscaping

```go
func (l *Logger) ParsingAndEscaping(disable bool) *Logger
```

Controls message parsing and character escaping.

**Parameters:**
- `disable`: `true` to disable parsing and escaping

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
logger.ParsingAndEscaping(false)  // Enable parsing/escaping
logger.ParsingAndEscaping(true)   // Disable parsing/escaping
```

#### AddHook

```go
func (l *Logger) AddHook(hook constant.Hook) *Logger
```

Adds a hook to process log entries before writing.

**Parameters:**
- `hook`: A hook implementing the `constant.Hook` interface

**Returns:**
- `*Logger`: Returns itself for method chaining

**Example:**

```go
type MyHook struct{}

func (h *MyHook) OnWrite(entry interface{}) interface{} {
    e := entry.(*log.Entry)
    e.Message = "[HOOK] " + e.Message
    return e
}

logger.AddHook(&MyHook{})
```

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Creates a deep copy of the logger with independent configuration.

**Returns:**
- `*Logger`: A new logger instance

**Example:**

```go
original := log.New().SetLevel(log.InfoLevel)
cloned := original.Clone()

// Cloned logger has independent configuration
cloned.SetLevel(log.DebugLevel)  // Doesn't affect original
```

#### Sync

```go
func (l *Logger) Sync()
```

Flushes all buffered log entries to their output destinations.

**Example:**

```go
logger.Info("Before sync")
logger.Sync()
logger.Info("After sync")
```

### Logging Methods

#### Log

```go
func (l *Logger) Log(level Level, args ...interface{})
```

Logs a message at the specified level.

**Parameters:**
- `level`: The log level
- `args`: The message arguments (formatted with `fmt.Sprint`)

**Example:**

```go
logger.Log(log.InfoLevel, "Info message")
logger.Log(log.WarnLevel, "Warning:", "something happened")
```

#### Logf

```go
func (l *Logger) Logf(level Level, format string, args ...interface{})
```

Logs a formatted message at the specified level.

**Parameters:**
- `level`: The log level
- `format`: The format string
- `args`: The format arguments

**Example:**

```go
logger.Logf(log.InfoLevel, "User %s logged in", "admin")
logger.Logf(log.ErrorLevel, "Failed to connect: %v", err)
```

#### Trace

```go
func (l *Logger) Trace(args ...interface{})
```

Logs a trace-level message.

**Example:**

```go
logger.Trace("Detailed trace information")
```

#### Debug

```go
func (l *Logger) Debug(args ...interface{})
```

Logs a debug-level message.

**Example:**

```go
logger.Debug("Debug information")
```

#### Info

```go
func (l *Logger) Info(args ...interface{})
```

Logs an info-level message.

**Example:**

```go
logger.Info("Information message")
```

#### Warn

```go
func (l *Logger) Warn(args ...interface{})
```

Logs a warning-level message.

**Example:**

```go
logger.Warn("Warning message")
```

#### Error

```go
func (l *Logger) Error(args ...interface{})
```

Logs an error-level message.

**Example:**

```go
logger.Error("Error occurred")
```

#### Fatal

```go
func (l *Logger) Fatal(args ...interface{})
```

Logs a fatal-level message and exits the program.

**Example:**

```go
logger.Fatal("Fatal error, exiting")
```

#### Panic

```go
func (l *Logger) Panic(args ...interface{})
```

Logs a panic-level message and panics.

**Example:**

```go
logger.Panic("Panic condition")
```

#### Tracef, Debugf, Infof, Warnf, Errorf, Fatalf, Panicf

Formatted versions of the logging methods.

```go
func (l *Logger) Tracef(format string, args ...interface{})
func (l *Logger) Debugf(format string, args ...interface{})
func (l *Logger) Infof(format string, args ...interface{})
func (l *Logger) Warnf(format string, args ...interface{})
func (l *Logger) Errorf(format string, args ...interface{})
func (l *Logger) Fatalf(format string, args ...interface{})
func (l *Logger) Panicf(format string, args ...interface{})
```

**Example:**

```go
logger.Infof("User %s logged in from %s", username, ip)
logger.Errorf("Failed to connect: %v", err)
```

#### Structured Logging Methods

##### Tracew, Debugw, Infow, Warnw, Errorw, Fatalw, Panicw

Structured logging methods with key-value pairs.

```go
func (l *Logger) Tracew(msg string, args ...interface{})
func (l *Logger) Debugw(msg string, args ...interface{})
func (l *Logger) Infow(msg string, args ...interface{})
func (l *Logger) Warnw(msg string, args ...interface{})
func (l *Logger) Errorw(msg string, args ...interface{})
func (l *Logger) Fatalw(msg string, args ...interface{})
func (l *Logger) Panicw(msg string, args ...interface{})
```

**Parameters:**
- `msg`: The log message
- `args`: Key-value pairs (key1, value1, key2, value2, ...)

**Example:**

```go
logger.Infow("User logged in",
    "user_id", 12345,
    "username", "admin",
    "ip", "192.168.1.100",
    "success", true,
)

// Output: ... User logged in user_id=12345 username=admin ip=192.168.1.100 success=true
```

## Global Functions

Package-level functions that use the standard logger instance.

```go
func New() *Logger
func SetLevel(level Level) *Logger
func GetLevel() Level
func Sync()
func Clone() *Logger
func SetCallerDepth(callerDepth int) *Logger
func SetPrefixMsg(prefixMsg string) *Logger
func AppendPrefixMsg(prefixMsg string) *Logger
func SetSuffixMsg(suffixMsg string) *Logger
func AppendSuffixMsg(suffixMsg string) *Logger
func ParsingAndEscaping(disable bool) *Logger
func Caller(disable bool) *Logger

func Trace(args ...interface{})
func Debug(args ...interface{})
func Info(args ...interface{})
func Warn(args ...interface{})
func Error(args ...interface{})
func Fatal(args ...interface{})
func Panic(args ...interface{})

func Tracef(format string, args ...interface{})
func Debugf(format string, args ...interface{})
func Infof(format string, args ...interface{})
func Warnf(format string, args ...interface{})
func Errorf(format string, args ...interface{})
func Fatalf(format string, args ...interface{})
func Panicf(format string, args ...interface{})

func Tracew(msg string, args ...interface{})
func Debugw(msg string, args ...interface{})
func Infow(msg string, args ...interface{})
func Warnw(msg string, args ...interface{})
func Errorw(msg string, args ...interface{})
func Fatalw(msg string, args ...interface{})
func Panicw(msg string, args ...interface{})

func Pid() int
```

**Example:**

```go
import "github.com/lazygophers/log"

func main() {
    // Use global logger
    log.Info("Application started")
    log.SetLevel(log.WarnLevel)
    log.Debug("This won't be logged")
    log.Warn("This will be logged")
}
```

## Log Levels

### Level Constants

```go
const (
    PanicLevel Level = iota  // Highest
    FatalLevel
    ErrorLevel
    WarnLevel
    InfoLevel
    DebugLevel
    TraceLevel               // Lowest
)
```

### Level Comparison

Log levels are ordered from lowest to highest verbosity:

```
Trace (most verbose)
  ↓
Debug
  ↓
Info
  ↓
Warn
  ↓
Error
  ↓
Fatal
  ↓
Panic (least verbose)
```

When you set the log level to `InfoLevel`, only `Info`, `Warn`, `Error`, `Fatal`, and `Panic` messages will be logged. `Debug` and `Trace` messages will be ignored.

### Level Methods

```go
func (level Level) String() string
func (level Level) MarshalText() ([]byte, error)
```

**Example:**

```go
level := log.InfoLevel
fmt.Println(level.String())  // "info"

data, _ := level.MarshalText()
fmt.Println(string(data))    // "info"
```

## Structured Logging

Structured logging allows you to attach key-value pairs to log messages for better querying and analysis.

### Basic Usage

```go
logger.Infow("User action",
    "user_id", 12345,
    "action", "login",
    "success", true,
)
```

### Multiple Fields

```go
logger.Errorw("Database error",
    "error", err,
    "query", "SELECT * FROM users",
    "duration_ms", 150,
    "retry_count", 3,
)
```

### With Dynamic Fields

```go
fields := []interface{}{
    "request_id", uuid.New().String(),
    "user_id", userID,
    "method", r.Method,
    "path", r.URL.Path,
}

logger.Infow("HTTP request", fields...)
```

## Formatters

### Formatter Interface

```go
type Format interface {
    Format(entry interface{}) []byte
}

type FormatFull interface {
    Format
    ParsingAndEscaping(disable bool)
    Caller(disable bool)
    Clone() Format
}
```

### Text Formatter

Default text formatter with color support.

```go
type Formatter struct {
    DisableParsingAndEscaping bool
    DisableCaller             bool
}
```

**Example:**

```go
formatter := &log.Formatter{
    DisableParsingAndEscaping: true,
    DisableCaller:             false,
}

logger.Format = formatter
```

### JSON Formatter

JSON output formatter.

```go
type JSONFormatter struct {
    EnablePrettyPrint bool
    DisableCaller     bool
    DisableTrace      bool
}
```

**Example:**

```go
// Compact JSON
logger.Format = &log.JSONFormatter{}

// Pretty printed JSON
logger.Format = &log.JSONFormatter{EnablePrettyPrint: true}

// JSON without caller information
logger.Format = &log.JSONFormatter{DisableCaller: true}
```

### Custom Formatter

```go
type MyFormatter struct{}

func (f *MyFormatter) Format(entry interface{}) []byte {
    e := entry.(*log.Entry)
    output := fmt.Sprintf("[%s] %s: %s\n",
        e.Time.Format(time.RFC3339),
        e.Level,
        e.Message,
    )
    return []byte(output)
}

// Usage
logger.Format = &MyFormatter{}
```

## Hooks

### Hook Interface

```go
type Hook interface {
    OnWrite(entry interface{}) interface{}
}
```

### HookFunc Convenience Type

```go
type HookFunc func(entry interface{}) interface{}

func (h HookFunc) OnWrite(entry interface{}) interface{} {
    return h(entry)
}
```

### Hook Examples

#### Add Global Fields

```go
type EnvironmentHook struct {
    Env string
}

func (h *EnvironmentHook) OnWrite(entry interface{}) interface{} {
    e := entry.(*log.Entry)
    e.Fields = append(e.Fields, log.KV{
        Key:   "environment",
        Value: h.Env,
    })
    return e
}

logger.AddHook(&EnvironmentHook{Env: "production"})
```

#### Filter Sensitive Data

```go
type SensitiveHook struct{}

func (h *SensitiveHook) OnWrite(entry interface{}) interface{} {
    e := entry.(*log.Entry)
    if strings.Contains(e.Message, "password") {
        return nil // Filter out
    }
    return e
}

logger.AddHook(&SensitiveHook{})
```

#### Modify Content

```go
type PrefixHook struct {
    Prefix string
}

func (h *PrefixHook) OnWrite(entry interface{}) interface{} {
    e := entry.(*log.Entry)
    e.Message = h.Prefix + e.Message
    return e
}

logger.AddHook(&PrefixHook{Prefix: "[HOOK] "})
```

## Output Writers

### GetOutputWriterHourly

```go
func GetOutputWriterHourly(logDir string, maxSize int64, maxFiles int) io.Writer
```

Creates an hourly rotating log writer with size-based sharding.

**Parameters:**
- `logDir`: Directory for log files
- `maxSize`: Maximum size per file in bytes
- `maxFiles`: Maximum number of files to keep

**Returns:**
- `io.Writer`: A writer that handles rotation

**Example:**

```go
writer := log.GetOutputWriterHourly(
    "/var/log/app",
    100*1024*1024,  // 100 MB per file
    168,            // Keep 7 days (168 hours)
)

logger.SetOutput(os.Stdout, writer)
```

### Custom Writer

```go
type MyWriter struct{}

func (w *MyWriter) Write(p []byte) (n int, err error) {
    // Custom write logic
    return os.Stdout.Write(p)
}

logger.SetOutput(&MyWriter{})
```

## Build Tags

### Default

```go
// +build !debug,!release,!discard
```

Full functionality with all features enabled.

### Debug

```go
// +build debug
```

Enhanced debugging information and verbose output.

### Release

```go
// +build release
```

Production-optimized with debug messages disabled.

### Discard

```go
// +build discard
```

Maximum performance - all log operations are no-ops.

**Usage:**

```bash
# Default build
go build

# Debug build
go build -tags=debug

# Release build
go build -tags=release

# Discard build
go build -tags=discard
```

## Examples

### Basic Logging

```go
package main

import "github.com/lazygophers/log"

func main() {
    log.Trace("Trace message")
    log.Debug("Debug message")
    log.Info("Info message")
    log.Warn("Warning message")
    log.Error("Error message")
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
    logger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetPrefixMsg("[MyApp] ").
        SetOutput(
            os.Stdout,
            log.GetOutputWriterHourly("/var/log/app.log", 100*1024*1024, 168),
        )

    logger.Info("Application started")
    logger.Infow("User logged in", "user_id", 12345)
}
```

### JSON Output

```go
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New()
    logger.Format = &log.JSONFormatter{EnablePrettyPrint: true}

    logger.Infow("Service started",
        "port", 8080,
        "environment", "production",
    )
}
```

### With Hooks

```go
package main

import (
    "github.com/lazygophers/log"
    "github.com/lazygophers/log/constant"
)

type MyHook struct{}

func (h *MyHook) OnWrite(entry interface{}) interface{} {
    e := entry.(*log.Entry)
    e.Fields = append(e.Fields, log.KV{
        Key:   "processed",
        Value: "true",
    })
    return e
}

func main() {
    logger := log.New()
    logger.AddHook(&MyHook{})

    logger.Info("This will include the 'processed' field")
}
```

## Best Practices

1. **Use appropriate log levels**: Choose the right level for each message
2. **Structure your logs**: Use `*w` methods for key-value pairs
3. **Avoid sensitive data**: Use hooks to filter sensitive information
4. **Set reasonable caller depth**: Default is 4, adjust for wrapper functions
5. **Use object pools**: The library handles this automatically
6. **Consider performance**: Disable caller/trace in production if not needed
7. **Test your hooks**: Ensure hooks don't introduce performance issues

## See Also

- [Architecture Documentation](architecture.md)
- [Hook Guide](hooks_guide.md)
- [README](../README.md)
