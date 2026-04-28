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

## Hydra Orchestration Toolkit

Hydra is a Lead-driven orchestration toolkit. You (the Lead) make strategic
decisions at decision points; Hydra handles operational management.
`result.json` is the only completion evidence.

Why this design (vs. other coding-agent products):
- **SWF decider pattern, specialized for LLM deciders.** Hydra is the AWS SWF / Cadence / Temporal decider pattern. `hydra watch` is `PollForDecisionTask`; the Lead is the decider; `lead_terminal_id` enforces single-decider semantics.
- **Parallel-first, not bolted on.** `dispatch` + worktree + `merge` are first-class. Lead sequences nodes manually and passes context explicitly via `--context-ref`. Other products treat parallelism as open research; Hydra makes it the default.
- **Typed result contract.** Workers publish a schema-validated `result.json` (`outcome: completed | stuck | error`, optional `stuck_reason: needs_clarification | needs_credentials | needs_context | blocked_technical`). Other products return free-text final messages and require downstream parsing.
- **Lead intervention points.** `hydra reset --feedback` lets the Lead actually intervene at decision points instead of being block-and-join. A stale or wrong run is one `reset` away.

Core rules:
- Root cause first. Fix the implementation problem before changing tests.
- Do not hack tests, fixtures, or mocks to force a green result.
- Do not add silent fallbacks or swallowed errors.
- An assignment run is only complete when `result.json` exists and passes schema validation.

Workflow patterns:
1. Do the task directly when it is simple, local, or clearly faster without workflow overhead.
2. Use Hydra for ambiguous, risky, parallel, or multi-step work:
   ```
   hydra init --intent "<task>" --repo .
   hydra dispatch --workbench W --dispatch <id> --role <role> --intent "<desc>" --repo .
   hydra watch --workbench W --repo .
   # → DecisionPoint returned, decide next step
   hydra complete --workbench W --repo .
   ```
3. Use a direct isolated worker when only a separate worker is needed:
   `hydra spawn --task "<specific task>" --repo . [--worktree .]`

Agent launch rule:
- When dispatching Claude/Codex through TermCanvas CLI, start a fresh agent terminal with `termcanvas terminal create --prompt "..."`
- Do not use `termcanvas terminal input` for task dispatch; it is not a supported automation path

TermCanvas Computer Use:
- TermCanvas may dynamically inject a Computer Use MCP server into Claude/Codex terminals; it does not have to appear in static MCP settings files.
- For local macOS desktop apps or system UI, check for TermCanvas Computer Use before assuming only shell, browser, or Playwright tools are available.
- If available, call `status` first, then `setup` if permissions or helper health are missing, then `get_instructions` for the current operating protocol.
- Do not manually start `computer-use-helper`, write its state file, launch the MCP server, or hand-write JSON-RPC unless explicitly debugging Computer Use itself.

Workflow control:
- After dispatching, always call `hydra watch`. It returns at decision points.
1. Watch until decision point: `hydra watch --workbench <workbenchId> --repo .`
2. Inspect structured state: `hydra status --workbench <workbenchId> --repo .`
3. Reset a dispatch for rework: `hydra reset --workbench W --dispatch N --feedback "..." --repo .`
4. Approve a dispatch's output: `hydra approve --workbench W --dispatch N --repo .`
5. Merge parallel branches: `hydra merge --workbench W --dispatches A,B --repo .`
6. View event log: `hydra ledger --workbench <workbenchId> --repo .`
7. Clean up: `hydra cleanup --workbench <workbenchId> --repo .`

Telemetry polling:
1. Treat `hydra watch` as the main polling loop; do not infer progress from terminal prose alone.
2. Before deciding wait / retry / takeover, query:
   - `termcanvas telemetry get --workbench <workbenchId> --repo .`
   - `termcanvas telemetry get --terminal <terminalId>`
   - `termcanvas telemetry events --terminal <terminalId> --limit 20`
3. Trust `derived_status` and `task_status` as the primary decision signals.

`result.json` must contain (slim, schema_version `hydra/result/v0.1`):
- `schema_version`, `workbench_id`, `assignment_id`, `run_id` (passthrough IDs)
- `outcome` (completed/stuck/error — Hydra routes on this)
- `report_file` (path to a `report.md` written alongside `result.json`)

All human-readable content (summary, outputs, evidence, reflection) lives in
`report.md`. Hydra rejects any extra fields in `result.json`. Write `report.md`
first, then publish `result.json` atomically as the final artifact of the run.

When NOT to use: simple fixes, high-certainty tasks, or work that is faster to do directly in the current agent.

## TermCanvas Pin System

TermCanvas has a first-class pin store. Pins are persistent records of work
the user wants done — captured when the user expresses intent, not when the
work happens. Use the `termcanvas pin` CLI to read and write them. Any agent
terminal can record, read, and update pins.

When to record a pin:
- User says "记一下", "回头处理", "帮我留意", "later", "todo this", or any phrasing that defers the work.
- User describes a problem or idea but isn't asking you to fix it right now.
- User pastes a GitHub issue URL and asks you to track it (record the URL via `--link`).

Do NOT silently nod — capture the pin with `termcanvas pin add` so it survives the session.

Recording a pin:
```
termcanvas pin add --title "<short imperative>" --body "<detail>" [--link <url>]
```
- `--title`: short, scannable. Rephrase the user's words into imperative mood.
- `--body`: preserve enough context for a future agent or the user to resume without re-asking basic questions. Do not store only the user's raw sentence unless it is truly just a lightweight memo.
- For bugs, feature requests, research threads, design feedback, or follow-up engineering work, write the body like a compact issue. Prefer sections such as:
  `Background`: what prompted this and where it came from.
  `Observed / Request`: the concrete symptom, ask, or idea.
  `Expected / Goal`: what should be true when this is handled.
  `Evidence / References`: user quote, screenshot, link, file path, command output, or code location if available.
  `Next action`: the first useful step when someone picks it up.
  `Why pinned`: why this is being saved instead of handled immediately.
  `Unknowns`: missing decisions or facts that still need confirmation.
- If the information is thin, choose deliberately:
  If local context can answer it cheaply, inspect the relevant code, state, logs, or files before recording and include what you found.
  If the missing information changes scope, product behavior, security, or architecture, ask the user one concise question before recording.
  If the user is clearly deferring and cannot answer now, record the pin anyway but mark assumptions and unknowns explicitly.
- If it is only a personal memo or reminder, a short body is acceptable, but still include why it matters or when to revisit it if that is known.
- For multi-line bodies, pass real newlines. In shell commands, use ANSI-C quoting such as
  `--body $'line 1\nline 2'`; do not put literal `\n` sequences inside ordinary quotes.
- `--link <url>`: attach an external reference (GitHub issue, doc, etc.). Use `--link-type github_issue` for issue URLs.
- Repo defaults to cwd. Pass `--repo <path>` only if you need a different one.

Reading and updating pins:
- `termcanvas pin list` — list pins for the current repo (filter `--status done` etc.)
- `termcanvas pin show <id>` — read a single pin before acting on it
- `termcanvas pin update <id> --status done` — mark complete after finishing the work
- `termcanvas pin update <id> --body "..."` — refine the description as you learn more

Rules:
- Pins belong to the user. Don't invent pins the user didn't ask for.
- One pin per intent. Three deferred items = three `pin add` calls.
- After completing work that originated from a pin, call `pin update <id> --status done`.
- The pin store is local to TermCanvas. It does NOT auto-sync to GitHub. If the user wants something on GitHub, they will say so explicitly.
- Status values: `open` (default), `done`, `dropped`. Pick `dropped` (not delete) when a pin is abandoned, so the history is preserved.
