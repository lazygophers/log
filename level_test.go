package log

import (
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
		{Level(999), "trace"}, // 未知级别返回 "trace"
	}
	
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.level.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestLevel_MarshalText(t *testing.T) {
	tests := []struct {
		level       Level
		expected    string
		shouldError bool
	}{
		{TraceLevel, "trace", false},
		{DebugLevel, "debug", false},
		{InfoLevel, "info", false},
		{WarnLevel, "warning", false}, // MarshalText 返回 "warning"
		{ErrorLevel, "error", false},
		{FatalLevel, "fatal", false},
		{PanicLevel, "panic", false},
		{Level(999), "", true}, // 未知级别返回错误
	}
	
	for _, tt := range tests {
		name := tt.expected
		if name == "" {
			name = "UNKNOWN"
		}
		t.Run(name, func(t *testing.T) {
			result, err := tt.level.MarshalText()
			
			if tt.shouldError {
				if err == nil {
					t.Error("Expected MarshalText to return error for unknown level")
				}
				return
			}
			
			if err != nil {
				t.Errorf("MarshalText returned error: %v", err)
			}
			
			if string(result) != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, string(result))
			}
		})
	}
}

func TestLevel_Values(t *testing.T) {
	// 测试所有级别的数值（基于实际定义的顺序）
	expectedValues := map[Level]int{
		PanicLevel: 0,
		FatalLevel: 1,
		ErrorLevel: 2,
		WarnLevel:  3,
		InfoLevel:  4,
		DebugLevel: 5,
		TraceLevel: 6,
	}
	
	for level, expectedValue := range expectedValues {
		if int(level) != expectedValue {
			t.Errorf("Expected level %s to have value %d, got %d", level.String(), expectedValue, int(level))
		}
	}
}

func TestLevel_Ordering(t *testing.T) {
	// 测试级别的顺序关系（按实际顺序：PanicLevel < FatalLevel < ... < TraceLevel）
	levels := []Level{PanicLevel, FatalLevel, ErrorLevel, WarnLevel, InfoLevel, DebugLevel, TraceLevel}
	
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] >= levels[i+1] {
			t.Errorf("Level %s should be less than %s", levels[i].String(), levels[i+1].String())
		}
	}
}

func TestLevel_Comparison(t *testing.T) {
	// 测试级别比较（基于实际值：数值越小优先级越高）
	if !(ErrorLevel < WarnLevel) {
		t.Error("ErrorLevel should be less than WarnLevel (higher priority)")
	}
	
	if !(PanicLevel < FatalLevel) {
		t.Error("PanicLevel should be less than FatalLevel (higher priority)")
	}
	
	if !(InfoLevel < DebugLevel) {
		t.Error("InfoLevel should be less than DebugLevel (higher priority)")
	}
	
	if !(DebugLevel < TraceLevel) {
		t.Error("DebugLevel should be less than TraceLevel (higher priority)")
	}
}