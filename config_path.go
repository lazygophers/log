package log

import (
	"os"
	"path/filepath"
)

// ReleaseLogPath is the default log file path.
var ReleaseLogPath = filepath.Join(os.TempDir(), "lazygophers", "log")
