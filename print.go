package log

// Trace logs a message at Trace level (most verbose).
func Trace(args ...interface{}) {
	std.Trace(args...)
}

// Log logs a message at the specified level.
func Log(level Level, args ...interface{}) {
	std.Log(level, args...)
}

// Info logs a message at Info level.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Warn logs a message at Warn level.
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Error logs a message at Error level.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Panic logs a message at Panic level, then panics.
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal logs a message at Fatal level, then calls os.Exit(1).
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Logf logs a formatted message at the specified level.
func Logf(level Level, format string, args ...interface{}) {
	std.Logf(level, format, args...)
}

// Infof logs a formatted message at Info level.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warnf logs a formatted message at Warn level.
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Errorf logs a formatted message at Error level.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Panicf logs a formatted message at Panic level, then panics.
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Fatalf logs a formatted message at Fatal level, then calls os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Debug logs a message at Debug level.
//
// The message is only output when the global log level is set to DebugLevel or lower.
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Debugf logs a formatted message at Debug level.
//
// The message is only output when the global log level is set to DebugLevel or lower.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// StartMsg outputs the application startup message.
func StartMsg() {
	std.StartMsg()
}
