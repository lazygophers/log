package log

import (
	"bytes"
	"testing"
)

func TestGetBuffer(t *testing.T) {
	buf := GetBuffer()
	if buf == nil {
		t.Fatal("expected to get a buffer, but got nil")
	}
	if buf.Len() != 0 {
		t.Errorf("expected buffer length 0, but got %d", buf.Len())
	}
	// Note: The capacity of a buffer from a sync.Pool is not guaranteed to be 0.
	// The buffer is reused, so it retains its allocated memory. We only need to
	// ensure its length is 0.
}

func TestPutBuffer(t *testing.T) {
	// 1. 调用 GetBuffer() 获取一个缓冲区实例
	buf := GetBuffer()
	// 2. 向缓冲区写入一些数据
	buf.WriteString("test data")

	// 3. 调用 PutBuffer() 将缓冲区归还
	PutBuffer(buf)

	// 4. 再次调用 GetBuffer() 获取一个缓冲区，并检查其内容是否为空
	buf2 := GetBuffer()
	if buf2.Len() != 0 {
		t.Errorf("expected buffer to be reset, but got length %d", buf2.Len())
	}

	// 5. 测试 PutBuffer(nil) 的情况，确保程序不会因此崩溃
	PutBuffer(nil)
}

func BenchmarkPool(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := GetBuffer()
			buf.WriteString("benchmark data")
			PutBuffer(buf)
		}
	})
}

func BenchmarkAlloc(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = new(bytes.Buffer)
		}
	})
}
