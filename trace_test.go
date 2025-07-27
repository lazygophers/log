package log

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetSetDelTrace tests basic Get/Set/Del operations for the current goroutine.
func TestGetSetDelTrace(t *testing.T) {
	// Initially, trace should be empty
	assert.Empty(t, GetTrace(), "Initial trace should be empty")

	// Set a trace ID
	SetTrace("test-id")
	assert.Equal(t, "test-id", GetTrace(), "Trace ID should be 'test-id'")

	// Delete the trace
	DelTrace()
	assert.Empty(t, GetTrace(), "Trace should be empty after deletion")

	// Set an empty trace ID, should generate a new one
	SetTrace()
	assert.NotEmpty(t, GetTrace(), "A new trace ID should be generated when set with no arguments")
	assert.Len(t, GetTrace(), 16, "Generated trace ID should have a length of 16")
	DelTrace()
}

// TestGetSetDelTraceWithGID tests basic Get/Set/Del operations for a specific goroutine ID.
func TestGetSetDelTraceWithGID(t *testing.T) {
	const gid int64 = 12345

	// Initially, trace for GID should be empty
	assert.Empty(t, GetTraceWithGID(gid), "Initial trace for GID should be empty")

	// Set a trace ID for GID
	SetTraceWithGID(gid, "test-id-gid")
	assert.Equal(t, "test-id-gid", GetTraceWithGID(gid), "Trace ID for GID should be 'test-id-gid'")

	// Delete the trace for GID
	DelTraceWithGID(gid)
	assert.Empty(t, GetTraceWithGID(gid), "Trace for GID should be empty after deletion")

	// Set an empty trace ID for GID, should generate a new one
	SetTraceWithGID(gid)
	assert.NotEmpty(t, GetTraceWithGID(gid), "A new trace ID should be generated for GID")
	assert.Len(t, GetTraceWithGID(gid), 16, "Generated trace ID for GID should have a length of 16")
	DelTraceWithGID(gid)
}

// TestDisableTrace tests the functionality of the DisableTrace switch.
func TestDisableTrace(t *testing.T) {
	DisableTrace = true
	defer func() { DisableTrace = false }() // Reset after test

	SetTrace("no-trace")
	assert.Empty(t, GetTrace(), "GetTrace should return empty when tracing is disabled")
}

// TestGenTraceId tests the trace ID generation.
func TestGenTraceId(t *testing.T) {
	traceId := GenTraceId()
	assert.Len(t, traceId, 16, "Generated trace ID should be 16 characters long")
}

// TestConcurrencySafety tests the thread safety of traceMap.
func TestConcurrencySafety(t *testing.T) {
	var wg sync.WaitGroup
	numGoroutines := 100
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()

			// Each goroutine sets its own trace ID
			traceID := fmt.Sprintf("goroutine-%d", i)
			SetTrace(traceID)

			// Verify the trace ID is correctly set for this goroutine
			assert.Equal(t, traceID, GetTrace(), "Trace ID should be correctly set for each goroutine")

			// Simulate some work and re-verify
			assert.Equal(t, traceID, GetTrace(), "Trace ID should remain consistent within the same goroutine")

			DelTrace()
			assert.Empty(t, GetTrace(), "Trace ID should be empty after deletion in concurrent goroutine")
		}(i)
	}

	wg.Wait()
}

// BenchmarkSetTrace 测试 SetTrace 函数的性能
func BenchmarkSetTrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetTrace("benchmark-trace-id")
	}
}

// BenchmarkGetTrace 测试 GetTrace 函数的性能
func BenchmarkGetTrace(b *testing.B) {
	SetTrace("benchmark-trace-id") // 先设置一个值
	b.ResetTimer()                 // 重置计时器
	for i := 0; i < b.N; i++ {
		GetTrace()
	}
}

// BenchmarkSetGetTraceParallel 测试并发场景下 SetTrace 和 GetTrace 的性能
func BenchmarkSetGetTraceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SetTrace()
			GetTrace()
		}
	})
}

// Example_package 展示了 log 包的基本用法。
//
// 这个示例演示了如何设置一个自定义的追踪ID，然后获取它，
// 最后再将它删除，并验证它已被成功移除。
func Example_package() {
	// 1. 设置一个自定义的追踪ID
	myTraceID := "my-awesome-trace-id"
	SetTrace(myTraceID)

	// 2. 获取并打印当前的追踪ID
	retrievedTraceID := GetTrace()
	fmt.Println(retrievedTraceID)

	// 3. 清理追踪ID
	DelTrace()

	// 4. 再次获取，确认已被删除
	afterDeleteTraceID := GetTrace()
	fmt.Printf("Trace ID after deletion: '%s'\n", afterDeleteTraceID)

	// Output:
	// my-awesome-trace-id
	// Trace ID after deletion: ''
}
