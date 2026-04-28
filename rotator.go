package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	cleanupHourlyRotatorOnce sync.Once
)

// HourlyRotator implements hourly log rotation
type HourlyRotator struct {
	mu          sync.Mutex
	logDir      string
	linkName    string
	currentFile *os.File
	currentHour string
	currentSize int64 // Tracked file size to avoid Stat() calls
	maxSize     int64 // Maximum file size in bytes
	maxFiles    int   // Maximum number of files to keep
}

// NewHourlyRotator creates a new hourly rotating log writer
func NewHourlyRotator(logDir string, maxSize int64, maxFiles int) *HourlyRotator {
	r := &HourlyRotator{
		logDir:   logDir,                               // This is now the directory path
		linkName: filepath.Join(logDir, "current.log"), // Link inside the directory
		maxSize:  maxSize,
		maxFiles: maxFiles,
	}

	cleanupHourlyRotatorOnce.Do(func() {
		go func() {
			ticker := time.NewTicker(10 * time.Minute)
			defer ticker.Stop()

			for _, rotator := range rotatorInstances {
				rotator.cleanupOldFiles()
			}
		}()
	})

	return r
}

// Write implements io.Writer interface
func (r *HourlyRotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := r.rotate(); err != nil {
		return 0, err
	}

	if r.currentFile == nil {
		return 0, fmt.Errorf("no current file")
	}

	n, err = r.currentFile.Write(p)
	r.currentSize += int64(n)
	return n, err
}

// rotate checks if rotation is needed and performs it
func (r *HourlyRotator) rotate() error {
	now := time.Now()
	hour := now.Format("2006010215")

	// todo 对于单小时超过日志文件大小的自动分片的支持
	// Check if rotation needed (hour change or file size limit exceeded)
	needRotate := r.currentHour != hour
	if r.currentFile != nil && !needRotate {
		if r.currentSize >= r.maxSize {
			needRotate = true
		}
	}

	if needRotate || r.currentFile == nil {
		return r.doRotate(hour)
	}

	return nil
}

// doRotate performs actual rotation operation
func (r *HourlyRotator) doRotate(hour string) error {
	// Close current file
	if r.currentFile != nil {
		_ = r.currentFile.Close()
	}

	// Ensure directory exists
	ensureDir(r.logDir)

	// Generate new filename (timestamp.log inside the directory)
	newFilename := filepath.Join(r.logDir, hour+".log")

	// Open new file
	file, err := os.OpenFile(newFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) // #nosec G304
	if err != nil {
		return err
	}

	r.currentFile = file
	r.currentHour = hour
	r.currentSize = 0
	if stat, err := file.Stat(); err == nil {
		r.currentSize = stat.Size()
	}

	// Update soft link
	r.updateLink(newFilename)

	return nil
}

// updateLink updates soft link pointing to the latest log file
func (r *HourlyRotator) updateLink(target string) {
	// Remove old link
	_ = os.Remove(r.linkName)

	// Create new link (ignore errors, soft link creation failure should not affect logging)
	_ = os.Symlink(filepath.Base(target), r.linkName)
}

// cleanupOldFiles cleans up expired log files
func (r *HourlyRotator) cleanupOldFiles() {
	// r.filename is now the directory path
	dir := r.logDir

	// Read all files in directory
	files, err := os.ReadDir(dir)
	if err != nil {
		// 找不到的情况，没有必要输出信息
		if os.IsNotExist(err) {
			return
		}

		fmt.Printf("Error reading log directory %s for cleanup: %v\n", dir, err)
		return
	}

	// Filter log files
	var logFiles []string
	for _, file := range files {
		if !file.IsDir() {
			// todo: 兼容对于单小时的多文件的情况
			name := file.Name()
			// Match format: YYYYMMDDHH + .log (e.g., "2026040114.log")
			// Length: 10 digits (YYYYMMDDHH) + 4 (".log") = 14
			if strings.HasSuffix(name, ".log") && len(name) == 14 {
				// Verify that the prefix (without .log) is all digits (timestamp)
				timestamp := strings.TrimSuffix(name, ".log")
				if len(timestamp) == 10 && isAllDigits(timestamp) {
					logFiles = append(logFiles, name)
				}
			}
		}
	}

	// Sort files by name descending (newest first)
	sort.Slice(logFiles, func(i, j int) bool {
		return logFiles[i] > logFiles[j]
	})

	// Delete old files exceeding retention count
	for i, filename := range logFiles {
		if i < r.maxFiles {
			continue
		}

		err = os.Remove(filepath.Join(dir, filename))
		if err != nil {
			fmt.Printf("Failed to delete old log file %v\n", err)
		}
	}
}

// Sync syncs current file content to disk
func (r *HourlyRotator) Sync() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentFile != nil {
		return r.currentFile.Sync()
	}

	return nil
}

// Close closes the rotator and cleans up resources
func (r *HourlyRotator) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentFile != nil {
		return r.currentFile.Close()
	}

	return nil
}

// isAllDigits checks if a string consists only of digits
func isAllDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
