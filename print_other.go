//go:build !debug && !release && !discard

package log

import "fmt"

func Debug(args ...interface{}) {
	if std.level >= DebugLevel {
		std.Debug(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if std.level >= DebugLevel {
		std.Debug(fmt.Sprintf(format, args...))
	}
}
