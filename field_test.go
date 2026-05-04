package log

import (
	"testing"
	"time"
)

func TestLogger_Infow_Basic(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Infow("User logged in",
		"user", "alice",
		"age", 30,
	)
}

func TestLogger_Infow_Single(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Infow("Action", "action", "login")
}

func TestLogger_MultipleFields(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Infow("Complex log",
		"component", "auth",
		"action", "login",
		"user", "bob",
		"attempts", 3,
	)
}

func TestLogger_Errorw(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Errorw("Request failed",
		"path", "/api/users",
		"status", 500,
	)
}

func TestLogger_Debugw(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Debugw("Debug info",
		"key", "value",
		"count", 100,
	)
}

func TestLogger_AllTypes(t *testing.T) {
	now := time.Now()

	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Infow("All field types",
		"str", "test",
		"int", 42,
		"int8", int8(8),
		"int16", int16(16),
		"int32", int32(32),
		"int64", int64(64),
		"uint", uint(42),
		"uint8", uint8(8),
		"uint16", uint16(16),
		"uint32", uint32(32),
		"uint64", uint64(64),
		"float32", float32(3.14),
		"float64", 6.28,
		"bool", true,
		"time", now,
		"duration", time.Second,
		"err", error(nil),
		"any", map[string]int{"a": 1},
	)
}

func TestLogger_Infow_Empty(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Infow("No fields")
	logger.Infow("Also no fields")
}

func TestLogger_AllLevelw(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Tracew("Trace", "level", "trace")
	logger.Debugw("Debug", "level", "debug")
	logger.Infow("Info", "level", "info")
	logger.Warnw("Warn", "level", "warn")
	logger.Errorw("Error", "level", "error")
}

func TestLogger_Escaping(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Infow("Special characters",
		"msg", "value with \"quotes\"",
		"path", "C:\\Users\\test",
	)
}

func TestLogger_OddArgs(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Odd number of args - last one should be printed as-is
	logger.Infow("Odd args",
		"key1", "value1",
		"key2", // missing value
	)
}

func TestLogger_NilValues(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	logger.Infow("Nil values",
		"err", nil,
		"ptr", (*int)(nil),
	)
}

func BenchmarkLogger_Infow(b *testing.B) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Infow("Performance test",
			"field0", "value0",
			"field1", "value1",
			"field2", "value2",
			"field3", "value3",
			"field4", "value4",
			"field5", "value5",
			"field6", "value6",
			"field7", "value7",
			"field8", "value8",
			"field9", "value9",
		)
	}
}
