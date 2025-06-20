package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewEntry 测试日志条目创建功能
func TestNewEntry(t *testing.T) {
	entry := NewEntry()
	assert.Equal(t, pid, entry.Pid, "进程ID应正确初始化")
	assert.Equal(t, int64(0), entry.Gid, "协程ID应默认为0")
	assert.Empty(t, entry.TraceId, "追踪ID应为空字符串")
	assert.True(t, entry.Time.Before(time.Now()) || entry.Time.Equal(time.Now()), "时间戳应在当前时间附近")
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

	assert.Equal(t, int64(0), entry.Gid, "协程ID应重置为0")
	assert.Empty(t, entry.TraceId, "追踪ID应重置为空")
	assert.Empty(t, entry.File, "文件路径应重置为空")
	assert.Empty(t, entry.Message, "消息内容应重置为空")
	assert.Empty(t, entry.CallerName, "调用函数全名应重置为空")
	assert.Empty(t, entry.CallerDir, "调用文件目录应重置为空")
	assert.Empty(t, entry.CallerFunc, "调用函数名应重置为空")
	assert.Len(t, entry.PrefixMsg, 0, "PrefixMsg应重置为空切片")
	assert.Len(t, entry.SuffixMsg, 0, "SuffixMsg应重置为空切片")
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
