package log

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJSONFormatter_Format_Basic(t *testing.T) {
	entry := &Entry{
		Level:   InfoLevel,
		Message: "test message",
		Pid:     12345,
		Time:    time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{}
	output := f.Format(entry)

	if len(output) == 0 {
		t.Fatal("empty output")
	}

	// Verify it's valid JSON
	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if je.Level != "info" {
		t.Errorf("expected level 'info', got '%s'", je.Level)
	}
	if je.Message != "test message" {
		t.Errorf("expected message 'test message', got '%s'", je.Message)
	}
	if je.Pid != 12345 {
		t.Errorf("expected pid 12345, got %d", je.Pid)
	}
}

func TestJSONFormatter_Format_WithTrace(t *testing.T) {
	entry := &Entry{
		Level:   DebugLevel,
		Message: "debug with trace",
		Pid:     12345,
		Gid:     67890,
		TraceId: "trace-123",
		Time:    time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{}
	output := f.Format(entry)

	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if je.TraceID != "trace-123" {
		t.Errorf("expected trace_id 'trace-123', got '%s'", je.TraceID)
	}
	if je.Gid != 67890 {
		t.Errorf("expected gid 67890, got %d", je.Gid)
	}
}

func TestJSONFormatter_Format_WithCaller(t *testing.T) {
	entry := &Entry{
		Level:      ErrorLevel,
		Message:    "error with caller",
		Pid:        12345,
		File:       "logger_test.go",
		CallerLine: 42,
		CallerFunc: "TestJSONFormatter",
		CallerDir:  "log",
		CallerName: "log.TestJSONFormatter",
		Time:       time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{}
	output := f.Format(entry)

	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if je.CallerFile != "logger_test.go" {
		t.Errorf("expected caller_file 'logger_test.go', got '%s'", je.CallerFile)
	}
	if je.CallerLine != 42 {
		t.Errorf("expected caller_line 42, got %d", je.CallerLine)
	}
	if je.CallerFunc != "TestJSONFormatter" {
		t.Errorf("expected caller_func 'TestJSONFormatter', got '%s'", je.CallerFunc)
	}
}

func TestJSONFormatter_PrettyPrint(t *testing.T) {
	entry := &Entry{
		Level:   InfoLevel,
		Message: "pretty test",
		Pid:     12345,
		Time:    time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{EnablePrettyPrint: true}
	output := f.Format(entry)

	// Pretty printed JSON should contain newlines/indentation
	if len(output) == 0 {
		t.Fatal("empty output")
	}

	// Should contain indentation
	if !contains(output, []byte("  ")) {
		t.Error("pretty print should contain indentation")
	}
}

func TestJSONFormatter_DisableTimestamp(t *testing.T) {
	entry := &Entry{
		Level:   InfoLevel,
		Message: "no timestamp",
		Pid:     12345,
		Time:    time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{DisableTimestamp: true}
	output := f.Format(entry)

	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if je.Time != "" {
		t.Errorf("expected empty time, got '%s'", je.Time)
	}
}

func TestJSONFormatter_DisableCaller(t *testing.T) {
	entry := &Entry{
		Level:      InfoLevel,
		Message:    "no caller",
		Pid:        12345,
		File:       "test.go",
		CallerLine: 10,
		Time:       time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{DisableCaller: true}
	output := f.Format(entry)

	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if je.CallerFile != "" {
		t.Errorf("expected empty caller_file, got '%s'", je.CallerFile)
	}
}

func TestJSONFormatter_DisableTrace(t *testing.T) {
	entry := &Entry{
		Level:   InfoLevel,
		Message: "no trace",
		Pid:     12345,
		Gid:     67890,
		TraceId: "trace-abc",
		Time:    time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{DisableTrace: true}
	output := f.Format(entry)

	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if je.TraceID != "" {
		t.Errorf("expected empty trace_id, got '%s'", je.TraceID)
	}
	if je.Gid != 0 {
		t.Errorf("expected zero gid, got %d", je.Gid)
	}
}

func TestJSONFormatter_WithPrefixSuffix(t *testing.T) {
	entry := &Entry{
		Level:    InfoLevel,
		Message:   "with prefix/suffix",
		Pid:      12345,
		PrefixMsg: []byte("[PREFIX]"),
		SuffixMsg: []byte("[SUFFIX]"),
		Time:     time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{}
	output := f.Format(entry)

	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if je.PrefixMsg != "[PREFIX]" {
		t.Errorf("expected prefix_msg '[PREFIX]', got '%s'", je.PrefixMsg)
	}
	if je.SuffixMsg != "[SUFFIX]" {
		t.Errorf("expected suffix_msg '[SUFFIX]', got '%s'", je.SuffixMsg)
	}
}

func TestJSONFormatter_Escaping(t *testing.T) {
	entry := &Entry{
		Level:   InfoLevel,
		Message: "test with \"quotes\" and\nnewlines",
		Pid:     12345,
		Time:    time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC),
	}

	f := &JSONFormatter{}
	output := f.Format(entry)

	// Should be valid JSON
	var je jsonEntry
	if err := json.Unmarshal(output, &je); err != nil {
		t.Fatalf("invalid JSON with special chars: %v", err)
	}
}

func TestJSONFormatter_AllLevels(t *testing.T) {
	levels := []Level{
		TraceLevel, DebugLevel, InfoLevel, WarnLevel,
		ErrorLevel, FatalLevel, PanicLevel,
	}

	time := time.Date(2026, 5, 4, 12, 0, 0, 0, time.UTC)

	for _, level := range levels {
		entry := &Entry{
			Level:   level,
			Message: level.String() + " message",
			Pid:     12345,
			Time:    time,
		}

		f := &JSONFormatter{}
		output := f.Format(entry)

		var je jsonEntry
		if err := json.Unmarshal(output, &je); err != nil {
			t.Errorf("level %s: invalid JSON: %v", level, err)
		}

		expectedLevel := level.String()
		if je.Level != expectedLevel {
			t.Errorf("expected level '%s', got '%s'", expectedLevel, je.Level)
		}
	}
}

func contains(data []byte, substr []byte) bool {
	for i := 0; i <= len(data)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
		if i+j >= len(data) || data[i+j] != substr[j] {
			match = false
			break
		}
	}
		if match {
			return true
		}
	}
	return false
}
