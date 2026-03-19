---
titleSuffix: " | LazyGophers Log - API Documentation"
---

# 📚 API Documentation

Complete API reference for the LazyGophers Log library.

## Core Types

### Logger

Main logging structure with configurable options.

```go
type Logger struct {
    // Configuration fields
    level    Level
    output   io.Writer
    formatter Formatter
    // ... additional fields
}
```

#### Constructor

```go
func New() *Logger
func New() *Logger {
    return &Logger{
        level:    DebugLevel,
        output:   os.Stdout,
        formatter: DefaultFormatter{},
    }
}
```

#### Methods

```go
// Logging methods
func (l *Logger) Trace(msg string) *Logger
func (l *Logger) Debug(msg string) *Logger
func (l *Logger) Info(msg string) *Logger
func (l *Logger) Warn(msg string) *Logger
func (l *Logger) Error(msg string) *Logger
func (l *Logger) Fatal(msg string) *Logger
func (l *Logger) Panic(msg string) *Logger

// Formatted logging
func (l *Logger) Tracef(format string, args ...any) *Logger
func (l *Logger) Debugf(format string, args ...any) *Logger
func (l *Logger) Infof(format string, args ...any) *Logger
func (l *Logger) Warnf(format string, args ...any) *Logger
func (l *Logger) Errorf(format string, args ...any) *Logger
func (l *Logger) Fatalf(format string, args ...any) *Logger
func (l *Logger) Panicf(format string, args ...any) *Logger

// Configuration methods
func (l *Logger) SetLevel(level Level) *Logger
func (l *Logger) EnableCaller(enable bool) *Logger
func (l *Logger) EnableTrace(enable bool) *Logger
func (l *Logger) SetCallerDepth(depth int) *Logger
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *logger) SetSuffixMsg(suffix string) *Logger
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
func (l *Logger) CloneToCtx(ctx context.Context) context.Context
```

### Entry

Represents a single log record with metadata.

```go
type Entry struct {
    Timestamp time.Time
    Level     Level
    Message   string
    Pid       int
    Gid       int
    TraceID   string
    Caller    string
    // ... additional fields
}
```

### Level

Log levels from PanicLevel to TraceLevel.

```go
type Level int

const (
    PanicLevel Level = iota
    FatalLevel
    ErrorLevel
    WarnLevel
    InfoLevel
    DebugLevel
    TraceLevel
)
```

### Formatter

Interface for custom formatting.

```go
type Formatter interface {
    Format(entry *Entry) []byte
    FormatFull(entry *Entry, parseable bool) ([]byte, error)
}
```

## Global Functions

```go
// Alias for New()
func New() *Logger

// Global logging functions
func Trace(msg string) *Logger
func Debug(msg string) *Logger
func Info(msg string) *Logger
func Warn(msg string) *Logger
func Error(msg string) *Logger
func Fatal(msg string)
func Panic(msg string)

// Formatted versions
func Tracef(format string, args ...any) *Logger
func Debugf(format string, args ...any) *Logger
func Infof(format string, args ...any) *Logger
func Warnf(format string, args ...any) *Logger
func Errorf(format string, args ...any) *Logger
func Fatalf(format string, args ...any)
func Panicf(format string, args ...any)

// Global configuration
func SetLevel(level Level)
func EnableCaller(enable bool)
func EnableTrace(enable bool)
```

## Writers

### Output Writers

```go
func GetOutputWriterHourly(filename string) io.Writer
func GetOutputWriterDaily(filename string) io.Writer
func GetOutputWriterSize(filename string, maxSize int64) io.Writer
```

### Async Writer

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) io.Writer
```

## Constants

```go
const (
    // Build tags
    DefaultBuildTag  = "default"
    DebugBuildTag    = "debug"
    ReleaseBuildTag  = "release"
    DiscardBuildTag  = "discard"
    
    // Default values
    DefaultLevel     = DebugLevel
    DefaultDepth     = 2
)
```

## See Also

-   [README](README.md) - Overview and quick start
-   [Examples](../examples/) - Usage examples
-   [Benchmark Results](../docs/benchmark.md) - Performance metrics
