package log

import "time"

// Entry 结构体表示日志条目
// 包含日志的元数据和上下文信息
type Entry struct {
	// Pid: 进程ID
	// 用于标识生成日志的进程
	Pid int

	// Gid: 协程ID
	// 用于标识生成日志的goroutine
	Gid int64

	// TraceId: 分布式追踪ID
	// 用于分布式系统中的请求追踪
	TraceId string

	// Time: 日志时间戳
	// 记录日志生成的精确时间
	Time time.Time

	// Level: 日志级别
	// 表示日志的严重程度（DEBUG/INFO/WARN/ERROR等）
	Level Level

	// File: 调用文件路径
	// 存储生成日志的源文件完整路径
	File string

	// Message: 日志消息主体
	// 核心日志内容，经过格式化处理
	Message string

	// CallerName: 调用函数全名
	// 包含包路径的完整函数名称
	CallerName string

	// CallerLine: 调用行号
	// 标识日志生成的具体代码行
	CallerLine int

	// CallerDir: 调用文件所在目录
	// 存储源文件的目录路径
	CallerDir string

	// CallerFunc: 调用函数名
	// 仅包含函数名称（不含包路径）
	CallerFunc string

	// PrefixMsg: 消息前缀
	// 存储日志格式化前缀字节数据
	PrefixMsg []byte

	// SuffixMsg: 消息后缀
	// 存储日志格式化后缀字节数据
	SuffixMsg []byte
}

// NewEntry 创建一个新的日志条目实例
// 初始化时自动设置进程ID
// 返回带有默认初始值的Entry指针
func NewEntry() *Entry {
	return &Entry{Pid: pid}
}

// Reset 重置日志条目字段为默认值
// 用于复用Entry对象时清除所有状态
// 注意：PrefixMsg和SuffixMsg使用切片重置而非重新分配
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
