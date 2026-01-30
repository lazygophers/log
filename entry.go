package log

import (
	"sync"
	"time"
)

// Entry represents a log entry
type Entry struct {
	Pid        int       // Process ID
	Gid        int64     // Goroutine ID
	TraceId    string    // Trace ID for distributed tracing
	Time       time.Time // Log timestamp
	Level      Level     // Log level
	File       string    // Source file name
	Message    string    // Core log message
	CallerName string    // Caller package name
	CallerLine int       // Caller line number
	CallerDir  string    // Caller directory
	CallerFunc string    // Caller function name
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
