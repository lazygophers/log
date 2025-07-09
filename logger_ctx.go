package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"io"
)

func (p *Logger) CloneToCtx() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: *p.Clone(),
	}
}

type LoggerWithCtx struct {
	Logger
}

// newLogger 创建默认日志记录器实例
func newLoggerWithCtx() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: *newLogger(),
	}
}

// 返回: 当前Logger实例（支持链式调用）
func (p *LoggerWithCtx) SetCallerDepth(callerDepth int) *LoggerWithCtx {
	p.SetCallerDepth(callerDepth)
	return p
}

// SetPrefixMsg 设置消息前缀
// 参数 prefixMsg: 字符串，将被转换为字节切片
// 返回: 当前Logger实例（支持链式调用）
func (p *LoggerWithCtx) SetPrefixMsg(prefixMsg string) *LoggerWithCtx {
	p.SetPrefixMsg(prefixMsg)
	return p
}

// AppendPrefixMsg 追加消息前缀
// 参数 prefixMsg: 字符串，将被追加到现有前缀
// 返回: 当前Logger实例（支持链式调用）
func (p *LoggerWithCtx) AppendPrefixMsg(prefixMsg string) *LoggerWithCtx {
	p.AppendPrefixMsg(prefixMsg)
	return p
}

// SetSuffixMsg 设置消息后缀
// 参数 suffixMsg: 字符串，将被转换为字节切片
// 返回: 当前Logger实例（支持链式调用）
func (p *LoggerWithCtx) SetSuffixMsg(suffixMsg string) *LoggerWithCtx {
	p.SetSuffixMsg(suffixMsg)
	return p
}

// AppendSuffixMsg 追加消息后缀
// 参数 suffixMsg: 字符串，将被追加到现有后缀
// 返回: 当前Logger实例（支持链式调用）
func (p *LoggerWithCtx) AppendSuffixMsg(suffixMsg string) *LoggerWithCtx {
	p.AppendSuffixMsg(suffixMsg)
	return p
}

// Clone 创建Logger的深拷贝实例
// 返回: 新的Logger指针
func (p *LoggerWithCtx) Clone() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: *p.Logger.Clone(),
	}
}

// SetLevel 设置日志级别
// 参数 level: 目标日志级别
// 返回: 当前Logger实例（支持链式调用）
func (p *LoggerWithCtx) SetLevel(level Level) *LoggerWithCtx {
	p.SetLevel(level)
	return p
}

// SetOutput 设置日志输出目标
// 参数 writes: 一个或多个io.Writer接口实现
// 返回: 当前Logger实例（支持链式调用）
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

func (p *LoggerWithCtx) Log(ctx context.Context, level Level, args ...interface{}) {
	p.log(level, fmt.Sprint(args...))
}

func (p *LoggerWithCtx) Logf(ctx context.Context, level Level, format string, args ...interface{}) {
	p.log(level, fmt.Sprintf(format, args...))
}

func (p *LoggerWithCtx) Trace(ctx context.Context, args ...interface{}) {
	p.Log(ctx, TraceLevel, args...)
}

func (p *LoggerWithCtx) Debug(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

func (p *LoggerWithCtx) Print(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

func (p *LoggerWithCtx) Info(ctx context.Context, args ...interface{}) {
	p.Log(ctx, InfoLevel, args...)
}

func (p *LoggerWithCtx) Warn(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

func (p *LoggerWithCtx) Warning(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

func (p *LoggerWithCtx) Error(ctx context.Context, args ...interface{}) {
	p.Log(ctx, ErrorLevel, args...)
}

func (p *LoggerWithCtx) Panic(ctx context.Context, args ...interface{}) {
	p.Log(ctx, PanicLevel, args...)
}

func (p *LoggerWithCtx) Fatal(ctx context.Context, args ...interface{}) {
	p.Log(ctx, FatalLevel, args...)
}

func (p *LoggerWithCtx) Tracef(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, TraceLevel, format, args...)
}

func (p *LoggerWithCtx) Printf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, DebugLevel, format, args...)
}

func (p *LoggerWithCtx) Debugf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, DebugLevel, format, args...)
}

func (p *LoggerWithCtx) Infof(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, InfoLevel, format, args...)
}

func (p *LoggerWithCtx) Warnf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, WarnLevel, format, args...)
}

func (p *LoggerWithCtx) Warningf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, WarnLevel, format, args...)
}

func (p *LoggerWithCtx) Errorf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, ErrorLevel, format, args...)
}

func (p *LoggerWithCtx) Fatalf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, FatalLevel, format, args...)
}

func (p *LoggerWithCtx) Panicf(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, PanicLevel, format, args...)
}

func (p *LoggerWithCtx) ParsingAndEscaping(disable bool) *LoggerWithCtx {
	p.ParsingAndEscaping(disable)
	return p
}

func (p *LoggerWithCtx) Caller(disable bool) *LoggerWithCtx {
	p.Caller(disable)
	return p
}
