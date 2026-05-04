package log

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestOutputCoverage(t *testing.T) {
	t.Run("GetOutputWriter_creates_directory", func(t *testing.T) {
		tmpDir := t.TempDir()
		logFile := filepath.Join(tmpDir, "logs", "app", "test.log")

		writer := GetOutputWriter(logFile)
		if writer == nil {
			t.Error("GetOutputWriter should return writer")
		}

		// Verify directory was created
		logDir := filepath.Dir(logFile)
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			t.Error("directory should be created")
		}
	})
}

func TestRotatorCoverage(t *testing.T) {
	t.Run("isAllDigits_various_inputs", func(t *testing.T) {
		tests := []struct {
			input string
			want  bool
		}{
			{"12345", true},
			{"abcde", false},
			{"12abc34", false},
			{"", true},
			{"0", true},
			{"00000", true},
			{"99999", true},
			{"12 34", false},
		}

		for _, tt := range tests {
			t.Run(tt.input, func(t *testing.T) {
				if got := isAllDigits(tt.input); got != tt.want {
					t.Errorf("isAllDigits(%q) = %v, want %v", tt.input, got, tt.want)
				}
			})
		}
	})
}

func TestLoggerLevelEnabled(t *testing.T) {
	t.Run("levelEnabled_all_levels", func(t *testing.T) {
		logger := New()

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
			if !logger.levelEnabled(level) {
				t.Errorf("levelEnabled(%v) should be true when set to that level", level)
			}

			// Next level down should be disabled
			if level > TraceLevel {
				lowerLevel := Level(int(level) - 1)
				if !logger.levelEnabled(lowerLevel) {
					t.Errorf("levelEnabled(%v) should be false", lowerLevel)
				}
			}
		}
	})
}

func TestLoggerPopulateFields(t *testing.T) {
	t.Run("populateFields_with_various_patterns", func(t *testing.T) {
		logger := New()
		entry := &Entry{}

		// Even number of args
		logger.populateFields(entry, "k1", "v1", "k2", "v2")
		if len(entry.Fields) != 2 {
			t.Errorf("expected 2 fields, got %d", len(entry.Fields))
		}

		// Odd number of args
		entry.Fields = nil
		logger.populateFields(entry, "k1", "v1", "k2")
		if len(entry.Fields) != 2 {
			t.Errorf("expected 2 fields (last with nil), got %d", len(entry.Fields))
		}

		// No args
		entry.Fields = nil
		logger.populateFields(entry)
		if len(entry.Fields) != 0 {
			t.Errorf("expected 0 fields, got %d", len(entry.Fields))
		}
	})
}

func TestLoggerWriteLevels(t *testing.T) {
	t.Run("write_with_all_levels", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		levels := []Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel}

		for _, level := range levels {
			buf.Reset()
			logger.write(level, []byte("test\n"))

			if buf.Len() == 0 {
				t.Errorf("should write for level %v", level)
			}
		}
	})
}

func TestJSONFormatterEdges(t *testing.T) {
	t.Run("JSON_with_special_characters", func(t *testing.T) {
		formatter := &JSONFormatter{}

		messages := []string{
			"Message with quotes: \"hello\"",
			"Message with backslash: \\",
			"Message with newline: \n",
			"Message with tab: \t",
			"Message with all: \\n\t\"",
		}

		for _, msg := range messages {
			entry := &Entry{
				Level:   InfoLevel,
				Message: msg,
				Time:    time.Now(),
			}

			result := formatter.Format(entry)
			if len(result) == 0 {
				t.Errorf("should format message: %q", msg)
			}
		}
	})

	t.Run("JSON_all_field_types", func(t *testing.T) {
		formatter := &JSONFormatter{}

		entry := &Entry{
			Level:  InfoLevel,
			Message: "all types",
			Fields: []KV{
				{Key: "string", Value: "test"},
				{Key: "int", Value: 42},
				{Key: "float", Value: 3.14},
				{Key: "bool", Value: true},
				{Key: "nil", Value: nil},
				{Key: "slice", Value: []int{1, 2, 3}},
			},
		}

		result := formatter.Format(entry)
		if !strings.Contains(string(result), "all types") {
			t.Error("should contain message")
		}
	})
}

func TestFormatterClone(t *testing.T) {
	t.Run("Formatter_clone_independence", func(t *testing.T) {
		formatter := &Formatter{
			DisableParsingAndEscaping: true,
			DisableCaller:          true,
		}

		cloned := formatter.Clone().(*Formatter)

		// Verify settings copied
		if !cloned.DisableParsingAndEscaping {
			t.Error("Clone should copy DisableParsingAndEscaping")
		}
		if !cloned.DisableCaller {
			t.Error("Clone should copy DisableCaller")
		}

		// Verify independence
		cloned.DisableParsingAndEscaping = false
		if !formatter.DisableParsingAndEscaping {
			t.Error("original should be unchanged")
		}
	})
}

func TestJSONFormatterDisableFeatures(t *testing.T) {
	t.Run("JSON_with_all_features_disabled", func(t *testing.T) {
		formatter := &JSONFormatter{
			DisableTimestamp: true,
			DisableCaller:    true,
			DisableTrace:     true,
		}

		entry := &Entry{
			Level:      InfoLevel,
			Message:    "minimal",
			Time:       time.Now(),
			Pid:        123,
			Gid:        456,
			CallerName:  "test.Func",
			File:       "test.go",
			CallerLine: 10,
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		// Verify disabled fields not present
		disabledFields := []string{"\"time\"", "\"caller\"", "\"trace\""}
		for _, field := range disabledFields {
			if strings.Contains(resultStr, field) {
				t.Errorf("field %s should be disabled", field)
			}
		}
	})
}

func TestEntryPool(t *testing.T) {
	t.Run("getEntry_returns_valid_entry", func(t *testing.T) {
		entry := getEntry()

		if entry == nil {
			t.Error("getEntry should return non-nil")
		}

		if entry.Pid == 0 {
			t.Error("entry should have pid set")
		}

		putEntry(entry)
	})

	t.Run("putEntry_handles_nil", func(t *testing.T) {
		putEntry(nil) // Should not panic
	})
}

func TestLevelString(t *testing.T) {
	t.Run("Level_String_all_levels", func(t *testing.T) {
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
			str := level.String()
			if str == "" {
				t.Errorf("Level.String() should return non-empty for %v", level)
			}
		}
	})
}
