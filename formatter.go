// Package log 提供日志记录功能
package log

import (
	"fmt"
	"path"
	"strconv"
	"strings"
)

// Format 定义日志条目格式化的基本接口
type Format interface {
	// Format 将日志条目格式化为字节切片
	Format(entry *Entry) []byte
}

// FormatFull 扩展的格式化接口，提供额外控制选项
type FormatFull interface {
	Format
	// ParsingAndEscaping 控制是否解析转义字符（默认开启）
	ParsingAndEscaping(disable bool)

	// Caller 控制是否显示调用者信息（默认开启）
	Caller(disable bool)

	// Clone 创建格式化器的副本
	Clone() Format
}

// Formatter 实现 Format 接口，提供带颜色的日志格式化功能
type Formatter struct {
	Module string // 模块名称前缀

	DisableParsingAndEscaping bool // 禁用解析和转义处理
	DisableCaller             bool // 禁用调用者信息显示
}

// format 内部格式化实现，处理单个日志条目
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

// Format 实现 Format 接口，处理多行日志消息
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

// ParsingAndEscaping 设置是否禁用解析和转义处理
func (p *Formatter) ParsingAndEscaping(disable bool) {
	p.DisableParsingAndEscaping = disable
}

// Caller 设置是否禁用调用者信息显示
func (p *Formatter) Caller(disable bool) {
	p.DisableCaller = disable
}

// Clone 创建 Formatter 的深拷贝
func (p *Formatter) Clone() Format {
	return &Formatter{
		Module:                    p.Module,
		DisableParsingAndEscaping: p.DisableParsingAndEscaping,
		DisableCaller:             p.DisableCaller,
	}
}

var (
	colorBlack   = []byte("\u001B[30m") // 黑色ANSI代码
	colorRed     = []byte("\u001B[31m") // 红色ANSI代码
	colorGreen   = []byte("\u001B[32m") // 绿色ANSI代码
	colorYellow  = []byte("\u001B[33m") // 黄色ANSI代码
	colorBlue    = []byte("\u001B[34m") // 蓝色ANSI代码
	colorMagenta = []byte("\u001B[35m") // 品红色ANSI代码
	colorCyan    = []byte("\u001B[36m") // 青色ANSI代码
	colorGray    = []byte("\u001B[37m") // 灰色ANSI代码
	colorWhite   = []byte("\u001B[38m") // 白色ANSI代码

	colorEnd = []byte("\u001B[0m") // ANSI重置代码
)

// getColorByLevel 根据日志级别返回对应的ANSI颜色代码
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

// SplitPackageName 拆分完整函数名，返回包路径和函数名
// 示例: "github.com/user/pkg.Function" -> ("github.com/user/pkg", "Function")
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
