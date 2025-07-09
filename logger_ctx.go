package log

import (
	"fmt"
	"io"

	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
)

func (p *Logger) CloneToCtx() *LoggerWithCtx {
	return &LoggerWithCtx{
		Logger: *p.Clone(),
	}
}

// LoggerWithCtx 带上下文的日志记录器
// 继承自 Logger，添加了支持 context.Context 的日志方法
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

// Log 记录日志
// 参数 ctx: 上下文
// 参数 level: 日志级别
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Log(ctx context.Context, level Level, args ...interface{}) {
	p.log(level, fmt.Sprint(args...))
}

// Logf 记录格式化日志
// 参数 ctx: 上下文
// 参数 level: 日志级别
// 参数 format: 格式化字符串
// 参数 args: 格式化参数
func (p *LoggerWithCtx) Logf(ctx context.Context, level Level, format string, args ...interface{}) {
	p.log(level, fmt.Sprintf(format, args...))
}

// Trace 记录TRACE级别日志
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Trace(ctx context.Context, args ...interface{}) {
	p.Log(ctx, TraceLevel, args...)
}

// Debug 记录DEBUG级别日志
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Debug(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

// Print 记录DEBUG级别日志（Print别名）
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Print(ctx context.Context, args ...interface{}) {
	p.Log(ctx, DebugLevel, args...)
}

// Info 记录INFO级别日志
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Info(ctx context.Context, args ...interface{}) {
	p.Log(ctx, InfoLevel, args...)
}

// Warn 记录WARN级别日志
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Warn(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

// Warning 记录WARN级别日志（Warning别名）
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Warning(ctx context.Context, args ...interface{}) {
	p.Log(ctx, WarnLevel, args...)
}

// Error 记录ERROR级别日志
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Error(ctx context.Context, args ...interface{}) {
	p.Log(ctx, ErrorLevel, args...)
}

// Panic 记录PANIC级别日志（会触发panic）
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Panic(ctx context.Context, args ...interface{}) {
	p.Log(ctx, PanicLevel, args...)
}

// Fatal 记录FATAL级别日志（会触发os.Exit(1)）
// 参数 ctx: 上下文
// 参数 args: 日志内容参数
func (p *LoggerWithCtx) Fatal(ctx context.Context, args ...interface{}) {
	p.Log(ctx, FatalLevel, args...)
}

func (p *LoggerWithCtx) Tracef(ctx context.Context, format string, args ...interface{}) {
	p.Logf(ctx, TraceLevel, format, args...)
}

// Printf 记录格式化DEBUG级别日志（Printf别名）
// 参数 ctx: 上下文
// 参数 format: 格式化字符串
// 参数 args: 格式化参数
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

// ParsingAndEscaping 设置是否禁用解析和转义
// 参数 disable: true表示禁用
// 返回: 当前LoggerWithCtx实例（支持链式调用）
func (p *LoggerWithCtx) ParsingAndEscaping(disable bool) *LoggerWithCtx {
	p.ParsingAndEscaping(disable)
	return p
}

func (p *LoggerWithCtx) Caller(disable bool) *LoggerWithCtx {
	p.Caller(disable)
	return p
}
