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

type logLine struct {
}

type Logger struct {
	level Level

	outList []zapcore.WriteSyncer

	Format Format

	callerDepth int

	PrefixMsg []byte
	SuffixMsg []byte
}

func newLogger() *Logger {
	return &Logger{
		level:       DebugLevel,
		outList:     []zapcore.WriteSyncer{os.Stdout},
		Format:      &Formatter{},
		callerDepth: 4,
	}
}

func (p *Logger) SetCallerDepth(callerDepth int) *Logger {
	p.callerDepth = callerDepth
	return p
}

func (p *Logger) SetPrefixMsg(prefixMsg string) *Logger {
	p.PrefixMsg = []byte(prefixMsg)
	return p
}

func (p *Logger) SetSuffixMsg(suffixMsg string) *Logger {
	p.SuffixMsg = []byte(suffixMsg)
	return p
}

func (p *Logger) Clone() *Logger {
	l := Logger{
		level:       p.level,
		outList:     p.outList,
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

func (p *Logger) SetLevel(level Level) *Logger {
	p.level = level
	return p
}

func (p *Logger) Level() Level {
	return p.level
}

func (p *Logger) SetOutput(writes ...io.Writer) *Logger {
	var ws []zapcore.WriteSyncer
	for _, write := range writes {
		ws = append(ws, zapcore.AddSync(write))
	}
	p.outList = ws
	return p
}

func (p *Logger) AddOutput(writes ...io.Writer) *Logger {
	for _, write := range writes {
		p.outList = append(p.outList, zapcore.AddSync(write))
	}
	return p
}

func (p *Logger) Log(level Level, args ...interface{}) {
	p.log(level, fmt.Sprint(args...))
}

func (p *Logger) Logf(level Level, format string, args ...interface{}) {
	p.log(level, fmt.Sprintf(format, args...))
}

var entryPool = sync.Pool{
	New: func() any {
		return NewEntry()
	},
}

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
	for _, out := range p.outList {
		_, _ = out.Write(buf)
	}

	if level == PanicLevel {
		p.Sync()
		panic(buf)
	} else if level == FatalLevel {
		p.Sync()
		os.Exit(-1)
	}
}

func (p *Logger) levelEnabled(level Level) bool {
	return p.level >= level
}

func (p *Logger) Trace(args ...interface{}) {
	p.Log(TraceLevel, args...)
}

func (p *Logger) Debug(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

func (p *Logger) Print(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

func (p *Logger) Info(args ...interface{}) {
	p.Log(InfoLevel, args...)
}

func (p *Logger) Warn(args ...interface{}) {
	p.Log(WarnLevel, args...)
}

func (p *Logger) Error(args ...interface{}) {
	p.Log(ErrorLevel, args...)
}

func (p *Logger) Panic(args ...interface{}) {
	p.Log(PanicLevel, args...)
}

func (p *Logger) Fatal(args ...interface{}) {
	p.Log(FatalLevel, args...)
}

func (p *Logger) Tracef(format string, args ...interface{}) {
	p.Logf(TraceLevel, format, args...)
}

func (p *Logger) Printf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

func (p *Logger) Debugf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

func (p *Logger) Infof(format string, args ...interface{}) {
	p.Logf(InfoLevel, format, args...)
}

func (p *Logger) Warnf(format string, args ...interface{}) {
	p.Logf(WarnLevel, format, args...)
}

func (p *Logger) Errorf(format string, args ...interface{}) {
	p.Logf(ErrorLevel, format, args...)
}

func (p *Logger) Fatalf(format string, args ...interface{}) {
	p.Logf(FatalLevel, format, args...)
}

func (p *Logger) Panicf(format string, args ...interface{}) {
	p.Logf(PanicLevel, format, args...)
}

func (p *Logger) Sync() {
	for _, out := range p.outList {
		_ = out.Sync()
	}
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
