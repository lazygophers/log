package log

import (
	"os"
)

var (
	// std is the standard logger instance
	std = newLogger()

	// pid stores current process ID
	pid = os.Getpid()
)

// Pid returns current process ID
func Pid() int {
	return pid
}

// New creates and returns a new Logger instance
func New() *Logger {
	return newLogger()
}

// SetLevel sets the standard logger's log level
func SetLevel(level Level) *Logger {
	return std.SetLevel(level)
}

// GetLevel returns the standard logger's current log level
func GetLevel() Level {
	return std.Level()
}

// Sync flushes all buffered log entries to their output destinations
func Sync() {
	std.Sync()
}

// Clone creates a deep copy of the standard logger
func Clone() *Logger {
	return std.Clone()
}

// CloneToCtx creates a context-aware logger copy
func CloneToCtx() *LoggerWithCtx {
	return std.CloneToCtx()
}

// SetCallerDepth sets the caller stack depth
func SetCallerDepth(callerDepth int) *Logger {
	return std.SetCallerDepth(callerDepth)
}

// SetPrefixMsg sets the log message prefix
func SetPrefixMsg(prefixMsg string) *Logger {
	return std.SetPrefixMsg(prefixMsg)
}

// AppendPrefixMsg appends to the log message prefix
func AppendPrefixMsg(prefixMsg string) *Logger {
	return std.AppendPrefixMsg(prefixMsg)
}

// SetSuffixMsg sets the log message suffix
func SetSuffixMsg(suffixMsg string) *Logger {
	return std.SetSuffixMsg(suffixMsg)
}

// AppendSuffixMsg appends to the log message suffix
func AppendSuffixMsg(suffixMsg string) *Logger {
	return std.AppendSuffixMsg(suffixMsg)
}

// ParsingAndEscaping controls JSON parsing and character escaping of log messages
func ParsingAndEscaping(disable bool) *Logger {
	return std.ParsingAndEscaping(disable)
}

// Caller controls whether to include caller information in log entries
func Caller(disable bool) *Logger {
	return std.Caller(disable)
}
