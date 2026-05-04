package log

import (
	"encoding/json"
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
	Level      string `json:"level"`
	Time       string `json:"time,omitempty"`
	Message    string `json:"message"`
	Pid        int    `json:"pid"`
	Gid        int64  `json:"gid,omitempty"`

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

// Format formats log entry to JSON
func (f *JSONFormatter) Format(entry *Entry) []byte {
	b := GetBuffer()
	defer PutBuffer(b)

	je := f.toJSONEntry(entry)

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
		b.WriteString(jsonEscapeString(entry.Message))
		b.WriteString(`"}`)
	} else {
		b.Write(data)
	}

	b.WriteByte('\n')
	return b.Bytes()
}

// toJSONEntry converts Entry to JSON-serializable format
func (f *JSONFormatter) toJSONEntry(entry *Entry) jsonEntry {
	je := jsonEntry{
		Level:   entry.Level.String(),
		Message: entry.Message,
		Pid:     entry.Pid,
	}

	// Timestamp
	if !f.DisableTimestamp && !entry.Time.IsZero() {
		je.Time = entry.Time.Format(time.RFC3339Nano)
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

	return je
}

// jsonEscapeString escapes special characters for JSON strings
func jsonEscapeString(s string) string {
	var result []byte
	for _, r := range s {
		switch r {
		case '"':
			result = append(result, `\"`...)
		case '\\':
			result = append(result, `\\`...)
		case '\n':
			result = append(result, `\n`...)
		case '\r':
			result = append(result, `\r`...)
		case '\t':
			result = append(result, `\t`...)
		default:
			if r < 32 {
				result = append(result, `\u00`...)
				result = append(result, hexByte(byte(r)>>4))
				result = append(result, hexByte(byte(r)&0x0F))
			} else {
				result = append(result, string(r)...)
			}
		}
	}
	return string(result)
}

// hexByte converts a byte to its hex character
func hexByte(b byte) byte {
	if b < 10 {
		return '0' + b
	}
	return 'a' + b - 10
}
