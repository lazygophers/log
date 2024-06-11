package log

import (
	"fmt"
	"path"
	"strconv"
	"strings"
)

type Format interface {
	Format(entry *Entry) []byte
}

type FormatFull interface {
	Format
	// ParsingAndEscaping Default should be on
	ParsingAndEscaping(disable bool)

	// DisableCaller Default should be on
	Caller(disable bool)

	Clone() Format
}

type Formatter struct {
	Module string

	DisableParsingAndEscaping bool
	DisableCaller             bool
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

func (p *Formatter) ParsingAndEscaping(disable bool) {
	p.DisableParsingAndEscaping = disable
}

func (p *Formatter) Caller(disable bool) {
	p.DisableCaller = disable
}

func (p *Formatter) Clone() Format {
	return &Formatter{
		Module:                    p.Module,
		DisableParsingAndEscaping: p.DisableParsingAndEscaping,
		DisableCaller:             p.DisableCaller,
	}
}

var (
	colorBlack   = []byte("\u001B[30m")
	colorRed     = []byte("\u001B[31m")
	colorGreen   = []byte("\u001B[32m")
	colorYellow  = []byte("\u001B[33m")
	colorBlue    = []byte("\u001B[34m")
	colorMagenta = []byte("\u001B[35m")
	colorCyan    = []byte("\u001B[36m")
	colorGray    = []byte("\u001B[37m")
	colorWhite   = []byte("\u001B[38m")

	colorEnd = []byte("\u001B[0m")
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
