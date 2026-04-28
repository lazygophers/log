package logctx

import (
	"bytes"
	"context"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := New()
	if logger == nil {
		t.Fatal("New() returned nil")
	}
}

func TestLogger_SetLevel(t *testing.T) {
	logger := New()
	logger.SetLevel(ErrorLevel)
	
	// Should log at Error level
	logger.Info(context.Background(), "test") // should not output
	
	// Verify level was set
	if logger.level != ErrorLevel {
		t.Errorf("Expected level ErrorLevel, got %v", logger.level)
	}
}

func TestLogger_SetOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	
	logger.Info(context.Background(), "test message")
	
	if buf.Len() == 0 {
		t.Error("Expected output, got empty buffer")
	}
}

func TestLogger_SetPrefixMsg(t *testing.T) {
	logger := New()
	prefix := "TEST: "
	logger.SetPrefixMsg(prefix)
	
	if logger.prefixMsg != prefix {
		t.Errorf("Expected prefix %q, got %q", prefix, logger.prefixMsg)
	}
}

func TestLogger_SetSuffixMsg(t *testing.T) {
	logger := New()
	suffix := "[END]"
	logger.SetSuffixMsg(suffix)
	
	if logger.suffixMsg != suffix {
		t.Errorf("Expected suffix %q, got %q", suffix, logger.suffixMsg)
	}
}

func TestLogger_Clone(t *testing.T) {
	logger := New()
	logger.SetLevel(ErrorLevel)
	logger.SetPrefixMsg("ORIGINAL: ")
	
	clone := logger.Clone()
	
	if clone == logger {
		t.Error("Clone should return a different instance")
	}
	
	if clone.level != ErrorLevel {
		t.Errorf("Cloned logger should have same level")
	}
	
	if clone.prefixMsg != "ORIGINAL: " {
		t.Errorf("Cloned logger should have same prefix")
	}
}

func TestLogger_LogWithContext(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(InfoLevel)
	
	ctx := context.Background()
	logger.Info(ctx, "test message")
	
	if buf.Len() == 0 {
		t.Error("Expected output with valid context")
	}
	
	// Test with canceled context
	canceled, cancel := context.WithCancel(ctx)
	cancel()
	
	var buf2 bytes.Buffer
	logger.SetOutput(&buf2)
	logger.Info(canceled, "should not log")
	
	if buf2.Len() > 0 {
		t.Error("Should not log with canceled context")
	}
}

func TestLogger_Logf(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(InfoLevel)
	
	logger.Infof(context.Background(), "formatted %s", "message")
	
	output := buf.String()
	if !contains(output, "formatted message") {
		t.Errorf("Expected formatted output, got %q", output)
	}
}

func TestLogger_Levels(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)
	
	ctx := context.Background()
	
	// Test all level methods
	logger.Trace(ctx, "trace")
	logger.Debug(ctx, "debug")
	logger.Info(ctx, "info")
	logger.Warn(ctx, "warn")
	logger.Error(ctx, "error")
	
	output := buf.String()
	if !contains(output, "TRACE") || !contains(output, "DEBUG") || 
	   !contains(output, "INFO") || !contains(output, "WARN") || !contains(output, "ERROR") {
		t.Error("Expected all log levels in output")
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && 
		   (s == substr || len(s) > len(substr) && 
		    (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || 
		     indexOf(s, substr) >= 0))
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
