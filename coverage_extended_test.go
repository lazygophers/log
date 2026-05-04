package log

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestJSONFormatterSpecialCases(t *testing.T) {
	t.Run("jsonEscapeString_with_control_chars", func(t *testing.T) {
		formatter := &JSONFormatter{}

		// Test various control characters that trigger hex encoding
		testCases := []struct {
			input    string
			contains string
		}{
			{"\x01", "\\u00"},     // SOH
			{"\x1F", "\\u00"},     // US
			{"Hello\x00World", "\\u00"}, // NULL in middle
		}

		for _, tc := range testCases {
			entry := &Entry{
				Level:   InfoLevel,
				Message: tc.input,
				Time:    time.Now(),
			}

			result := formatter.Format(entry)
			resultStr := string(result)

			if !strings.Contains(resultStr, tc.contains) {
				t.Errorf("Expected %q to contain %q, got %q", tc.input, tc.contains, resultStr)
			}
		}
	})

	t.Run("jsonEscapeString_all_escapes", func(t *testing.T) {
		formatter := &JSONFormatter{}

		// Test all escape sequences
		testCases := []string{
			"Quote: \"",
			"Backslash: \\",
			"Newline: \n",
			"Carriage Return: \r",
			"Tab: \t",
		}

		for _, msg := range testCases {
			entry := &Entry{
				Level:   InfoLevel,
				Message: msg,
				Time:    time.Now(),
			}

			result := formatter.Format(entry)
			if len(result) == 0 {
				t.Errorf("Should format message with escape: %q", msg)
			}
		}
	})

	t.Run("jsonMarshal_fallback", func(t *testing.T) {
		formatter := &JSONFormatter{}

		// Create a channel (unmarshallable type) to trigger JSON marshaling error
		ch := make(chan int)
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
			Time:    time.Now(),
			Fields: []KV{
				{Key: "unmarshallable", Value: ch},
			},
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		// Should fallback to error message with jsonEscapeString
		if !strings.Contains(resultStr, "JSON marshaling failed") {
			t.Error("Should include marshaling error in fallback")
		}
		if !strings.Contains(resultStr, "test") {
			t.Error("Should include original message in fallback")
		}
	})

	t.Run("jsonMarshal_fallback_with_control_chars", func(t *testing.T) {
		formatter := &JSONFormatter{}

		// Create a channel (unmarshallable type) to trigger JSON marshaling error
		ch := make(chan int)
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test\x01message\x02", // Include control characters
			Time:    time.Now(),
			Fields: []KV{
				{Key: "unmarshallable", Value: ch},
			},
		}

		result := formatter.Format(entry)
		resultStr := string(result)

		// Should fallback to error message with jsonEscapeString (including hexByte)
		if !strings.Contains(resultStr, "JSON marshaling failed") {
			t.Error("Should include marshaling error in fallback")
		}
		// Control characters should be escaped as \u00
		if !strings.Contains(resultStr, "\\u00") {
			t.Error("Should escape control characters using hexByte")
		}
	})
}

func TestLoggerWMethodsFull(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	methods := []struct {
		name   string
		method func(string, ...interface{})
	}{
		{"Debugw", logger.Debugw},
		{"Infow", logger.Infow},
		{"Warnw", logger.Warnw},
		{"Errorw", logger.Errorw},
	}

	for _, m := range methods {
		t.Run(m.name, func(t *testing.T) {
			buf.Reset()
			m.method("test message", "key1", "value1", "key2", "value2")

			output := buf.String()
			if !strings.Contains(output, "test message") {
				t.Errorf("%s should log message", m.name)
			}
			if !strings.Contains(output, "key1") {
				t.Errorf("%s should include key1 field", m.name)
			}
		})
	}
}

func TestLoggerWMethodsDisabled(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(ErrorLevel) // Only Error and above
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	methods := []struct {
		name   string
		method func(string, ...interface{})
	}{
		{"Debugw", logger.Debugw},
		{"Infow", logger.Infow},
		{"Warnw", logger.Warnw},
	}

	for _, m := range methods {
		t.Run(m.name, func(t *testing.T) {
			buf.Reset()
			m.method("should not log", "key", "value")

			if buf.Len() != 0 {
				t.Errorf("%s should not log when level is disabled", m.name)
			}
		})
	}
}
