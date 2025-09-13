package log

import (
	"context"
	"fmt"
	"io"

	"go.uber.org/zap/zapcore"
)

// CloneToCtx 将一个 Logger 克隆到一个新的 LoggerWithCtx 中。
// 这个新的 LoggerWithCtx 实例将继承原 Logger 的所有设置。
func (p *Logger) CloneToCtx() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: p.Clone(),
	}
}

// LoggerWithCtx 是一个带有 context.Context 的日志记录器。
// 它内嵌了 Logger，并扩展了其功能，使其所有日志记录方法都接受 context.Context 作为第一个参数。
type LoggerWithCtx struct {
	*Logger
}

// newLoggerWithCtx 创建并返回一个新的 LoggerWithCtx 实例。
func newLoggerWithCtx() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: newLogger(),
	}
}

// SetCallerDepth 设置调用者深度，用于正确显示调用日志的文件和行号。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) SetCallerDepth(callerDepth int) *LoggerWithCtx {
	p.Logger.SetCallerDepth(callerDepth)
	return p
}

// SetPrefixMsg 设置日志消息的前缀。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) SetPrefixMsg(prefixMsg string) *LoggerWithCtx {
	p.Logger.SetPrefixMsg(prefixMsg)
	return p
}

// AppendPrefixMsg 在现有前缀的末尾追加内容。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) AppendPrefixMsg(prefixMsg string) *LoggerWithCtx {
	p.Logger.AppendPrefixMsg(prefixMsg)
	return p
}

// SetSuffixMsg 设置日志消息的后缀。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) SetSuffixMsg(suffixMsg string) *LoggerWithCtx {
	p.Logger.SetSuffixMsg(suffixMsg)
	return p
}

// AppendSuffixMsg 在现有后缀的末尾追加内容。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) AppendSuffixMsg(suffixMsg string) *LoggerWithCtx {
	p.Logger.AppendSuffixMsg(suffixMsg)
	return p
}

// Clone 创建并返回一个当前 LoggerWithCtx 实例的深拷贝。
func (p *LoggerWithCtx) Clone() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: p.Logger.Clone(),
	}
}

// SetLevel 设置日志记录级别。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) SetLevel(level Level) *LoggerWithCtx {
	p.Logger.SetLevel(level)
	return p
}

// SetOutput 设置日志输出目标。
// 可以接受一个或多个 io.Writer。
// 如果没有提供 writer，输出将被禁用。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) SetOutput(writes ...io.Writer) *LoggerWithCtx {
	var ws []zapcore.WriteSyncer
	for _, write := range writes {
		if write == nil {
			continue
		}
		ws = append(ws, zapcore.AddSync(write))
	}

	if len(ws) == 0 {
		p.out = nil
	} else if len(ws) == 1 {
		p.out = ws[0]
	} else {
		p.out = zapcore.NewMultiWriteSyncer(ws...)
	}

	return p
}

// Log 记录一条通用日志。
func (p *LoggerWithCtx) Log(ctx context.Context, level Level, args ...interface{}) {
	if !p.levelEnabled(level) {
		return
	}

	p.log(level, fmt.Sprint(args...))
}

// Logf 记录一条格式化的通用日志。
func (p *LoggerWithCtx) Logf(ctx context.Context, level Level, format string, args ...interface{}) {
	if !p.levelEnabled(level) {
		return
	}

	p.log(level, fmt.Sprintf(format, args...))
}

// Trace 记录 Trace级别的日志。
func (p *LoggerWithCtx) Trace(ctx context.Context, args ...interface{}) {
	p.Log(ctx, TraceLevel, args...)
}

// Debug 记录 Debug级别的日志。
func (p *LoggerWithCtx) Debug(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

// Print 记录 Debug级别的日志, 是 Debug 的别名。
func (p *LoggerWithCtx) Print(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

// Info 记录 Info级别的日志。
func (p *LoggerWithCtx) Info(ctx context.Context, args ...interface{}) {
	p.Log(ctx, InfoLevel, args...)
}

// Warn 记录 Warn级别的日志。
func (p *LoggerWithCtx) Warn(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

// Warning 记录 Warn级别的日志, 是 Warn 的别名。
func (p *LoggerWithCtx) Warning(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

// Error 记录 Error级别的日志。
func (p *LoggerWithCtx) Error(ctx context.Context, args ...interface{}) {
	p.Log(ctx, ErrorLevel, args...)
}

// Panic 记录 Panic级别的日志, 记录日志后会引发 panic。
func (p *LoggerWithCtx) Panic(ctx context.Context, args ...interface{}) {
	p.Log(ctx, PanicLevel, args...)
}

// Fatal 记录 Fatal级别的日志, 记录日志后会调用 os.Exit(1)。
func (p *LoggerWithCtx) Fatal(ctx context.Context, args ...interface{}) {
	p.Log(ctx, FatalLevel, args...)
}

// Tracef 记录一条格式化的 Trace级别的日志。
func (p *LoggerWithCtx) Tracef(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, TraceLevel, format, args...)
}

// Printf 记录一条格式化的 Debug级别的日志, 是 Debugf 的别名。
func (p *LoggerWithCtx) Printf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, DebugLevel, format, args...)
}

// Debugf 记录一条格式化的 Debug级别的日志。
func (p *LoggerWithCtx) Debugf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, DebugLevel, format, args...)
}

// Infof 记录一条格式化的 Info级别的日志。
func (p *LoggerWithCtx) Infof(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, InfoLevel, format, args...)
}

// Warnf 记录一条格式化的 Warn级别的日志。
func (p *LoggerWithCtx) Warnf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, WarnLevel, format, args...)
}

// Warningf 记录一条格式化的 Warn级别的日志, 是 Warnf 的别名。
func (p *LoggerWithCtx) Warningf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, WarnLevel, format, args...)
}

// Errorf 记录一条格式化的 Error级别的日志。
func (p *LoggerWithCtx) Errorf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, ErrorLevel, format, args...)
}

// Fatalf 记录一条格式化的 Fatal级别的日志, 记录日志后会调用 os.Exit(1)。
func (p *LoggerWithCtx) Fatalf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, FatalLevel, format, args...)
}

// Panicf 记录一条格式化的 Panic级别的日志, 记录日志后会引发 panic。
func (p *LoggerWithCtx) Panicf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, PanicLevel, format, args...)
}

// ParsingAndEscaping 控制是否禁用HTML转义。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) ParsingAndEscaping(disable bool) *LoggerWithCtx {
	p.Logger.ParsingAndEscaping(disable)
	return p
}

// Caller 控制是否在日志中记录调用者信息。
// 返回 p 本身，以支持链式调用。
func (p *LoggerWithCtx) Caller(disable bool) *LoggerWithCtx {
	p.Logger.Caller(disable)
	return p
}
