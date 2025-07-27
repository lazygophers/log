package log

import (
	"bytes"
	"testing"
)

// TestGetBuffer 测试从池中获取缓冲区。
func TestGetBuffer(t *testing.T) {
	// 从池中获取一个缓冲区。
	buf := GetBuffer()
	// 验证获取到的缓冲区不应为 nil。
	if buf == nil {
		t.Fatal("期望获取一个缓冲区，但得到的是 nil")
	}
	// 验证缓冲区的初始长度应为 0。
	if buf.Len() != 0 {
		t.Errorf("期望缓冲区长度为 0，但得到的是 %d", buf.Len())
	}
	// 注意：从 sync.Pool 获取的缓冲区的容量（Capacity）不保证为 0。
	// 因为缓冲区是复用的，它会保留之前分配的内存。我们只需要确保其长度（Length）为 0 即可。
}

// TestPutBuffer 测试将缓冲区归还到池中。
func TestPutBuffer(t *testing.T) {
	// 1. 从池中获取一个缓冲区实例。
	buf := GetBuffer()
	// 2. 向缓冲区写入一些测试数据。
	buf.WriteString("test data")

	// 3. 将缓冲区归还到池中。
	PutBuffer(buf)

	// 4. 再次从池中获取一个缓冲区，并检查其内容是否已被重置。
	buf2 := GetBuffer()
	if buf2.Len() != 0 {
		t.Errorf("期望缓冲区被重置，但其长度为 %d", buf2.Len())
	}

	// 5. 测试归还 nil 缓冲区的情况，确保程序不会因此引发 panic。
	PutBuffer(nil)
}

// BenchmarkPool 对使用 sync.Pool 的缓冲区分配进行基准测试。
func BenchmarkPool(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 从池中获取缓冲区，写入数据，然后归还。
			buf := GetBuffer()
			buf.WriteString("benchmark data")
			PutBuffer(buf)
		}
	})
}

// BenchmarkAlloc 对不使用池，直接通过 new(bytes.Buffer) 进行内存分配的方式进行基准测试。
func BenchmarkAlloc(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// 每次迭代都创建一个新的缓冲区实例。
			_ = new(bytes.Buffer)
		}
	})
}
