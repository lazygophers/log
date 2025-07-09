package log

import (
	"fmt"
	"path"
	"strconv"
	"strings"
)

// Format defines the interface for formatting a log entry.
type Format interface {
	// Format takes a log entry and returns the formatted byte slice.
	Format(entry *Entry) []byte
}

// FormatFull extends the Format interface with additional configuration methods.
type FormatFull interface {
	Format
	// ParsingAndEscaping enables/disables message parsing and escaping
	ParsingAndEscaping(disable bool)
	// Caller enables/disables caller information in log output
	Caller(disable bool)
	// Clone creates a copy of the formatter
	Clone() Format
}

type Formatter struct {
	Module                    string // Name of the logging module
	DisableParsingAndEscaping bool   // If true, disables message parsing/escaping
	DisableCaller             bool   // If true, disables caller information
}

func (p *Formatter) format(entry *Entry) []byte {
	b := GetBuffer()
	defer PutBuffer(b)

	if len(entry.PrefixMsg) > 0 {
		b.Write(entry.PrefixMsg)
		b.Write([]byte(" "))
	}

	b.WriteString(fmt.Sprintf("(%d.%d) ", entry.Pid, entry.Gid))
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999Z07:00"))
	color := getColorByLevel(entry.Level)
	b.Write(color)
	b.Write([]byte(" ["))
	b.WriteString(entry.Level.String())
	b.Write([]byte("] "))
	b.Write(colorEnd)
	b.WriteString(strings.TrimSpace(entry.Message))

	if !p.DisableCaller || len(entry.TraceId) > 0 {
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

	if len(entry.SuffixMsg) > 0 {
		b.Write(entry.SuffixMsg)
	}
	b.WriteByte('\n')
	return b.Bytes()
}

// Format implements the Format interface. It processes the log entry and returns formatted bytes.
// If parsing is enabled, it splits multi-line messages into separate log entries.
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

// ParsingAndEscaping controls whether log messages are parsed and escaped.
// When disabled, messages are logged as-is without processing.
func (p *Formatter) ParsingAndEscaping(disable bool) {
	p.DisableParsingAndEscaping = disable
}

// Caller controls whether caller information (file, line, function) is included in logs.
func (p *Formatter) Caller(disable bool) {
	p.DisableCaller = disable
}

// Clone creates a deep copy of the Formatter instance.
func (p *Formatter) Clone() Format {
	return &Formatter{
		Module:                    p.Module,
		DisableParsingAndEscaping: p.DisableParsingAndEscaping,
		DisableCaller:             p.DisableCaller,
	}
}

var (
	colorBlack   = []byte("\u001B[30m") // ANSI black
	colorRed     = []byte("\u001B[31m") // ANSI red
	colorGreen   = []byte("\u001B[32m") // ANSI green
	colorYellow  = []byte("\u001B[33m") // ANSI yellow
	colorBlue    = []byte("\u001B[34m") // ANSI blue
	colorMagenta = []byte("\u001B[35m") // ANSI magenta
	colorCyan    = []byte("\u001B[36m") // ANSI cyan
	colorGray    = []byte("\u001B[37m") // ANSI gray
	colorWhite   = []byte("\u001B[38m") // ANSI white
	colorEnd     = []byte("\u001B[0m")  // ANSI reset
)

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

// SplitPackageName splits a fully qualified function name into directory and function parts.
// Input: fully qualified function name (e.g., "github.com/user/pkg.Function")
// Output:
//   - callDir: package path (e.g., "github.com/user/pkg")
//   - callFunc: function name (e.g., "Function")
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
