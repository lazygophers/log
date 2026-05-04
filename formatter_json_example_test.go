package log

import (
	"testing"

	"github.com/lazygophers/log/constant"
)

func TestJSONFormatter_Usage(t *testing.T) {
	// Example 1: Basic JSON logging
	logger := New()
	logger.Format = &JSONFormatter{}

	// This will output JSON format instead of text
	logger.Info("JSON log message")

	// Example 2: Pretty print JSON
	logger.Format = &JSONFormatter{EnablePrettyPrint: true}
	logger.Info("Pretty JSON log")

	// Example 3: JSON with disabled timestamp
	logger.Format = &JSONFormatter{
		DisableTimestamp: true,
	}
	logger.Info("JSON without timestamp")

	// Example 4: JSON with only essential fields
	logger.Format = &JSONFormatter{
		DisableCaller: true,
		DisableTrace:  true,
	}
	logger.Error("Minimal JSON error log")
}

func TestJSONFormatter_Integration(t *testing.T) {
	// Test integration with standard logger
	logger := New()
	logger.SetLevel(DebugLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Switch to JSON output
	logger.Format = &JSONFormatter{}

	// Log at different levels
	logger.Trace("Trace message in JSON")
	logger.Debug("Debug message in JSON")
	logger.Info("Info message in JSON")
	logger.Warn("Warning message in JSON")
	logger.Error("Error message in JSON")

	// Verify Format interface compatibility
	var _ constant.Format = (*JSONFormatter)(nil)
}

func BenchmarkJSONFormatter_Basic(b *testing.B) {
	logger := New()
	logger.Format = &JSONFormatter{}
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	entry := &Entry{
		Level:   InfoLevel,
		Message: "benchmark message",
		Pid:     12345,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = logger.Format.Format(entry)
	}
}

func BenchmarkJSONFormatter_WithAllFields(b *testing.B) {
	logger := New()
	logger.Format = &JSONFormatter{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entry := &Entry{
			Level:      InfoLevel,
			Message:    "benchmark with all fields",
			Pid:        12345,
			Gid:        67890,
			TraceId:    "trace-benchmark",
			File:       "logger_test.go",
			CallerLine: 42,
			CallerFunc: "BenchmarkJSONFormatter",
		}
		_ = logger.Format.Format(entry)
	}
}

func BenchmarkJSONFormatter_PrettyPrint(b *testing.B) {
	logger := New()
	logger.Format = &JSONFormatter{EnablePrettyPrint: true}

	entry := &Entry{
		Level:   InfoLevel,
		Message: "pretty print benchmark",
		Pid:     12345,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = logger.Format.Format(entry)
	}
}
