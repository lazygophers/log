package log

import (
	"fmt"
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

	// Write prefix message
	if len(entry.PrefixMsg) > 0 {
		b.Write(entry.PrefixMsg)
		b.Write([]byte(" "))
	}

	// Write process and goroutine IDs
	b.WriteString(fmt.Sprintf("(%d.%d) ", entry.Pid, entry.Gid))
	// Write formatted timestamp
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999Z07:00"))
	// Get color by log level
	color := getColorByLevel(entry.Level)
	// Write color code
	b.Write(color)
	// Write log level
	b.Write([]byte(" ["))
	b.WriteString(entry.Level.String())
	b.Write([]byte("] "))
	// Write color end code
	b.Write(colorEnd)
	// Write log message
	b.WriteString(strings.TrimSpace(entry.Message))

	// Write additional info if caller info enabled or TraceID exists
	if !p.DisableCaller || entry.TraceId != "" {
		b.Write(colorCyan)
		b.WriteString(" [ ")
		// Write caller info if not disabled
		if !p.DisableCaller {
			b.WriteString(path.Join(entry.CallerDir, path.Base(entry.File)))
			b.Write([]byte(":"))
			b.WriteString(strconv.Itoa(entry.CallerLine))
			b.Write([]byte(" "))
			b.WriteString(entry.CallerFunc)
			b.Write([]byte(" "))
		}
		// Write TraceID if exists
		if entry.TraceId != "" {
			b.WriteString(entry.TraceId)
			b.Write([]byte(" "))
		}
		b.Write([]byte("]"))
		b.Write(colorEnd)
	} else if !p.DisableCaller {
		b.WriteString(" ")
		b.WriteString(path.Join(entry.CallerDir, path.Base(entry.File)))
		b.Write([]byte(":"))
		b.WriteString(strconv.Itoa(entry.CallerLine))
		b.Write([]byte(" "))
		b.WriteString(entry.CallerFunc)
		b.Write([]byte(" "))
	}

	// Write suffix message
	if len(entry.SuffixMsg) > 0 {
		b.Write(entry.SuffixMsg)
	}
	// Write newline
	b.WriteByte('\n')
	return b.Bytes()
}

// Format implements Format interface, handles multi-line messages
func (p *Formatter) Format(entry *Entry) []byte {
	if p.DisableParsingAndEscaping {
		return p.format(entry)
	}
	b := GetBuffer()
	defer PutBuffer(b)
	for _, msg := range strings.Split(entry.Message, "\n") {
		entry.Message = msg
		b.Write(p.format(entry))
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
func SplitPackageName(f string) (callDir string, callFunc string) {
	slashIndex := strings.LastIndex(f, "/")
	if slashIndex > 0 {
		idx := strings.Index(f[slashIndex:], ".") + slashIndex
		callDir, callFunc = f[:idx], f[idx+1:]
	} else {
		slashIndex = strings.Index(f, ".")
		if slashIndex > 0 {
			callDir, callFunc = f[:slashIndex], f[slashIndex+1:]
		} else {
			callDir, callFunc = f, ""
		}
	}
	callDir = strings.TrimPrefix(callDir, "github.com/")
	callDir = strings.TrimPrefix(callDir, "lazygophers/")
	return
}
