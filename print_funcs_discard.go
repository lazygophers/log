//go:build discard

package log

import "io"

// init sets the standard logger output to io.Discard during package initialization.
// io.Discard consumes and discards all written data, enabling zero-cost logging.
func init() {
	SetOutput(io.Discard)
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
