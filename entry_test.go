package log

import (
	"testing"
	"time"
)

// TestNewEntry 测试日志条目创建功能
func TestNewEntry(t *testing.T) {
	entry := NewEntry()
	if entry.Pid != pid {
		t.Errorf("进程ID应正确初始化: got %v, want %v", entry.Pid, pid)
	}
	if entry.Gid != int64(0) {
		t.Errorf("协程ID应默认为0: got %v", entry.Gid)
	}
	if entry.TraceId != "" {
		t.Errorf("追踪ID应为空字符串: got %q", entry.TraceId)
	}
	if !entry.Time.Before(time.Now()) && !entry.Time.Equal(time.Now()) {
		t.Errorf("时间戳应在当前时间附近: got %v", entry.Time)
	}
}

// TestEntryReset 测试日志条目重置功能
func TestEntryReset(t *testing.T) {
	entry := &Entry{
		Pid:        123,
		Gid:        456,
		TraceId:    "test-trace",
		Time:       time.Now(),
		Level:      InfoLevel,
		File:       "/test/file.go",
		Message:    "test message",
		CallerName: "main.testFunc",
		CallerLine: 42,
		CallerDir:  "/test",
		CallerFunc: "testFunc",
		PrefixMsg:  []byte("prefix"),
		SuffixMsg:  []byte("suffix"),
	}

	entry.Reset()

	if entry.Gid != int64(0) {
		t.Errorf("协程ID应重置为0: got %v", entry.Gid)
	}
	if entry.TraceId != "" {
		t.Errorf("追踪ID应重置为空: got %q", entry.TraceId)
	}
	if entry.File != "" {
		t.Errorf("文件路径应重置为空: got %q", entry.File)
	}
	if entry.Message != "" {
		t.Errorf("消息内容应重置为空: got %q", entry.Message)
	}
	if entry.CallerName != "" {
		t.Errorf("调用函数全名应重置为空: got %q", entry.CallerName)
	}
	if entry.CallerDir != "" {
		t.Errorf("调用文件目录应重置为空: got %q", entry.CallerDir)
	}
	if entry.CallerFunc != "" {
		t.Errorf("调用函数名应重置为空: got %q", entry.CallerFunc)
	}
	if len(entry.PrefixMsg) != 0 {
		t.Errorf("PrefixMsg应重置为空切片: got %d elements", len(entry.PrefixMsg))
	}
	if len(entry.SuffixMsg) != 0 {
		t.Errorf("SuffixMsg应重置为空切片: got %d elements", len(entry.SuffixMsg))
	}
}

// BenchmarkNewEntry 测试日志条目创建性能
func BenchmarkNewEntry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewEntry()
	}
}

// BenchmarkEntryReset 测试日志条目重置性能
func BenchmarkEntryReset(b *testing.B) {
	entry := NewEntry()
	for i := 0; i < b.N; i++ {
		entry.Reset()
	}
}
