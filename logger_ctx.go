package log

import (
	"context"
	"io"
)

// todo: 移动到独立的子模块，使得当前的 log 包的模块复杂度减少

// CloneToCtx clones a Logger into a new LoggerWithCtx instance.
// The new LoggerWithCtx inherits all settings from the original Logger.
func (p *Logger) CloneToCtx() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: p.Clone(),
	}
}

// LoggerWithCtx is a context-aware logger.
// It embeds Logger and extends it so that all logging methods accept context.Context as the first parameter.
type LoggerWithCtx struct {
	*Logger
}

// newLoggerWithCtx creates and returns a new LoggerWithCtx instance.
func newLoggerWithCtx() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: newLogger(),
	}
}

// SetCallerDepth sets the caller depth for accurate file and line reporting.
// Returns p for method chaining.
func (p *LoggerWithCtx) SetCallerDepth(callerDepth int) *LoggerWithCtx {
	p.Logger.SetCallerDepth(callerDepth)
	return p
}

// SetPrefixMsg sets the log message prefix.
// Returns p for method chaining.
func (p *LoggerWithCtx) SetPrefixMsg(prefixMsg string) *LoggerWithCtx {
	p.Logger.SetPrefixMsg(prefixMsg)
	return p
}

// AppendPrefixMsg appends content to the existing prefix.
// Returns p for method chaining.
func (p *LoggerWithCtx) AppendPrefixMsg(prefixMsg string) *LoggerWithCtx {
	p.Logger.AppendPrefixMsg(prefixMsg)
	return p
}

// SetSuffixMsg sets the log message suffix.
// Returns p for method chaining.
func (p *LoggerWithCtx) SetSuffixMsg(suffixMsg string) *LoggerWithCtx {
	p.Logger.SetSuffixMsg(suffixMsg)
	return p
}

// AppendSuffixMsg appends content to the existing suffix.
// Returns p for method chaining.
func (p *LoggerWithCtx) AppendSuffixMsg(suffixMsg string) *LoggerWithCtx {
	p.Logger.AppendSuffixMsg(suffixMsg)
	return p
}

// Clone creates and returns a deep copy of the current LoggerWithCtx instance.
func (p *LoggerWithCtx) Clone() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: p.Logger.Clone(),
	}
}

// SetLevel sets the logging level.
// Returns p for method chaining.
func (p *LoggerWithCtx) SetLevel(level Level) *LoggerWithCtx {
	p.Logger.SetLevel(level)
	return p
}

// SetOutput sets the log output targets.
// Accepts one or more io.Writer instances.
// If no writer is provided, output will be disabled.
// Returns p for method chaining.
func (p *LoggerWithCtx) SetOutput(writes ...io.Writer) *LoggerWithCtx {
	p.Logger.SetOutput(writes...)
	return p
}

// Log logs a message at the given level.
// Skips logging if the context is canceled or timed out.
func (p *LoggerWithCtx) Log(ctx context.Context, level Level, args ...interface{}) {
	if ctx.Err() != nil {
		return
	}
	if !p.levelEnabled(level) {
		return
	}

	p.log(level, fastSprint(args...))
}

// Logf logs a formatted message at the given level.
// Skips logging if the context is canceled or timed out.
func (p *LoggerWithCtx) Logf(ctx context.Context, level Level, format string, args ...interface{}) {
	if ctx.Err() != nil {
		return
	}
	if !p.levelEnabled(level) {
		return
	}

	p.log(level, fastSprintf(format, args...))
}

// Trace logs a message at Trace level.
func (p *LoggerWithCtx) Trace(ctx context.Context, args ...interface{}) {
	p.Log(ctx, TraceLevel, args...)
}

// Debug logs a message at Debug level.
func (p *LoggerWithCtx) Debug(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

// Print logs a message at Debug level. It is an alias for Debug.
func (p *LoggerWithCtx) Print(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

// Info logs a message at Info level.
func (p *LoggerWithCtx) Info(ctx context.Context, args ...interface{}) {
	p.Log(ctx, InfoLevel, args...)
}

// Warn logs a message at Warn level.
func (p *LoggerWithCtx) Warn(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

// Warning logs a message at Warn level. It is an alias for Warn.
func (p *LoggerWithCtx) Warning(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

// Error logs a message at Error level.
func (p *LoggerWithCtx) Error(ctx context.Context, args ...interface{}) {
	p.Log(ctx, ErrorLevel, args...)
}

// Panic logs a message at Panic level, then panics.
func (p *LoggerWithCtx) Panic(ctx context.Context, args ...interface{}) {
	p.Log(ctx, PanicLevel, args...)
}

// Fatal logs a message at Fatal level, then calls os.Exit(1).
func (p *LoggerWithCtx) Fatal(ctx context.Context, args ...interface{}) {
	p.Log(ctx, FatalLevel, args...)
}

// Tracef logs a formatted message at Trace level.
func (p *LoggerWithCtx) Tracef(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, TraceLevel, format, args...)
}

// Printf logs a formatted message at Debug level. It is an alias for Debugf.
func (p *LoggerWithCtx) Printf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, DebugLevel, format, args...)
}

// Debugf logs a formatted message at Debug level.
func (p *LoggerWithCtx) Debugf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, DebugLevel, format, args...)
}

// Infof logs a formatted message at Info level.
func (p *LoggerWithCtx) Infof(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, InfoLevel, format, args...)
}

// Warnf logs a formatted message at Warn level.
func (p *LoggerWithCtx) Warnf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, WarnLevel, format, args...)
}

// Warningf logs a formatted message at Warn level. It is an alias for Warnf.
func (p *LoggerWithCtx) Warningf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, WarnLevel, format, args...)
}

// Errorf logs a formatted message at Error level.
func (p *LoggerWithCtx) Errorf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, ErrorLevel, format, args...)
}

// Fatalf logs a formatted message at Fatal level, then calls os.Exit(1).
func (p *LoggerWithCtx) Fatalf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, FatalLevel, format, args...)
}

// Panicf logs a formatted message at Panic level, then panics.
func (p *LoggerWithCtx) Panicf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, PanicLevel, format, args...)
}

// ParsingAndEscaping controls whether HTML escaping is disabled.
// Returns p for method chaining.
func (p *LoggerWithCtx) ParsingAndEscaping(disable bool) *LoggerWithCtx {
	p.Logger.ParsingAndEscaping(disable)
	return p
}

// Caller controls whether caller information is included in log output.
// Returns p for method chaining.
func (p *LoggerWithCtx) Caller(disable bool) *LoggerWithCtx {
	p.Logger.Caller(disable)
	return p
}
