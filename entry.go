package log

import (
	"sync"

	"github.com/lazygophers/log/constant"
)

// Entry re-exports constant.Entry for convenience
type Entry = constant.Entry

// KV re-exports constant.KV for convenience
type KV = constant.KV

// NewEntry creates a new Entry instance from object pool
func NewEntry() *Entry {
	return &Entry{Pid: pid}
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
	return entryPool.Get().(*Entry)
}

// putEntry returns Entry instance to object pool for reuse
//
//go:inline
func putEntry(entry *Entry) {
	if entry == nil {
		return
	}
	entry.Reset()
	entryPool.Put(entry)
}

