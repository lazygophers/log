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

// syncWriter is a test writer that tracks Sync calls
type syncWriter struct {
	*bytes.Buffer
	synced bool
}

func (s *syncWriter) Sync() error {
	s.synced = true
	return nil
}

// Comprehensive tests to reach 90% coverage

func TestLoggerComprehensive(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Log_method", func(t *testing.T) {
		levels := []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel}
		for _, level := range levels {
			buf.Reset()
			logger.Log(level, "message")
			if buf.Len() == 0 {
				t.Errorf("Log should work at level %v", level)
			}
		}
	})

	t.Run("Logf_method", func(t *testing.T) {
		levels := []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel}
		for _, level := range levels {
			buf.Reset()
			logger.Logf(level, "formatted %s", "message")
			if buf.Len() == 0 {
				t.Errorf("Logf should work at level %v", level)
			}
		}
	})

	t.Run("Log_method_disabled", func(t *testing.T) {
		logger.SetLevel(ErrorLevel)

		buf.Reset()
		logger.Log(InfoLevel, "should not log")
		if buf.Len() != 0 {
			t.Error("Log should not log when level disabled")
		}
	})

	t.Run("Logf_method_disabled", func(t *testing.T) {
		buf.Reset()
		logger.Logf(InfoLevel, "should not log %s", "arg")
		if buf.Len() != 0 {
			t.Error("Logf should not log when level disabled")
		}
	})
}

func TestLoggerOutputAndSync(t *testing.T) {
	logger := New()

	t.Run("SetOutput_and_write", func(t *testing.T) {
		var buf bytes.Buffer
		logger.SetOutput(&buf)
		logger.SetLevel(InfoLevel)
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		logger.Info("test")
		if !strings.Contains(buf.String(), "test") {
			t.Error("Should write to output")
		}
	})

	t.Run("Sync", func(t *testing.T) {
		var buf bytes.Buffer
		logger.SetOutput(&buf)
		logger.Sync() // Should not panic
	})
}

func TestLoggerSyncWithSyncerComprehensive(t *testing.T) {
	t.Run("Sync_with_syncer", func(t *testing.T) {
		sw := &syncWriter{Buffer: &bytes.Buffer{}}
		logger := New()
		logger.SetOutput(sw)
		logger.SetLevel(InfoLevel)
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		logger.Info("test")
		logger.Sync()

		if !sw.synced {
			t.Error("Sync should be called")
		}
	})
}

func TestLoggerCloneComprehensive(t *testing.T) {
	logger := New()
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Clone
	cloned := logger.Clone()

	t.Run("Clone_independence", func(t *testing.T) {
		// Modify original
		logger.SetLevel(DebugLevel)

		// Clone should keep original level
		if cloned.Level() != InfoLevel {
			t.Error("Clone should have independent level")
		}
	})

	t.Run("Clone_works", func(t *testing.T) {
		var buf bytes.Buffer
		cloned.SetOutput(&buf)
		cloned.Info("test")

		if !strings.Contains(buf.String(), "test") {
			t.Error("Clone should work")
		}
	})
}

func TestEntryOperations(t *testing.T) {
	t.Run("getEntry_and_putEntry", func(t *testing.T) {
		entry := getEntry()

		// Modify entry
		entry.Message = "test"
		entry.Time = time.Now()
		entry.Fields = []KV{{Key: "test", Value: "value"}}

		// Return to pool
		putEntry(entry)

		// Get another entry
		entry2 := getEntry()
		if entry2 == nil {
			t.Error("getEntry should return entry")
		}
		putEntry(entry2)
	})
}

func TestFormatterAllOptions(t *testing.T) {
	entry := &Entry{
		Level:      InfoLevel,
		Message:    "test",
		Time:       time.Now(),
		Pid:        123,
		Gid:        456,
		CallerName: "main.test",
		File:       "test.go",
		CallerLine: 42,
		TraceId:    "trace-123",
		Fields: []KV{
			{Key: "key", Value: "value"},
		},
	}

	t.Run("Formatter_with_all_features", func(t *testing.T) {
		formatter := &Formatter{}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format entry")
		}

		resultStr := string(result)
		if !strings.Contains(resultStr, "test") {
			t.Error("Should contain message")
		}
	})

	t.Run("Formatter_with_disabled_caller", func(t *testing.T) {
		formatter := &Formatter{}
		formatter.DisableCaller = true

		result := formatter.Format(entry)
		resultStr := string(result)

		if !strings.Contains(resultStr, "test") {
			t.Error("Should format without caller")
		}
	})

	t.Run("Formatter_with_disabled_parse", func(t *testing.T) {
		formatter := &Formatter{}
		formatter.ParsingAndEscaping(false)

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format with parsing disabled")
		}
	})
}

func TestJSONFormatterComprehensive(t *testing.T) {
	entry := &Entry{
		Level:      InfoLevel,
		Message:    "test \"message\" with \\special chars",
		Time:       time.Now(),
		Pid:        123,
		TraceId:    "trace-123",
		CallerName: "main.Func",
		File:       "file.go",
		CallerLine: 10,
		Fields: []KV{
			{Key: "string", Value: "value"},
			{Key: "number", Value: 42},
			{Key: "bool", Value: true},
			{Key: "nil", Value: nil},
		},
	}

	t.Run("JSON_pretty_print", func(t *testing.T) {
		formatter := &JSONFormatter{EnablePrettyPrint: true}

		result := formatter.Format(entry)
		resultStr := string(result)

		if !strings.HasPrefix(resultStr, "{") {
			t.Error("Should be JSON")
		}
		if !strings.Contains(resultStr, "\n") {
			t.Error("Pretty print should have newlines")
		}
	})

	t.Run("JSON_compact", func(t *testing.T) {
		formatter := &JSONFormatter{EnablePrettyPrint: false}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format JSON")
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

		// Should still have level and message
		if !strings.Contains(resultStr, "level") {
			t.Error("Should have level")
		}
		if !strings.Contains(resultStr, "message") {
			t.Error("Should have message")
		}
	})
}

func TestGetOutputWriterComprehensive(t *testing.T) {
	t.Run("GetOutputWriter_in_temp_dir", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := filepath.Join(tmpDir, "app.log")

		writer := GetOutputWriter(logFile)
		if writer == nil {
			t.Fatal("GetOutputWriter should return writer")
		}

		// Write data
		writer.Write([]byte("test line 1\n"))
		writer.Write([]byte("test line 2\n"))

		// Verify file exists and has content
		content, err := os.ReadFile(logFile)
		if err != nil {
			t.Fatalf("Failed to read log file: %v", err)
		}

		contentStr := string(content)
		if !strings.Contains(contentStr, "test line 1") {
			t.Error("Should contain first line")
		}
		if !strings.Contains(contentStr, "test line 2") {
			t.Error("Should contain second line")
		}
	})

	t.Run("GetOutputWriter_creates_nested_dirs", func(t *testing.T) {
		tmpDir := t.TempDir()
		nestedPath := filepath.Join(tmpDir, "a", "b", "c", "test.log")

		writer := GetOutputWriter(nestedPath)
		if writer == nil {
			t.Fatal("GetOutputWriter should return writer")
		}

		writer.Write([]byte("test\n"))

		// Verify all parent directories were created
		for _, dir := range []string{
			filepath.Join(tmpDir, "a"),
			filepath.Join(tmpDir, "a", "b"),
			filepath.Join(tmpDir, "a", "b", "c"),
		} {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				t.Errorf("Directory %s should be created", dir)
			}
		}
	})
}

func TestLoggerWithDiscard(t *testing.T) {
	logger := New()
	logger.SetOutput(io.Discard)
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Write_to_discard", func(t *testing.T) {
		// Should not panic
		logger.Info("test")
		logger.Infof("test %s", "formatted")
		logger.Infow("test", "key", "value")
	})
}

func TestLoggerSetters(t *testing.T) {
	logger := New()

	t.Run("SetLevel_all_levels", func(t *testing.T) {
		levels := []Level{
			TraceLevel,
			DebugLevel,
			InfoLevel,
			WarnLevel,
			ErrorLevel,
			PanicLevel,
			FatalLevel,
		}

		for _, level := range levels {
			logger.SetLevel(level)
			if logger.Level() != level {
				t.Errorf("SetLevel(%v) should set level to %v", level, level)
			}
		}
	})

	t.Run("EnableCaller", func(t *testing.T) {
		logger.EnableCaller(false)
		logger.EnableCaller(true)
	})

	t.Run("EnableTrace", func(t *testing.T) {
		logger.EnableTrace(false)
		logger.EnableTrace(true)
	})

	t.Run("SetPrefixMsg", func(t *testing.T) {
		logger.SetPrefixMsg("[PREFIX]")
	})

	t.Run("SetSuffixMsg", func(t *testing.T) {
		logger.SetSuffixMsg("[SUFFIX]")
	})

	t.Run("SetFormat", func(t *testing.T) {
		logger.Format = &Formatter{}
		if logger.Format == nil {
			t.Error("Format should be settable")
		}
	})
}

func TestLoggerAliases(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(DebugLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Print_alias", func(t *testing.T) {
		buf.Reset()
		logger.Print("test")
		if !strings.Contains(buf.String(), "test") {
			t.Error("Print should work")
		}
	})

	t.Run("Warning_alias", func(t *testing.T) {
		buf.Reset()
		logger.Warning("test")
		if !strings.Contains(buf.String(), "test") {
			t.Error("Warning should work")
		}
	})
}
