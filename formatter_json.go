package log

import (
	"bytes"
	"encoding/json"
)

// JSONFormatter implements Format interface for JSON output
type JSONFormatter struct {
	EnablePrettyPrint bool // Enable pretty print with indentation
	DisableCaller    bool // Disable caller information
	DisableTrace     bool // Disable trace information
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

	// Apply conditional modifications by creating a shallow copy
	serializeEntry := *e

	if f.DisableTrace {
		serializeEntry.Gid = 0
		serializeEntry.TraceId = ""
	}

	if f.DisableCaller {
		serializeEntry.File = ""
		serializeEntry.CallerLine = 0
		serializeEntry.CallerFunc = ""
		serializeEntry.CallerDir = ""
		serializeEntry.CallerName = ""
	}

	var data []byte
	var err error

	if f.EnablePrettyPrint {
		data, err = json.MarshalIndent(&serializeEntry, "", "  ")
	} else {
		data, err = json.Marshal(&serializeEntry)
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

// jsonEscapeString escapes special characters for JSON strings
// Optimized with bytes.Buffer and pre-allocation to reduce allocations
func jsonEscapeString(s string) string {
	if s == "" {
		return ""
	}

	// Pre-allocate capacity to reduce reallocations
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
