package log

import (
	"sync"
	"time"
)

// Entry 表示一个日志条目，定义了日志输出的完整信息。
type Entry struct {
	Pid        int       // 进程ID (Process ID)
	Gid        int64     // 协程ID (Goroutine ID)
	TraceId    string    // 分布式追踪ID (Trace ID)，用于在分布式系统中串联日志
	Time       time.Time // 日志记录的时间戳
	Level      Level     // 日志级别 (e.g., Debug, Info, Error)
	File       string    // 产生日志的源文件名
	Message    string    // 核心日志消息内容
	CallerName string    // 调用者的包名
	CallerLine int       // 调用者在源文件中的行号
	CallerDir  string    // 调用者所在的目录路径
	CallerFunc string    // 调用者的函数名
	PrefixMsg  []byte    // 在核心消息前添加的前缀内容
	SuffixMsg  []byte    // 在核心消息后追加的后缀内容
}

// NewEntry 使用对象池创建一个新的 Entry 实例。
// 这样做可以复用对象，减少垃圾回收的压力，提高性能。
// 返回值:
// - *Entry: 一个初始化后的日志条目指针，其中 Pid 已经被设置为当前进程ID。
func NewEntry() *Entry {
	return &Entry{Pid: pid}
}

// Reset 将一个 Entry 对象的状态重置为初始值，以便安全地回收到对象池中。
// 这避免了旧数据在对象复用时被意外使用。
// 优化：使用高性能重置策略，减少条件判断
//go:inline
func (p *Entry) Reset() {
	// 批量重置数值字段
	p.Gid = 0
	p.CallerLine = 0
	p.Level = 0
	
	// 批量重置字符串字段（编译器会优化连续赋值）
	p.TraceId, p.File, p.Message = "", "", ""
	p.CallerName, p.CallerDir, p.CallerFunc = "", "", ""
	
	// 高效清空切片，保留容量（无条件重置，减少分支）
	p.PrefixMsg = p.PrefixMsg[:0]
	p.SuffixMsg = p.SuffixMsg[:0]
}

// entryPool 是一个 sync.Pool，用于缓存和复用 Entry 对象。
// 当需要新的 Entry 时，会先尝试从池中获取，如果池为空，则通过 New 字段中指定的函数创建一个新的。
// 极大地提升了在高并发场景下日志记录的性能。
var entryPool = sync.Pool{
	New: func() any {
		return &Entry{Pid: pid}
	},
}

// getEntry 从对象池中获取一个 Entry 实例
// 优化：高性能实现，减少函数调用开销
//go:inline
func getEntry() *Entry {
	if entry := entryPool.Get(); entry != nil {
		return entry.(*Entry)
	}
	return &Entry{Pid: pid}
}

// putEntry 将 Entry 实例放回对象池中以供复用  
// 优化：高性能实现，减少分支判断开销
//go:inline
func putEntry(entry *Entry) {
	if entry != nil {
		entry.Reset()
		entryPool.Put(entry)
	}
}

// FastGetEntry 快速获取Entry的内联版本 (已被getEntry优化版本替代)
// 保留作为备用实现
func FastGetEntry() *Entry {
	return getEntry()
}

// FastPutEntry 快速归还Entry的内联版本 (已被putEntry优化版本替代) 
// 保留作为备用实现
func FastPutEntry(entry *Entry) {
	putEntry(entry)
}
