package log

import (
	"testing"
	"time"
)

// TestNewEntry 用于测试 NewEntry 函数的正确性
func TestNewEntry(t *testing.T) {
	// 调用 NewEntry() 函数创建一个 Entry 实例
	entry := NewEntry()

	// 1. 验证返回的 *Entry 实例不为 nil
	if entry == nil {
		t.Fatal("NewEntry() returned nil, expected a non-nil entry")
	}

	// 2. 验证 entry.Pid 字段是否已正确初始化 (等于 log.pid)
	// 由于测试文件和源文件在同一个包下，可以直接访问包内变量 pid
	if entry.Pid != pid {
		t.Errorf("entry.Pid = %d; want %d", entry.Pid, pid)
	}
}

// TestEntry_Reset 用于测试 Reset 方法是否能正确重置 Entry 的字段
func TestEntry_Reset(t *testing.T) {
	// 1. 创建一个 Entry 实例并为其字段赋上非零值
	entry := &Entry{
		Gid:        12345,
		TraceId:    "test-trace-id",
		File:       "test_file.go",
		Message:    "this is a test message",
		CallerName: "main",
		CallerDir:  "/tmp/test",
		CallerFunc: "main.main",
		PrefixMsg:  []byte("prefix"),
		SuffixMsg:  []byte("suffix"),
		// 以下字段不应被 Reset 修改
		Pid:        pid,
		Time:       time.Now(),
		Level:      InfoLevel,
		CallerLine: 42,
	}

	// 2. 调用 Reset() 方法
	entry.Reset()

	// 3. 验证可重置字段是否已变回其零值
	if entry.Gid != 0 {
		t.Errorf("entry.Gid = %d; want 0", entry.Gid)
	}
	if entry.TraceId != "" {
		t.Errorf("entry.TraceId = %q; want \"\"", entry.TraceId)
	}
	if entry.File != "" {
		t.Errorf("entry.File = %q; want \"\"", entry.File)
	}
	if entry.Message != "" {
		t.Errorf("entry.Message = %q; want \"\"", entry.Message)
	}
	if entry.CallerName != "" {
		t.Errorf("entry.CallerName = %q; want \"\"", entry.CallerName)
	}
	if entry.CallerDir != "" {
		t.Errorf("entry.CallerDir = %q; want \"\"", entry.CallerDir)
	}
	if entry.CallerFunc != "" {
		t.Errorf("entry.CallerFunc = %q; want \"\"", entry.CallerFunc)
	}
	if len(entry.PrefixMsg) != 0 {
		t.Errorf("len(entry.PrefixMsg) = %d; want 0", len(entry.PrefixMsg))
	}
	if len(entry.SuffixMsg) != 0 {
		t.Errorf("len(entry.SuffixMsg) = %d; want 0", len(entry.SuffixMsg))
	}

	// 4. 验证不应被重置的字段保持不变
	if entry.Pid != pid {
		t.Errorf("entry.Pid was reset to %d; should have remained %d", entry.Pid, pid)
	}
	if entry.Time.IsZero() {
		t.Error("entry.Time was reset; should not have been")
	}
	if entry.Level != InfoLevel {
		t.Errorf("entry.Level was reset to %v; should have remained %v", entry.Level, InfoLevel)
	}
	if entry.CallerLine != 42 {
		t.Errorf("entry.CallerLine was reset to %d; should have remained 42", entry.CallerLine)
	}
}

// BenchmarkNewEntry is a benchmark test for the NewEntry function.
func BenchmarkNewEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		NewEntry()
	}
}

// BenchmarkEntry_Reset a benchmark for the Reset method.
func BenchmarkEntry_Reset(b *testing.B) {
	entry := &Entry{
		Gid:        12345,
		TraceId:    "test-trace-id",
		File:       "test_file.go",
		Message:    "this is a test message",
		CallerName: "main",
		CallerDir:  "/tmp/test",
		CallerFunc: "main.main",
		PrefixMsg:  []byte("prefix"),
		SuffixMsg:  []byte("suffix"),
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entry.Reset()
	}
}

// BenchmarkEntryPool a benchmark for the entryPool.
func BenchmarkEntryPool(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			entry := entryPool.Get().(*Entry)
			entry.Message = "benchmark test"
			entryPool.Put(entry)
		}
	})
}
