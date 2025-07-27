package log

import (
	"os"
	"testing"
)

func TestNewEntry(t *testing.T) {
	entry := NewEntry()

	if entry == nil {
		t.Fatal("NewEntry() should not return nil")
	}

	expectedPid := os.Getpid()
	if entry.Pid != expectedPid {
		t.Errorf("Expected Pid to be %d, but got %d", expectedPid, entry.Pid)
	}
}

func TestEntry_Reset(t *testing.T) {
	entry := &Entry{
		Gid:        123,
		TraceId:    "test-trace-id",
		File:       "test_file.go",
		Message:    "test message",
		CallerName: "test.caller",
		CallerDir:  "/path/to/caller",
		CallerFunc: "TestFunction",
		PrefixMsg:  []byte("prefix"),
		SuffixMsg:  []byte("suffix"),
	}

	entry.Reset()

	if entry.Gid != 0 {
		t.Errorf("Expected Gid to be 0, but got %d", entry.Gid)
	}
	if entry.TraceId != "" {
		t.Errorf("Expected TraceId to be empty, but got %s", entry.TraceId)
	}
	if entry.File != "" {
		t.Errorf("Expected File to be empty, but got %s", entry.File)
	}
	if entry.Message != "" {
		t.Errorf("Expected Message to be empty, but got %s", entry.Message)
	}
	if entry.CallerName != "" {
		t.Errorf("Expected CallerName to be empty, but got %s", entry.CallerName)
	}
	if entry.CallerDir != "" {
		t.Errorf("Expected CallerDir to be empty, but got %s", entry.CallerDir)
	}
	if entry.CallerFunc != "" {
		t.Errorf("Expected CallerFunc to be empty, but got %s", entry.CallerFunc)
	}
	if len(entry.PrefixMsg) != 0 {
		t.Errorf("Expected PrefixMsg to be empty, but got %v", entry.PrefixMsg)
	}
	if len(entry.SuffixMsg) != 0 {
		t.Errorf("Expected SuffixMsg to be empty, but got %v", entry.SuffixMsg)
	}
}

func BenchmarkEntryPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entry := entryPool.Get().(*Entry)
		entryPool.Put(entry)
	}
}
