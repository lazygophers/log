package log

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetSetDelTrace 测试针对当前 goroutine 的基本 Get/Set/Del 操作。
func TestGetSetDelTrace(t *testing.T) {
	// 初始时，追踪信息应为空
	assert.Empty(t, GetTrace(), "Initial trace should be empty")

	// 设置一个追踪ID
	SetTrace("test-id")
	assert.Equal(t, "test-id", GetTrace(), "Trace ID should be 'test-id'")

	// 删除追踪信息
	DelTrace()
	assert.Empty(t, GetTrace(), "Trace should be empty after deletion")

	// 设置一个空的追踪ID，此时应自动生成一个新的
	SetTrace()
	assert.NotEmpty(t, GetTrace(), "A new trace ID should be generated when set with no arguments")
	assert.Len(t, GetTrace(), 16, "Generated trace ID should have a length of 16")
	DelTrace()
}

// TestGetSetDelTraceWithGID 测试针对指定 goroutine ID 的基本 Get/Set/Del 操作。
func TestGetSetDelTraceWithGID(t *testing.T) {
	const gid int64 = 12345

	// 初始时，指定 GID 的追踪信息应为空
	assert.Empty(t, GetTraceWithGID(gid), "Initial trace for GID should be empty")

	// 为指定 GID 设置一个追踪ID
	SetTraceWithGID(gid, "test-id-gid")
	assert.Equal(t, "test-id-gid", GetTraceWithGID(gid), "Trace ID for GID should be 'test-id-gid'")

	// 删除指定 GID 的追踪信息
	DelTraceWithGID(gid)
	assert.Empty(t, GetTraceWithGID(gid), "Trace for GID should be empty after deletion")

	// 为指定 GID 设置一个空的追踪ID，此时应自动生成一个新的
	SetTraceWithGID(gid)
	assert.NotEmpty(t, GetTraceWithGID(gid), "A new trace ID should be generated for GID")
	assert.Len(t, GetTraceWithGID(gid), 16, "Generated trace ID for GID should have a length of 16")
	DelTraceWithGID(gid)
}

// TestDisableTrace 测试禁用追踪功能的开关是否有效。
func TestDisableTrace(t *testing.T) {
	DisableTrace = true
	defer func() { DisableTrace = false }() // 测试结束后恢复

	SetTrace("no-trace")
	assert.Empty(t, GetTrace(), "GetTrace should return empty when tracing is disabled")
}

// TestGenTraceId 测试追踪ID的生成函数。
func TestGenTraceId(t *testing.T) {
	traceId := GenTraceId()
	assert.Len(t, traceId, 16, "Generated trace ID should be 16 characters long")
}

// TestConcurrencySafety 测试 traceMap 的并发安全性。
func TestConcurrencySafety(t *testing.T) {
	var wg sync.WaitGroup
	numGoroutines := 100
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()

			// 每个 goroutine 设置自己的追踪ID
			traceID := fmt.Sprintf("goroutine-%d", i)
			SetTrace(traceID)

			// 验证当前 goroutine 的追踪ID是否设置正确
			assert.Equal(t, traceID, GetTrace(), "Trace ID should be correctly set for each goroutine")

			// 模拟一些工作并再次验证
			assert.Equal(t, traceID, GetTrace(), "Trace ID should remain consistent within the same goroutine")

			DelTrace()
			assert.Empty(t, GetTrace(), "Trace ID should be empty after deletion in concurrent goroutine")
		}(i)
	}

	wg.Wait()
}

// BenchmarkSetTrace 测试 SetTrace 函数的性能。
func BenchmarkSetTrace(b *testing.B) {
	// 循环 b.N 次，b.N 由 testing 框架自动调整
	for i := 0; i < b.N; i++ {
		SetTrace("benchmark-trace-id")
	}
}

// BenchmarkGetTrace 测试 GetTrace 函数的性能。
func BenchmarkGetTrace(b *testing.B) {
	SetTrace("benchmark-trace-id") // 准备测试数据
	b.ResetTimer()                 // 重置计时器，忽略准备数据的时间
	for i := 0; i < b.N; i++ {
		GetTrace()
	}
}

// BenchmarkSetGetTraceParallel 测试在并发场景下 SetTrace 和 GetTrace 的综合性能。
func BenchmarkSetGetTraceParallel(b *testing.B) {
	// RunParallel 会创建多个 goroutine 并发执行测试
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SetTrace() // 设置自动生成的追踪ID
			GetTrace() // 获取追踪ID
		}
	})
}

// Example_package 展示了 trace 功能的基本用法。
//
// 这个可执行的示例代码演示了如何设置一个自定义追踪ID，
// 接着获取并验证它，最后删除该ID并确认其已被移除。
func Example_package() {
	// 步骤 1: 设置一个自定义的追踪ID
	myTraceID := "my-awesome-trace-id"
	SetTrace(myTraceID)

	// 步骤 2: 获取并打印当前 goroutine 的追踪ID
	retrievedTraceID := GetTrace()
	fmt.Println(retrievedTraceID)

	// 步骤 3: 清理当前 goroutine 的追踪ID
	DelTrace()

	// 步骤 4: 再次获取，以确认追踪ID已被成功删除
	afterDeleteTraceID := GetTrace()
	fmt.Printf("Trace ID after deletion: '%s'\n", afterDeleteTraceID)

	// Output:
	// my-awesome-trace-id
	// Trace ID after deletion: ''
}
