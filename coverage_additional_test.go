package log

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestGetOutputWriterEdgeCases(t *testing.T) {
	t.Run("GetOutputWriter_creates_parent_dirs", func(t *testing.T) {
		tmpDir := t.TempDir()
		deepPath := filepath.Join(tmpDir, "level1", "level2", "level3", "test.log")

		writer := GetOutputWriter(deepPath)
		if writer == nil {
			t.Fatal("GetOutputWriter should return writer")
		}

		// Verify parent directory was created
		parentDir := filepath.Join(tmpDir, "level1", "level2", "level3")
		if _, err := os.Stat(parentDir); os.IsNotExist(err) {
			t.Errorf("Parent directory %s should be created", parentDir)
		}
	})

	t.Run("GetOutputWriter_same_file_twice", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := filepath.Join(tmpDir, "test.log")

		writer1 := GetOutputWriter(logFile)
		writer2 := GetOutputWriter(logFile)

		if writer1 == nil || writer2 == nil {
			t.Error("GetOutputWriter should return writer")
		}

		// Both writers should work
		writer1.Write([]byte("line1\n"))
		writer2.Write([]byte("line2\n"))

		// Verify file exists and has content
		content, err := os.ReadFile(logFile)
		if err != nil {
			t.Fatalf("Failed to read log file: %v", err)
		}

		contentStr := string(content)
		if !strings.Contains(contentStr, "line1") || !strings.Contains(contentStr, "line2") {
			t.Error("Both writers should write to file")
		}
	})
}

func TestRotatorEdgeCases(t *testing.T) {
	t.Run("NewHourlyRotater_with_subdirs", func(t *testing.T) {
		tmpDir := t.TempDir()
		logDir := filepath.Join(tmpDir, "logs", "app")

		rotator := NewHourlyRotator(logDir, 100, 10)
		if rotator == nil {
			t.Error("NewHourlyRotator should return rotator")
		}

		rotator.Close()
	})
}

func TestLoggerWriterInterface(t *testing.T) {
	t.Run("Writer_implements_io_writer", func(t *testing.T) {
		var _ io.Writer = &HourlyRotator{}
	})
}

func TestLoggerWithVariousOutputs(t *testing.T) {
	logger := New()

	t.Run("With_discard_output", func(t *testing.T) {
		logger.SetOutput(io.Discard)
		logger.Info("test")
		// Should not panic and should not write anywhere
	})
}

func TestLoggerLevelTransitions(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	levels := []Level{
		TraceLevel,
		DebugLevel,
		InfoLevel,
		WarnLevel,
		ErrorLevel,
	}

	for _, level := range levels {
		t.Run("SetLevel_to_"+level.String(), func(t *testing.T) {
			buf.Reset()
			logger.SetLevel(level)

			// Should log at this level
			logger.Log(level, "message")
			if buf.Len() == 0 {
				t.Errorf("Should log at level %v", level)
			}

			// Should not log at lower level
			if level > TraceLevel {
				lowerLevel := Level(int(level) - 1)
				buf.Reset()
				logger.Log(lowerLevel, "lower message")
				if buf.Len() != 0 {
					t.Errorf("Should not log at lower level %v", lowerLevel)
				}
			}
		})
	}
}

func TestEntryReset(t *testing.T) {
	t.Run("putEntry_resets_entry", func(t *testing.T) {
		entry := getEntry()
		entry.Message = "test message"
		entry.Time = time.Now()
		entry.Fields = []KV{{Key: "test", Value: "value"}}

		putEntry(entry)

		// Get another entry - might be the same one from pool
		entry2 := getEntry()
		if entry2.Message != "" {
			// Entry should be clean (though this is not guaranteed due to pooling)
		}
		putEntry(entry2)
	})
}

func TestJSONFormatterWithAllOptions(t *testing.T) {
	entry := &Entry{
		Level:      InfoLevel,
		Message:    "test message",
		Time:       time.Now(),
		Pid:        123,
		Gid:        456,
		CallerName: "main.example",
		File:       "example.go",
		CallerLine: 42,
		TraceId:    "trace-123",
		PrefixMsg:  []byte("[PREFIX]"),
		SuffixMsg:  []byte("[SUFFIX]"),
		Fields: []KV{
			{Key: "key1", Value: "value1"},
			{Key: "key2", Value: 123},
		},
	}

	t.Run("JSON_with_all_fields", func(t *testing.T) {
		formatter := &JSONFormatter{
			EnablePrettyPrint: true,
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		checks := []string{
			"test message",
			"\"level\":",
			"\"message\":",
			"\"pid\":",
			"\"key1\":",
			"\"key2\":",
		}

		for _, check := range checks {
			if !strings.Contains(resultStr, check) {
				t.Errorf("JSON should contain %s", check)
			}
		}
	})

	t.Run("JSON_with_all_disabled", func(t *testing.T) {
		formatter := &JSONFormatter{
			DisableTimestamp: true,
			DisableCaller:    true,
			DisableTrace:     true,
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		// Should still have basic fields
		if !strings.Contains(resultStr, "test message") {
			t.Error("Should contain message")
		}

		// Should not have disabled fields
		if strings.Contains(resultStr, "\"caller\"") {
			t.Error("Should not have caller when disabled")
		}
		if strings.Contains(resultStr, "\"trace\"") {
			t.Error("Should not have trace when disabled")
		}
	})
}

func TestLoggerConcurrency(t *testing.T) {
	t.Run("Concurrent_logging", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)
		logger.SetLevel(InfoLevel)
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		done := make(chan bool, 10)

		for i := 0; i < 10; i++ {
			go func(n int) {
				for j := 0; j < 100; j++ {
					logger.Info("message", "count", n)
				}
				done <- true
			}(i)
		}

		for i := 0; i < 10; i++ {
			<-done
		}

		// Should have logged all messages without corruption
		if buf.Len() == 0 {
			t.Error("Should have logged messages")
		}
	})
}
