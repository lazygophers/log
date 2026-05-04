package log

import (
	"testing"
	"time"
)

// TestToJSONEntryCoverage 测试 toJSONEntry 的所有分支
func TestToJSONEntryCoverage(t *testing.T) {
	t.Run("toJSONEntry_with_zero_time", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level: InfoLevel,
			Message: "test",
			Time:  time.Time{}, // 零值时间
			Pid:   123,
		}

		je := formatter.toJSONEntry(entry)

		if je.Time != "" {
			t.Error("Time should be empty for zero time")
		}
	})

	t.Run("toJSONEntry_with_disabled_timestamp", func(t *testing.T) {
		formatter := &JSONFormatter{DisableTimestamp: true}
		entry := &Entry{
			Level: InfoLevel,
			Message: "test",
			Time:  time.Now(),
			Pid:   123,
		}

		je := formatter.toJSONEntry(entry)

		if je.Time != "" {
			t.Error("Time should be empty when timestamp disabled")
		}
	})

	t.Run("toJSONEntry_with_gid_and_traceid", func(t *testing.T) {
		formatter := &JSONFormatter{} // 默认不禁用 trace
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
			Gid:     456,
			TraceId: "trace-123",
		}

		je := formatter.toJSONEntry(entry)

		if je.Gid != 456 {
			t.Error("Gid should be set")
		}
		if je.TraceID != "trace-123" {
			t.Error("TraceID should be set")
		}
	})

	t.Run("toJSONEntry_with_disabled_trace", func(t *testing.T) {
		formatter := &JSONFormatter{DisableTrace: true}
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
			Gid:     456,
			TraceId: "trace-123",
		}

		je := formatter.toJSONEntry(entry)

		if je.Gid != 0 {
			t.Error("Gid should be zero when trace disabled")
		}
		if je.TraceID != "" {
			t.Error("TraceID should be empty when trace disabled")
		}
	})

	t.Run("toJSONEntry_with_zero_gid_and_empty_traceid", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
			Gid:     0,   // 零值
			TraceId: "", // 空字符串
		}

		je := formatter.toJSONEntry(entry)

		if je.Gid != 0 {
			t.Error("Gid should be zero")
		}
		if je.TraceID != "" {
			t.Error("TraceID should be empty")
		}
	})

	t.Run("toJSONEntry_with_empty_file", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:      InfoLevel,
			Message:    "test",
			Time:       time.Now(),
			Pid:        123,
			File:       "", // 空文件名
			CallerLine: 42,
			CallerFunc: "main.Func",
		}

		je := formatter.toJSONEntry(entry)

		if je.CallerFile != "" {
			t.Error("CallerFile should be empty when File is empty")
		}
	})

	t.Run("toJSONEntry_with_disabled_caller", func(t *testing.T) {
		formatter := &JSONFormatter{DisableCaller: true}
		entry := &Entry{
			Level:      InfoLevel,
			Message:    "test",
			Time:       time.Now(),
			Pid:        123,
			File:       "file.go",
			CallerLine: 42,
			CallerFunc: "main.Func",
		}

		je := formatter.toJSONEntry(entry)

		if je.CallerFile != "" {
			t.Error("CallerFile should be empty when caller disabled")
		}
	})

	t.Run("toJSONEntry_with_full_caller_info", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:      InfoLevel,
			Message:    "test",
			Time:       time.Now(),
			Pid:        123,
			File:       "path/to/file.go",
			CallerLine: 42,
			CallerFunc: "main.Func",
			CallerDir:  "path/to",
			CallerName: "Func",
		}

		je := formatter.toJSONEntry(entry)

		if je.CallerFile != "path/to/file.go" {
			t.Error("CallerFile should be set")
		}
		if je.CallerLine != 42 {
			t.Error("CallerLine should be set")
		}
		if je.CallerFunc != "main.Func" {
			t.Error("CallerFunc should be set")
		}
		if je.CallerDir != "path/to" {
			t.Error("CallerDir should be set")
		}
		if je.CallerName != "Func" {
			t.Error("CallerName should be set")
		}
	})

	t.Run("toJSONEntry_with_prefix_and_suffix", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:     InfoLevel,
			Message:   "test",
			Time:      time.Now(),
			Pid:       123,
			PrefixMsg: []byte("[PREFIX]"),
			SuffixMsg: []byte("[SUFFIX]"),
		}

		je := formatter.toJSONEntry(entry)

		if je.PrefixMsg != "[PREFIX]" {
			t.Error("PrefixMsg should be set")
		}
		if je.SuffixMsg != "[SUFFIX]" {
			t.Error("SuffixMsg should be set")
		}
	})

	t.Run("toJSONEntry_with_empty_prefix_suffix", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:     InfoLevel,
			Message:   "test",
			Time:      time.Now(),
			Pid:       123,
			PrefixMsg: []byte{},
			SuffixMsg: []byte{},
		}

		je := formatter.toJSONEntry(entry)

		if je.PrefixMsg != "" {
			t.Error("PrefixMsg should be empty")
		}
		if je.SuffixMsg != "" {
			t.Error("SuffixMsg should be empty")
		}
	})

	t.Run("toJSONEntry_with_fields", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
			Fields: []KV{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: 123},
				{Key: "key3", Value: true},
			},
		}

		je := formatter.toJSONEntry(entry)

		if len(je.Fields) != 3 {
			t.Errorf("Fields should have 3 entries, got %d", len(je.Fields))
		}
		if je.Fields["key1"] != "value1" {
			t.Error("key1 should be value1")
		}
		if je.Fields["key2"] != 123 {
			t.Error("key2 should be 123")
		}
		if je.Fields["key3"] != true {
			t.Error("key3 should be true")
		}
	})

	t.Run("toJSONEntry_with_empty_fields", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
			Fields:  []KV{},
		}

		je := formatter.toJSONEntry(entry)

		if je.Fields != nil {
			t.Error("Fields should be nil when empty")
		}
	})

	t.Run("toJSONEntry_with_all_fields_enabled", func(t *testing.T) {
		formatter := &JSONFormatter{} // 所有字段默认启用
		entry := &Entry{
			Level:      InfoLevel,
			Message:    "test message",
			Time:       time.Now(),
			Pid:        123,
			Gid:        456,
			TraceId:    "trace-789",
			File:       "path/to/file.go",
			CallerLine: 42,
			CallerFunc: "main.TestFunc",
			CallerDir:  "path/to",
			CallerName: "TestFunc",
			PrefixMsg:  []byte("[P]"),
			SuffixMsg:  []byte("[S]"),
			Fields: []KV{
				{Key: "field1", Value: "value1"},
			},
		}

		je := formatter.toJSONEntry(entry)

		// 验证所有字段都被设置
		if je.Time == "" {
			t.Error("Time should be set")
		}
		if je.Gid != 456 {
			t.Error("Gid should be set")
		}
		if je.TraceID != "trace-789" {
			t.Error("TraceID should be set")
		}
		if je.CallerFile != "path/to/file.go" {
			t.Error("CallerFile should be set")
		}
		if je.PrefixMsg != "[P]" {
			t.Error("PrefixMsg should be set")
		}
		if je.SuffixMsg != "[S]" {
			t.Error("SuffixMsg should be set")
		}
		if len(je.Fields) != 1 {
			t.Error("Fields should have 1 entry")
		}
	})
}
