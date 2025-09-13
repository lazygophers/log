package log

import (
	"runtime"
	"testing"
	"unsafe"
)

// 深度分析Entry结构体的内存布局和性能瓶颈
func TestEntry_MemoryLayout_Analysis(t *testing.T) {
	entry := &Entry{}

	t.Logf("Entry struct size: %d bytes", unsafe.Sizeof(*entry))
	t.Logf("Entry pointer size: %d bytes", unsafe.Sizeof(entry))

	// 分析各字段的内存偏移
	t.Logf("Field offsets:")
	t.Logf("  Pid: %d", unsafe.Offsetof(entry.Pid))
	t.Logf("  Gid: %d", unsafe.Offsetof(entry.Gid))
	t.Logf("  TraceId: %d", unsafe.Offsetof(entry.TraceId))
	t.Logf("  Time: %d", unsafe.Offsetof(entry.Time))
	t.Logf("  Level: %d", unsafe.Offsetof(entry.Level))
	t.Logf("  File: %d", unsafe.Offsetof(entry.File))
	t.Logf("  Message: %d", unsafe.Offsetof(entry.Message))
	t.Logf("  CallerName: %d", unsafe.Offsetof(entry.CallerName))
	t.Logf("  CallerLine: %d", unsafe.Offsetof(entry.CallerLine))
	t.Logf("  CallerDir: %d", unsafe.Offsetof(entry.CallerDir))
	t.Logf("  CallerFunc: %d", unsafe.Offsetof(entry.CallerFunc))
	t.Logf("  PrefixMsg: %d", unsafe.Offsetof(entry.PrefixMsg))
	t.Logf("  SuffixMsg: %d", unsafe.Offsetof(entry.SuffixMsg))

	// 计算内存对齐浪费
	totalFieldSize := int(unsafe.Sizeof(entry.Pid)) +
		int(unsafe.Sizeof(entry.Gid)) +
		int(unsafe.Sizeof(entry.TraceId)) +
		int(unsafe.Sizeof(entry.Time)) +
		int(unsafe.Sizeof(entry.Level)) +
		int(unsafe.Sizeof(entry.File)) +
		int(unsafe.Sizeof(entry.Message)) +
		int(unsafe.Sizeof(entry.CallerName)) +
		int(unsafe.Sizeof(entry.CallerLine)) +
		int(unsafe.Sizeof(entry.CallerDir)) +
		int(unsafe.Sizeof(entry.CallerFunc)) +
		int(unsafe.Sizeof(entry.PrefixMsg)) +
		int(unsafe.Sizeof(entry.SuffixMsg))

	actualSize := int(unsafe.Sizeof(*entry))
	padding := actualSize - totalFieldSize

	t.Logf("Total field sizes: %d bytes", totalFieldSize)
	t.Logf("Actual struct size: %d bytes", actualSize)
	t.Logf("Padding/alignment waste: %d bytes (%.1f%%)", padding, float64(padding)/float64(actualSize)*100)
}

// 测试不同内存布局的性能影响
func BenchmarkEntry_MemoryAccess_Patterns(b *testing.B) {
	entry := getEntry()
	defer putEntry(entry)

	b.Run("SequentialAccess", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// 按字段定义顺序访问（缓存友好）
			entry.Pid = i
			entry.Gid = int64(i)
			entry.TraceId = "trace"
			// time.Time 跳过，避免时间调用开销
			entry.Level = InfoLevel
			entry.File = "test.go"
			entry.Message = "test"
			entry.CallerName = "test"
			entry.CallerLine = i
			entry.CallerDir = "test"
			entry.CallerFunc = "test"
		}
	})

	b.Run("RandomAccess", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// 随机顺序访问（缓存不友好）
			entry.Message = "test"
			entry.Pid = i
			entry.CallerFunc = "test"
			entry.Gid = int64(i)
			entry.File = "test.go"
			entry.CallerLine = i
			entry.TraceId = "trace"
			entry.Level = InfoLevel
			entry.CallerName = "test"
			entry.CallerDir = "test"
		}
	})
}

// 分析sync.Pool的性能开销
func BenchmarkEntry_SyncPool_Overhead(b *testing.B) {
	b.Run("DirectAllocation", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := &Entry{Pid: pid}
			entry.Message = "test"
			entry.TraceId = "trace"
			_ = entry
		}
	})

	b.Run("SyncPoolGet", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := entryPool.Get().(*Entry)
			entry.Message = "test"
			entry.TraceId = "trace"
			entryPool.Put(entry)
		}
	})

	b.Run("OptimizedGetPut", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := getEntry()
			entry.Message = "test"
			entry.TraceId = "trace"
			putEntry(entry)
		}
	})
}

// 分析Reset方法的各种实现方式
func BenchmarkEntry_Reset_Strategies(b *testing.B) {
	entry := getEntry()
	defer putEntry(entry)

	// 设置测试数据
	setupEntry := func() {
		entry.Gid = 12345
		entry.CallerLine = 42
		entry.Level = InfoLevel
		entry.TraceId = "trace-12345"
		entry.File = "test.go"
		entry.Message = "test message"
		entry.CallerName = "TestFunc"
		entry.CallerDir = "/path/to/test"
		entry.CallerFunc = "main.test"
		entry.PrefixMsg = []byte("prefix")
		entry.SuffixMsg = []byte("suffix")
	}

	b.Run("CurrentOptimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			setupEntry()
			entry.Reset()
		}
	})

	b.Run("IndividualAssign", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			setupEntry()
			// 逐个赋值（原始方式）
			entry.Gid = 0
			entry.CallerLine = 0
			entry.Level = 0
			entry.TraceId = ""
			entry.File = ""
			entry.Message = ""
			entry.CallerName = ""
			entry.CallerDir = ""
			entry.CallerFunc = ""
			entry.PrefixMsg = entry.PrefixMsg[:0]
			entry.SuffixMsg = entry.SuffixMsg[:0]
		}
	})

	b.Run("StructReset", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			setupEntry()
			// 结构体重置（会分配新内存）
			pid := entry.Pid
			prefixMsg := entry.PrefixMsg[:0]
			suffixMsg := entry.SuffixMsg[:0]
			*entry = Entry{
				Pid:       pid,
				PrefixMsg: prefixMsg,
				SuffixMsg: suffixMsg,
			}
		}
	})
}

// 分析类型断言开销
func BenchmarkEntry_TypeAssertion_Overhead(b *testing.B) {
	pool := &entryPool

	b.Run("WithTypeAssertion", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := pool.Get().(*Entry)
			pool.Put(entry)
		}
	})

	b.Run("DirectTyped", func(b *testing.B) {
		// 无法避免类型断言，这里只是为了对比
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			obj := pool.Get()
			entry, ok := obj.(*Entry)
			if !ok {
				b.Fatal("type assertion failed")
			}
			pool.Put(entry)
		}
	})
}

// 分析字符串重置的不同策略
func BenchmarkEntry_StringReset_Strategies(b *testing.B) {
	entry := getEntry()
	defer putEntry(entry)

	b.Run("BatchAssignment", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// 当前的批量赋值方式
			entry.TraceId, entry.File, entry.Message = "", "", ""
			entry.CallerName, entry.CallerDir, entry.CallerFunc = "", "", ""
		}
	})

	b.Run("IndividualAssignment", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// 逐个赋值
			entry.TraceId = ""
			entry.File = ""
			entry.Message = ""
			entry.CallerName = ""
			entry.CallerDir = ""
			entry.CallerFunc = ""
		}
	})
}

// CPU 缓存行分析
func TestEntry_CacheLine_Analysis(t *testing.T) {
	const cacheLineSize = 64 // 大多数现代CPU的缓存行大小

	entrySize := unsafe.Sizeof(Entry{})
	entriesPerCacheLine := cacheLineSize / entrySize

	t.Logf("Entry size: %d bytes", entrySize)
	t.Logf("Cache line size: %d bytes", cacheLineSize)
	t.Logf("Entries per cache line: %d", entriesPerCacheLine)

	if entrySize > cacheLineSize {
		t.Logf("WARNING: Entry size exceeds cache line size by %d bytes", entrySize-cacheLineSize)
		t.Logf("This may cause cache misses during field access")
	}
}

// 内存分配模式分析
func BenchmarkEntry_AllocationPatterns(b *testing.B) {
	b.Run("StackAllocation", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// 栈分配（编译器可能优化为栈分配）
			entry := Entry{Pid: pid}
			entry.Message = "test"
			_ = entry
		}
	})

	b.Run("HeapAllocation", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			// 强制堆分配
			entry := &Entry{Pid: pid}
			entry.Message = "test"
			runtime.KeepAlive(entry)
		}
	})

	b.Run("PooledAllocation", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := getEntry()
			entry.Message = "test"
			putEntry(entry)
		}
	})
}
