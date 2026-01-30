package log

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

// 测试原版 vs 优化版的性能对比
func BenchmarkEntry_Performance_Comparison(b *testing.B) {
	b.Run("OriginalGetPut", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := getEntry()
			entry.Message = "test"
			entry.TraceId = "trace"
			putEntry(entry)
		}
	})

	b.Run("FastGetPut", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			entry := getEntry()
			entry.Message = "test"
			entry.TraceId = "trace"
			putEntry(entry)
		}
	})
}

// 测试Reset方法的性能
func BenchmarkEntry_Reset_Performance(b *testing.B) {
	entry := getEntry()
	// 预设数据以测试重置性能
	entry.TraceId = "test-trace-id-12345"
	entry.File = "/path/to/source/file.go"
	entry.Message = "This is a test log message with some content"
	entry.CallerName = "github.com/test/package"
	entry.CallerDir = "/path/to/caller/directory"
	entry.CallerFunc = "TestFunction"
	entry.PrefixMsg = []byte("PREFIX: ")
	entry.SuffixMsg = []byte(" :SUFFIX")
	entry.Gid = 12345
	entry.CallerLine = 42
	entry.Level = InfoLevel

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entry.Reset()
		// 重新设置数据以保持一致的测试条件
		entry.TraceId = "test-trace-id-12345"
		entry.File = "/path/to/source/file.go"
		entry.Message = "This is a test log message with some content"
		entry.CallerName = "github.com/test/package"
		entry.CallerDir = "/path/to/caller/directory"
		entry.CallerFunc = "TestFunction"
		entry.Gid = 12345
		entry.CallerLine = 42
		entry.Level = InfoLevel
	}
}

// 并发性能测试
func BenchmarkEntry_Concurrent_Performance(b *testing.B) {
	b.Run("OriginalConcurrent", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				entry := getEntry()
				entry.Message = "concurrent test"
				entry.TraceId = "trace-concurrent"
				entry.Gid = 12345
				putEntry(entry)
			}
		})
	})

	b.Run("FastConcurrent", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				entry := getEntry()
				entry.Message = "concurrent test"
				entry.TraceId = "trace-concurrent"
				entry.Gid = 12345
				putEntry(entry)
			}
		})
	})
}

// 内存分配测试
func BenchmarkEntry_Memory_Allocation(b *testing.B) {
	b.Run("PoolOperations", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			entry := getEntry()
			entry.Message = "memory test"
			entry.TraceId = "trace-memory"
			entry.File = "test.go"
			putEntry(entry)
		}

		runtime.GC()
		runtime.ReadMemStats(&m2)

		allocDiff := m2.TotalAlloc - m1.TotalAlloc
		b.ReportMetric(float64(allocDiff)/float64(b.N), "bytes/op")
	})

	b.Run("FastPoolOperations", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			entry := getEntry()
			entry.Message = "memory test"
			entry.TraceId = "trace-memory"
			entry.File = "test.go"
			putEntry(entry)
		}

		runtime.GC()
		runtime.ReadMemStats(&m2)

		allocDiff := m2.TotalAlloc - m1.TotalAlloc
		b.ReportMetric(float64(allocDiff)/float64(b.N), "bytes/op")
	})
}

// 高负载压力测试
func BenchmarkEntry_HighLoad_Performance(b *testing.B) {
	numGoroutines := runtime.NumCPU() * 2
	entriesPerGoroutine := 1000

	b.Run("OriginalHighLoad", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup

			for j := 0; j < numGoroutines; j++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()

					for k := 0; k < entriesPerGoroutine; k++ {
						entry := getEntry()
						entry.Message = "high load test"
						entry.TraceId = "trace-load"
						entry.Gid = int64(id*1000 + k)
						entry.CallerLine = k
						putEntry(entry)
					}
				}(j)
			}

			wg.Wait()
		}
	})

	b.Run("FastHighLoad", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup

			for j := 0; j < numGoroutines; j++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()

					for k := 0; k < entriesPerGoroutine; k++ {
						entry := getEntry()
						entry.Message = "high load test"
						entry.TraceId = "trace-load"
						entry.Gid = int64(id*1000 + k)
						entry.CallerLine = k
						putEntry(entry)
					}
				}(j)
			}

			wg.Wait()
		}
	})
}

// 切片重用性能测试
func BenchmarkEntry_SliceReuse_Performance(b *testing.B) {
	entry := getEntry()
	testData := []byte("test prefix data for slice reuse benchmark")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// 模拟切片使用
		entry.PrefixMsg = append(entry.PrefixMsg, testData...)
		entry.SuffixMsg = append(entry.SuffixMsg, testData...)

		// 重置（会保留切片容量）
		entry.Reset()
	}

	putEntry(entry)
}

// 真实场景模拟测试
func BenchmarkEntry_RealWorld_Scenario(b *testing.B) {
	scenarios := []struct {
		name     string
		message  string
		traceId  string
		file     string
		function string
		level    Level
	}{
		{"ShortLog", "OK", "t1", "a.go", "fn1", InfoLevel},
		{"MediumLog", "Processing request for user authentication", "trace-auth-12345", "auth.go", "AuthenticateUser", InfoLevel},
		{"LongLog", "Failed to connect to database after 3 retries, connection timeout exceeded, fallback to read-only replica", "trace-db-error-78901", "/path/to/db/connection.go", "ConnectWithRetry", ErrorLevel},
	}

	for _, scenario := range scenarios {
		b.Run("Original_"+scenario.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				entry := getEntry()
				entry.Message = scenario.message
				entry.TraceId = scenario.traceId
				entry.File = scenario.file
				entry.CallerFunc = scenario.function
				entry.Level = scenario.level
				entry.Time = time.Now()
				entry.Gid = int64(i)
				putEntry(entry)
			}
		})

		b.Run("Fast_"+scenario.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				entry := getEntry()
				entry.Message = scenario.message
				entry.TraceId = scenario.traceId
				entry.File = scenario.file
				entry.CallerFunc = scenario.function
				entry.Level = scenario.level
				entry.Time = time.Now()
				entry.Gid = int64(i)
				putEntry(entry)
			}
		})
	}
}

// CPU 使用效率测试
func BenchmarkEntry_CPU_Efficiency(b *testing.B) {
	b.Run("OriginalCPUEfficiency", func(b *testing.B) {
		b.ReportAllocs()

		start := time.Now()

		for i := 0; i < b.N; i++ {
			entry := getEntry()

			// 模拟常见的日志字段设置
			entry.Message = "CPU efficiency test message"
			entry.TraceId = "trace-cpu-test"
			entry.File = "cpu_test.go"
			entry.CallerFunc = "TestCPUEfficiency"
			entry.CallerLine = 100
			entry.Level = InfoLevel
			entry.Gid = int64(i)

			// 模拟一些字符串操作
			if i%2 == 0 {
				entry.PrefixMsg = append(entry.PrefixMsg, []byte("CPU: ")...)
			}

			putEntry(entry)
		}

		duration := time.Since(start)
		b.ReportMetric(float64(duration.Nanoseconds())/float64(b.N), "ns/op")
	})

	b.Run("FastCPUEfficiency", func(b *testing.B) {
		b.ReportAllocs()

		start := time.Now()

		for i := 0; i < b.N; i++ {
			entry := getEntry()

			// 模拟常见的日志字段设置
			entry.Message = "CPU efficiency test message"
			entry.TraceId = "trace-cpu-test"
			entry.File = "cpu_test.go"
			entry.CallerFunc = "TestCPUEfficiency"
			entry.CallerLine = 100
			entry.Level = InfoLevel
			entry.Gid = int64(i)

			// 模拟一些字符串操作
			if i%2 == 0 {
				entry.PrefixMsg = append(entry.PrefixMsg, []byte("CPU: ")...)
			}

			putEntry(entry)
		}

		duration := time.Since(start)
		b.ReportMetric(float64(duration.Nanoseconds())/float64(b.N), "ns/op")
	})
}
