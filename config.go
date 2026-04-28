package log

import (
	"os"
	"path/filepath"
)

var (
	// ReleaseLogDir 如果是存储到文件，默认的文件存储目录
	ReleaseLogDir = filepath.Join(os.TempDir(), "lazygophers", "log")
	// DefaultMaxFiles 默认的文件保留的数量，如果超过的话，会自动清理
	DefaultMaxFiles = 48
	// DefaultMaxFileSize 默认单文件的大小，如果超过的话，会自动文件分割
	DefaultMaxFileSize = int64(1024 * 1024 * 8 * 100) // 800MB
)
