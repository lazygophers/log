// Package log 提供日志初始化和管理功能
package log

import (
	"os"
)

var (
	// std 是标准日志记录器实例
	std = newLogger()

	// pid 存储当前进程ID
	pid = os.Getpid()
)

// Pid 返回当前进程ID
func Pid() int {
	return pid
}

// New 创建并返回一个新的Logger实例
func New() *Logger {
	return newLogger()
}

// SetLevel 设置标准日志记录器的日志级别
// 参数:
//
//	level: 要设置的日志级别
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetLevel(level Level) *Logger {
	return std.SetLevel(level)
}

// GetLevel 返回标准日志记录器的当前日志级别
// 返回值:
//
//	Level: 当前日志级别
func GetLevel() Level {
	return std.Level()
}

// Sync 将标准日志记录器(std)中所有缓冲的日志条目写入其输出目的地。
//
// 在应用程序即将退出时，调用此函数至关重要。由于日志库通常为了性能而采用异步或缓冲写入策略，
// 若不显式调用 Sync，可能会有部分日志尚在内存缓冲区中，未及持久化，从而导致日志丢失。
//
// 此函数确保了日志的完整性和持久性，是优雅关闭(graceful shutdown)流程中的一个关键步骤。
func Sync() {
	std.Sync()
}

// Clone 创建并返回一个标准日志记录器(std)的深度副本。
//
// "深度副本"意味着新创建的Logger实例拥有与原始std实例完全独立的配置，
// 包括日志级别、输出目的地、格式化选项等。后续对新实例的任何修改(如通过SetLevel, SetSuffixMsg等)
// 都不会影响到原始的std实例，反之亦然。
//
// 此功能对于需要基于通用配置创建临时或特定用途的日志记录器非常有用，
// 例如，在特定的业务模块中增加额外的日志前缀或调整日志级别，而无需修改全局日志配置。
//
// 返回值:
//   - *Logger: 一个新的、与std配置相同但完全独立的日志记录器实例。
func Clone() *Logger {
	return std.Clone()
}

// CloneToCtx 创建并返回一个支持上下文(context)传播的标准日志记录器副本。
//
// 这个函数不仅创建了一个独立的Logger实例(类似于Clone)，更将其封装在一个`LoggerWithCtx`结构中。
// `LoggerWithCtx`的设计目的是为了与Go的`context.Context`机制深度集成，
// 从而可以方便地将日志记录器实例与请求作用域或其它上下文关联起来。
//
// 这使得在复杂的并发或分布式系统中，可以轻松地传递和获取与当前上下文绑定的日志记录器，
// 极大地简化了日志的追踪和管理。
//
// 返回值:
//   - *LoggerWithCtx: 一个新的、支持上下文的日志记录器实例。
func CloneToCtx() *LoggerWithCtx {
	return std.CloneToCtx()
}

// SetCallerDepth 设置日志调用堆栈深度
// 参数:
//
//	callerDepth: 调用堆栈深度
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetCallerDepth(callerDepth int) *Logger {
	return std.SetCallerDepth(callerDepth)
}

// SetPrefixMsg 设置日志消息前缀
// 参数:
//
//	prefixMsg: 前缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetPrefixMsg(prefixMsg string) *Logger {
	return std.SetPrefixMsg(prefixMsg)
}

// AppendPrefixMsg 追加日志消息前缀
// 参数:
//
//	prefixMsg: 要追加的前缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func AppendPrefixMsg(prefixMsg string) *Logger {
	return std.AppendPrefixMsg(prefixMsg)
}

// SetSuffixMsg 设置日志消息后缀
// 参数:
//
//	suffixMsg: 后缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func SetSuffixMsg(suffixMsg string) *Logger {
	return std.SetSuffixMsg(suffixMsg)
}

// AppendSuffixMsg 追加日志消息后缀
// 参数:
//
//	suffixMsg: 要追加的后缀消息
//
// 返回值:
//
//	*Logger: 标准日志记录器实例(支持链式调用)
func AppendSuffixMsg(suffixMsg string) *Logger {
	return std.AppendSuffixMsg(suffixMsg)
}

// ParsingAndEscaping 控制是否对日志消息进行JSON解析和特殊字符转义。
//
// 当启用时(disable=false)，日志库会尝试将字符串形式的日志消息解析为JSON对象，
// 并对其中的特殊字符(如引号、反斜杠)进行转义，以确保最终输出的日志是格式正确的JSON。
//
// 当禁用时(disable=true)，日志消息将作为普通字符串处理，不进行任何解析或转义。
// 这在需要记录非JSON格式的原始文本或二进制数据时非常有用。
//
// 参数:
//   - disable: 传入`true`以禁用该功能，`false`则启用。
//
// 返回值:
//   - *Logger: 返回标准日志记录器实例(std)，以支持链式调用。
func ParsingAndEscaping(disable bool) *Logger {
	return std.ParsingAndEscaping(disable)
}

// Caller 控制是否在日志条目中包含调用者的信息（文件名和行号）。
//
// 开启此功能 (disable=false) 对于调试非常有用，它可以帮助快速定位日志输出的具体代码位置。
// 但是，获取调用者信息会带来一定的性能开销，因为它需要运行时检查调用栈。
//
// 在生产环境中，尤其是在高并发或对性能有极致要求的场景下，建议禁用此功能 (disable=true) 以获得最佳性能。
//
// 参数:
//   - disable: 传入`true`以禁用调用者信息，`false`则启用。
//
// 返回值:
//   - *Logger: 返回标准日志记录器实例(std)，以支持链式调用。
func Caller(disable bool) *Logger {
	return std.Caller(disable)
}
