package log

import (
	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
)

// ZapHook 实现 zapcore.WriteSyncer 接口，用于将 zap 日志重定向到我们的日志系统
type ZapHook struct {
	logger *Logger
}

// NewZapHook 创建一个新的 ZapHook 实例
func NewZapHook(log *Logger) *ZapHook {
	return &ZapHook{logger: log}
}

// Write 实现 io.Writer 接口
func (zh *ZapHook) Write(p []byte) (n int, err error) {
	// 直接写入到底层 logger
	return zh.logger.out.Write(p)
}

// Sync 实现 zapcore.WriteSyncer 接口
func (zh *ZapHook) Sync() error {
	zh.logger.Sync()
	return nil
}

// createZapHook 创建一个 zap core hook，用于将 zap 的日志条目适配到当前日志系统中。
// 这个钩子函数会从池中获取一个 Entry 实例，填充必要信息，然后写入目标日志。
func createZapHook(log *Logger) func(entry zapcore.Entry) error {
	return func(entry zapcore.Entry) error {
		// 从对象池中获取一个 Entry 对象，以减少内存分配
		logEntry := entryPool.Get().(*Entry)
		// 使用 defer 确保 Entry 对象在函数结束时被重置并放回池中
		defer func() {
			logEntry.Reset()
			entryPool.Put(logEntry)
		}()

		// 填充日志条目的基础信息
		logEntry.Gid = goid.Get()                 // 获取当前 goroutine ID
		logEntry.TraceId = getTrace(logEntry.Gid) // 获取追踪 ID
		logEntry.Time = entry.Time                // 设置日志时间
		logEntry.Message = entry.Message          // 设置日志消息
		logEntry.SuffixMsg = log.SuffixMsg        // 设置日志后缀
		logEntry.PrefixMsg = log.PrefixMsg        // 设置日志前缀

		// 填充调用者信息
		logEntry.File = entry.Caller.File           // 文件路径
		logEntry.CallerLine = entry.Caller.Line     // 行号
		logEntry.CallerName = entry.Caller.Function // 函数名

		// 解析包名和函数名
		logEntry.CallerDir, logEntry.CallerFunc = SplitPackageName(entry.Caller.Function)

		// 将 zap 的日志级别转换为内部定义的日志级别
		switch entry.Level {
		case zapcore.DebugLevel:
			logEntry.Level = DebugLevel
		case zapcore.InfoLevel:
			logEntry.Level = InfoLevel
		case zapcore.WarnLevel:
			logEntry.Level = WarnLevel
		case zapcore.ErrorLevel:
			logEntry.Level = ErrorLevel
		case zapcore.DPanicLevel:
			logEntry.Level = PanicLevel
		case zapcore.PanicLevel:
			logEntry.Level = PanicLevel
		case zapcore.FatalLevel:
			logEntry.Level = FatalLevel
		default:
			// 对于未知的日志级别，默认为 ErrorLevel
			logEntry.Level = ErrorLevel
		}

		// 格式化日志条目并写入
		log.write(logEntry.Level, log.Format.Format(logEntry))

		return nil
	}
}
