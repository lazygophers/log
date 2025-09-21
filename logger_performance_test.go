package log

import (
	"bytes"
	"io"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Logger 性能基准测试

// 测试基础日志记录性能
func BenchmarkLogger_BasicLogging(b *testing.B) {
	buf := &bytes.Buffer{}
	logger := newLogger()
	logger.SetOutput(buf)
	logger.SetLevel(InfoLevel)

	b.Run("Info", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("Infof", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Infof("test message %d", i)
		}
	})

	b.Run("Debug_Filtered", func(b *testing.B) {
		logger.SetLevel(InfoLevel) // Debug会被过滤
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Debug("debug message")
		}
	})
}

// 测试不同输出目标的性能影响
func BenchmarkLogger_OutputTargets(b *testing.B) {
	b.Run("BytesBuffer", func(b *testing.B) {
		buf := &bytes.Buffer{}
		logger := newLogger()
		logger.SetOutput(buf)
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("DevNull", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(io.Discard)
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("NilWriter", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(&NilWriter{})
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})
}

// 测试不同格式化器的性能
func BenchmarkLogger_Formatters(b *testing.B) {
	buf := &bytes.Buffer{}

	b.Run("DefaultFormatter", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(buf)
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
			buf.Reset()
		}
	})

	b.Run("DisabledCallerFormatter", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(buf)
		logger.SetLevel(InfoLevel)
		logger.Caller(false) // 禁用调用者信息

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
			buf.Reset()
		}
	})

	b.Run("DisabledParsingFormatter", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(buf)
		logger.SetLevel(InfoLevel)
		logger.ParsingAndEscaping(false)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
			buf.Reset()
		}
	})
}

// 测试级别检查性能
func BenchmarkLogger_LevelChecking(b *testing.B) {
	logger := newLogger()
	logger.SetOutput(io.Discard)

	b.Run("EnabledLevel", func(b *testing.B) {
		logger.SetLevel(InfoLevel)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if logger.levelEnabled(InfoLevel) {
				// 模拟日志处理 - 这里故意为空以测试级别检查性能
				_ = i
			}
		}
	})

	b.Run("DisabledLevel", func(b *testing.B) {
		logger.SetLevel(ErrorLevel)
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if logger.levelEnabled(InfoLevel) {
				// 模拟日志处理 - 这里故意为空以测试级别检查性能
				_ = i
			}
		}
	})
}

// 并发性能测试
func BenchmarkLogger_Concurrent(b *testing.B) {
	logger := newLogger()
	logger.SetOutput(&NilWriter{})
	logger.SetLevel(InfoLevel)

	b.Run("SingleGoroutine", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("MultipleGoroutines", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info("concurrent test")
			}
		})
	})
}

// 内存分配模式测试
func BenchmarkLogger_MemoryPatterns(b *testing.B) {
	b.Run("ShortMessages", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(&NilWriter{})
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("OK")
		}
	})

	b.Run("MediumMessages", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(&NilWriter{})
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("Processing user request with authentication token validation")
		}
	})

	b.Run("LongMessages", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(&NilWriter{})
		logger.SetLevel(InfoLevel)
		longMsg := "This is a very long log message that contains detailed information about the operation being performed, including multiple parameters, execution context, and diagnostic data that might be useful for debugging and monitoring purposes."

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info(longMsg)
		}
	})
}

// 测试前缀和后缀对性能的影响
func BenchmarkLogger_PrefixSuffix(b *testing.B) {
	logger := newLogger()
	logger.SetOutput(&NilWriter{})
	logger.SetLevel(InfoLevel)

	b.Run("NoPrefixSuffix", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("WithPrefix", func(b *testing.B) {
		logger.SetPrefixMsg("PREFIX: ")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("WithSuffix", func(b *testing.B) {
		logger.SetPrefixMsg("")
		logger.SetSuffixMsg(" :SUFFIX")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("WithBoth", func(b *testing.B) {
		logger.SetPrefixMsg("PREFIX: ")
		logger.SetSuffixMsg(" :SUFFIX")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})
}

// 复杂场景性能测试
func BenchmarkLogger_ComplexScenarios(b *testing.B) {
	b.Run("HighVolumeLogging", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(&NilWriter{})
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()

		var wg sync.WaitGroup
		numGoroutines := runtime.NumCPU()

		for i := 0; i < b.N; i++ {
			for j := 0; j < numGoroutines; j++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					for k := 0; k < 100; k++ {
						logger.Infof("goroutine %d iteration %d", id, k)
					}
				}(j)
			}
			wg.Wait()
		}
	})

	b.Run("MixedLevels", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(&NilWriter{})
		logger.SetLevel(DebugLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Debug("debug info")
			logger.Info("processing")
			logger.Warn("warning occurred")
			logger.Error("error encountered")
		}
	})
}

// 实际使用场景模拟
func BenchmarkLogger_RealWorldScenarios(b *testing.B) {
	logger := newLogger()
	logger.SetOutput(&NilWriter{})
	logger.SetLevel(InfoLevel)

	b.Run("WebRequestLogging", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Infof("HTTP %s %s %d %dms", "GET", "/api/users", 200, 25)
		}
	})

	b.Run("DatabaseOperations", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Infof("DB Query: %s, Duration: %dms, Rows: %d", "SELECT * FROM users WHERE id = ?", 15, 1)
		}
	})

	b.Run("ErrorReporting", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Errorf("Failed to connect to database: %s, retry %d/3", "connection timeout", 1)
		}
	})
}

// 内存使用分析
func BenchmarkLogger_MemoryUsage(b *testing.B) {
	logger := newLogger()
	logger.SetOutput(&NilWriter{})
	logger.SetLevel(InfoLevel)

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		logger.Info("memory usage test")
	}

	runtime.GC()
	runtime.ReadMemStats(&m2)

	allocDiff := m2.TotalAlloc - m1.TotalAlloc
	b.ReportMetric(float64(allocDiff)/float64(b.N), "bytes/op")
}

// CPU使用效率测试
func BenchmarkLogger_CPUEfficiency(b *testing.B) {
	logger := newLogger()
	logger.SetOutput(&NilWriter{})
	logger.SetLevel(InfoLevel)

	b.ReportAllocs()

	start := time.Now()
	for i := 0; i < b.N; i++ {
		logger.Infof("CPU efficiency test %d", i)
	}
	duration := time.Since(start)

	b.ReportMetric(float64(duration.Nanoseconds())/float64(b.N), "ns/op")
}

// NilWriter 用于性能测试，丢弃所有写入
type NilWriter struct{}

func (w *NilWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (w *NilWriter) Sync() error {
	return nil
}

// BufferedWriter 用于测试缓冲写入性能
type BufferedWriter struct {
	buf []byte
}

func (w *BufferedWriter) Write(p []byte) (n int, err error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}

func (w *BufferedWriter) Reset() {
	w.buf = w.buf[:0]
}

func (w *BufferedWriter) Len() int {
	return len(w.buf)
}

// 测试不同Writer实现的性能
func BenchmarkLogger_WriterImplementations(b *testing.B) {
	b.Run("NilWriter", func(b *testing.B) {
		logger := newLogger()
		logger.SetOutput(&NilWriter{})
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
		}
	})

	b.Run("BufferedWriter", func(b *testing.B) {
		logger := newLogger()
		writer := &BufferedWriter{}
		logger.SetOutput(writer)
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
			if i%1000 == 0 {
				writer.Reset() // 定期清空缓冲区
			}
		}
	})

	b.Run("BytesBuffer", func(b *testing.B) {
		logger := newLogger()
		buf := &bytes.Buffer{}
		logger.SetOutput(buf)
		logger.SetLevel(InfoLevel)

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			logger.Info("test message")
			if i%1000 == 0 {
				buf.Reset() // 定期清空缓冲区
			}
		}
	})
}
