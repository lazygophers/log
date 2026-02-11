package log

import (
	"sync"
	"time"
)

// Entry represents a log entry
//
// Field layout is optimized for cache performance:
// - Hot path fields (accessed on every log) are grouped at the beginning
// - Fields of similar sizes are grouped together to minimize padding
// - Total size: 200 bytes (4 bytes padding waste = 2.0%)
type Entry struct {
	// Hot path fields - accessed on every log call
	Level      Level     // Log level (1 byte + 7 padding)
	Pid        int       // Process ID (8 bytes)
	Gid        int64     // Goroutine ID (8 bytes)
	CallerLine int       // Caller line number (8 bytes)

	// Timestamp - accessed frequently but less than core fields
	Time       time.Time // Log timestamp (24 bytes)

	// String fields (16 bytes each) - ordered by access frequency
	Message    string    // Core log message (highest frequency)
	File       string    // Source file name
	CallerFunc string    // Caller function name
	CallerDir  string    // Caller directory
	CallerName string    // Caller package name
	TraceId    string    // Trace ID for distributed tracing

	// Byte slices (24 bytes each) - lower frequency
	PrefixMsg  []byte    // Prefix message
	SuffixMsg  []byte    // Suffix message
}

// NewEntry creates a new Entry instance from object pool
func NewEntry() *Entry {
	return &Entry{Pid: pid}
}

// Reset resets Entry to initial values for safe pool reuse
//
//go:inline
func (p *Entry) Reset() {
	// Reset numeric fields
	p.Gid = 0
	p.CallerLine = 0
	p.Level = 0

	// Reset string fields (compiler optimizes consecutive assignments)
	p.TraceId, p.File, p.Message = "", "", ""
	p.CallerName, p.CallerDir, p.CallerFunc = "", "", ""

	// Clear slices efficiently while retaining capacity
	p.PrefixMsg = p.PrefixMsg[:0]
	p.SuffixMsg = p.SuffixMsg[:0]
}

// entryPool is a sync.Pool for caching and reusing Entry objects
var entryPool = sync.Pool{
	New: func() any {
		return &Entry{Pid: pid}
	},
}

// getEntry gets Entry instance from object pool
//
//go:inline
func getEntry() *Entry {
	if entry := entryPool.Get(); entry != nil {
		return entry.(*Entry)
	}
	return &Entry{Pid: pid}
}

// putEntry returns Entry instance to object pool for reuse
//
//go:inline
func putEntry(entry *Entry) {
	if entry != nil {
		entry.Reset()
		entryPool.Put(entry)
	}
}
