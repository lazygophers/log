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
