package log

import (
	"fmt"
	"testing"
)

// BenchmarkFastStringify_String tests fast path for string
func BenchmarkFastStringify_String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fastStringify("test message")
	}
}

// BenchmarkFastStringify_Int tests fast path for int
func BenchmarkFastStringify_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fastStringify(12345)
	}
}

// BenchmarkFastStringify_Error tests fast path for error
func BenchmarkFastStringify_Error(b *testing.B) {
	b.ReportAllocs()
	err := fmt.Errorf("test error")
	for i := 0; i < b.N; i++ {
		_ = fastStringify(err)
	}
}

// BenchmarkFmtSprint_String tests fmt.Sprint for string (baseline)
func BenchmarkFmtSprint_String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint("test message")
	}
}

// BenchmarkFmtSprint_Int tests fmt.Sprint for int (baseline)
func BenchmarkFmtSprint_Int(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(12345)
	}
}

// BenchmarkFastSprint_SingleArg tests fastSprint with single argument
func BenchmarkFastSprint_SingleArg(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fastSprint("test message")
	}
}

// BenchmarkFastSprint_MultipleArgs tests fastSprint with multiple args
func BenchmarkFastSprint_MultipleArgs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fastSprint("test", 123, "message")
	}
}

// BenchmarkFmtSprint_MultipleArgs tests fmt.Sprint with multiple args
func BenchmarkFmtSprint_MultipleArgs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint("test", 123, "message")
	}
}

// BenchmarkFastSprintf_NoArgs tests fastSprintf with no args
func BenchmarkFastSprintf_NoArgs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fastSprintf("test message")
	}
}

// BenchmarkFastSprintf_SimplePercentS tests fastSprintf with %s
func BenchmarkFastSprintf_SimplePercentS(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fastSprintf("%s", "test message")
	}
}

// BenchmarkFmtSprintf_SimplePercentS tests fmt.Sprintf with %s
func BenchmarkFmtSprintf_SimplePercentS(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", "test message")
	}
}
