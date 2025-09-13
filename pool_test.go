package log

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetPutBuffer(t *testing.T) {
	// 测试获取和返回缓冲区
	buf := GetBuffer()

	if buf == nil {
		t.Fatal("GetBuffer returned nil")
	}

	// 写入一些数据
	buf.WriteString("test data")

	if buf.String() != "test data" {
		t.Errorf("Expected 'test data', got %q", buf.String())
	}

	// 返回到池中
	PutBuffer(buf)

	// 再次获取，应该是重置后的缓冲区
	buf2 := GetBuffer()

	if buf2.Len() != 0 {
		t.Error("Buffer should be reset when returned from pool")
	}

	// 清理
	PutBuffer(buf2)
}

func TestPutBuffer_Nil_Coverage(t *testing.T) {
	// 测试PutBuffer处理nil的情况（pool.go:38-40）

	// 这应该不会panic或出现错误
	PutBuffer(nil)

	// 测试通过说明nil检查工作正常
}

func TestBufferPool_Reuse(t *testing.T) {
	// 测试缓冲区池的重用功能
	buffers := make([]*bytes.Buffer, 10)

	// 获取多个缓冲区
	for i := 0; i < 10; i++ {
		buffers[i] = GetBuffer()
		if buffers[i] == nil {
			t.Fatalf("GetBuffer returned nil at index %d", i)
		}
	}

	// 向每个缓冲区写入唯一数据
	for i, buf := range buffers {
		buf.WriteString(fmt.Sprintf("data_%d", i))
	}

	// 验证数据
	for i, buf := range buffers {
		expected := fmt.Sprintf("data_%d", i)
		if buf.String() != expected {
			t.Errorf("Expected %q, got %q", expected, buf.String())
		}
	}

	// 将所有缓冲区返回池中
	for _, buf := range buffers {
		PutBuffer(buf)
	}

	// 再次获取缓冲区，验证已被重置
	for i := 0; i < 5; i++ {
		buf := GetBuffer()
		if buf.Len() != 0 {
			t.Errorf("Buffer %d should be reset, but has length %d", i, buf.Len())
		}
		PutBuffer(buf)
	}
}
