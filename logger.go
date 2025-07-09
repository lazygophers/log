package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
)

// Logger 管理日志记录器的配置和行为
type Logger struct {
	// level 控制输出日志的详细程度
	level Level

	// out 日志输出目标（支持多路复用）
	out zapcore.WriteSyncer

	// Format 定义日志输出格式
	Format Format

	// callerDepth 用于确定日志中显示的调用位置
	callerDepth int

	// PrefixMsg 日志消息前缀
	PrefixMsg []byte
	// SuffixMsg 日志消息后缀
	SuffixMsg []byte
}

// newLogger 创建默认日志记录器实例
func newLogger() *Logger {
	return &Logger{
		level:       DebugLevel,
		out:         os.Stdout,
		Format:      &Formatter{},
		callerDepth: 4,
	}
}

// SetCallerDepth 设置调用者深度
// 参数 callerDepth: 整数，表示调用栈深度
// 返回: 当前Logger实例（支持链式调用）
func (p *Logger) SetCallerDepth(callerDepth int) *Logger {
	p.callerDepth = callerDepth
	return p
}

// SetPrefixMsg 设置消息前缀
// 参数 prefixMsg: 字符串，将被转换为字节切片
// 返回: 当前Logger实例（支持链式调用）
func (p *Logger) SetPrefixMsg(prefixMsg string) *Logger {
	p.PrefixMsg = []byte(prefixMsg)
	return p
}

// AppendPrefixMsg 追加消息前缀
// 参数 prefixMsg: 字符串，将被追加到现有前缀
// 返回: 当前Logger实例（支持链式调用）
func (p *Logger) AppendPrefixMsg(prefixMsg string) *Logger {
	p.PrefixMsg = []byte(string(p.PrefixMsg) + prefixMsg)
	return p
}

// SetSuffixMsg 设置消息后缀
// 参数 suffixMsg: 字符串，将被转换为字节切片
// 返回: 当前Logger实例（支持链式调用）
func (p *Logger) SetSuffixMsg(suffixMsg string) *Logger {
	p.SuffixMsg = []byte(suffixMsg)
	return p
}

// AppendSuffixMsg 追加消息后缀
// 参数 suffixMsg: 字符串，将被追加到现有后缀
// 返回: 当前Logger实例（支持链式调用）
func (p *Logger) AppendSuffixMsg(suffixMsg string) *Logger {
	p.SuffixMsg = []byte(string(p.SuffixMsg) + suffixMsg)
	return p
}

// Clone 创建Logger的深拷贝实例
// 返回: 新的Logger指针
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
// 参数 level: 目标日志级别
// 返回: 当前Logger实例（支持链式调用）
func (p *Logger) SetLevel(level Level) *Logger {
	p.level = level
	return p
}

// Level 获取当前日志级别
// 返回: 当前配置的日志级别
func (p *Logger) Level() Level {
	return p.level
}

// SetOutput 设置日志输出目标
// 参数 writes: 一个或多个io.Writer接口实现
// 返回: 当前Logger实例（支持链式调用）
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

// Log 通用日志记录方法
// 参数 level: 日志级别
// 参数 args: 任意类型参数，将被转换为字符串
func (p *Logger) Log(level Level, args ...interface{}) {
	p.log(level, fmt.Sprint(args...))
}

// Logf 格式化日志记录方法
// 参数 level: 日志级别
// 参数 format: 格式化字符串
// 参数 args: 格式化参数
func (p *Logger) Logf(level Level, format string, args ...interface{}) {
	p.log(level, fmt.Sprintf(format, args...))
}

var entryPool = sync.Pool{
	New: func() any {
		return NewEntry()
	},
}

// log 内部日志处理方法
// 参数 level: 日志级别
// 参数 msg: 已格式化的日志消息
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

// levelEnabled 检查指定级别是否启用
func (p *Logger) levelEnabled(level Level) bool {
	return p.level >= level
}

// Trace 记录TRACE级别日志
func (p *Logger) Trace(args ...interface{}) {
	p.Log(TraceLevel, args...)
}

// Debug 记录DEBUG级别日志
func (p *Logger) Debug(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

// Print 记录DEBUG级别日志（Print的别名）
func (p *Logger) Print(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

// Info 记录INFO级别日志
func (p *Logger) Info(args ...interface{}) {
	p.Log(InfoLevel, args...)
}

// Warn 记录WARN级别日志
func (p *Logger) Warn(args ...interface{}) {
	p.Log(WarnLevel, args...)
}

// Warning 记录WARN级别日志（Warn的别名）
func (p *Logger) Warning(args ...interface{}) {
	p.Log(WarnLevel, args...)
}

// Error 记录ERROR级别日志
func (p *Logger) Error(args ...interface{}) {
	p.Log(ErrorLevel, args...)
}

// Panic 记录PANIC级别日志并触发panic
func (p *Logger) Panic(args ...interface{}) {
	p.Log(PanicLevel, args...)
}

// Fatal 记录FATAL级别日志并退出程序
func (p *Logger) Fatal(args ...interface{}) {
	p.Log(FatalLevel, args...)
}

// Tracef 格式化记录TRACE级别日志
func (p *Logger) Tracef(format string, args ...interface{}) {
	p.Logf(TraceLevel, format, args...)
}

// Printf 格式化记录DEBUG级别日志
func (p *Logger) Printf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

// Debugf 格式化记录DEBUG级别日志
func (p *Logger) Debugf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

// Infof 格式化记录INFO级别日志
func (p *Logger) Infof(format string, args ...interface{}) {
	p.Logf(InfoLevel, format, args...)
}

// Warnf 格式化记录WARN级别日志
func (p *Logger) Warnf(format string, args ...interface{}) {
	p.Logf(WarnLevel, format, args...)
}

// Warningf 格式化记录WARN级别日志（Warnf的别名）
func (p *Logger) Warningf(format string, args ...interface{}) {
	p.Logf(WarnLevel, format, args...)
}

// Errorf 格式化记录ERROR级别日志
func (p *Logger) Errorf(format string, args ...interface{}) {
	p.Logf(ErrorLevel, format, args...)
}

// Fatalf 格式化记录FATAL级别日志并退出程序
func (p *Logger) Fatalf(format string, args ...interface{}) {
	p.Logf(FatalLevel, format, args...)
}

// Panicf 格式化记录PANIC级别日志并触发panic
func (p *Logger) Panicf(format string, args ...interface{}) {
	p.Logf(PanicLevel, format, args...)
}

func (p *Logger) Sync() {
	_ = p.out.Sync()
}

func (p *Logger) ParsingAndEscaping(disable bool) *Logger {
	switch f := p.Format.(type) {
	case FormatFull:
		f.ParsingAndEscaping(disable)
	default:
		Panicf("%v is not interface log.FormatFull", f)
	}
	return p
}

func (p *Logger) Caller(disable bool) *Logger {
	switch f := p.Format.(type) {
	case FormatFull:
		f.Caller(disable)
	default:
		Panicf("%v is not interface log.FormatFull", f)
	}
	return p
}

func (p *Logger) StartMsg() {
	Infof("========== start new log ==========")
}
