package log

import "time"

type Entry struct {
	Pid     int
	Gid     int64
	TraceId []byte

	Time  time.Time
	Level Level
	File  string

	Message    string
	CallerName string
	CallerLine int
	CallerDir  string
	CallerFunc string

	PrefixMsg []byte
	SuffixMsg []byte
}

func NewEntry() *Entry {
	return &Entry{Pid: pid}
}

func (p *Entry) Reset() {
	p.Gid = 0
	p.TraceId = p.TraceId[:0]
	p.File = ""
	p.Message = ""
	p.CallerName = ""
	p.CallerDir = ""
	p.CallerFunc = ""
	p.PrefixMsg = p.PrefixMsg[:0]
	p.SuffixMsg = p.SuffixMsg[:0]
}
