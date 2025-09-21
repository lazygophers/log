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

// HourlyRotator implements hourly log rotation
type HourlyRotator struct {
	mu          sync.Mutex
	filename    string
	linkName    string
	currentFile *os.File
	currentHour string
	maxSize     int64 // Maximum file size in bytes
	maxFiles    int   // Maximum number of files to keep
	cleanupOnce sync.Once
}

// NewHourlyRotator creates a new hourly rotating log writer
func NewHourlyRotator(filename string, maxSize int64, maxFiles int) *HourlyRotator {
	r := &HourlyRotator{
		filename: filename,
		linkName: filename + ".log",
		maxSize:  maxSize,
		maxFiles: maxFiles,
	}

	// Start cleanup goroutine
	r.cleanupOnce.Do(func() {
		go r.cleanup()
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

	return r.currentFile.Write(p)
}

// rotate checks if rotation is needed and performs it
func (r *HourlyRotator) rotate() error {
	now := time.Now()
	hour := now.Format("2006010215")

	// Check if rotation needed (hour change or file size limit exceeded)
	needRotate := r.currentHour != hour
	if r.currentFile != nil && !needRotate {
		// Check file size
		if stat, err := r.currentFile.Stat(); err == nil && stat.Size() >= r.maxSize {
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
	ensureDir(filepath.Dir(r.filename))

	// Generate new filename
	newFilename := r.filename + hour + ".log"

	// Open new file
	file, err := os.OpenFile(newFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) // #nosec G304
	if err != nil {
		return err
	}

	r.currentFile = file
	r.currentHour = hour

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

// cleanup periodically cleans up old log files
func (r *HourlyRotator) cleanup() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		r.cleanupOldFiles()
	}
}

// cleanupOldFiles cleans up expired log files
func (r *HourlyRotator) cleanupOldFiles() {
	dir := filepath.Dir(r.filename)
	base := filepath.Base(r.filename)

	// Read all files in directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading log directory %s for cleanup: %v\n", dir, err)
		return
	}

	// Filter log files
	var logFiles []string
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			// Match format: base + YYYYMMDDHH + .log
			if strings.HasPrefix(name, base) && strings.HasSuffix(name, ".log") && len(name) == len(base)+14 {
				logFiles = append(logFiles, name)
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

		fullPath := filepath.Join(dir, filename)
		fmt.Printf("Deleting old log file: %s\n", fullPath)
		if err := os.Remove(fullPath); err != nil {
			fmt.Printf("Failed to delete old log file %s: %v\n", fullPath, err)
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
