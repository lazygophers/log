// Package log 提供灵活可配置的日志记录功能，支持多级别日志输出、自定义格式和输出目标。
package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
)

// Logger 是日志记录器核心结构，负责日志的输出控制和格式配置
type Logger struct {
	level Level

	out zapcore.WriteSyncer

	Format Format

	callerDepth int

	PrefixMsg []byte
	SuffixMsg []byte
}

// newLogger 创建一个新的 Logger 实例，并设置默认值。
// 默认日志级别为 DebugLevel，输出到 os.Stdout。
func newLogger() *Logger {
	return &Logger{
		level: DebugLevel,
		out:   os.Stdout,
		Format: &Formatter{
			DisableParsingAndEscaping: true,
		},
		callerDepth: 4,
	}
}

// SetCallerDepth 设置日志调用栈深度
// callerDepth: 调用栈深度（从当前函数开始计算）
// 返回: Logger指针用于链式调用
func (p *Logger) SetCallerDepth(callerDepth int) *Logger {
	p.callerDepth = callerDepth
	return p
}

// SetPrefixMsg 设置日志消息前缀
// prefixMsg: 要添加的前缀字符串
// 返回: Logger指针用于链式调用
func (p *Logger) SetPrefixMsg(prefixMsg string) *Logger {
	p.PrefixMsg = []byte(prefixMsg)
	return p
}

// AppendPrefixMsg 追加日志消息前缀
// prefixMsg: 要追加的前缀字符串
// 返回: Logger指针用于链式调用
func (p *Logger) AppendPrefixMsg(prefixMsg string) *Logger {
	p.PrefixMsg = []byte(string(p.PrefixMsg) + prefixMsg)
	return p
}

// SetSuffixMsg 设置日志消息后缀
// suffixMsg: 要添加的后缀字符串
// 返回: Logger指针用于链式调用
func (p *Logger) SetSuffixMsg(suffixMsg string) *Logger {
	p.SuffixMsg = []byte(suffixMsg)
	return p
}

// AppendSuffixMsg 追加日志消息后缀
// suffixMsg: 要追加的后缀字符串
// 返回: Logger指针用于链式调用
func (p *Logger) AppendSuffixMsg(suffixMsg string) *Logger {
	p.SuffixMsg = []byte(string(p.SuffixMsg) + suffixMsg)
	return p
}

// Clone 创建当前Logger的深度拷贝
// 返回: 新的Logger实例
func (p *Logger) Clone() *Logger {
	l := Logger{
		level:       p.level,
		out:         p.out,
		callerDepth: p.callerDepth,
		PrefixMsg:   p.PrefixMsg,
		SuffixMsg:   p.SuffixMsg,
	}

	switch f := p.Format.(type) {
	case FormatFull:
		l.Format = f.Clone()
	default:
		l.Format = f
	}

	return &l
}

// SetLevel 设置日志级别
// level: 日志级别（TraceLevel/DebugLevel/InfoLevel等）
// 返回: Logger指针用于链式调用
func (p *Logger) SetLevel(level Level) *Logger {
	p.level = level
	return p
}

// Level 获取当前日志级别
// 返回: 当前日志级别
func (p *Logger) Level() Level {
	return p.level
}

// SetOutput 设置日志输出目标
// writes: 一个或多个io.Writer输出目标
// 返回: Logger指针用于链式调用
func (p *Logger) SetOutput(writes ...io.Writer) *Logger {
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

// Log 记录指定级别的日志
// level: 日志级别
// args: 日志内容参数
func (p *Logger) Log(level Level, args ...interface{}) {
	p.log(level, fmt.Sprint(args...))
}

// Logf 记录格式化日志
// level: 日志级别
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Logf(level Level, format string, args ...interface{}) {
	p.log(level, fmt.Sprintf(format, args...))
}

// log 是内部核心日志记录函数。
// 它首先检查指定的日志级别是否启用，如果未启用则直接返回。
// 然后，它从对象池获取一个 Entry 对象，并填充日志信息，
// 包括Gid、时间、级别、消息、前后缀、TraceId以及调用者信息。
// 最后，它使用指定的格式化器格式化 Entry，并将结果写入输出。
// 完成后，Entry 对象被重置并放回对象池。
func (p *Logger) log(level Level, msg string) {
	if !p.levelEnabled(level) {
		return
	}

	entry := entryPool.Get().(*Entry)
	defer func() {
		entry.Reset()
		entryPool.Put(entry)
	}()

	entry.Gid = goid.Get()
	entry.Time = time.Now()
	entry.Level = level
	entry.Message = msg
	entry.SuffixMsg = p.SuffixMsg
	entry.PrefixMsg = p.PrefixMsg
	entry.TraceId = getTrace(entry.Gid)

	var pc uintptr
	var ok bool
	pc, entry.File, entry.CallerLine, ok = runtime.Caller(p.callerDepth)
	if ok {
		entry.CallerName = runtime.FuncForPC(pc).Name()
	}

	entry.CallerDir, entry.CallerFunc = SplitPackageName(entry.CallerName)

	p.write(level, p.Format.Format(entry))
}

// write 将格式化后的日志字节流写入输出。
// 如果日志级别是 PanicLevel，它会在写入后调用 Sync() 并触发 panic。
// 如果日志级别是 FatalLevel，它会在写入后调用 Sync() 并以状态码 -1 退出程序。
func (p *Logger) write(level Level, buf []byte) {
	_, _ = p.out.Write(buf)

	if level == PanicLevel {
		p.Sync()
		panic(buf)
	} else if level == FatalLevel {
		p.Sync()
		os.Exit(-1)
	}
}

// levelEnabled 检查当前日志级别是否允许输出指定级别的日志。
func (p *Logger) levelEnabled(level Level) bool {
	return p.level >= level
}

// Trace 记录TRACE级别日志
// args: 日志内容参数
func (p *Logger) Trace(args ...interface{}) {
	p.Log(TraceLevel, args...)
}

// Debug 记录DEBUG级别日志
// args: 日志内容参数
func (p *Logger) Debug(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

// Print 记录DEBUG级别日志（Print的别名）
// args: 日志内容参数
func (p *Logger) Print(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

// Info 记录INFO级别日志
// args: 日志内容参数
func (p *Logger) Info(args ...interface{}) {
	p.Log(InfoLevel, args...)
}

// Warn 记录WARN级别日志。
// args: 日志内容参数
func (p *Logger) Warn(args ...interface{}) {
	p.Log(WarnLevel, args...)
}

// Warning 是 Warn 的别名，记录WARN级别日志。
// args: 日志内容参数
func (p *Logger) Warning(args ...interface{}) {
	p.Log(WarnLevel, args...)
}

// Error 记录ERROR级别日志。
// args: 日志内容参数
func (p *Logger) Error(args ...interface{}) {
	p.Log(ErrorLevel, args...)
}

// Panic 记录PANIC级别日志，并触发panic。
// args: 日志内容参数
func (p *Logger) Panic(args ...interface{}) {
	p.Log(PanicLevel, args...)
}

// Fatal 记录FATAL级别日志，并终止程序。
// args: 日志内容参数
func (p *Logger) Fatal(args ...interface{}) {
	p.Log(FatalLevel, args...)
}

// Tracef 记录格式化的TRACE级别日志。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Tracef(format string, args ...interface{}) {
	p.Logf(TraceLevel, format, args...)
}

// Printf 记录格式化的DEBUG级别日志（Printf是Debugf的别名）。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Printf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

// Debugf 记录格式化的DEBUG级别日志。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Debugf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

// Infof 记录格式化的INFO级别日志。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Infof(format string, args ...interface{}) {
	p.Logf(InfoLevel, format, args...)
}

// Warnf 记录格式化的WARN级别日志。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Warnf(format string, args ...interface{}) {
	p.Logf(WarnLevel, format, args...)
}

// Warningf 是 Warnf 的别名，记录格式化的WARN级别日志。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Warningf(format string, args ...interface{}) {
	p.Logf(WarnLevel, format, args...)
}

// Errorf 记录格式化的ERROR级别日志。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Errorf(format string, args ...interface{}) {
	p.Logf(ErrorLevel, format, args...)
}

// Fatalf 记录格式化的FATAL级别日志，并终止程序。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Fatalf(format string, args ...interface{}) {
	p.Logf(FatalLevel, format, args...)
}

// Panicf 记录格式化的PANIC级别日志，并触发panic。
// format: 格式化字符串
// args: 格式化参数
func (p *Logger) Panicf(format string, args ...interface{}) {
	p.Logf(PanicLevel, format, args...)
}

// Sync 将缓冲区的日志刷新到磁盘。
func (p *Logger) Sync() {
	_ = p.out.Sync()
}

// ParsingAndEscaping 控制是否禁用日志内容的解析和转义。
// 这只对实现了 FormatFull 接口的格式化器有效。
// disable: true表示禁用，false表示启用
// 返回: Logger指针用于链式调用
func (p *Logger) ParsingAndEscaping(disable bool) *Logger {
	switch f := p.Format.(type) {
	case FormatFull:
		f.ParsingAndEscaping(disable)
	default:
		Panicf("%v is not interface log.FormatFull", f)
	}
	return p
}

// Caller 控制是否在日志中禁用调用者信息（文件名、行号、函数名）。
// 这只对实现了 FormatFull 接口的格式化器有效。
// disable: true表示禁用，false表示启用
// 返回: Logger指针用于链式调用
func (p *Logger) Caller(disable bool) *Logger {
	switch f := p.Format.(type) {
	case FormatFull:
		f.Caller(disable)
	default:
		Panicf("%v is not interface log.FormatFull", f)
	}
	return p
}

// StartMsg 记录一条表示新日志开始的INFO级别消息。
func (p *Logger) StartMsg() {
	Infof("========== start new log ==========")
}
