package log

import (
	"io"
	"os"
	"path/filepath"
	"sync"
)

// SetOutput configures output targets for the standard logger
func SetOutput(writes ...io.Writer) *Logger {
	return std.SetOutput(writes...)
}

// GetOutputWriter creates a basic file log writer
func GetOutputWriter(filename string) io.Writer {
	// Ensure log file directory exists
	ensureDir(filepath.Dir(filename))

	// Open file for writing
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) // #nosec G304
	if err != nil {
		// Panic if creating log writer fails (critical functionality)
		std.Panicf("创建日志文件写入器 %s 失败: %v", filename, err)
	}
	return file
}

// ensureDir ensures the specified directory exists
func ensureDir(dir string) {
	// Create directory if not current dir and path doesn't exist as directory
	if dir != "." && !isDir(dir) {
		// Create directory with full permissions, log error if fails
		if err := os.MkdirAll(dir, 0750); err != nil {
			std.Errorf("无法创建日志目录 %s: %v", dir, err)
		}
	}
}

// isDir checks if the given path is a valid directory
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Output interface defines log writer behavior
type Output interface {
	io.Writer
}

var (
	// rotatorInstances stores created rotator instances to avoid duplicates
	rotatorInstances = make(map[string]*HourlyRotator)

	// rotatorMutex protects concurrent access to rotatorInstances
	rotatorMutex = &sync.Mutex{}
)

// GetOutputWriterHourly creates an hourly rotating log writer with auto-cleanup
func GetOutputWriterHourly(filename string) Writer {
	rotatorMutex.Lock()
	defer rotatorMutex.Unlock()

	// Check if rotator instance already exists for this file
	if rotator, exists := rotatorInstances[filename]; exists {
		return rotator
	}

	// Create new rotator instance
	rotator := NewHourlyRotator(
		filename,
		1024*1024*8*100, // 800MB size limit
		12,              // Keep latest 12 files
	)

	// Store instance for future use
	rotatorInstances[filename] = rotator

	return rotator
}
