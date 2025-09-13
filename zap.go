package log

import (
	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
)

// ZapHook implements zapcore.WriteSyncer interface to redirect zap logs to our logging system
type ZapHook struct {
	logger *Logger
}

// NewZapHook creates a new ZapHook instance
func NewZapHook(log *Logger) *ZapHook {
	return &ZapHook{logger: log}
}

// Write implements io.Writer interface
func (zh *ZapHook) Write(p []byte) (n int, err error) {
	// Write directly to underlying logger
	return zh.logger.out.Write(p)
}

// Sync implements zapcore.WriteSyncer interface
func (zh *ZapHook) Sync() error {
	zh.logger.Sync()
	return nil
}

// createZapHook creates a zap core hook to adapt zap log entries to current logging system
func createZapHook(log *Logger) func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		// Get Entry object from pool to reduce memory allocation
		logEntry := entryPool.Get().(*Entry)
		// Use defer to ensure Entry object is reset and returned to pool
		defer func() {
			logEntry.Reset()
			entryPool.Put(logEntry)
		}()

		// Fill basic information of log entry
		logEntry.Gid = goid.Get()                 // Get current goroutine ID
		logEntry.TraceId = getTrace(logEntry.Gid) // Get trace ID
		logEntry.Time = entry.Time                // Set log time
		logEntry.Message = entry.Message          // Set log message
		logEntry.SuffixMsg = log.SuffixMsg        // Set log suffix
		logEntry.PrefixMsg = log.PrefixMsg        // Set log prefix

		// Fill caller information
		logEntry.File = entry.Caller.File           // File path
		logEntry.CallerLine = entry.Caller.Line     // Line number
		logEntry.CallerName = entry.Caller.Function // Function name

		// Parse package name and function name
		logEntry.CallerDir, logEntry.CallerFunc = SplitPackageName(entry.Caller.Function)

		// Convert zap log level to internal log level
		switch entry.Level {
		case zapcore.DebugLevel:
			logEntry.Level = DebugLevel
		case zapcore.InfoLevel:
			logEntry.Level = InfoLevel
		case zapcore.WarnLevel:
			logEntry.Level = WarnLevel
		case zapcore.ErrorLevel:
			logEntry.Level = ErrorLevel
		case zapcore.DPanicLevel:
			logEntry.Level = PanicLevel
		case zapcore.PanicLevel:
			logEntry.Level = PanicLevel
		case zapcore.FatalLevel:
			logEntry.Level = FatalLevel
		default:
			// Default to ErrorLevel for unknown log levels
			logEntry.Level = ErrorLevel
		}

		// Format log entry and write
		log.write(logEntry.Level, log.Format.Format(logEntry))

		return nil
	}
}
