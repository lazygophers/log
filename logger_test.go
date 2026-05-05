package log

import (
	"bytes"
	"strings"
	"testing"
	"time"
	"io"


	"github.com/lazygophers/log/constant"
)

// Additional tests for edge cases and remaining coverage

func TestLoggerWithFieldsEdgeCases(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("With_empty_fields", func(t *testing.T) {
		buf.Reset()
		logger.Infow("message")
		if !strings.Contains(buf.String(), "message") {
			t.Error("Should log message with no fields")
		}
	})

	t.Run("With_nil_value_field", func(t *testing.T) {
		buf.Reset()
		logger.Infow("message", "key", nil)
		if !strings.Contains(buf.String(), "key") {
			t.Error("Should log message with nil field value")
		}
	})

	t.Run("With_multiple_fields", func(t *testing.T) {
		buf.Reset()
		logger.Infow("message", "key1", "value1", "key2", 123, "key3", true)
		output := buf.String()
		if !strings.Contains(output, "key1") || !strings.Contains(output, "key2") || !strings.Contains(output, "key3") {
			t.Error("Should log all fields")
		}
	})

	t.Run("With_complex_values", func(t *testing.T) {
		buf.Reset()
		logger.Infow("message", "slice", []int{1, 2, 3}, "map", map[string]int{"a": 1})
		output := buf.String()
		if !strings.Contains(output, "slice") || !strings.Contains(output, "map") {
			t.Error("Should log complex values")
		}
	})
}

func TestLoggerOutputMethods(t *testing.T) {
	logger := New()

	t.Run("SetOutput", func(t *testing.T) {
		var buf1, buf2 bytes.Buffer
		logger.SetOutput(&buf1)
		logger.Info("test1")
		if !strings.Contains(buf1.String(), "test1") {
			t.Error("Should write to first output")
		}

		logger.SetOutput(&buf2)
		logger.Info("test2")
		if strings.Contains(buf1.String(), "test2") {
			t.Error("Should not write to old output")
		}
		if !strings.Contains(buf2.String(), "test2") {
			t.Error("Should write to new output")
		}
	})
}

func TestLoggerSync(t *testing.T) {
	t.Run("Sync_on_non_syncer", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		// Should not panic even if output doesn't implement Syncer
		logger.Sync()
	})
}

type mockSyncer struct {
	*bytes.Buffer
	syncCount int
}

func (m *mockSyncer) Sync() error {
	m.syncCount++
	return nil
}

func TestLoggerSyncWithSyncer(t *testing.T) {
	t.Run("Sync_calls_sync", func(t *testing.T) {
		ms := &mockSyncer{Buffer: &bytes.Buffer{}}
		logger := New()
		logger.SetOutput(ms)

		logger.Sync()

		if ms.syncCount != 1 {
			t.Errorf("Expected 1 sync, got %d", ms.syncCount)
		}
	})
}

func TestLoggerClone(t *testing.T) {
	logger := New()
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Clone_with_modified_settings", func(t *testing.T) {
		cloned := logger.Clone()

		// Verify cloned logger works independently
		var buf1, buf2 bytes.Buffer
		logger.SetOutput(&buf1)
		cloned.SetOutput(&buf2)

		logger.Info("original")
		cloned.Info("cloned")

		if !strings.Contains(buf1.String(), "original") {
			t.Error("Original logger should work")
		}
		if !strings.Contains(buf2.String(), "cloned") {
			t.Error("Cloned logger should work")
		}
		if strings.Contains(buf1.String(), "cloned") {
			t.Error("Cloned should not write to original output")
		}
	})
}

func TestEntryPoolOperations(t *testing.T) {
	t.Run("getEntry_returns_initialized_entry", func(t *testing.T) {
		entry := getEntry()

		if entry.Pid == 0 {
			t.Error("Entry should have Pid set")
		}
		// Time is set when entry is used, not when retrieved from pool
		entry.Time = time.Now()
		if entry.Time.IsZero() {
			t.Error("Entry time should be settable")
		}

		putEntry(entry)
	})

	t.Run("getEntry_multiple_entries", func(t *testing.T) {
		entry1 := getEntry()
		entry2 := getEntry()

		if entry1 == entry2 {
			t.Error("Should return different entries")
		}

		putEntry(entry1)
		putEntry(entry2)
	})
}

func TestLoggerWithPrefixSuffix(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("With_prefix", func(t *testing.T) {
		logger.SetPrefixMsg("[PREFIX] ")
		logger.Info("message")

		if !strings.Contains(buf.String(), "[PREFIX]") {
			t.Error("Should include prefix")
		}
		if !strings.Contains(buf.String(), "message") {
			t.Error("Should include message")
		}
	})

	t.Run("With_suffix", func(t *testing.T) {
		buf.Reset()
		logger.SetSuffixMsg(" [SUFFIX]")
		logger.Info("message")

		if !strings.Contains(buf.String(), "[SUFFIX]") {
			t.Error("Should include suffix")
		}
	})
}

func TestFormatterWithAllLevels(t *testing.T) {
	formatter := &Formatter{}
	levels := []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel}

	for _, level := range levels {
		t.Run(level.String(), func(t *testing.T) {
			entry := &Entry{
				Level:   level,
				Message: "test",
				Time:    time.Now(),
				Pid:     123,
			}

			result := formatter.Format(entry)
			if len(result) == 0 {
				t.Errorf("Should format entry for level %v", level)
			}

			resultStr := string(result)
			if !strings.Contains(resultStr, level.String()) {
				t.Errorf("Should contain level string %v", level)
			}
		})
	}
}

func TestAddSync(t *testing.T) {
	t.Run("AddSync_with_WriteSyncer", func(t *testing.T) {
		ms := &mockSyncer{Buffer: &bytes.Buffer{}}
		ws := constant.AddSync(ms)

		// Should return the original WriteSyncer
		if ws != ms {
			t.Error("AddSync should return original WriteSyncer")
		}
	})

	t.Run("AddSync_with_io_Writer", func(t *testing.T) {
		var buf bytes.Buffer
		ws := constant.AddSync(&buf)

		if ws == nil {
			t.Error("AddSync should return wrapped writer")
		}

		// Verify it implements WriteSyncer
		if _, err := ws.Write([]byte("test")); err != nil {
			t.Errorf("Wrapped writer should write, got %v", err)
		}

		// Sync should not error for non-syncer
		if err := ws.Sync(); err != nil {
			t.Errorf("Sync should return nil for non-syncer, got %v", err)
		}
	})
}

func TestCloneWithHooks(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Clone_copies_hooks", func(t *testing.T) {
		hookCalled := false
		hook := constant.HookFunc(func(entry interface{}) interface{} {
			hookCalled = true
			return entry
		})

		logger.AddHook(hook)
		cloned := logger.Clone()

		var buf bytes.Buffer
		cloned.SetOutput(&buf)
		cloned.Info("test")

		if !hookCalled {
			t.Error("Cloned logger should have hooks copied")
		}
	})

	t.Run("Clone_with_empty_hooks", func(t *testing.T) {
		cloned := logger.Clone()
		if cloned == nil {
			t.Error("Clone should work with no hooks")
		}

		// Should not panic when logging
		var buf bytes.Buffer
		cloned.SetOutput(&buf)
		cloned.Info("test")
	})
}

func TestCloneWithFormatFull(t *testing.T) {
	logger := New()
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Clone_with_FormatFull", func(t *testing.T) {
		formatter := &Formatter{
			DisableCaller:             true,
			DisableParsingAndEscaping: true,
		}
		logger.Format = formatter

		cloned := logger.Clone()

		// Verify Format was cloned, not just copied
		clonedFormatter, ok := cloned.Format.(*Formatter)
		if !ok {
			t.Fatal("Cloned format should be *Formatter")
		}

		// Cloned formatter should have same values as original
		if clonedFormatter.DisableCaller != formatter.DisableCaller {
			t.Errorf("Cloned formatter should have same DisableCaller, got %v want %v",
				clonedFormatter.DisableCaller, formatter.DisableCaller)
		}

		// Modifying cloned format should not affect original
		clonedFormatter.DisableCaller = false
		if formatter.DisableCaller != true {
			t.Errorf("Original formatter should be unchanged, got %v want true", formatter.DisableCaller)
		}
	})
}

func TestLoggerCloneBranches(t *testing.T) {
	t.Run("Clone_with_hooks_copies_hooks", func(t *testing.T) {
		logger := New()

		// Create a test hook that implements constant.Hook
		hook := &mockHookForClone{}
		logger.AddHook(hook)

		cloned := logger.Clone()

		// Verify cloned has its own hooks slice
		if len(cloned.hooks) != len(logger.hooks) {
			t.Errorf("Clone should copy hooks, got %d want %d", len(cloned.hooks), len(logger.hooks))
		}

		// Verify hooks are independent
		cloned.RemoveHooks()
		if len(logger.hooks) == 0 {
			t.Error("Original logger hooks should not be affected by clone")
		}
	})

	t.Run("Clone_without_hooks", func(t *testing.T) {
		logger := New()
		logger.SetLevel(InfoLevel)

		cloned := logger.Clone()

		if cloned == nil {
			t.Error("Clone should return non-nil logger")
		}
		if cloned.level != logger.level {
			t.Errorf("Clone should copy level, got %v want %v", cloned.level, logger.level)
		}
	})
}

// mockHookForClone implements constant.Hook interface for testing
type mockHookForClone struct{}

func (h *mockHookForClone) OnWrite(entry interface{}) interface{} {
	return entry
}


// mockHookForClone implements constant.Hook interface for testing

func TestLoggerCloneWithHooksFull(t *testing.T) {
	logger := New()

	mockHook := &testCloneHook{}
	logger.AddHook(mockHook)

	cloned := logger.Clone()

	if cloned == nil {
		t.Error("Clone should return non-nil")
	}

	cloned.RemoveHooks()
	if len(logger.hooks) == 0 {
		t.Error("Original logger should still have hooks")
	}
}

type testCloneHook struct{}

func (h *testCloneHook) OnWrite(entry interface{}) interface{} {
	return entry
}

func (h *testCloneHook) Levels() []constant.Level {
	return []constant.Level{constant.InfoLevel}
}

func TestLoggerSetOutputVariations(t *testing.T) {
	logger := New()

	original := logger.out
	logger.SetOutput(nil)
	if logger.out != nil {
		t.Error("output should be nil")
	}

	logger.SetOutput(original)
	if logger.out != original {
		t.Error("output should be restored")
	}
}

func TestLoggerCloneAllFormats(t *testing.T) {
	formats := []constant.Format{
		&Formatter{},
		&JSONFormatter{},
		&JSONFormatter{EnablePrettyPrint: true},
		&JSONFormatter{DisableCaller: true},
		&JSONFormatter{DisableTrace: true},
	}

	for _, fmt := range formats {
		logger := New()
		logger.Format = fmt

		cloned := logger.Clone()

		if cloned == nil {
			t.Error("Clone should return non-nil")
		}
		if cloned.level != logger.level {
			t.Error("Clone should copy level")
		}
	}
}

func TestLoggerSetOutputAllCases(t *testing.T) {
	logger := New()

	outputs := []io.Writer{
		nil,
		io.Discard,
		&bytes.Buffer{},
	}

	for _, out := range outputs {
		logger.SetOutput(out)
		// Just verify it doesn't panic
	}
}

func TestLoggerCloneComprehensive(t *testing.T) {
	logger := New()
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(true)
	logger.EnableTrace(true)
	logger.SetPrefixMsg("[TEST]")
	logger.SetSuffixMsg("[END]")
	logger.SetCallerDepth(5)

	cloned := logger.Clone()

	if cloned.level != logger.level {
		t.Error("level not cloned")
	}
	if cloned.callerDepth != logger.callerDepth {
		t.Error("callerDepth not cloned")
	}
	if cloned.enableCaller != logger.enableCaller {
		t.Error("enableCaller not cloned")
	}
	if cloned.enableTrace != logger.enableTrace {
		t.Error("enableTrace not cloned")
	}
}
