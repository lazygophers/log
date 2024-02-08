//go:build release && discard

package log

import "io"

func init() {
	SetOutput(io.Discard)
}

func Debug(args ...interface{}) {
}

func Debugf(format string, args ...interface{}) {
}
