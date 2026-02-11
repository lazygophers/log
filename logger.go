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

// WriteSyncerWrapper implements zapcore.WriteSyncer interface
type WriteSyncerWrapper struct {
	writer io.Writer
}

// Write implements io.Writer interface
func (w *WriteSyncerWrapper) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

// Sync implements zapcore.WriteSyncer interface
func (w *WriteSyncerWrapper) Sync() error {
	// Call Sync if writer implements it
	if syncer, ok := w.writer.(interface{ Sync() error }); ok {
		return syncer.Sync()
	}
	// Return nil for standard writers
	return nil
}

// wrapWriter wraps io.Writer as zapcore.WriteSyncer
func wrapWriter(w io.Writer) zapcore.WriteSyncer {
	// Return directly if already WriteSyncer
	if ws, ok := w.(zapcore.WriteSyncer); ok {
		return ws
	}
	// Otherwise wrap it
	return &WriteSyncerWrapper{writer: w}
}

// Logger is the core logging structure
type Logger struct {
	level       Level
	out         zapcore.WriteSyncer
	Format      Format
	callerDepth int
	PrefixMsg   []byte
	SuffixMsg   []byte

	// Performance optimization fields
	enableCaller bool
	enableTrace  bool
}

// newLogger creates a new Logger instance with default values
func newLogger() *Logger {
	var out io.Writer = os.Stdout

	// Use hourly rotating file in release mode if path specified
	if ReleaseLogPath != "" {
		out = GetOutputWriterHourly(ReleaseLogPath)
	}

	logger := &Logger{
		level: DebugLevel,
		out:   wrapWriter(out),
		Format: &Formatter{
			DisableParsingAndEscaping: true,
		},
		callerDepth:  3,
		enableCaller: true,
		enableTrace:  true,
	}

	return logger
}

// SetCallerDepth sets the caller stack depth
func (p *Logger) SetCallerDepth(callerDepth int) *Logger {
	p.callerDepth = callerDepth
	return p
}

// SetPrefixMsg sets the log message prefix
func (p *Logger) SetPrefixMsg(prefixMsg string) *Logger {
	p.PrefixMsg = []byte(prefixMsg)
	return p
}

// AppendPrefixMsg appends to the log message prefix
func (p *Logger) AppendPrefixMsg(prefixMsg string) *Logger {
	p.PrefixMsg = []byte(string(p.PrefixMsg) + prefixMsg)
	return p
}

// SetSuffixMsg sets the log message suffix
func (p *Logger) SetSuffixMsg(suffixMsg string) *Logger {
	p.SuffixMsg = []byte(suffixMsg)
	return p
}

// AppendSuffixMsg appends to the log message suffix
func (p *Logger) AppendSuffixMsg(suffixMsg string) *Logger {
	p.SuffixMsg = []byte(string(p.SuffixMsg) + suffixMsg)
	return p
}

// EnableCaller controls caller information
func (p *Logger) EnableCaller(enable bool) *Logger {
	p.enableCaller = enable
	return p
}

// EnableTrace controls trace information
func (p *Logger) EnableTrace(enable bool) *Logger {
	p.enableTrace = enable
	return p
}

// Clone creates a deep copy of current Logger
func (p *Logger) Clone() *Logger {
	l := Logger{
		level:        p.level,
		out:          p.out,
		callerDepth:  p.callerDepth,
		PrefixMsg:    p.PrefixMsg,
		SuffixMsg:    p.SuffixMsg,
		enableCaller: p.enableCaller,
		enableTrace:  p.enableTrace,
	}

	switch f := p.Format.(type) {
	case FormatFull:
		l.Format = f.Clone()
	default:
		l.Format = f
	}

	return &l
}

// SetLevel sets the logging level
func (p *Logger) SetLevel(level Level) *Logger {
	p.level = level
	return p
}

// Level returns the current logging level
func (p *Logger) Level() Level {
	return p.level
}

// SetOutput sets the log output targets
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

// Log records a log with specified level
func (p *Logger) Log(level Level, args ...interface{}) {
	if !p.levelEnabled(level) {
		return
	}
	p.log(level, fmt.Sprint(args...))
}

// Logf records a formatted log
func (p *Logger) Logf(level Level, format string, args ...interface{}) {
	if !p.levelEnabled(level) {
		return
	}
	p.log(level, fmt.Sprintf(format, args...))
}

// populateEntry sets basic fields on the log entry
//
//go:inline
func (p *Logger) populateEntry(entry *Entry, level Level, msg string) {
	entry.Level = level
	entry.Message = msg
	entry.Time = time.Now()
}

// fillTraceInfo conditionally sets trace information
//
//go:inline
func (p *Logger) fillTraceInfo(entry *Entry) {
	if p.enableTrace {
		entry.Gid = goid.Get()
		entry.TraceId = getTrace(entry.Gid)
	}
}

// fillCallerInfo conditionally sets caller information
//
//go:inline
func (p *Logger) fillCallerInfo(entry *Entry) {
	if !p.enableCaller {
		return
	}

	var pc uintptr
	var ok bool
	pc, entry.File, entry.CallerLine, ok = runtime.Caller(p.callerDepth)
	if ok && pc != 0 {
		if fn := runtime.FuncForPC(pc); fn != nil {
			entry.CallerName = fn.Name()
			entry.CallerDir, entry.CallerFunc = SplitPackageName(entry.CallerName)
		}
	}
}

// fillPrefixSuffix sets prefix and suffix messages
//
//go:inline
func (p *Logger) fillPrefixSuffix(entry *Entry) {
	if len(p.PrefixMsg) > 0 {
		entry.PrefixMsg = p.PrefixMsg
	}
	if len(p.SuffixMsg) > 0 {
		entry.SuffixMsg = p.SuffixMsg
	}
}

// log is the internal core logging function
//
//go:noinline
func (p *Logger) log(level Level, msg string) {
	entry := getEntry()

	p.populateEntry(entry, level, msg)
	p.fillTraceInfo(entry)
	p.fillCallerInfo(entry)
	p.fillPrefixSuffix(entry)

	// Format and write
	formatted := p.Format.Format(entry)
	p.write(level, formatted)

	putEntry(entry)
}

// write writes formatted log bytes to output
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

// levelEnabled checks if the level should be logged
func (p *Logger) levelEnabled(level Level) bool {
	return p.level >= level
}

// Trace logs at TRACE level
func (p *Logger) Trace(args ...interface{}) {
	if !p.levelEnabled(TraceLevel) {
		return
	}
	p.log(TraceLevel, fmt.Sprint(args...))
}

// Debug logs at DEBUG level
func (p *Logger) Debug(args ...interface{}) {
	if !p.levelEnabled(DebugLevel) {
		return
	}
	p.log(DebugLevel, fmt.Sprint(args...))
}

// Print logs at DEBUG level (alias for Debug)
func (p *Logger) Print(args ...interface{}) {
	if !p.levelEnabled(DebugLevel) {
		return
	}
	p.log(DebugLevel, fmt.Sprint(args...))
}

// Info logs at INFO level
func (p *Logger) Info(args ...interface{}) {
	if !p.levelEnabled(InfoLevel) {
		return
	}
	p.log(InfoLevel, fmt.Sprint(args...))
}

// Warn logs at WARN level
func (p *Logger) Warn(args ...interface{}) {
	if !p.levelEnabled(WarnLevel) {
		return
	}
	p.log(WarnLevel, fmt.Sprint(args...))
}

// Warning is an alias for Warn
func (p *Logger) Warning(args ...interface{}) {
	if !p.levelEnabled(WarnLevel) {
		return
	}
	p.log(WarnLevel, fmt.Sprint(args...))
}

// Error logs at ERROR level
func (p *Logger) Error(args ...interface{}) {
	if !p.levelEnabled(ErrorLevel) {
		return
	}
	p.log(ErrorLevel, fmt.Sprint(args...))
}

// Panic logs at PANIC level and panics
func (p *Logger) Panic(args ...interface{}) {
	if !p.levelEnabled(PanicLevel) {
		return
	}
	p.log(PanicLevel, fmt.Sprint(args...))
}

// Fatal logs at FATAL level and exits
func (p *Logger) Fatal(args ...interface{}) {
	if !p.levelEnabled(FatalLevel) {
		return
	}
	p.log(FatalLevel, fmt.Sprint(args...))
}

// Tracef logs formatted TRACE level message
func (p *Logger) Tracef(format string, args ...interface{}) {
	if !p.levelEnabled(TraceLevel) {
		return
	}
	p.log(TraceLevel, fmt.Sprintf(format, args...))
}

// Printf logs formatted DEBUG level message (alias for Debugf)
func (p *Logger) Printf(format string, args ...interface{}) {
	if !p.levelEnabled(DebugLevel) {
		return
	}
	p.log(DebugLevel, fmt.Sprintf(format, args...))
}

// Debugf logs formatted DEBUG level message
func (p *Logger) Debugf(format string, args ...interface{}) {
	if !p.levelEnabled(DebugLevel) {
		return
	}
	p.log(DebugLevel, fmt.Sprintf(format, args...))
}

// Infof logs formatted INFO level message
func (p *Logger) Infof(format string, args ...interface{}) {
	if !p.levelEnabled(InfoLevel) {
		return
	}
	p.log(InfoLevel, fmt.Sprintf(format, args...))
}

// Warnf logs formatted WARN level message
func (p *Logger) Warnf(format string, args ...interface{}) {
	if !p.levelEnabled(WarnLevel) {
		return
	}
	p.log(WarnLevel, fmt.Sprintf(format, args...))
}

// Warningf is an alias for Warnf
func (p *Logger) Warningf(format string, args ...interface{}) {
	if !p.levelEnabled(WarnLevel) {
		return
	}
	p.log(WarnLevel, fmt.Sprintf(format, args...))
}

// Errorf logs formatted ERROR level message
func (p *Logger) Errorf(format string, args ...interface{}) {
	if !p.levelEnabled(ErrorLevel) {
		return
	}
	p.log(ErrorLevel, fmt.Sprintf(format, args...))
}

// Fatalf logs formatted FATAL level message and exits
func (p *Logger) Fatalf(format string, args ...interface{}) {
	if !p.levelEnabled(FatalLevel) {
		return
	}
	p.log(FatalLevel, fmt.Sprintf(format, args...))
}

// Panicf logs formatted PANIC level message and panics
func (p *Logger) Panicf(format string, args ...interface{}) {
	if !p.levelEnabled(PanicLevel) {
		return
	}
	p.log(PanicLevel, fmt.Sprintf(format, args...))
}

// Sync flushes buffered logs to disk
func (p *Logger) Sync() {
	_ = p.out.Sync()
}

// ParsingAndEscaping controls log content parsing and escaping
func (p *Logger) ParsingAndEscaping(disable bool) *Logger {
	switch f := p.Format.(type) {
	case FormatFull:
		f.ParsingAndEscaping(disable)
	default:
		Panicf("%v is not interface log.FormatFull", f)
	}
	return p
}

// Caller controls caller information in logs
func (p *Logger) Caller(disable bool) *Logger {
	switch f := p.Format.(type) {
	case FormatFull:
		f.Caller(disable)
	default:
		Panicf("%v is not interface log.FormatFull", f)
	}
	return p
}

// StartMsg logs a new log start message
func (p *Logger) StartMsg() {
	p.Infof("========== start new log ==========")
}
