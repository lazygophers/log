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


func TestFastStringifyCoverage(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
	}{
		{"string", "hello"},
		{"[]byte", []byte("world")},
		{"int", 42},
		{"int8", int8(8)},
		{"int16", int16(16)},
		{"int32", int32(32)},
		{"int64", int64(64)},
		{"uint", uint(10)},
		{"uint8", uint8(8)},
		{"uint16", uint16(16)},
		{"uint32", uint32(32)},
		{"uint64", uint64(64)},
		{"float32", float32(3.14)},
		{"float64", float64(6.28)},
		{"bool true", true},
		{"bool false", false},
		{"error", fmt.Errorf("test error")},
		{"nil error", error(nil)},
		{"nil", nil},
		{"custom type", struct{ Name string }{"test"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just test that fastStringify doesn't panic for all types
			_ = fastStringify(tt.input)
		})
	}
}
