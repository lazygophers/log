package log

import (
	"fmt"
	"path"
	"strconv"
	"strings"
)

// Format 定义了日志格式化接口。
type Format interface {
	// Format 将日志条目格式化为字节数组
	// entry: 需要格式化的日志条目
	Format(entry *Entry) []byte
}

// FormatFull 扩展了基础格式化接口，提供更多控制功能。
type FormatFull interface {
	Format // 继承基础格式化接口

	// ParsingAndEscaping 控制是否禁用消息解析和转义
	// disable: true 表示禁用，false 表示启用
	ParsingAndEscaping(disable bool)

	Caller(disable bool)

	Clone() Format
}

// Formatter 实现了 FormatFull 接口，提供默认日志格式化功能。
type Formatter struct {
	Module                    string // 日志所属模块名
	DisableParsingAndEscaping bool   // 是否禁用消息解析和转义
	DisableCaller             bool   // 是否禁用调用者信息
}

// format 是内部格式化方法，处理单行日志格式。
// entry: 需要格式化的日志条目
// 返回: 格式化后的字节数组
func (p *Formatter) format(entry *Entry) []byte {
	// 从缓冲池获取一个字节缓冲，用完后放回
	b := GetBuffer()
	defer PutBuffer(b)

	// 写入前缀消息
	if len(entry.PrefixMsg) > 0 {
		b.Write(entry.PrefixMsg)
		b.Write([]byte(" "))
	}

	// 写入进程ID和协程ID
	b.WriteString(fmt.Sprintf("(%d.%d) ", entry.Pid, entry.Gid))
	// 写入格式化的时间戳
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999Z07:00"))
	// 根据日志级别获取颜色
	color := getColorByLevel(entry.Level)
	// 写入颜色代码
	b.Write(color)
	// 写入日志级别
	b.Write([]byte(" ["))
	b.WriteString(entry.Level.String())
	b.Write([]byte("] "))
	// 写入颜色结束代码
	b.Write(colorEnd)
	// 写入日志正文
	b.WriteString(strings.TrimSpace(entry.Message))

	// 如果未禁用调用者信息或存在TraceID，则写入额外信息
	if !p.DisableCaller || entry.TraceId != "" {
		b.Write(colorCyan)
		b.WriteString(" [ ")
		// 如果未禁用调用者信息，则写入文件名、行号和函数名
		if !p.DisableCaller {
			b.WriteString(path.Join(entry.CallerDir, path.Base(entry.File)))
			b.Write([]byte(":"))
			b.WriteString(strconv.Itoa(entry.CallerLine))
			b.Write([]byte(" "))
			b.WriteString(entry.CallerFunc)
			b.Write([]byte(" "))
		}
		// 如果存在TraceID，则写入
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

	// 写入后缀消息
	if len(entry.SuffixMsg) > 0 {
		b.Write(entry.SuffixMsg)
	}
	// 写入换行符
	b.WriteByte('\n')
	return b.Bytes()
}

// Format 实现 Format 接口，处理多行日志消息。
// entry: 需要格式化的日志条目
// 返回: 格式化后的字节数组
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

// ParsingAndEscaping 设置是否禁用消息解析和转义。
// disable: true 表示禁用，false 表示启用
func (p *Formatter) ParsingAndEscaping(disable bool) {
	p.DisableParsingAndEscaping = disable
}

// Caller 设置是否禁用调用者信息显示。
func (p *Formatter) Caller(disable bool) {
	p.DisableCaller = disable
}

// Clone 创建 Formatter 的深拷贝。
// 返回: 新的 Formatter 实例
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
	colorEnd     = []byte("\u001B[0m")
)

// getColorByLevel 根据日志级别获取对应的终端颜色代码。
// level: 日志级别
// 返回: 终端颜色代码字节数组
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

// SplitPackageName 分割完整的包路径为目录路径和函数名。
// f: 完整的包路径字符串 (如 "github.com/user/pkg.Function")
// 返回:
//
//	callDir - 包目录路径
//	callFunc - 函数名
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
