//go:build debug && !discard

package log

import "fmt"

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debug(fmt.Sprintf(format, args...))
}
