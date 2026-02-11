package log

import (
	"bytes"
	"path"
	"strconv"
	"strings"
)

// Format defines log formatting interface
type Format interface {
	// Format formats log entry to byte array
	Format(entry *Entry) []byte
}

// FormatFull extends Format with additional controls
type FormatFull interface {
	Format // Inherits basic formatting interface

	// ParsingAndEscaping controls message parsing and escaping
	ParsingAndEscaping(disable bool)

	Caller(disable bool)

	Clone() Format
}

// Formatter implements FormatFull interface with default formatting
type Formatter struct {
	Module                    string // Log module name
	DisableParsingAndEscaping bool   // Disable message parsing and escaping
	DisableCaller             bool   // Disable caller information
}

// format handles single-line log formatting
func (p *Formatter) format(entry *Entry) []byte {
	// Get byte buffer from pool, return after use
	b := GetBuffer()
	defer PutBuffer(b)

	p.formatPrefix(b, entry)
	p.formatTimestamp(b, entry)
	p.formatLevel(b, entry)
	b.WriteString(strings.TrimSpace(entry.Message))
	p.formatCallerAndTrace(b, entry)
	p.formatSuffix(b, entry)

	return b.Bytes()
}

// formatPrefix writes prefix message and process/goroutine IDs
//
//go:inline
func (p *Formatter) formatPrefix(b *bytes.Buffer, entry *Entry) {
	if len(entry.PrefixMsg) > 0 {
		b.Write(entry.PrefixMsg)
		b.WriteByte(' ')
	}
	b.WriteByte('(')
	b.WriteString(strconv.Itoa(entry.Pid))
	b.WriteByte('.')
	b.WriteString(strconv.FormatInt(entry.Gid, 10))
	b.WriteString(") ")
}

// formatTimestamp writes formatted timestamp
//
//go:inline
func (p *Formatter) formatTimestamp(b *bytes.Buffer, entry *Entry) {
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999Z07:00"))
}

// formatLevel writes colored log level
//
//go:inline
func (p *Formatter) formatLevel(b *bytes.Buffer, entry *Entry) {
	color := getColorByLevel(entry.Level)
	b.Write(color)
	b.Write([]byte(" ["))
	b.WriteString(entry.Level.String())
	b.Write([]byte("] "))
	b.Write(colorEnd)
}

// formatCallerAndTrace writes caller and trace information
//
//go:inline
func (p *Formatter) formatCallerAndTrace(b *bytes.Buffer, entry *Entry) {
	// Skip if both disabled and no trace ID
	if p.DisableCaller && entry.TraceId == "" {
		return
	}

	b.Write(colorCyan)
	b.WriteString(" [ ")

	if !p.DisableCaller {
		b.WriteString(path.Join(entry.CallerDir, path.Base(entry.File)))
		b.Write([]byte(":"))
		b.WriteString(strconv.Itoa(entry.CallerLine))
		b.Write([]byte(" "))
		b.WriteString(entry.CallerFunc)
		b.Write([]byte(" "))
	}

	if entry.TraceId != "" {
		b.WriteString(entry.TraceId)
		b.Write([]byte(" "))
	}

	b.Write([]byte("]"))
	b.Write(colorEnd)
}

// formatSuffix writes suffix message and newline
//
//go:inline
func (p *Formatter) formatSuffix(b *bytes.Buffer, entry *Entry) {
	if len(entry.SuffixMsg) > 0 {
		b.Write(entry.SuffixMsg)
	}
	b.WriteByte('\n')
}

// Format implements Format interface, handles multi-line messages
func (p *Formatter) Format(entry *Entry) []byte {
	if p.DisableParsingAndEscaping {
		return p.format(entry)
	}
	b := GetBuffer()
	defer PutBuffer(b)

	// Manual iteration to avoid strings.Split allocations
	// This creates fewer intermediate objects compared to strings.Split
	msg := entry.Message
	start := 0
	for {
		// Find next newline in the remaining message
		idx := strings.IndexByte(msg[start:], '\n')
		if idx == -1 {
			// Last line (or only line if no \n found)
			entry.Message = msg[start:]
			b.Write(p.format(entry))
			break
		}
		// idx is relative to msg[start:], so we add start to get absolute position
		absIdx := start + idx
		// Extract line without the newline character
		entry.Message = msg[start:absIdx]
		b.Write(p.format(entry))
		start = absIdx + 1 // Move past the newline
	}

	return b.Bytes()
}

// ParsingAndEscaping sets message parsing and escaping
func (p *Formatter) ParsingAndEscaping(disable bool) {
	p.DisableParsingAndEscaping = disable
}

// Caller sets caller information display
func (p *Formatter) Caller(disable bool) {
	p.DisableCaller = disable
}

// Clone creates a deep copy of Formatter
func (p *Formatter) Clone() Format {
	return &Formatter{
		Module:                    p.Module,
		DisableParsingAndEscaping: p.DisableParsingAndEscaping,
		DisableCaller:             p.DisableCaller,
	}
}

var (
	colorRed    = []byte("\u001B[31m")
	colorGreen  = []byte("\u001B[32m")
	colorYellow = []byte("\u001B[33m")
	colorCyan   = []byte("\u001B[36m")
	colorEnd    = []byte("\u001B[0m")
)

// getColorByLevel gets terminal color code by log level
func getColorByLevel(level Level) []byte {
	switch level {
	case DebugLevel, TraceLevel:
		return colorGreen
	case WarnLevel:
		return colorYellow
	case ErrorLevel, FatalLevel, PanicLevel:
		return colorRed
	default:
		return colorGreen
	}
}

// SplitPackageName splits full package path into directory and function name
// Optimized to minimize string operations and allocations
func SplitPackageName(f string) (callDir string, callFunc string) {
	// Find last slash using IndexByte (faster than LastIndex with string)
	lastSlash := strings.LastIndexByte(f, '/')
	if lastSlash == -1 {
		// No slash: look for first dot
		dotIdx := strings.IndexByte(f, '.')
		if dotIdx > 0 {
			return f[:dotIdx], f[dotIdx+1:]
		}
		return f, ""
	}

	// Find first dot after the last slash
	dotIdx := strings.IndexByte(f[lastSlash+1:], '.')
	if dotIdx == -1 {
		// No dot after slash
		return f, ""
	}
	dotIdx += lastSlash + 1 // Adjust to absolute position

	// Extract directory and function
	callDir = f[:dotIdx]
	callFunc = f[dotIdx+1:]

	// Trim known prefixes efficiently (single pass)
	// Check "github.com/" first (more specific)
	if strings.HasPrefix(callDir, "github.com/") {
		callDir = callDir[11:] // len("github.com/") = 11
	} else if strings.HasPrefix(callDir, "lazygophers/") {
		callDir = callDir[12:] // len("lazygophers/") = 12
	}

	return
}
