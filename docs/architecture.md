# Constant Package Architecture

## Overview

The `constant` package defines all core interfaces for the logging library, providing a clean separation between interface definitions and implementations.

## Package Structure

```
constant/
├── go.mod          # Independent module definition
└── interface.go    # Core interface definitions
```

## Interfaces

### Hook

Hook interface allows intercepting, modifying, filtering, or enriching log entries.

```go
type Hook interface {
    OnWrite(entry interface{}) interface{}
}
```

**Usage:**
- Return `nil` to filter out the log entry
- Return modified `entry` to change the log
- Return original `entry` to allow logging

**Example:**
```go
hook := constant.HookFunc(func(entry interface{}) interface{} {
    if e, ok := entry.(*log.Entry); ok {
        if strings.Contains(e.Message, "secret") {
            return nil // Filter out
        }
    }
    return entry
})

logger.AddHook(hook)
```

### Format

Format interface defines how log entries are converted to bytes.

```go
type Format interface {
    Format(entry interface{}) []byte
}
```

**Usage:**
- Implement this interface to create custom formatters
- Use `interface{}` to avoid circular dependencies
- Type assert to `*log.Entry` inside your implementation

**Example:**
```go
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry interface{}) []byte {
    if e, ok := entry.(*log.Entry); ok {
        return []byte(fmt.Sprintf("[%s] %s\n", e.Level, e.Message))
    }
    return nil
}

logger.Format = &CustomFormatter{}
```

### FormatFull

FormatFull extends Format with additional controls.

```go
type FormatFull interface {
    Format

    ParsingAndEscaping(disable bool)
    Caller(disable bool)
    Clone() Format
}
```

**Methods:**
- `ParsingAndEscaping(disable bool)`: Control message parsing
- `Caller(disable bool)`: Control caller information display
- `Clone()`: Create a deep copy of the formatter

### Writer

Writer interface defines output destination.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

## Module Dependencies

```
┌─────────────────────────────────┐
│  constant (independent)         │
│  - Hook, Format, FormatFull     │
│  - No dependencies              │
└──────────┬──────────────────────┘
           │
           ├─────────────────┐
           │                 │
           ▼                 ▼
┌──────────────────┐  ┌──────────────────┐
│  log             │  │  log/hooks       │
│  - Logger        │  │  - Hook impls    │
│  - Entry         │  │  - Uses constant │
│  - Uses constant │  └──────────────────┘
└──────────────────┘
```

## Benefits

1. **No Circular Dependencies**: Interfaces defined separately from implementations
2. **Independent Versioning**: Each package can version independently
3. **Clean API**: Consumers interact through well-defined interfaces
4. **Extensibility**: Third parties can implement interfaces without modifying core

## Implementation Guidelines

When implementing interfaces:

1. **Always use `interface{}`** in interface signatures
2. **Type assert inside implementations** to concrete types
3. **Handle type assertion failures** gracefully (return `nil` or error)
4. **Document expected concrete types** in interface comments

Example pattern:
```go
// MyHook implements constant.Hook
type MyHook struct{}

func (h *MyHook) OnWrite(entry interface{}) interface{} {
    // Type assert to expected concrete type
    e, ok := entry.(*log.Entry)
    if !ok {
        return entry // Handle unexpected type
    }

    // Your logic here
    e.Message = "[MODIFIED] " + e.Message

    return e
}
```

## Migration Notes

If migrating from old API:

**Before:**
```go
type Format interface {
    Format(entry *Entry) []byte
}
```

**After:**
```go
type Format interface {
    Format(entry interface{}) []byte
}
```

Update implementations to include type assertion:
```go
func (f *MyFormatter) Format(entry interface{}) []byte {
    e, ok := entry.(*Entry)
    if !ok {
        return nil
    }
    // ... rest of implementation
}
```
