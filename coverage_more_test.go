package log

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test for additional logger methods to improve coverage

func TestLoggerMethodsDisabledLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(ErrorLevel) // Only Error and above enabled
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Log_disabled", func(t *testing.T) {
		buf.Reset()
		logger.Log(InfoLevel, "should not log")
		if buf.Len() != 0 {
			t.Error("Log should not log when level is disabled")
		}
	})

	t.Run("Print_disabled", func(t *testing.T) {
		buf.Reset()
		logger.Print("should not log")
		if buf.Len() != 0 {
			t.Error("Print should not log when Debug is disabled")
		}
	})

	t.Run("Warn_disabled", func(t *testing.T) {
		buf.Reset()
		logger.Warn("should not log")
		if buf.Len() != 0 {
			t.Error("Warn should not log when level is disabled")
		}
	})

	t.Run("Warning_disabled", func(t *testing.T) {
		buf.Reset()
		logger.Warning("should not log")
		if buf.Len() != 0 {
			t.Error("Warning should not log when level is disabled")
		}
	})

	t.Run("Error_enabled", func(t *testing.T) {
		buf.Reset()
		logger.Error("should log")
		if buf.Len() == 0 {
			t.Error("Error should log when level is enabled")
		}
		if !strings.Contains(buf.String(), "should log") {
			t.Error("Error should contain message")
		}
	})
}

func TestLoggerLogfAndPrintf(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(DebugLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Logf_with_format", func(t *testing.T) {
		buf.Reset()
		logger.Logf(InfoLevel, "formatted %s %d", "message", 42)
		if !strings.Contains(buf.String(), "formatted message 42") {
			t.Error("Logf should format message")
		}
	})

	t.Run("Printf_same_as_Debugf", func(t *testing.T) {
		buf.Reset()
		logger.Printf("printf test %s", "arg")
		if !strings.Contains(buf.String(), "printf test arg") {
			t.Error("Printf should format message")
		}
	})
}

func TestGetOutputWriterCoverage(t *testing.T) {
	t.Run("GetOutputWriter_with_extension", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := filepath.Join(tmpDir, "test.log")

		writer := GetOutputWriter(logFile)
		if writer == nil {
			t.Fatal("GetOutputWriter should return writer")
		}

		// Write something to verify it works
		n, err := writer.Write([]byte("test\n"))
		if err != nil {
			t.Errorf("Write failed: %v", err)
		}
		if n != 5 {
			t.Errorf("Expected 5 bytes written, got %d", n)
		}

		// Verify file was created
		if _, err := os.Stat(logFile); os.IsNotExist(err) {
			t.Error("Log file should be created")
		}
	})
}
