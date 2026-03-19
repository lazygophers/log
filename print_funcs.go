//go:build !discard && (debug || canary || (!debug && !release && !canary))

package log

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
