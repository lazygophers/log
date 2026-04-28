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
