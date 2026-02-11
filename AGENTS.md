# AGENTS.md

Instructions for agentic coding agents working in the LazyGophers Log repository.

## Build Commands

```bash
# Basic build
go build ./...

# Build with specific tags
go build -tags="debug" ./...
go build -tags="release" ./...
go build -tags="discard" ./...
```

## Test Commands

```bash
# Run all tests (all build tags)
make test

# Quick test (default build only)
make test-simple
go test ./...

# Verbose output
go test -v ./...

# Run specific test
go test -run TestLogger_SetLevel ./...

# Run tests with specific build tag
go test -tags="debug" ./...
go test -tags="release" ./...
go test -tags="discard" ./...
go test -tags="debug,discard" ./...
go test -tags="release,discard" ./...

# Benchmarks
go test -bench=. -benchmem ./...
make benchmark

# Coverage
make coverage-all
make coverage-html
```

## Lint Commands

```bash
# Standard Go tools
go vet ./...
go fmt ./...

# Static analysis (if installed)
staticcheck ./...

# Format imports (if goimports installed)
goimports -w .
```

## Code Style Guidelines

### Formatting
- **Indentation**: Tabs (width: 4 spaces in display)
- **Line endings**: LF only
- **Charset**: UTF-8
- **Trim trailing whitespace**: Yes
- Use `gofmt` and `goimports` for formatting

### Imports Order
1. Standard library imports
2. Third-party imports
3. Local/package imports

Example:
```go
import (
    "fmt"
    "io"
    "sync"
    
    "github.com/petermattis/goid"
    "go.uber.org/zap/zapcore"
)
```

### Naming Conventions
- **Types**: PascalCase (e.g., `Logger`, `Entry`, `WriteSyncerWrapper`)
- **Interfaces**: PascalCase with verb/noun (e.g., `Format`, `FormatFull`)
- **Functions**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase for unexported, PascalCase for exported
- **Constants**: PascalCase or camelCase (use iota for enums)
- **Build tags files**: `print_<tag>.go` pattern

### Error Handling
- Return errors as the last return value
- Use descriptive error messages
- Check errors immediately after calls
- Example: `if err != nil { return nil, fmt.Errorf("context: %w", err) }`

### Performance Patterns
- Use `sync.Pool` for object reuse (see `entryPool`, `bufferPool`)
- Add `//go:inline` hint for hot path functions
- Add `//go:noinline` to prevent inlining where stack depth matters
- Check log level before expensive operations
- Use `defer` for resource cleanup

### Comments
- Start with `// ` (space after slashes)
- Document all exported types and functions
- Use complete sentences
- For Chinese comments, use full-width punctuation

### Build Tags
Available tags:
- `debug`: Development mode with stdout output
- `release`: Production mode with file rotation
- `discard`: Discard all log output (for testing)
- Combinations: `debug,discard`, `release,discard`

Tag format at file top:
```go
//go:build debug && !discard
//go:build release && !discard
```

### Testing Patterns
- Use table-driven tests
- Mock external dependencies
- Test all build tag combinations
- Name tests descriptively: `Test<Type>_<Method>_<Scenario>`
- Use `t.TempDir()` for temporary files
- Check both success and error cases

### Project Structure
- Tests co-located: `*_test.go` next to source files
- Benchmarks: `*_performance_test.go`
- Build tag variants: `print_*.go` files
- No `internal/` packages - all exported

## Key Dependencies
- `go.uber.org/zap` - WriteSyncer and zapcore
- `github.com/petermattis/goid` - Goroutine ID
- `github.com/lestrrat-go/file-rotatelogs` - File rotation (indirect)

## Pre-commit Checklist
1. Run `go fmt ./...`
2. Run `go vet ./...`
3. Run `make test-quick` (all build tags)
4. Ensure no new lint warnings
5. Add tests for new functionality
6. Update documentation if needed
