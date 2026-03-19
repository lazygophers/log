package log

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"testing"
)

// 1. fastStringify 性能验证
func BenchmarkFastStringify_Types(b *testing.B) {
	tests := map[string]interface{}{
		"string":  "hello world",
		"int":     42,
		"int64":   int64(42),
		"float64": 3.14159,
		"bool":    true,
		"error":   errors.New("test error"),
	}

	for name, v := range tests {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = fastStringify(v)
			}
		})
	}
}

func BenchmarkFastStringify_VSFmt(b *testing.B) {
	value := 42

	b.Run("fastStringify", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fastStringify(value)
		}
	})

	b.Run("fmt.Sprint", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprint(value)
		}
	})
}

// 2. Trace 操作性能
func BenchmarkTrace_Operations(b *testing.B) {
	b.Run("SetTrace", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			SetTrace("test-trace-id")
		}
	})

	b.Run("GetTrace", func(b *testing.B) {
		SetTrace("test-trace-id")
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = GetTrace()
		}
	})

	b.Run("DelTrace", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			SetTrace("test")
			DelTrace()
		}
	})

	b.Run("Concurrent", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				if i%2 == 0 {
					SetTrace(fmt.Sprintf("id-%d", i))
				} else {
					_ = GetTrace()
				}
				i++
			}
		})
	})
}

// 3. Rotator Write 性能
func BenchmarkRotator_Write(b *testing.B) {
	tmpDir := b.TempDir()
	logPath := filepath.Join(tmpDir, "app.log")

	rotator := NewHourlyRotator(logPath, 10*1024*1024, 24)
	defer rotator.Close()

	data := []byte("test log message with some content\n")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = rotator.Write(data)
	}
}

// 4. AsyncWriter 吞吐量
func BenchmarkAsyncWriter_Throughput(b *testing.B) {
	data := []byte("log message\n")

	b.Run("SingleThread", func(b *testing.B) {
		buf := &WriterWrapper{&bytes.Buffer{}}
		asyncWriter := NewAsyncWriter(buf)
		defer asyncWriter.Close()

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = asyncWriter.Write(data)
		}
	})

	b.Run("MultiThread", func(b *testing.B) {
		buf := &WriterWrapper{&bytes.Buffer{}}
		asyncWriter := NewAsyncWriter(buf)
		defer asyncWriter.Close()

		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, _ = asyncWriter.Write(data)
			}
		})
	})
}

// 5. fillCallerInfo 性能
func BenchmarkFillCallerInfo(b *testing.B) {
	logger := newLogger()
	logger.enableCaller = true

	entry := getEntry()
	defer putEntry(entry)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.fillCallerInfo(entry)
	}
}
