package constant

import (
	"encoding/json"
	"time"
)

// KV represents a key-value pair for structured logging
type KV struct {
	Key   string
	Value interface{}
}

// Entry represents a log entry
//
// Field layout is optimized for cache performance:
// - Hot path fields (accessed on every log) are grouped at the beginning
// - Fields of similar sizes are grouped together to minimize padding
// - Total size: 200 bytes (4 bytes padding waste = 2.0%)
type Entry struct {
	// Hot path fields - accessed on every log call
	Level      Level     `json:"level"`
	Pid        int       `json:"pid"`
	Gid        int64     `json:"gid,omitempty"`
	CallerLine int       `json:"caller_line,omitempty"`

	// Timestamp - accessed frequently but less than core fields
	Time       time.Time `json:"-"`
	TimeStr    string    `json:"time,omitempty"` // Cached formatted timestamp for performance
	TimeStrSet bool      `json:"-"`              // Internal flag

	// String fields (16 bytes each) - ordered by access frequency
	Message    string `json:"message"`
	File       string `json:"caller_file,omitempty"`
	CallerFunc string `json:"caller_func,omitempty"`
	CallerDir  string `json:"caller_dir,omitempty"`
	CallerName string `json:"caller_name,omitempty"`
	TraceId    string `json:"trace_id,omitempty"`

	// Byte slices (24 bytes each) - lower frequency
	PrefixMsg []byte `json:"prefix_msg,omitempty"`
	SuffixMsg []byte `json:"suffix_msg,omitempty"`

	// Structured fields (key-value pairs)
	Fields []KV `json:"fields,omitempty"`
}

// MarshalJSON implements json.Marshaler interface for custom JSON serialization
func (e *Entry) MarshalJSON() ([]byte, error) {
	// Create a map for conditional JSON serialization
	m := make(map[string]interface{}, 12)
	m["level"] = e.Level.String()
	m["message"] = e.Message
	m["pid"] = e.Pid

	if e.TimeStrSet {
		m["time"] = e.TimeStr
	}

	if e.Gid != 0 {
		m["gid"] = e.Gid
	}

	if e.TraceId != "" {
		m["trace_id"] = e.TraceId
	}

	if e.File != "" {
		m["caller_file"] = e.File
		m["caller_line"] = e.CallerLine
		m["caller_func"] = e.CallerFunc
		if e.CallerDir != "" {
			m["caller_dir"] = e.CallerDir
		}
		if e.CallerName != "" {
			m["caller_name"] = e.CallerName
		}
	}

	if len(e.PrefixMsg) > 0 {
		m["prefix_msg"] = string(e.PrefixMsg)
	}

	if len(e.SuffixMsg) > 0 {
		m["suffix_msg"] = string(e.SuffixMsg)
	}

	if len(e.Fields) > 0 {
		fields := make(map[string]interface{}, len(e.Fields))
		for _, f := range e.Fields {
			fields[f.Key] = f.Value
		}
		m["fields"] = fields
	}

	return json.Marshal(m)
}

// Reset resets Entry to initial values for safe pool reuse
func (p *Entry) Reset() {
	p.Gid = 0
	p.TraceId = ""
	p.Time = time.Time{}
	p.TimeStr = ""
	p.TimeStrSet = false
	p.Message = ""
	p.File = ""
	p.CallerLine = 0
	p.CallerName = ""
	p.CallerDir = ""
	p.CallerFunc = ""
	p.PrefixMsg = p.PrefixMsg[:0]
	p.SuffixMsg = p.SuffixMsg[:0]
}

