package log

import (
	"os"
	"path/filepath"
)

// ReleaseLogPath 默认日志文件路径
var ReleaseLogPath = filepath.Join(os.TempDir(), "lazygophers", "log")
