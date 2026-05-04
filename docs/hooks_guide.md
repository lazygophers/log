# Hooks System Guide

The Hooks system allows you to intercept, modify, filter, or enrich log entries before they are written.

## Architecture

The hooks system is built on a clean interface-based architecture:

- **constant** package: Defines the `Hook` interface
- **log/hooks** package: Provides built-in hook implementations
- **log** package: Uses hooks through the `Hook` interface

## Basic Usage

### Adding a Hook

```go
import (
    "github.com/lazygophers/log"
    "github.com/lazygophers/log/constant"
)

logger := log.New()

// Add a simple hook
hook := constant.HookFunc(func(entry interface{}) interface{} {
    // Process the entry
    return entry // Return nil to filter out
})

logger.AddHook(hook)
```

### Using Built-in Hooks

```go
import "github.com/lazygophers/log/hooks"

// Sensitive data masking
maskHook := hooks.NewSensitiveDataMaskHook()
logger.AddHook(maskHook)

// Level filtering
levelHook := hooks.NewLevelFilterHook(int(log.InfoLevel))
logger.AddHook(levelHook)

// Message filtering
filterHook := hooks.NewMessageFilterHook()
filterHook.AddAllowPattern("important")
logger.AddHook(filterHook)
```

## Available Hooks

### 1. SensitiveDataMaskHook

Masks sensitive data like emails, credit cards, passwords, etc.

```go
maskHook := hooks.NewSensitiveDataMaskHook()

// Add custom pattern
maskHook.AddPattern(`\b\d{3}-\d{3}-\d{4}\b`)

// Set custom mask
maskHook.SetMask("[REDACTED]")

logger.AddHook(maskHook)
logger.Info("Email: test@example.com")
// Output: Email: ***
```

### 2. ContextEnrichHook

Automatically adds contextual fields to all log entries.

```go
fields := map[string]interface{}{
    "service": "my-service",
    "version": "1.0.0",
}
enrichHook := hooks.NewContextEnrichHook(fields)
logger.AddHook(enrichHook)
```

### 3. LevelFilterHook

Filters logs by minimum level.

```go
// Only log INFO and above
levelHook := hooks.NewLevelFilterHook(int(log.InfoLevel))
logger.AddHook(levelHook)
```

### 4. MessageFilterHook

Filters logs based on message content using regex patterns.

```go
filterHook := hooks.NewMessageFilterHook()

// Allow only messages containing "important"
filterHook.AddAllowPattern("important")

// Deny messages containing "debug"
filterHook.AddDenyPattern("debug")

logger.AddHook(filterHook)
```

### 5. FieldFilterHook

Filters logs based on field values.

```go
fieldFilter := hooks.NewFieldFilterHook()

// Only allow logs with environment=production
fieldFilter.AllowField("environment", "production")

// Deny logs with user=admin
fieldFilter.DenyField("user", "admin")

logger.AddHook(fieldFilter)
```

### 6. MinLengthHook

Filters messages shorter than minimum length.

```go
minHook := hooks.NewMinLengthHook(10)
logger.AddHook(minHook)
```

### 7. MaxLengthHook

Truncates messages longer than maximum length.

```go
maxHook := hooks.NewMaxLengthHook(100)
maxHook.TruncateSuffix = "..."

logger.AddHook(maxHook)
```

### 8. PrefixHook

Adds a prefix to all messages.

```go
prefixHook := hooks.NewPrefixHook("[APP] ")
logger.AddHook(prefixHook)
```

### 9. SuffixHook

Adds a suffix to all messages.

```go
suffixHook := hooks.NewSuffixHook(" [END]")
logger.AddHook(suffixHook)
```

### 10. ConditionalHook

Executes a hook only when a condition is met.

```go
condition := func(entry interface{}) bool {
    if e, ok := entry.(*log.Entry); ok {
        return e.Level == log.ErrorLevel
    }
    return false
}

prefixHook := hooks.NewPrefixHook("[ERROR] ")
conditionalHook := hooks.NewConditionalHook(condition, prefixHook)
logger.AddHook(conditionalHook)
```

### 11. ChainHook

Executes multiple hooks in sequence.

```go
chain := hooks.NewChainHook(
    hooks.NewPrefixHook("[CHAIN1] "),
    hooks.NewPrefixHook("[CHAIN2] "),
    hooks.NewMaxLengthHook(50),
)
logger.AddHook(chain)
```

## Creating Custom Hooks

You can create custom hooks by implementing the `Hook` interface:

```go
type CustomHook struct {
    // Your fields
}

func (h *CustomHook) OnWrite(entry interface{}) interface{} {
    // Type assert to *log.Entry if needed
    if e, ok := entry.(*log.Entry); ok {
        // Modify the entry
        e.Message = "[CUSTOM] " + e.Message
    }

    // Return nil to filter out
    // Return entry to allow logging
    return entry
}

// Usage
logger.AddHook(&CustomHook{})
```

Or use `HookFunc` for simple cases:

```go
hook := constant.HookFunc(func(entry interface{}) interface{} {
    if e, ok := entry.(*log.Entry); ok {
        // Custom logic
        if someCondition(e) {
            return nil // Filter out
        }
    }
    return entry
})

logger.AddHook(hook)
```

## Hook Chaining

Hooks are executed in the order they are added:

```go
logger.AddHook(hook1)
logger.AddHook(hook2)
logger.AddHook(hook3)

// Execution order: hook1 -> hook2 -> hook3
```

If any hook returns `nil`, the log entry is filtered and subsequent hooks are not executed.

## Managing Hooks

```go
// Add a single hook
logger.AddHook(hook)

// Add multiple hooks
logger.AddHooks(hook1, hook2, hook3)

// Remove all hooks
logger.RemoveHooks()
```

## Accessing Entry Fields in Hooks

When implementing custom hooks, you can access entry fields through type assertion:

```go
type entryLike struct {
    Level      int
    Message    string
    Time       time.Time
    Fields     []log.KV
    File       string
    CallerLine int
    CallerName string
    CallerDir  string
    CallerFunc string
    PrefixMsg  []byte
    SuffixMsg  []byte
    Pid        int
    Gid        int64
    TraceId    uint64
}

func (h *MyHook) OnWrite(entry interface{}) interface{} {
    if e, ok := entry.(entryLike); ok {
        // Access fields
        level := e.Level
        message := e.Message
        // ... process entry
    }
    return entry
}
```

## Performance Considerations

- Hooks are executed for every log entry
- Keep hook logic minimal to avoid performance impact
- Use level filtering early to skip unnecessary processing
- Consider using conditional hooks for expensive operations

## Complete Example

```go
package main

import (
    "github.com/lazygophers/log"
    "github.com/lazygophers/log/hooks"
)

func main() {
    logger := log.New()
    logger.EnableCaller(false)
    logger.EnableTrace(false)

    // Add multiple hooks
    logger.AddHooks(
        // Mask sensitive data
        hooks.NewSensitiveDataMaskHook(),

        // Add context
        hooks.NewContextEnrichHook(map[string]interface{}{
            "service": "my-app",
            "env":     "production",
        }),

        // Add prefix
        hooks.NewPrefixHook("[PROD] "),

        // Truncate long messages
        hooks.NewMaxLengthHook(200),
    )

    logger.Info("Application started")
    logger.Error("Database connection failed: password=secret123")
}
```
