package log

import (
	"io"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// --- lazygophers/log Benchmarks ---

func BenchmarkLazyGophersLog_Simple(b *testing.B) {
	logger := New().
		SetLevel(InfoLevel).
		SetOutput(io.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("This is a test log message")
	}
}

func BenchmarkLazyGophersLog_With5Fields(b *testing.B) {
	logger := New().
		SetLevel(InfoLevel).
		SetOutput(io.Discard)
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

func BenchmarkLazyGophersLog_Infof(b *testing.B) {
	logger := New().
		SetLevel(InfoLevel).
		SetOutput(io.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("This is a formatted log message: %s %d", "test", 123)
	}
}

// --- zap.SugaredLogger Benchmarks ---

func BenchmarkZapSugared_Simple(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	logger := zap.New(core).Sugar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("This is a test log message")
	}
}

func BenchmarkZapSugared_With5Fields(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
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

func BenchmarkZapSugared_Infof(b *testing.B) {
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	logger := zap.New(core).Sugar()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infof("This is a formatted log message: %s %d", "test", 123)
	}
}
