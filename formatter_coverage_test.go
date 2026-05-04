package log

import (
	"strings"
	"testing"
	"time"
)

func TestJSONFormatterEdgeCases(t *testing.T) {
	t.Run("jsonEscapeString_control_chars_full_coverage", func(t *testing.T) {
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test\x00\x01\x1f\"\\\\",  // NULL, SOH, US, quote, backslash
			Time:    time.Now(),
			Pid:     123,
		}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format entry with control chars")
		}

		resultStr := string(result)
		if !strings.Contains(resultStr, "\\u00") {
			t.Error("Should escape control chars")
		}
		if !strings.Contains(resultStr, "\\\"") {
			t.Error("Should escape quotes")
		}
		if !strings.Contains(resultStr, "\\\\") {
			t.Error("Should escape backslash")
		}
	})

	t.Run("jsonEscapeString_all_special_chars", func(t *testing.T) {
		testCases := []struct {
			input    string
			contains []string
		}{
			{"\n", []string{"\\n"}},
			{"\r", []string{"\\r"}},
			{"\t", []string{"\\t"}},
			{"\"", []string{"\\\""}},
			{"\\", []string{"\\\\"}},
		}

		for _, tc := range testCases {
			formatter := &JSONFormatter{}
			entry := &Entry{
				Level:   InfoLevel,
				Message: tc.input,
				Time:    time.Now(),
				Pid:     123,
			}

			result := formatter.Format(entry)
			if len(result) == 0 {
				t.Errorf("Should format message with %q", tc.input)
			}

			for _, c := range tc.contains {
				if !strings.Contains(string(result), c) {
					t.Errorf("Should escape %s in message %q", c, tc.input)
				}
			}
		}
	})

	t.Run("jsonEscapeString_hexByte_coverage", func(t *testing.T) {
		// Test hexByte function coverage by including chars that trigger it
		formatter := &JSONFormatter{}
		entry := &Entry{
			Level:   InfoLevel,
			Message: "test\x01\x02\x03\x10\x11\x1f",  // Control chars that trigger hexByte
			Time:    time.Now(),
			Pid:     123,
		}

		result := formatter.Format(entry)
		if len(result) == 0 {
			t.Error("Should format entry with hex byte chars")
		}

		// Verify hex escaping happened
		resultStr := string(result)
		if !strings.Contains(resultStr, "\\u00") {
			t.Error("Should use hexByte for control chars")
		}
	})
}
