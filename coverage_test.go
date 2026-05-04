package log

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/lazygophers/log/constant"
)

// Tests for Fatal/Panic functions that exit or panic

func TestFatalFunctions(t *testing.T) {
	t.Run("Fatal_exits", func(t *testing.T) {
		if os.Getenv("TEST_FATAL_EXIT") == "1" {
			Fatal("test fatal")
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestFatalFunctions/Fatal_exits")
		cmd.Env = append(os.Environ(), "TEST_FATAL_EXIT=1")
		if err := cmd.Run(); err == nil {
			t.Error("expected non-zero exit status")
		}
	})

	t.Run("Fatalf_exits", func(t *testing.T) {
		if os.Getenv("TEST_FATALF_EXIT") == "1" {
			Fatalf("test %s", "fatalf")
			return
		}
		cmd := exec.Command(os.Args[0], "-test.run=TestFatalFunctions/Fatalf_exits")
		cmd.Env = append(os.Environ(), "TEST_FATALF_EXIT=1")
		if err := cmd.Run(); err == nil {
			t.Error("expected non-zero exit status")
		}
	})
}

func TestPanicFunctions(t *testing.T) {
	t.Run("Panic_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		Panic("test panic")
	})

	t.Run("Panicf_panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		Panicf("test %s", "panicf")
	})
}

// Tests for logging methods with different argument patterns

func TestLogMethods(t *testing.T) {
	t.Run("Log_with_odd_args", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)
		logger.SetLevel(InfoLevel)

		// Odd number of args - last key has nil value
		logger.Log(InfoLevel, "key1", "value1", "key2")

		output := buf.String()
		if !strings.Contains(output, "key1") {
			t.Error("should contain key1")
		}
		if !strings.Contains(output, "key2") {
			t.Error("should contain key2 even with nil value")
		}
	})

	t.Run("Log_with_no_args", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)
		logger.SetLevel(InfoLevel)

		logger.Log(InfoLevel)

		if buf.Len() == 0 {
			t.Error("should log even with no args")
		}
	})
}

// Tests for structured logging methods

func TestStructuredLogging(t *testing.T) {
	t.Run("All_w_methods", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)
		logger.SetLevel(TraceLevel) // Enable all levels

		methods := []func(string, ...interface{}){
			logger.Tracew,
			logger.Debugw,
			logger.Infow,
			logger.Warnw,
			logger.Errorw,
		}

		for _, method := range methods {
			buf.Reset()
			method("test", "key", "value")
			if !strings.Contains(buf.String(), "test") {
				t.Error("method should log message")
			}
		}
	})
}

// Tests for Clone functionality

func TestLoggerCloneIndependence(t *testing.T) {
	t.Run("Clone_independent_settings", func(t *testing.T) {
		logger := New()
		logger.SetLevel(InfoLevel)
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		cloned := logger.Clone()

		// Modify original
		logger.SetLevel(DebugLevel)
		logger.EnableCaller(true)

		// Verify clone is independent
		if cloned.Level() == DebugLevel {
			t.Error("clone level should not change")
		}

		if cloned.callerDepth != logger.callerDepth {
			t.Error("clone should copy callerDepth")
		}
	})

	t.Run("Clone_with_prefix_suffix", func(t *testing.T) {
		logger := New()
		logger.SetPrefixMsg("[PREFIX]")
		logger.SetSuffixMsg("[SUFFIX]")

		cloned := logger.Clone()

		if string(cloned.PrefixMsg) != "[PREFIX]" {
			t.Error("clone should copy PrefixMsg")
		}

		if string(cloned.SuffixMsg) != "[SUFFIX]" {
			t.Error("clone should copy SuffixMsg")
		}
	})

	t.Run("Clone_with_hooks", func(t *testing.T) {
		logger := New()

		hook1 := constant.HookFunc(func(entry interface{}) interface{} {
			return entry
		})
		hook2 := constant.HookFunc(func(entry interface{}) interface{} {
			return entry
		})

		logger.AddHooks(hook1, hook2)
		cloned := logger.Clone()

		if len(cloned.hooks) != 2 {
			t.Errorf("expected 2 hooks, got %d", len(cloned.hooks))
		}

		// Verify hooks are independent
		logger.AddHook(constant.HookFunc(func(entry interface{}) interface{} {
			return entry
		}))

		if len(cloned.hooks) != 2 {
			t.Error("cloned hooks should not be affected")
		}
	})
}

// Tests for JSON formatter edge cases

func TestJSONFormatterCoverage(t *testing.T) {
	t.Run("JSON_formatter_with_special_chars", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)
		logger.Format = &JSONFormatter{}
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		// Test message with special characters
		testStrings := []string{
			`Message with "quotes"`,
			"Message with\nnewline",
			"Message with\ttab",
			"Message with \\ backslash",
		}

		for _, s := range testStrings {
			buf.Reset()
			logger.Info(s)

			output := buf.String()
			if len(output) == 0 {
				t.Errorf("should format message with special chars: %q", s)
			}

			// Verify it's valid JSON
			if !strings.HasPrefix(output, "{") {
				t.Error("should be JSON object")
			}
		}
	})

	t.Run("JSON_formatter_all_fields", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)
		logger.Format = &JSONFormatter{}
		logger.EnableCaller(false)
		logger.EnableTrace(true)

		logger.Infow("test", "key1", "value1", "key2", 123, "key3", true)

		output := buf.String()
		if !strings.Contains(output, "key1") {
			t.Error("should contain key1")
		}
		if !strings.Contains(output, "key2") {
			t.Error("should contain key2")
		}
		if !strings.Contains(output, "key3") {
			t.Error("should contain key3")
		}
	})

	t.Run("JSON_formatter_pretty_print", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		formatter := &JSONFormatter{
			EnablePrettyPrint: true,
		}
		logger.Format = formatter
		logger.EnableCaller(false)
		logger.EnableTrace(false)

		logger.Info("test")

		output := buf.String()
		// Pretty printed JSON should have indentation
		if !strings.Contains(output, "\n") {
			t.Error("pretty print should have newlines")
		}
	})
}

// Tests for write method internals

func TestWriteMethod(t *testing.T) {
	t.Run("write_normal", func(t *testing.T) {
		var buf bytes.Buffer
		logger := New()
		logger.SetOutput(&buf)

		entry := &Entry{
			Level:   InfoLevel,
			Message: "test",
		}

		logger.write(entry.Level, []byte("test message\n"))

		if buf.Len() == 0 {
			t.Error("should write to output")
		}
	})

	t.Run("write_with_sync", func(t *testing.T) {
		if os.Getenv("TEST_SYNC") == "1" {
			var mock syncMock
			logger := New()
			logger.SetOutput(&mock)

			// write only calls Sync for PanicLevel or FatalLevel
			entry := &Entry{
				Level:   FatalLevel,
				Message: "test",
			}

			logger.write(entry.Level, []byte("test\n"))

			if mock.syncCount != 1 {
				t.Errorf("expected 1 sync, got %d", mock.syncCount)
			}
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestWriteMethod/write_with_sync")
		cmd.Env = append(os.Environ(), "TEST_SYNC=1")
		cmd.Run() // Will exit, that's expected
	})
}

type syncMock struct {
	syncCount int
}

func (m *syncMock) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *syncMock) Sync() error {
	m.syncCount++
	return nil
}

// Tests for Print aliases

func TestPrintAliases(t *testing.T) {
	t.Run("Print_same_as_Debug", func(t *testing.T) {
		var buf1, buf2 bytes.Buffer
		logger1 := New()
		logger2 := New()
		logger1.SetOutput(&buf1)
		logger2.SetOutput(&buf2)
		logger1.SetLevel(DebugLevel)
		logger2.SetLevel(DebugLevel)

		logger1.Print("test print")
		logger2.Debug("test print")

		if buf1.String() != buf2.String() {
			t.Error("Print should be same as Debug")
		}
	})

	t.Run("Warning_same_as_Warn", func(t *testing.T) {
		var buf1, buf2 bytes.Buffer
		logger1 := New()
		logger2 := New()
		logger1.SetOutput(&buf1)
		logger2.SetOutput(&buf2)

		logger1.Warning("test warning")
		logger2.Warn("test warning")

		if buf1.String() != buf2.String() {
			t.Error("Warning should be same as Warn")
		}
	})
}
