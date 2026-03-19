package log

import (
	"crypto/rand"
	"encoding/hex"
	"sync"

	"github.com/petermattis/goid"
)

// traceMap stores trace IDs using sync.Map for high-concurrency read/write performance.
var traceMap sync.Map

// DisableTrace is a global switch to disable tracing.
var DisableTrace bool

// getTrace returns the trace ID for the specified goroutine.
//
//go:inline
func getTrace(gid int64) string {
	if val, ok := traceMap.Load(gid); ok {
		return val.(string)
	}
	return ""
}

// setTrace sets the trace ID for the specified goroutine. An empty string triggers auto-generation.
//
//go:inline
func setTrace(gid int64, traceId string) {
	if DisableTrace {
		return
	}
	if traceId == "" {
		traceId = fastGenTraceId()
	}
	traceMap.Store(gid, traceId)
}

// delTrace removes the trace ID for the specified goroutine.
//
//go:inline
func delTrace(gid int64) {
	traceMap.Delete(gid)
}

// GetTrace returns the trace ID of the current goroutine.
func GetTrace() string {
	return getTrace(goid.Get())
}

// GetTraceWithGID returns the trace ID for the specified goroutine ID.
func GetTraceWithGID(gid int64) string {
	return getTrace(gid)
}

// SetTrace sets the trace ID for the current goroutine. An empty argument triggers auto-generation.
func SetTrace(traceId ...string) {
	currentGid := goid.Get()
	if len(traceId) > 0 {
		setTrace(currentGid, traceId[0])
		return
	}
	setTrace(currentGid, "")
}

// SetTraceWithGID sets the trace ID for the specified goroutine ID.
func SetTraceWithGID(gid int64, traceId ...string) {
	if len(traceId) > 0 {
		setTrace(gid, traceId[0])
		return
	}
	setTrace(gid, "")
}

// DelTrace removes the trace ID of the current goroutine.
func DelTrace() {
	delTrace(goid.Get())
}

// DelTraceWithGID removes the trace ID for the specified goroutine ID.
func DelTraceWithGID(gid int64) {
	delTrace(gid)
}

// fastGenTraceId generates a high-performance trace ID.
//
//go:inline
func fastGenTraceId() string {
	var buf [8]byte
	_, _ = rand.Read(buf[:])
	return hex.EncodeToString(buf[:])
}

// GenTraceId generates a 16-character unique trace ID.
func GenTraceId() string {
	return fastGenTraceId()
}

// clearTraceMapForTest is a test helper that clears the trace map.
func clearTraceMapForTest() {
	traceMap.Range(func(key, _ any) bool {
		traceMap.Delete(key)
		return true
	})
}

func loadTraceForTest(gid int64) (string, bool) {
	val, ok := traceMap.Load(gid)
	if !ok {
		return "", false
	}
	return val.(string), true
}
