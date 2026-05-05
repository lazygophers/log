package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Final comprehensive tests to reach 90% coverage

func TestRotatorDetailedCoverage(t *testing.T) {
	t.Run("NewHourlyRotator_basic_write", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 5)

		rotator.Write([]byte("line 1\n"))
		rotator.Write([]byte("line 2\n"))
		rotator.Write([]byte("line 3\n"))

		rotator.Close()

		// Verify file was created
		files, _ := os.ReadDir(tmpDir)
		if len(files) == 0 {
			t.Error("Should create log file")
		}
	})

	t.Run("NewHourlyRotator_empty_write", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 5)

		_, err := rotator.Write([]byte{})
		if err != nil {
			t.Errorf("Empty write failed: %v", err)
		}

		rotator.Close()
	})

	t.Run("NewHourlyRotator_large_write", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 1000, 10)

		largeData := strings.Repeat("test line\n", 1000)
		_, err := rotator.Write([]byte(largeData))
		if err != nil {
			t.Errorf("Large write failed: %v", err)
		}

		rotator.Close()
	})

	t.Run("NewHourlyRotator_sync_operations", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 5)

		rotator.Write([]byte("test\n"))

		for i := 0; i < 10; i++ {
			err := rotator.Sync()
			if err != nil {
				t.Errorf("Sync %d failed: %v", i, err)
			}
		}

		rotator.Close()
	})

	t.Run("NewHourlyRotator_close_operations", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 5)

		// Close without write
		err := rotator.Close()
		if err != nil {
			t.Errorf("Close failed: %v", err)
		}
	})
}

func TestFormatterDetailedCoverage(t *testing.T) {
	formatter := &Formatter{}

	t.Run("Format_with_various_messages", func(t *testing.T) {
		messages := []string{
			"",
			"simple",
			"with \"quotes\"",
			"with \\ backslash",
			"with \n newline",
			"with \t tab",
			"with \r return",
			"unicode: 你好世界",
			"emoji: 🌍🌎🌏",
			strings.Repeat("x", 1000),
		}

		for _, msg := range messages {
			entry := &Entry{
				Level:   InfoLevel,
				Message: msg,
				Time:    time.Now(),
				Pid:     123,
			}

			result := formatter.Format(entry)
			if len(result) == 0 {
				t.Errorf("Should format message: %q", msg)
			}
		}
	})

	t.Run("Format_with_all_fields", func(t *testing.T) {
		entry := &Entry{
			Level:      InfoLevel,
			Message:    "test",
			Time:       time.Now(),
			Pid:        123,
			Gid:        456,
			TraceId:    "trace-123",
			CallerName: "main.Func",
			File:       "file.go",
			CallerLine: 42,
			PrefixMsg:  []byte("[P]"),
			SuffixMsg:  []byte("[S]"),
			Fields: []KV{
				{Key: "key1", Value: "value1"},
				{Key: "key2", Value: 123},
				{Key: "key3", Value: true},
				{Key: "key4", Value: nil},
				{Key: "key5", Value: 3.14},
			},
		}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format entry with all fields")
		}
	})

	t.Run("Format_with_disabled_caller", func(t *testing.T) {
		formatter.DisableCaller = true

		entry := &Entry{
			Level:      InfoLevel,
			Message:    "test",
			Time:       time.Now(),
			Pid:        123,
			CallerName: "main.Func",
			File:       "file.go",
			CallerLine: 42,
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		if !strings.Contains(resultStr, "test") {
			t.Error("Should contain message")
		}
	})

	t.Run("Format_with_disabled_parsing", func(t *testing.T) {
		formatter.ParsingAndEscaping(false)

		entry := &Entry{
			Level:   InfoLevel,
			Message: "test with \"quotes\"",
			Time:    time.Now(),
			Pid:     123,
		}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format with parsing disabled")
		}
	})

	t.Run("Format_with_disabled_parse", func(t *testing.T) {
		formatter.Caller(false)

		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
		}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format with caller disabled")
		}
	})
}

func TestJSONFormatterDetailedCoverage(t *testing.T) {
	t.Run("JSON_with_empty_entry", func(t *testing.T) {
		formatter := &JSONFormatter{}

		entry := &Entry{}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format empty entry")
		}

		resultStr := string(result)
		if !strings.HasPrefix(resultStr, "{") {
			t.Error("Should be valid JSON object")
		}
	})

	t.Run("JSON_with_all_levels", func(t *testing.T) {
		formatter := &JSONFormatter{}

		levels := []Level{
			TraceLevel,
			DebugLevel,
			InfoLevel,
			WarnLevel,
			ErrorLevel,
		}

		for _, level := range levels {
			entry := &Entry{
				Level:   level,
				Message: "test",
				Time:    time.Now(),
				Pid:     123,
			}

			result := formatter.Format(entry)
			resultStr := string(result)

			if !strings.Contains(resultStr, level.String()) {
				t.Errorf("Should contain level %s", level.String())
			}
		}
	})

	t.Run("JSON_pretty_print", func(t *testing.T) {
		formatter := &JSONFormatter{EnablePrettyPrint: true}

		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		if !strings.Contains(resultStr, "\n") {
			t.Error("Pretty print should have newlines")
		}
		if !strings.HasPrefix(resultStr, "{") {
			t.Error("Should be valid JSON")
		}
	})

	t.Run("JSON_compact", func(t *testing.T) {
		formatter := &JSONFormatter{EnablePrettyPrint: false}

		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Pid:     123,
		}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format JSON")
		}
	})

	t.Run("JSON_with_all_disabled", func(t *testing.T) {
		formatter := &JSONFormatter{
			DisableCaller: true,
			DisableTrace:  true,
		}

		entry := &Entry{
			Level:      InfoLevel,
			Message:    "test",
			Time:       time.Now(),
			Pid:        123,
			TraceId:    "trace-123",
			CallerName: "main.Func",
			File:       "file.go",
			CallerLine: 42,
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		// Should still have message and level
		if !strings.Contains(resultStr, "test") {
			t.Error("Should contain message")
		}
		if !strings.Contains(resultStr, "level") {
			t.Error("Should contain level")
		}

		// Should not have disabled fields
		if strings.Contains(resultStr, "caller") {
			t.Error("Should not have caller when disabled")
		}
		if strings.Contains(resultStr, "trace") {
			t.Error("Should not have trace when disabled")
		}
	})

	t.Run("JSON_with_invalid_type", func(t *testing.T) {
		formatter := &JSONFormatter{}

		result := formatter.Format("not an entry")
		if result != nil {
			t.Error("Should return nil for non-entry type")
		}
	})

	t.Run("JSON_with_marshall_error", func(t *testing.T) {
		formatter := &JSONFormatter{}

		// Create entry that will cause JSON marshaling to fail
		ch := make(chan int)
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test message with control \x01 chars",
			Time:    time.Now(),
			Pid:     123,
			Fields: []KV{
				{Key: "invalid", Value: ch},
			},
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		// Should fall back to error message
		if !strings.Contains(resultStr, "JSON marshaling failed") {
			t.Error("Should include marshaling error")
		}
		if !strings.Contains(resultStr, "test message") {
			t.Error("Should include original message")
		}
	})
}

func TestGetOutputWriterDetailed(t *testing.T) {
	t.Run("GetOutputWriter_nested_paths", func(t *testing.T) {
		tmpDir := t.TempDir()

		paths := []string{
			filepath.Join(tmpDir, "test.log"),
			filepath.Join(tmpDir, "logs", "test.log"),
			filepath.Join(tmpDir, "logs", "app", "test.log"),
			filepath.Join(tmpDir, "logs", "app", "service", "test.log"),
		}

		for _, logFile := range paths {
			writer := GetOutputWriter(logFile)
			if writer == nil {
				t.Errorf("GetOutputWriter should return writer for %s", logFile)
			}

			writer.Write([]byte("test\n"))

			// Verify file exists
			if _, err := os.Stat(logFile); os.IsNotExist(err) {
				t.Errorf("Log file %s should be created", logFile)
			}
		}
	})

	t.Run("GetOutputWriter_multiple_writers_same_file", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := filepath.Join(tmpDir, "test.log")

		writer1 := GetOutputWriter(logFile)
		writer2 := GetOutputWriter(logFile)

		if writer1 == nil || writer2 == nil {
			t.Fatal("Both writers should be returned")
		}

		writer1.Write([]byte("line 1\n"))
		writer2.Write([]byte("line 2\n"))

		content, err := os.ReadFile(logFile)
		if err != nil {
			t.Fatalf("Failed to read log file: %v", err)
		}

		contentStr := string(content)
		if !strings.Contains(contentStr, "line 1") {
			t.Error("Should contain first line")
		}
		if !strings.Contains(contentStr, "line 2") {
			t.Error("Should contain second line")
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

func TestIsAllDigits(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"0", true},
		{"", true},           // Empty string returns true (no non-digit chars)
		{"abc", false},
		{"12a34", false},
		{"12.34", false},
		{"12 34", false},
		{"-123", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := isAllDigits(tc.input)
			if result != tc.expected {
				t.Errorf("isAllDigits(%q) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestCleanupOldFiles(t *testing.T) {
	t.Run("cleanupOldFiles_with_nonexistent_dir", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(filepath.Join(tmpDir, "nonexistent"), 100, 10)

		// Should not panic
		rotator.cleanupOldFiles()
		rotator.Close()
	})

	t.Run("cleanupOldFiles_with_no_log_files", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 10)

		// Should not panic
		rotator.cleanupOldFiles()
		rotator.Close()
	})

	t.Run("cleanupOldFiles_with_valid_log_files", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 10)

		// Create some test log files with proper naming
		timestamp := time.Now().Format("2006010215")
		for i := 0; i < 5; i++ {
			filename := filepath.Join(tmpDir, timestamp+".log."+fmt.Sprint(i))
			_ = os.WriteFile(filename, []byte("test"), 0644)
		}

		// Should not panic
		rotator.cleanupOldFiles()
		rotator.Close()
	})

	t.Run("cleanupOldFiles_with_mixed_files", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 10)

		// Create mix of log files and non-log files
		timestamp := time.Now().Format("2006010215")
		_ = os.WriteFile(filepath.Join(tmpDir, timestamp+".log"), []byte("log1"), 0644)
		_ = os.WriteFile(filepath.Join(tmpDir, "readme.txt"), []byte("readme"), 0644)
		_ = os.WriteFile(filepath.Join(tmpDir, "config.json"), []byte("{}"), 0644)

		// Should not panic and only process .log files
		rotator.cleanupOldFiles()
		rotator.Close()
	})

	t.Run("cleanupOldFiles_with_invalid_timestamp", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 100, 10)

		// Create log files with invalid timestamps
		_ = os.WriteFile(filepath.Join(tmpDir, "123.log"), []byte("short"), 0644)
		_ = os.WriteFile(filepath.Join(tmpDir, "12345678901.log"), []byte("long"), 0644)
		_ = os.WriteFile(filepath.Join(tmpDir, "notdigits.log"), []byte("invalid"), 0644)

		// Should not panic and skip invalid files
		rotator.cleanupOldFiles()
		rotator.Close()
	})

	t.Run("cleanupOldFiles_deletes_old_files", func(t *testing.T) {
		tmpDir := t.TempDir()
		maxFiles := 3
		rotator := NewHourlyRotator(tmpDir, 100, maxFiles)

		// Create more log files than maxFiles with different timestamps
		now := time.Now()
		for i := 0; i < 5; i++ {
			timestamp := now.Add(time.Duration(i) * time.Hour).Format("2006010215")
			filename := filepath.Join(tmpDir, timestamp+".log")
			_ = os.WriteFile(filename, []byte("test"), 0644)
		}

		// Cleanup should delete files exceeding maxFiles
		rotator.cleanupOldFiles()

		// Count remaining .log files
		files, _ := os.ReadDir(tmpDir)
		logFileCount := 0
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".log") {
				logFileCount++
			}
		}

		if logFileCount > maxFiles {
			t.Errorf("Should have at most %d log files, got %d", maxFiles, logFileCount)
		}

		rotator.Close()
	})
}

func TestRotatorSyncCoverage(t *testing.T) {
	t.Run("Sync_with_no_current_file", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 1024, 10)

		// Sync without writing anything (currentFile is nil)
		err := rotator.Sync()
		if err != nil {
			t.Errorf("Sync should return nil when no current file, got %v", err)
		}

		rotator.Close()
	})

	t.Run("Sync_with_current_file", func(t *testing.T) {
		tmpDir := t.TempDir()
		rotator := NewHourlyRotator(tmpDir, 1024, 10)

		// Write something to create current file
		rotator.Write([]byte("test\n"))

		// Now sync should work
		err := rotator.Sync()
		if err != nil {
			t.Errorf("Sync should succeed, got %v", err)
		}

		rotator.Close()
	})
}

func TestRotatorWriteEmptyCoverage(t *testing.T) {
	tmpDir := t.TempDir()
	rotator := NewHourlyRotator(tmpDir, 1024, 10)
	defer rotator.Close()

	// Test writing empty byte slice
	n, err := rotator.Write([]byte{})
	if n != 0 || err != nil {
		t.Errorf("Empty write should return (0, nil), got (%d, %v)", n, err)
	}
}
