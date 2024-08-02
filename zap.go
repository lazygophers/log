package log

import (
	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
)

func NewZapHook(log *Logger) func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		logEntry := entryPool.Get().(*Entry)
		defer func() {
			logEntry.Reset()
			entryPool.Put(logEntry)
		}()
		logEntry.Gid = goid.Get()
		logEntry.TraceId = getTrace(logEntry.Gid)
		logEntry.Time = entry.Time
		logEntry.Message = entry.Message

		//if entry.Stack != "" {
		//	logEntry.Message += " "
		//	logEntry.Message += entry.Stack
		//}

		logEntry.SuffixMsg = log.SuffixMsg
		logEntry.PrefixMsg = log.PrefixMsg

		logEntry.File = entry.Caller.File
		logEntry.CallerLine = entry.Caller.Line
		logEntry.CallerName = entry.Caller.Function

		logEntry.CallerDir, logEntry.CallerFunc = SplitPackageName(entry.Caller.Function)

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
			logEntry.Level = ErrorLevel
		}

		log.write(logEntry.Level, log.Format.Format(logEntry))

		return nil
	}
}
