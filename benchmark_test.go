package log

import (
	"io"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// --- lazygophers/log Benchmarks ---

// BenchmarkLazyGophersLog_Simple 基准测试：使用 lazygophers/log 记录一条简单的日志消息。
func BenchmarkLazyGophersLog_Simple(b *testing.B) {
	logger := New().
		SetLevel(InfoLevel).
		SetOutput(io.Discard) // 将日志输出到 io.Discard 以避免不必要的 IO 开销
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("This is a test log message")
	}
}

// BenchmarkLazyGophersLog_With5Fields 基准测试：使用 lazygophers/log 记录一条带 5 个字段的日志消息。
func BenchmarkLazyGophersLog_With5Fields(b *testing.B) {
	logger := New().
		SetLevel(InfoLevel).
		SetOutput(io.Discard) // 将日志输出到 io.Discard 以避免不必要的 IO 开销
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("This is a test log message with fields",
			"string_field", "test",
			"int_field", 123,
			"float_field", 123.456,
			"bool_field", true,
			"time_field", time.Now(),
		)
	}
}

// BenchmarkLazyGophersLog_Infof 基准测试：使用 lazygophers/log 记录一条格式化的日志消息。
func BenchmarkLazyGophersLog_Infof(b *testing.B) {
	logger := New().
		SetLevel(InfoLevel).
		SetOutput(io.Discard) // 将日志输出到 io.Discard 以避免不必要的 IO 开销
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("This is a formatted log message: %s %d", "test", 123)
	}
}

// --- zap.SugaredLogger Benchmarks ---

// BenchmarkZapSugared_Simple 基准测试：使用 zap.SugaredLogger 记录一条简单的日志消息。
func BenchmarkZapSugared_Simple(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard), // 将日志输出到 io.Discard 以避免不必要的 IO 开销
		zapcore.InfoLevel,
	)
	logger := zap.New(core).Sugar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("This is a test log message")
	}
}

// BenchmarkZapSugared_With5Fields 基准测试：使用 zap.SugaredLogger 记录一条带 5 个字段的日志消息。
func BenchmarkZapSugared_With5Fields(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard), // 将日志输出到 io.Discard 以避免不必要的 IO 开销
		zapcore.InfoLevel,
	)
	logger := zap.New(core).Sugar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infow("This is a test log message with fields",
			"string_field", "test",
			"int_field", 123,
			"float_field", 123.456,
			"bool_field", true,
			"time_field", time.Now(),
		)
	}
}

// BenchmarkZapSugared_Infof 基准测试：使用 zap.SugaredLogger 记录一条格式化的日志消息。
func BenchmarkZapSugared_Infof(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard), // 将日志输出到 io.Discard 以避免不必要的 IO 开销
		zapcore.InfoLevel,
	)
	logger := zap.New(core).Sugar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("This is a formatted log message: %s %d", "test", 123)
	}
}
