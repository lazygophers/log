package log

import (
	"os"
)

var (
	std = newLogger()

	pid = os.Getpid()
)

func Pid() int {
	return pid
}

func New() *Logger {
	return newLogger()
}

func SetLevel(level Level) *Logger {
	return std.SetLevel(level)
}

func GetLevel() Level {
	return std.Level()
}

func Sync() {
	std.Sync()
}

func Clone() *Logger {
	return std.Clone()
}

// CloneToCtx creates a contextual logger from the standard logger
func CloneToCtx() *LoggerWithCtx {
	return std.CloneToCtx()
}

// SetCallerDepth sets the number of stack frames to skip when capturing caller info
func SetCallerDepth(callerDepth int) *Logger {
	return std.SetCallerDepth(callerDepth)
}

// SetPrefixMsg sets a fixed prefix for all log messages
func SetPrefixMsg(prefixMsg string) *Logger {
	return std.SetPrefixMsg(prefixMsg)
}

// AppendPrefixMsg appends to the current message prefix
func AppendPrefixMsg(prefixMsg string) *Logger {
	return std.AppendPrefixMsg(prefixMsg)
}

// SetSuffixMsg sets a fixed suffix for all log messages
func SetSuffixMsg(suffixMsg string) *Logger {
	return std.SetSuffixMsg(suffixMsg)
}

// AppendSuffixMsg appends to the current message suffix
func AppendSuffixMsg(suffixMsg string) *Logger {
	return std.AppendSuffixMsg(suffixMsg)
}

// ParsingAndEscaping controls automatic parsing and escaping of log fields
// When disabled, fields are logged as-is without modification
func ParsingAndEscaping(disable bool) *Logger {
	return std.ParsingAndEscaping(disable)
}

// Caller enables/disables caller information (file and line number) in logs
func Caller(disable bool) *Logger {
	return std.Caller(disable)
}
