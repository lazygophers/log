package logctx

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

func (l Level) String() string {
	return []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC"}[l]
}

type Logger struct {
	mu           sync.Mutex
	level        Level
	output       io.Writer
	prefixMsg    string
	suffixMsg    string
	callerDepth  int
	enableCaller bool
}

func New() *Logger {
	return &Logger{level: InfoLevel, output: os.Stdout, callerDepth: 2, enableCaller: true}
}

func (l *Logger) SetLevel(level Level) *Logger { l.mu.Lock(); defer l.mu.Unlock(); l.level = level; return l }
func (l *Logger) SetOutput(w io.Writer) *Logger { l.mu.Lock(); defer l.mu.Unlock(); l.output = w; return l }
func (l *Logger) SetPrefixMsg(prefix string) *Logger { l.mu.Lock(); defer l.mu.Unlock(); l.prefixMsg = prefix; return l }
func (l *Logger) SetSuffixMsg(suffix string) *Logger { l.mu.Lock(); defer l.mu.Unlock(); l.suffixMsg = suffix; return l }
func (l *Logger) SetCallerDepth(depth int) *Logger { l.mu.Lock(); defer l.mu.Unlock(); l.callerDepth = depth; return l }
func (l *Logger) EnableCaller(enable bool) *Logger { l.mu.Lock(); defer l.mu.Unlock(); l.enableCaller = enable; return l }

func (l *Logger) Clone() *Logger {
	l.mu.Lock(); defer l.mu.Unlock()
	return &Logger{level: l.level, output: l.output, prefixMsg: l.prefixMsg, suffixMsg: l.suffixMsg, callerDepth: l.callerDepth, enableCaller: l.enableCaller}
}

func (l *Logger) Log(ctx context.Context, level Level, args ...interface{}) {
	if ctx.Err() != nil { return }
	l.mu.Lock(); defer l.mu.Unlock()
	if level < l.level { return }
	msg := fmt.Sprint(args...)
	if l.prefixMsg != "" { msg = l.prefixMsg + " " + msg }
	if l.suffixMsg != "" { msg += " " + l.suffixMsg }
	fmt.Fprintln(l.output, l.format(level, msg))
}

func (l *Logger) Logf(ctx context.Context, level Level, format string, args ...interface{}) {
	if ctx.Err() != nil { return }
	l.mu.Lock(); defer l.mu.Unlock()
	if level < l.level { return }
	msg := fmt.Sprintf(format, args...)
	if l.prefixMsg != "" { msg = l.prefixMsg + " " + msg }
	if l.suffixMsg != "" { msg += " " + l.suffixMsg }
	fmt.Fprintln(l.output, l.format(level, msg))
}

func (l *Logger) format(level Level, msg string) string {
	ts := time.Now().Format("2006-01-02 15:04:05")
	var caller string
	if l.enableCaller {
		_, file, line, _ := runtime.Caller(l.callerDepth)
		caller = fmt.Sprintf(" [%s:%d]", file, line)
	}
	return fmt.Sprintf("[%s] %s%s: %s", ts, level, caller, msg)
}

func (l *Logger) Trace(ctx context.Context, args ...interface{}) { l.Log(ctx, TraceLevel, args...) }
func (l *Logger) Debug(ctx context.Context, args ...interface{}) { l.Log(ctx, DebugLevel, args...) }
func (l *Logger) Print(ctx context.Context, args ...interface{}) { l.Log(ctx, DebugLevel, args...) }
func (l *Logger) Info(ctx context.Context, args ...interface{}) { l.Log(ctx, InfoLevel, args...) }
func (l *Logger) Warn(ctx context.Context, args ...interface{}) { l.Log(ctx, WarnLevel, args...) }
func (l *Logger) Warning(ctx context.Context, args ...interface{}) { l.Log(ctx, WarnLevel, args...) }
func (l *Logger) Error(ctx context.Context, args ...interface{}) { l.Log(ctx, ErrorLevel, args...) }
func (l *Logger) Fatal(ctx context.Context, args ...interface{}) { l.Log(ctx, FatalLevel, args...); os.Exit(1) }
func (l *Logger) Panic(ctx context.Context, args ...interface{}) { msg := fmt.Sprint(args...); l.Log(ctx, PanicLevel, msg); panic(msg) }

func (l *Logger) Tracef(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, TraceLevel, format, args...) }
func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, DebugLevel, format, args...) }
func (l *Logger) Printf(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, DebugLevel, format, args...) }
func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, InfoLevel, format, args...) }
func (l *Logger) Warnf(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, WarnLevel, format, args...) }
func (l *Logger) Warningf(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, WarnLevel, format, args...) }
func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, ErrorLevel, format, args...) }
func (l *Logger) Fatalf(ctx context.Context, format string, args ...interface{}) { l.Logf(ctx, FatalLevel, format, args...); os.Exit(1) }
func (l *Logger) Panicf(ctx context.Context, format string, args ...interface{}) { msg := fmt.Sprintf(format, args...); l.Logf(ctx, PanicLevel, msg); panic(msg) }

func (l *Logger) Sync() error {
	if f, ok := l.output.(interface{ Sync() error }); ok {
		return f.Sync()
	}
	return nil
}
