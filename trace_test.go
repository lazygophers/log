package log

import (
	"strings"
	"testing"

	"github.com/petermattis/goid"
)

func TestGetTrace(t *testing.T) {
	// 清空现有的 trace
	traceMap.Range(func(key, value interface{}) bool {
		traceMap.Delete(key)
		return true
	})
	
	// 测试不存在的 trace
	trace := GetTrace()
	if trace != "" {
		t.Errorf("Expected empty trace for new goroutine, got %q", trace)
	}
	
	// 设置 trace
	expectedTrace := "test-trace-123"
	SetTrace(expectedTrace)
	
	// 获取 trace
	trace = GetTrace()
	if trace != expectedTrace {
		t.Errorf("Expected %q, got %q", expectedTrace, trace)
	}
}

func TestGetTraceWithGID(t *testing.T) {
	gid := goid.Get()
	
	// 清空现有的 trace
	traceMap.Range(func(key, value interface{}) bool {
		traceMap.Delete(key)
		return true
	})
	
	// 测试不存在的 trace
	trace := GetTraceWithGID(gid)
	if trace != "" {
		t.Errorf("Expected empty trace for GID %d, got %q", gid, trace)
	}
	
	// 设置 trace
	expectedTrace := "test-trace-with-gid-456"
	SetTraceWithGID(gid, expectedTrace)
	
	// 获取 trace
	trace = GetTraceWithGID(gid)
	if trace != expectedTrace {
		t.Errorf("Expected %q, got %q", expectedTrace, trace)
	}
}

func TestSetTrace(t *testing.T) {
	gid := goid.Get()
	expectedTrace := "set-trace-test"
	
	SetTrace(expectedTrace)
	
	// 验证是否设置成功
	if trace := getTrace(gid); trace != expectedTrace {
		t.Errorf("Expected %q, got %q", expectedTrace, trace)
	}
	
	// 测试设置空字符串 - 根据实现，空字符串会自动生成新的trace ID
	SetTrace("")
	if trace := getTrace(gid); trace == "" {
		t.Error("Expected auto-generated trace after setting empty string, got empty")
	}
}

func TestSetTraceWithGID(t *testing.T) {
	gid := int64(99999) // 使用一个不太可能冲突的 GID
	expectedTrace := "set-trace-with-gid-test"
	
	SetTraceWithGID(gid, expectedTrace)
	
	// 验证是否设置成功
	if trace := getTrace(gid); trace != expectedTrace {
		t.Errorf("Expected %q, got %q", expectedTrace, trace)
	}
	
	// 测试设置空字符串 - 根据实现，空字符串会自动生成新的trace ID
	SetTraceWithGID(gid, "")
	if trace := getTrace(gid); trace == "" {
		t.Error("Expected auto-generated trace after setting empty string, got empty")
	}
}

func TestDelTrace(t *testing.T) {
	gid := goid.Get()
	expectedTrace := "del-trace-test"
	
	// 设置 trace
	SetTrace(expectedTrace)
	if trace := getTrace(gid); trace != expectedTrace {
		t.Errorf("Expected %q, got %q", expectedTrace, trace)
	}
	
	// 删除 trace
	DelTrace()
	if trace := getTrace(gid); trace != "" {
		t.Errorf("Expected empty trace after deletion, got %q", trace)
	}
}

func TestDelTraceWithGID(t *testing.T) {
	gid := int64(88888) // 使用一个不太可能冲突的 GID
	expectedTrace := "del-trace-with-gid-test"
	
	// 设置 trace
	SetTraceWithGID(gid, expectedTrace)
	if trace := getTrace(gid); trace != expectedTrace {
		t.Errorf("Expected %q, got %q", expectedTrace, trace)
	}
	
	// 删除 trace
	DelTraceWithGID(gid)
	if trace := getTrace(gid); trace != "" {
		t.Errorf("Expected empty trace after deletion, got %q", trace)
	}
}

func TestGenTraceId(t *testing.T) {
	trace1 := GenTraceId()
	trace2 := GenTraceId()
	
	// 检查生成的 trace ID 不为空
	if trace1 == "" {
		t.Error("Generated trace ID should not be empty")
	}
	
	if trace2 == "" {
		t.Error("Generated trace ID should not be empty")
	}
	
	// 检查两次生成的 trace ID 不相同
	if trace1 == trace2 {
		t.Error("Generated trace IDs should be unique")
	}
	
	// 检查 trace ID 的格式（根据实现，是16个字符的字符串，没有连字符）
	if len(trace1) != 16 {
		t.Errorf("Expected trace ID length 16, got %d", len(trace1))
	}
	
	if strings.Count(trace1, "-") != 0 { // 实现中移除了连字符
		t.Errorf("Expected 0 hyphens in trace ID, got %d", strings.Count(trace1, "-"))
	}
}

func TestGetTrace_Internal(t *testing.T) {
	gid := int64(77777)
	expectedTrace := "internal-get-trace-test"
	
	// 直接设置到 map 中
	setTrace(gid, expectedTrace)
	
	// 使用内部函数获取
	trace := getTrace(gid)
	if trace != expectedTrace {
		t.Errorf("Expected %q, got %q", expectedTrace, trace)
	}
	
	// 测试不存在的 GID
	nonExistentTrace := getTrace(int64(11111))
	if nonExistentTrace != "" {
		t.Errorf("Expected empty trace for non-existent GID, got %q", nonExistentTrace)
	}
}

func TestSetTrace_Internal(t *testing.T) {
	gid := int64(66666)
	expectedTrace := "internal-set-trace-test"
	
	// 使用内部函数设置
	setTrace(gid, expectedTrace)
	
	// 验证是否设置成功
	if value, ok := traceMap.Load(gid); !ok || value.(string) != expectedTrace {
		t.Errorf("Expected %q to be set for GID %d", expectedTrace, gid)
	}
	
	// 测试设置空字符串（根据实现，会自动生成新的trace ID）
	setTrace(gid, "")
	if value, ok := traceMap.Load(gid); !ok {
		t.Error("Expected auto-generated trace when setting empty string, but found none")
	} else if value.(string) == "" {
		t.Error("Expected auto-generated trace when setting empty string, got empty")
	}
}

func TestDelTrace_Internal(t *testing.T) {
	gid := int64(55555)
	expectedTrace := "internal-del-trace-test"
	
	// 设置 trace
	setTrace(gid, expectedTrace)
	if value, ok := traceMap.Load(gid); !ok || value.(string) != expectedTrace {
		t.Errorf("Expected %q to be set for GID %d", expectedTrace, gid)
	}
	
	// 使用内部函数删除
	delTrace(gid)
	
	// 验证是否删除成功
	if value, ok := traceMap.Load(gid); ok {
		t.Errorf("Expected trace to be deleted, but found %q", value.(string))
	}
}

func TestTraceConcurrency(t *testing.T) {
	// 测试并发安全性
	done := make(chan bool, 100)
	
	for i := 0; i < 100; i++ {
		go func(id int) {
			defer func() { done <- true }()
			
			traceID := GenTraceId()
			SetTrace(traceID)
			
			// 验证设置的 trace
			if getTrace := GetTrace(); getTrace != traceID {
				t.Errorf("Goroutine %d: Expected %q, got %q", id, traceID, getTrace)
			}
			
			DelTrace()
			
			// 验证删除后的 trace
			if getTrace := GetTrace(); getTrace != "" {
				t.Errorf("Goroutine %d: Expected empty trace after deletion, got %q", id, getTrace)
			}
		}(i)
	}
	
	// 等待所有 goroutine 完成
	for i := 0; i < 100; i++ {
		<-done
	}
}

func TestTraceWithContext(t *testing.T) {
	// 这里测试 trace 功能与上下文的配合使用
	originalTrace := "original-context-trace"
	
	// 设置当前 goroutine 的 trace
	SetTrace(originalTrace)
	
	// 验证当前 trace
	if trace := GetTrace(); trace != originalTrace {
		t.Errorf("Expected %q, got %q", originalTrace, trace)
	}
	
	// 在 goroutine 中设置不同的 trace
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		
		newTrace := "new-context-trace"
		SetTrace(newTrace)
		
		// 验证新的 trace
		if trace := GetTrace(); trace != newTrace {
			t.Errorf("Expected %q in new goroutine, got %q", newTrace, trace)
		}
	}()
	
	<-done
	
	// 验证原始 goroutine 的 trace 没有受影响
	if trace := GetTrace(); trace != originalTrace {
		t.Errorf("Expected original trace %q to be unchanged, got %q", originalTrace, trace)
	}
	
	// 清理
	DelTrace()
}

func TestSetTrace_DisableTrace_Coverage(t *testing.T) {
	// 测试 DisableTrace 标志的影响（trace.go:67-69）
	
	// 保存原始设置
	originalDisableTrace := DisableTrace
	defer func() {
		DisableTrace = originalDisableTrace
	}()
	
	// 测试设置 trace 时被禁用
	testGid := int64(99999)
	testTrace := "disabled-trace-test"
	
	// 先清理任何可能存在的trace
	delTrace(testGid)
	
	// 设置为禁用跟踪
	DisableTrace = true
	
	// 尝试设置 trace
	setTrace(testGid, testTrace)
	
	// 验证 trace 没有被设置
	if value, ok := traceMap.Load(testGid); ok {
		t.Errorf("Expected no trace when DisableTrace is true, but found %q", value.(string))
	}
	
	// 恢复设置，验证正常工作
	DisableTrace = false
	setTrace(testGid, testTrace)
	
	if value, ok := traceMap.Load(testGid); !ok || value.(string) != testTrace {
		t.Errorf("Expected trace %q to be set after enabling trace", testTrace)
	}
	
	// 清理
	delTrace(testGid)
}

func TestSetTrace_EmptyTraceId_Coverage(t *testing.T) {
	// 测试 SetTrace 使用空字符串时的自动生成（trace.go:130）
	
	// 保存原始 trace
	originalTrace := GetTrace()
	defer func() {
		if originalTrace != "" {
			SetTrace(originalTrace)
		} else {
			DelTrace()
		}
	}()
	
	// 使用空字符串调用 SetTrace（触发自动生成）
	SetTrace()
	
	// 获取生成的 trace ID
	generatedTrace := GetTrace()
	
	// 验证生成了 trace ID
	if generatedTrace == "" {
		t.Error("Expected SetTrace() to generate a trace ID, but got empty string")
	}
	
	// 验证生成的 trace ID 符合预期格式
	if len(generatedTrace) == 0 {
		t.Error("Generated trace ID should not be empty")
	}
}

func TestSetTraceWithGID_EmptyTraceId_Coverage(t *testing.T) {
	// 测试 SetTraceWithGID 使用空字符串时的自动生成（trace.go:145）
	
	testGid := int64(88888)
	
	// 清理可能存在的 trace
	delTrace(testGid)
	defer delTrace(testGid)
	
	// 使用空参数调用 SetTraceWithGID（触发自动生成）
	SetTraceWithGID(testGid)
	
	// 获取生成的 trace ID
	if value, ok := traceMap.Load(testGid); ok {
		generatedTrace := value.(string)
		
		// 验证生成了 trace ID
		if generatedTrace == "" {
			t.Error("Expected SetTraceWithGID to generate a trace ID, but got empty string")
		}
		
		// 验证生成的 trace ID 符合预期格式
		if len(generatedTrace) == 0 {
			t.Error("Generated trace ID should not be empty")
		}
	} else {
		t.Error("Expected trace ID to be set for GID, but none found")
	}
}