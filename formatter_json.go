package log

import (
	"bytes"
	"encoding/json"
	"sync"
	"time"
)

// JSONFormatter implements Format interface for JSON output
type JSONFormatter struct {
	EnablePrettyPrint bool // Enable pretty print with indentation
	DisableTimestamp bool // Disable timestamp field
	DisableCaller    bool // Disable caller information
	DisableTrace     bool // Disable trace information
}

// jsonEntry represents JSON-serializable log entry
type jsonEntry struct {
	Level      string                    `json:"level"`
	Time       string                    `json:"time,omitempty"`
	Message    string                    `json:"message"`
	Pid        int                       `json:"pid"`
	Gid        int64                     `json:"gid,omitempty"`
	Fields     map[string]interface{}    `json:"fields,omitempty"`

	// Caller information
	CallerFile     string `json:"caller_file,omitempty"`
	CallerLine     int    `json:"caller_line,omitempty"`
	CallerFunc     string `json:"caller_func,omitempty"`
	CallerDir      string `json:"caller_dir,omitempty"`
	CallerName     string `json:"caller_name,omitempty"`

	// Trace information
	TraceID string `json:"trace_id,omitempty"`

	// Prefix and suffix
	PrefixMsg string `json:"prefix_msg,omitempty"`
	SuffixMsg string `json:"suffix_msg,omitempty"`
}

// jsonEntryPool reuses jsonEntry objects to reduce memory allocations
var jsonEntryPool = sync.Pool{
	New: func() any {
		return &jsonEntry{
			Fields: make(map[string]interface{}, 8),
		}
	},
}

// getJSONEntry retrieves a jsonEntry from pool and resets it for reuse
func getJSONEntry() *jsonEntry {
	je := jsonEntryPool.Get().(*jsonEntry)
	je.Level = ""
	je.Message = ""
	je.Pid = 0
	je.Time = ""
	je.Gid = 0
	je.TraceID = ""
	je.CallerFile = ""
	je.CallerLine = 0
	je.CallerFunc = ""
	je.CallerDir = ""
	je.CallerName = ""
	je.PrefixMsg = ""
	je.SuffixMsg = ""
	for k := range je.Fields {
		delete(je.Fields, k)
	}
	return je
}

// Format formats log entry to JSON
func (f *JSONFormatter) Format(entry interface{}) []byte {
	// Type assert to *Entry
	e, ok := entry.(*Entry)
	if !ok {
		return nil
	}

	b := GetBuffer()
	defer PutBuffer(b)

	je := getJSONEntry()
	defer jsonEntryPool.Put(je)

	f.populateJSONEntry(je, e)

	var data []byte
	var err error

	if f.EnablePrettyPrint {
		data, err = json.MarshalIndent(je, "", "  ")
	} else {
		data, err = json.Marshal(je)
	}

	if err != nil {
		// Fallback to error message if JSON marshaling fails
		b.WriteString(`{"level":"error","message":"JSON marshaling failed: `)
		b.WriteString(err.Error())
		b.WriteString(`","original":"`)
		b.WriteString(jsonEscapeString(e.Message))
		b.WriteString(`"}`)
	} else {
		b.Write(data)
	}

	b.WriteByte('\n')
	return b.Bytes()
}

// populateJSONEntry populates jsonEntry from Entry (reuses pooled object)
func (f *JSONFormatter) populateJSONEntry(je *jsonEntry, entry *Entry) {
	je.Level = entry.Level.String()
	je.Message = entry.Message
	je.Pid = entry.Pid

	// Timestamp - use cached string if available
	if !f.DisableTimestamp && !entry.Time.IsZero() {
		if entry.timeStrSet {
			je.Time = entry.timeStr
		} else {
			je.Time = entry.Time.Format(time.RFC3339Nano)
		}
	}

	// Trace ID and GID
	if !f.DisableTrace {
		if entry.Gid != 0 {
			je.Gid = entry.Gid
		}
		if entry.TraceId != "" {
			je.TraceID = entry.TraceId
		}
	}

	// Caller information
	if !f.DisableCaller && entry.File != "" {
		je.CallerFile = entry.File
		je.CallerLine = entry.CallerLine
		je.CallerFunc = entry.CallerFunc
		je.CallerDir = entry.CallerDir
		je.CallerName = entry.CallerName
	}

	// Prefix and suffix
	if len(entry.PrefixMsg) > 0 {
		je.PrefixMsg = string(entry.PrefixMsg)
	}
	if len(entry.SuffixMsg) > 0 {
		je.SuffixMsg = string(entry.SuffixMsg)
	}

	// Structured fields - reuse existing map from pool
	if len(entry.Fields) > 0 {
		for _, field := range entry.Fields {
			je.Fields[field.Key] = field.Value
		}
	}
}

// toJSONEntry creates a new jsonEntry from Entry (for testing compatibility)
func (f *JSONFormatter) toJSONEntry(entry *Entry) jsonEntry {
	je := getJSONEntry()
	f.populateJSONEntry(je, entry)

	// Create a copy for return (since pool will be reused)
	result := jsonEntry{
		Level:      je.Level,
		Message:    je.Message,
		Pid:        je.Pid,
		Time:       je.Time,
		Gid:        je.Gid,
		TraceID:    je.TraceID,
		CallerFile: je.CallerFile,
		CallerLine: je.CallerLine,
		CallerFunc: je.CallerFunc,
		CallerDir:  je.CallerDir,
		CallerName: je.CallerName,
		PrefixMsg:  je.PrefixMsg,
		SuffixMsg:  je.SuffixMsg,
	}

	if len(je.Fields) > 0 {
		result.Fields = make(map[string]interface{}, len(je.Fields))
		for k, v := range je.Fields {
			result.Fields[k] = v
		}
	}

	// Return the pooled entry
	jsonEntryPool.Put(je)

	return result
}

// jsonEscapeString escapes special characters for JSON strings
// Optimized with bytes.Buffer and pre-allocation to reduce allocations
func jsonEscapeString(s string) string {
	if s == "" {
		return ""
	}

	// Pre-allocate capacity to reduce reallocations
	// Estimate: normal chars = 1 byte, special chars = 2-6 bytes
	estimatedLen := len(s) + len(s)/4
	var result bytes.Buffer
	result.Grow(estimatedLen)

	for _, r := range s {
		switch r {
		case '"':
			result.WriteString(`\"`)
		case '\\':
			result.WriteString(`\\`)
		case '\n':
			result.WriteString(`\n`)
		case '\r':
			result.WriteString(`\r`)
		case '\t':
			result.WriteString(`\t`)
		default:
			if r < 32 {
				result.WriteString(`\u00`)
				result.WriteByte(hexByte(byte(r) >> 4))
				result.WriteByte(hexByte(byte(r) & 0x0F))
			} else {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}

// hexByte converts a byte to its hex character
func hexByte(b byte) byte {
	if b < 10 {
		return '0' + b
	}
	return 'a' + b - 10
}
