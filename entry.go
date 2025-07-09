// Package log 提供日志条目的结构定义和基础操作方法
// 包含日志生成时的完整上下文信息，支持进程、协程、调用栈等元数据记录
package log

import "time"

// Entry 表示单条日志的完整上下文信息
type Entry struct {
	Pid        int       // Pid: 生成日志的进程ID
	Gid        int64     // Gid: 生成日志的goroutine ID
	TraceId    string    // TraceId: 分布式追踪ID，用于请求链路追踪
	Time       time.Time // Time: 日志生成时间戳
	Level      Level     // Level: 日志级别(DEBUG/INFO/WARN/ERROR/FATAL)
	File       string    // File: 产生日志的源文件完整路径
	Message    string    // Message: 格式化后的日志消息内容
	CallerName string    // CallerName: 调用函数的完整签名(包路径+函数名)
	CallerLine int       // CallerLine: 产生日志的代码行号
	CallerDir  string    // CallerDir: 源文件所在目录路径
	CallerFunc string    // CallerFunc: 函数名称(不含包路径)
	PrefixMsg  []byte    // PrefixMsg: 日志前缀字节数据(如时间戳/级别前缀)
	SuffixMsg  []byte    // SuffixMsg: 日志后缀字节数据(如换行符)
}

// NewEntry 创建并初始化新的日志条目
// 初始化进程ID(Pid)字段为当前进程ID
// 返回: *Entry - 初始化完成的日志条目指针
func NewEntry() *Entry {
	return &Entry{Pid: pid}
}

// Reset 重置日志条目状态以便复用
// 保留PrefixMsg/SuffixMsg底层数组避免内存重分配
// 注意:
//   - 进程ID(Pid)字段不会重置，保持创建时的初始值
//   - 重置后所有字段恢复零值状态(除Pid外)
func (p *Entry) Reset() {
	p.Gid = 0
	p.TraceId = ""
	p.File = ""
	p.Message = ""
	p.CallerName = ""
	p.CallerDir = ""
	p.CallerFunc = ""
	p.PrefixMsg = p.PrefixMsg[:0] // 切片重置(保留底层数组capacity)
	p.SuffixMsg = p.SuffixMsg[:0]
}
