package log

import (
	"bytes"
	"strings"
	"testing"
)

// Test all logger methods to improve coverage

func TestLoggerAllMethods(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Test simple methods
	simpleMethods := []struct {
		name string
		fn   func(...interface{})
	}{
		{"Trace", logger.Trace},
		{"Debug", logger.Debug},
		{"Info", logger.Info},
		{"Warn", logger.Warn},
		{"Error", logger.Error},
		{"Print", logger.Print},
		{"Warning", logger.Warning},
	}

	for _, m := range simpleMethods {
		t.Run(m.name, func(t *testing.T) {
			buf.Reset()
			m.fn("test message")
			if !strings.Contains(buf.String(), "test message") {
				t.Errorf("%s should log message", m.name)
			}
		})
	}

	// Test formatted methods
	formattedMethods := []struct {
		name string
		fn   func(string, ...interface{})
	}{
		{"Tracef", logger.Tracef},
		{"Debugf", logger.Debugf},
		{"Infof", logger.Infof},
		{"Warnf", logger.Warnf},
		{"Errorf", logger.Errorf},
		{"Printf", logger.Printf},
	}

	for _, m := range formattedMethods {
		t.Run(m.name, func(t *testing.T) {
			buf.Reset()
			m.fn("test %s", "formatted")
			if !strings.Contains(buf.String(), "test formatted") {
				t.Errorf("%s should format message", m.name)
			}
		})
	}

	// Test structured methods
	structuredMethods := []struct {
		name string
		fn   func(string, ...interface{})
	}{
		{"Tracew", logger.Tracew},
		{"Debugw", logger.Debugw},
		{"Infow", logger.Infow},
		{"Warnw", logger.Warnw},
		{"Errorw", logger.Errorw},
	}

	for _, m := range structuredMethods {
		t.Run(m.name, func(t *testing.T) {
			buf.Reset()
			m.fn("test", "key1", "value1", "key2", "value2")
			output := buf.String()
			if !strings.Contains(output, "test") {
				t.Errorf("%sw should log message", m.name)
			}
			if !strings.Contains(output, "key1") {
				t.Errorf("%sw should include fields", m.name)
			}
		})
	}
}

func TestLoggerMethodsDisabled(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(ErrorLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	// Test that methods respect level setting
	methods := []struct {
		name string
		fn   func(...interface{})
	}{
		{"Trace", logger.Trace},
		{"Debug", logger.Debug},
		{"Info", logger.Info},
		{"Warn", logger.Warn},
	}

	for _, m := range methods {
		t.Run(m.name+"_disabled", func(t *testing.T) {
			buf.Reset()
			m.fn("test")
			if buf.Len() != 0 {
				t.Errorf("%s should not log when disabled", m.name)
			}
		})
	}

	// Error should still work
	t.Run("Error_enabled", func(t *testing.T) {
		buf.Reset()
		logger.Error("test")
		if !strings.Contains(buf.String(), "test") {
			t.Error("Error should log when enabled")
		}
	})
}

func TestLoggerLogMethod(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel) // Enable all levels
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("Log_with_various_levels", func(t *testing.T) {
		levels := []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel}
		for _, level := range levels {
			buf.Reset()
			logger.Log(level, "test message")
			if buf.Len() == 0 {
				t.Errorf("Log should work at level %v", level)
			}
		}
	})

	t.Run("Log_with_fields", func(t *testing.T) {
		buf.Reset()
		logger.Log(InfoLevel, "test", "key1", "value1", "key2", "value2")
		output := buf.String()
		if !strings.Contains(output, "test") {
			t.Error("Log should include message")
		}
		if !strings.Contains(output, "key1") {
			t.Error("Log should include fields")
		}
	})

	t.Run("Log_disabled_level", func(t *testing.T) {
		logger.SetLevel(ErrorLevel)
		buf.Reset()
		logger.Log(InfoLevel, "test") // Info is disabled when level is Error
		if buf.Len() != 0 {
			t.Error("Log should not log when level is disabled")
		}
	})
}

func TestLoggerLevelEnabledCoverage(t *testing.T) {
	logger := New()

	t.Run("levelEnabled_various_levels", func(t *testing.T) {
		levels := []Level{
			TraceLevel,
			DebugLevel,
			InfoLevel,
			WarnLevel,
			ErrorLevel,
			PanicLevel,
			FatalLevel,
		}

		for _, setLevel := range levels {
			logger.SetLevel(setLevel)

			// Current level should be enabled
			if !logger.levelEnabled(setLevel) {
				t.Errorf("levelEnabled(%v) should be true when set to that level", setLevel)
			}

			// Next level up should be disabled (if exists)
			if setLevel < FatalLevel {
				higherLevel := Level(int(setLevel) + 1)
				if logger.levelEnabled(higherLevel) {
					t.Errorf("levelEnabled(%v) should be false when set to lower level", higherLevel)
				}
			}
		}
	})
}

func TestLoggerPopulateFieldsCoverage(t *testing.T) {
	logger := New()

	t.Run("populateFields_even_args", func(t *testing.T) {
		entry := &Entry{}
		logger.populateFields(entry, "k1", "v1", "k2", "v2")
		if len(entry.Fields) != 2 {
			t.Errorf("Expected 2 fields, got %d", len(entry.Fields))
		}
	})

	t.Run("populateFields_odd_args", func(t *testing.T) {
		entry := &Entry{}
		logger.populateFields(entry, "k1", "v1", "k2")
		if len(entry.Fields) != 2 {
			t.Errorf("Expected 2 fields (last with nil), got %d", len(entry.Fields))
		}
		if entry.Fields[1].Value != nil {
			t.Errorf("Last field value should be nil, got %v", entry.Fields[1].Value)
		}
	})

	t.Run("populateFields_no_args", func(t *testing.T) {
		entry := &Entry{}
		logger.populateFields(entry)
		if len(entry.Fields) != 0 {
			t.Errorf("Expected 0 fields, got %d", len(entry.Fields))
		}
	})

	t.Run("populateFields_many_fields", func(t *testing.T) {
		entry := &Entry{}
		logger.populateFields(entry,
			"k1", "v1",
			"k2", "v2",
			"k3", "v3",
			"k4", "v4",
			"k5", "v5",
			"k6", "v6",
		)
		if len(entry.Fields) != 6 {
			t.Errorf("Expected 6 fields, got %d", len(entry.Fields))
		}
	})
}

func TestLoggerWriteCoverage(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(InfoLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	t.Run("write_normal", func(t *testing.T) {
		buf.Reset()
		logger.write(InfoLevel, []byte("test message\n"))
		if buf.Len() == 0 {
			t.Error("write should write to output")
		}
		if !strings.Contains(buf.String(), "test message") {
			t.Error("write should write correct message")
		}
	})

	t.Run("write_empty", func(t *testing.T) {
		buf.Reset()
		logger.write(InfoLevel, []byte{})
		// Should not panic
	})
}
