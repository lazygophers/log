//go:build release && !discard

package log

import (
	"os"
	"path/filepath"
)

func init() {
	SetOutput((GetOutputWriterHourly(filepath.Join(os.TempDir(), "lazygophers", "log") + "/")))
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
