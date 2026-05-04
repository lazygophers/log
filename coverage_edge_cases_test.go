package log

import (
	"bytes"
	"strings"
	"testing"
	"time"
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
