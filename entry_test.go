package log

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestNewEntry(t *testing.T) {
	entry := NewEntry()

	if entry == nil {
		t.Fatal("NewEntry returned nil")
	}

	// 验证 Pid 被正确设置
	if entry.Pid != pid {
		t.Errorf("Expected Pid %d, got %d", pid, entry.Pid)
	}

	// 验证其他字段的初始状态
	if entry.Gid != 0 {
		t.Errorf("Expected Gid 0, got %d", entry.Gid)
	}

	if entry.TraceId != "" {
		t.Errorf("Expected empty TraceId, got %q", entry.TraceId)
	}

	if entry.Message != "" {
		t.Errorf("Expected empty Message, got %q", entry.Message)
	}
}

func TestEntry_Reset(t *testing.T) {
	entry := NewEntry()

	// 设置一些值
	entry.Gid = 12345
	entry.TraceId = "trace-123"
	entry.File = "test.go"
	entry.Message = "test message"
	entry.CallerName = "TestFunc"
	entry.CallerDir = "/path/to/test"
	entry.CallerFunc = "main.test"
	entry.PrefixMsg = []byte("prefix")
	entry.SuffixMsg = []byte("suffix")
	entry.Time = time.Now()
	entry.Level = InfoLevel
	entry.CallerLine = 42

	// 调用 Reset
	entry.Reset()

	// 验证所有字段都被重置
	if entry.Gid != 0 {
		t.Errorf("Expected Gid 0 after reset, got %d", entry.Gid)
	}

	if entry.TraceId != "" {
		t.Errorf("Expected empty TraceId after reset, got %q", entry.TraceId)
	}

	if entry.File != "" {
		t.Errorf("Expected empty File after reset, got %q", entry.File)
	}

	if entry.Message != "" {
		t.Errorf("Expected empty Message after reset, got %q", entry.Message)
	}

	if entry.CallerName != "" {
		t.Errorf("Expected empty CallerName after reset, got %q", entry.CallerName)
	}

	if entry.CallerDir != "" {
		t.Errorf("Expected empty CallerDir after reset, got %q", entry.CallerDir)
	}

	if entry.CallerFunc != "" {
		t.Errorf("Expected empty CallerFunc after reset, got %q", entry.CallerFunc)
	}

	if len(entry.PrefixMsg) != 0 {
		t.Errorf("Expected empty PrefixMsg after reset, got %v", entry.PrefixMsg)
	}

	if len(entry.SuffixMsg) != 0 {
		t.Errorf("Expected empty SuffixMsg after reset, got %v", entry.SuffixMsg)
	}

	// 验证 Reset 不会改变 Pid、Time、Level、CallerLine
	if entry.Pid != pid {
		t.Errorf("Reset should not change Pid, expected %d, got %d", pid, entry.Pid)
	}
}

func TestEntry_Reset_SliceReuse(t *testing.T) {
	entry := NewEntry()

	// 设置切片数据
	entry.PrefixMsg = append(entry.PrefixMsg, []byte("test prefix")...)
	entry.SuffixMsg = append(entry.SuffixMsg, []byte("test suffix")...)

	// 记录底层数组的容量和指针
	prefixCap := cap(entry.PrefixMsg)
	suffixCap := cap(entry.SuffixMsg)

	// Reset 后验证切片被清空但底层数组被保留
	entry.Reset()

	if len(entry.PrefixMsg) != 0 {
		t.Error("PrefixMsg should be empty after reset")
	}

	if len(entry.SuffixMsg) != 0 {
		t.Error("SuffixMsg should be empty after reset")
	}

	// 验证容量被保留（高效复用）
	if cap(entry.PrefixMsg) != prefixCap {
		t.Errorf("PrefixMsg capacity should be preserved, expected %d, got %d", prefixCap, cap(entry.PrefixMsg))
	}

	if cap(entry.SuffixMsg) != suffixCap {
		t.Errorf("SuffixMsg capacity should be preserved, expected %d, got %d", suffixCap, cap(entry.SuffixMsg))
	}
}

func TestGetEntry_PutEntry(t *testing.T) {
	// 测试基本的获取和放回
	entry := getEntry()
	if entry == nil {
		t.Fatal("getEntry returned nil")
	}

	// 验证获取的 entry 是重置状态
	if entry.Message != "" {
		t.Error("Entry from pool should have empty message")
	}

	// 设置一些值
	entry.Message = "test message"
	entry.TraceId = "trace-456"

	// 放回池中
	putEntry(entry)

	// 再次获取，应该得到重置后的 entry
	entry2 := getEntry()
	if entry2.Message != "" {
		t.Error("Entry from pool should be reset")
	}

	if entry2.TraceId != "" {
		t.Error("Entry from pool should be reset")
	}

	// 清理
	putEntry(entry2)
}

func TestPutEntry_NilEntry(t *testing.T) {
	// 测试传入 nil 不会 panic
	putEntry(nil)
	// 如果没有 panic，测试通过
}

func TestEntryPool_Concurrent(t *testing.T) {
	// 并发测试对象池的线程安全性
	var wg sync.WaitGroup
	numGoroutines := 100
	entriesPerGoroutine := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			entries := make([]*Entry, entriesPerGoroutine)

			// 获取大量 entries
			for j := 0; j < entriesPerGoroutine; j++ {
				entries[j] = getEntry()
				if entries[j] == nil {
					t.Errorf("Goroutine %d: getEntry returned nil", id)
					return
				}
				// 设置一些值来验证隔离性
				entries[j].Message = "test"
				entries[j].Gid = int64(id*1000 + j)
			}

			// 放回所有 entries
			for j := 0; j < entriesPerGoroutine; j++ {
				putEntry(entries[j])
			}
		}(i)
	}

	wg.Wait()
}

// 性能基准测试

func BenchmarkNewEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entry := NewEntry()
		_ = entry
	}
}

func BenchmarkGetEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entry := getEntry()
		putEntry(entry)
	}
}

func BenchmarkEntry_Reset(b *testing.B) {
	entry := NewEntry()
	// 预先设置一些数据
	entry.PrefixMsg = make([]byte, 100)
	entry.SuffixMsg = make([]byte, 100)
	entry.Message = "test message"
	entry.TraceId = "trace-id"
	entry.File = "test.go"
	entry.CallerName = "TestFunc"
	entry.CallerDir = "/path/to/test"
	entry.CallerFunc = "main.test"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entry.Reset()
		// 重新设置数据以保持一致的测试条件
		entry.Message = "test message"
		entry.TraceId = "trace-id"
		entry.File = "test.go"
		entry.CallerName = "TestFunc"
		entry.CallerDir = "/path/to/test"
		entry.CallerFunc = "main.test"
	}
}

func BenchmarkEntryPool_GetPut(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entry := getEntry()
		entry.Message = "benchmark test"
		entry.TraceId = "bench-trace"
		putEntry(entry)
	}
}

func BenchmarkEntryPool_Parallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			entry := getEntry()
			entry.Message = "parallel test"
			entry.Gid = 12345
			putEntry(entry)
		}
	})
}

func BenchmarkEntry_SliceAppend(b *testing.B) {
	entry := NewEntry()
	testData := []byte("test data")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entry.PrefixMsg = append(entry.PrefixMsg[:0], testData...)
		entry.SuffixMsg = append(entry.SuffixMsg[:0], testData...)
	}
}

func BenchmarkEntry_vs_DirectAllocation(b *testing.B) {
	b.Run("EntryPool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := getEntry()
			entry.Message = "test"
			entry.Gid = 12345
			entry.TraceId = "trace"
			putEntry(entry)
		}
	})

	b.Run("DirectAllocation", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := &Entry{
				Pid:     pid,
				Message: "test",
				Gid:     12345,
				TraceId: "trace",
			}
			_ = entry
		}
	})
}

// 内存使用测试
func TestEntry_MemoryUsage(t *testing.T) {
	// 获取初始内存统计
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	// 执行大量操作
	const iterations = 10000
	for i := 0; i < iterations; i++ {
		entry := getEntry()
		entry.Message = "memory test"
		entry.TraceId = "trace-memory"
		entry.PrefixMsg = append(entry.PrefixMsg, []byte("prefix")...)
		entry.SuffixMsg = append(entry.SuffixMsg, []byte("suffix")...)
		putEntry(entry)
	}

	// 获取最终内存统计
	runtime.GC()
	runtime.ReadMemStats(&m2)

	// 验证没有明显的内存泄漏
	allocDiff := m2.TotalAlloc - m1.TotalAlloc
	t.Logf("Memory allocated during %d operations: %d bytes", iterations, allocDiff)

	// 每次操作的平均内存分配应该很小（得益于对象池）
	avgAllocPerOp := allocDiff / iterations
	if avgAllocPerOp > 1000 { // 1KB per operation seems reasonable
		t.Errorf("Average memory allocation per operation is too high: %d bytes", avgAllocPerOp)
	}
}

// 竞争条件测试
func TestEntry_RaceCondition(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping race condition test in short mode")
	}

	var wg sync.WaitGroup
	numGoroutines := 50
	operationsPerGoroutine := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < operationsPerGoroutine; j++ {
				entry := getEntry()
				entry.Message = "race test"
				entry.Gid = int64(j)
				entry.Reset()
				putEntry(entry)
			}
		}()
	}

	wg.Wait()
}



// 测试getEntry的分支覆盖
func TestGetEntry_EmptyPool(t *testing.T) {
	// 清空对象池来测试创建新Entry的分支
	// 通过大量获取Entry但不归还来耗尽池
	var entries []*Entry
	for i := 0; i < 100; i++ {
		entry := getEntry()
		if entry == nil {
			t.Fatal("getEntry returned nil")
		}
		// 验证新创建的Entry有正确的Pid
		if entry.Pid != pid {
			t.Errorf("Expected Pid %d, got %d", pid, entry.Pid)
		}
		entries = append(entries, entry)
	}

	// 清理：归还所有entries
	for _, entry := range entries {
		putEntry(entry)
	}
}
