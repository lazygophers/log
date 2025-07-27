package log

import (
	"sync"
	"time"
)

// Entry 表示一个日志条目
type Entry struct {
	Pid        int       // 进程ID
	Gid        int64     // 协程ID
	TraceId    string    // 分布式追踪ID
	Time       time.Time // 日志时间戳
	Level      Level     // 日志级别
	File       string    // 产生日志的文件名
	Message    string    // 日志消息内容
	CallerName string    // 调用者包名
	CallerLine int       // 调用者行号
	CallerDir  string    // 调用者目录路径
	CallerFunc string    // 调用者函数名
	PrefixMsg  []byte    // 日志消息前缀
	SuffixMsg  []byte    // 日志消息后缀
}

// NewEntry 创建一个新的日志条目实例
// 返回：初始化后的日志条目指针
func NewEntry() *Entry {
	return &Entry{Pid: pid}
}

// Reset 重置日志条目所有字段为空状态
func (p *Entry) Reset() {
	p.Gid = 0
	p.TraceId = ""
	p.File = ""
	p.Message = ""
	p.CallerName = ""
	p.CallerDir = ""
	p.CallerFunc = ""
	p.PrefixMsg = p.PrefixMsg[:0]
	p.SuffixMsg = p.SuffixMsg[:0]
}

var entryPool = sync.Pool{
	New: func() any {
		return NewEntry()
	},
}
