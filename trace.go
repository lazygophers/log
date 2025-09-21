package log

import (
	"crypto/rand"
	"encoding/hex"
	"sync"

	"github.com/petermattis/goid"
)

// 追踪ID存储，优化为读多写少场景
var (
	traceMap   = make(map[int64]string, 64) // 预分配容量
	traceMapMu sync.RWMutex
)

// DisableTrace 全局开关，禁用追踪功能
var DisableTrace bool

// 获取指定 goroutine 的追踪ID
//
//go:inline
func getTrace(gid int64) string {
	traceMapMu.RLock()
	tid := traceMap[gid]
	traceMapMu.RUnlock()
	return tid
}

// 为指定 goroutine 设置追踪ID，空字符串自动生成
//
//go:inline
func setTrace(gid int64, traceId string) {
	if DisableTrace {
		return
	}
	if traceId == "" {
		traceId = fastGenTraceId()
	}
	traceMapMu.Lock()
	traceMap[gid] = traceId
	traceMapMu.Unlock()
}

// 删除指定 goroutine 的追踪ID
//
//go:inline
func delTrace(gid int64) {
	traceMapMu.Lock()
	delete(traceMap, gid)
	traceMapMu.Unlock()
}

// GetTrace 获取当前 goroutine 的追踪ID
func GetTrace() string {
	return getTrace(goid.Get())
}

// GetTraceWithGID 获取指定 GID 的追踪ID
func GetTraceWithGID(gid int64) string {
	return getTrace(gid)
}

// SetTrace 为当前 goroutine 设置追踪ID，空参数自动生成
func SetTrace(traceId ...string) {
	// 使用 goid.Get() 获取当前 Goroutine 的 ID
	currentGid := goid.Get()
	if len(traceId) > 0 {
		setTrace(currentGid, traceId[0])
		return
	}
	// 如果没有提供 traceId，则传入空字符串，由 setTrace 内部处理生成逻辑。
	setTrace(currentGid, "")
}

// SetTraceWithGID 为指定 GID 设置追踪ID
func SetTraceWithGID(gid int64, traceId ...string) {
	if len(traceId) > 0 {
		setTrace(gid, traceId[0])
		return
	}
	setTrace(gid, "")
}

// DelTrace 删除当前 goroutine 的追踪ID
func DelTrace() {
	delTrace(goid.Get())
}

// DelTraceWithGID 删除指定 GID 的追踪ID
func DelTraceWithGID(gid int64) {
	delTrace(gid)
}

// 高性能 TraceID 生成
//
//go:inline
func fastGenTraceId() string {
	var buf [8]byte
	_, _ = rand.Read(buf[:])
	return hex.EncodeToString(buf[:])
}

// GenTraceId 生成 16 字符唯一追踪ID
func GenTraceId() string {
	return fastGenTraceId()
}

// 测试辅助函数
func clearTraceMapForTest() {
	traceMapMu.Lock()
	for k := range traceMap {
		delete(traceMap, k)
	}
	traceMapMu.Unlock()
}

func loadTraceForTest(gid int64) (string, bool) {
	traceMapMu.RLock()
	trace, exists := traceMap[gid]
	traceMapMu.RUnlock()
	return trace, exists
}
