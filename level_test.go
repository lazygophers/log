package log

import (
	"fmt"
	"testing"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		level    Level
		expected string
	}{
		{TraceLevel, "trace"},
		{DebugLevel, "debug"},
		{InfoLevel, "info"},
		{WarnLevel, "warn"},
		{ErrorLevel, "error"},
		{FatalLevel, "fatal"},
		{PanicLevel, "panic"},
		{Level(100), "trace"}, // 默认回退
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.level), func(t *testing.T) {
			if got := tt.level.String(); got != tt.expected {
				t.Errorf("String() = %s, want %s", got, tt.expected)
			}
		})
	}
}

func TestLevel_MarshalText(t *testing.T) {
	tests := []struct {
		level    Level
		expected string
		hasError bool
	}{
		{TraceLevel, "trace", false},
		{DebugLevel, "debug", false},
		{InfoLevel, "info", false},
		{WarnLevel, "warning", false},
		{ErrorLevel, "error", false},
		{FatalLevel, "fatal", false},
		{PanicLevel, "panic", false},
		{Level(100), "", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.level), func(t *testing.T) {
			data, err := tt.level.MarshalText()
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("MarshalText() error = %v", err)
			}
			if got := string(data); got != tt.expected {
				t.Errorf("MarshalText() = %s, want %s", got, tt.expected)
			}
		})
	}
}

func BenchmarkLevel_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, l := range []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel} {
			_ = l.String()
		}
	}
}

func BenchmarkLevel_MarshalText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, l := range []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel} {
			_, _ = l.MarshalText()
		}
	}
}
