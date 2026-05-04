package log

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Tests for Fatalw/Panicw methods

func TestFatalwMethod(t *testing.T) {
	t.Run("Fatalw_exits", func(t *testing.T) {
		if os.Getenv("TEST_FATALW_EXIT") == "1" {
			logger := New()
			logger.SetLevel(FatalLevel)
			logger.Fatalw("fatal message", "key", "value")
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestFatalwMethod/Fatalw_exits")
		cmd.Env = append(os.Environ(), "TEST_FATALW_EXIT=1")
		err := cmd.Run()
		if err == nil {
			t.Error("expected non-zero exit status")
		}
	})
}

func TestPanicwMethod(t *testing.T) {
	t.Run("Panicw_panics", func(t *testing.T) {
		if os.Getenv("TEST_PANICW_PANIC") == "1" {
			logger := New()
			logger.SetLevel(PanicLevel)
			logger.Panicw("panic message", "key", "value")
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestPanicwMethod/Panicw_panics")
		cmd.Env = append(os.Environ(), "TEST_PANICW_PANIC=1")
		err := cmd.Run()
		if err == nil {
			t.Error("expected non-zero exit status from panic")
		}
	})
}

func TestFatalfAndPanicfMethods(t *testing.T) {
	t.Run("Fatalf_exits", func(t *testing.T) {
		if os.Getenv("TEST_FATALF_EXIT") == "1" {
			logger := New()
			logger.SetLevel(FatalLevel)
			logger.Fatalf("fatal %s", "message")
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestFatalfAndPanicfMethods/Fatalf_exits")
		cmd.Env = append(os.Environ(), "TEST_FATALF_EXIT=1")
		err := cmd.Run()
		if err == nil {
			t.Error("expected non-zero exit status")
		}
	})

	t.Run("Panicf_panics", func(t *testing.T) {
		if os.Getenv("TEST_PANICF_PANIC") == "1" {
			logger := New()
			logger.SetLevel(PanicLevel)
			logger.Panicf("panic %s", "message")
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestFatalfAndPanicfMethods/Panicf_panics")
		cmd.Env = append(os.Environ(), "TEST_PANICF_PANIC=1")
		err := cmd.Run()
		if err == nil {
			t.Error("expected non-zero exit status from panic")
		}
	})
}

func TestLoggerAllLevelMethods(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	methods := []struct {
		name string
		fn   func(...interface{})
	}{
		{"Trace", logger.Trace},
		{"Debug", logger.Debug},
		{"Info", logger.Info},
		{"Warn", logger.Warn},
		{"Error", logger.Error},
	}

	for _, m := range methods {
		t.Run(m.name+"_enabled", func(t *testing.T) {
			buf.Reset()
			m.fn("test message")
			if !strings.Contains(buf.String(), "test message") {
				t.Errorf("%s should log message", m.name)
			}
		})
	}

	// Test with disabled level
	logger.SetLevel(ErrorLevel)

	disabledMethods := []struct {
		name string
		fn   func(...interface{})
	}{
		{"Trace", logger.Trace},
		{"Debug", logger.Debug},
		{"Info", logger.Info},
		{"Warn", logger.Warn},
	}

	for _, m := range disabledMethods {
		t.Run(m.name+"_disabled", func(t *testing.T) {
			buf.Reset()
			m.fn("should not log")
			if buf.Len() != 0 {
				t.Errorf("%s should not log when disabled", m.name)
			}
		})
	}
}

func TestLoggerAllLevelfMethods(t *testing.T) {
	var buf bytes.Buffer
	logger := New()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)
	logger.EnableCaller(false)
	logger.EnableTrace(false)

	methods := []struct {
		name string
		fn   func(string, ...interface{})
	}{
		{"Tracef", logger.Tracef},
		{"Debugf", logger.Debugf},
		{"Infof", logger.Infof},
		{"Warnf", logger.Warnf},
		{"Errorf", logger.Errorf},
	}

	for _, m := range methods {
		t.Run(m.name+"_enabled", func(t *testing.T) {
			buf.Reset()
			m.fn("test %s", "formatted")
			if !strings.Contains(buf.String(), "test formatted") {
				t.Errorf("%s should log formatted message", m.name)
			}
		})
	}

	// Test with disabled level
	logger.SetLevel(ErrorLevel)

	disabledMethods := []struct {
		name string
		fn   func(string, ...interface{})
	}{
		{"Tracef", logger.Tracef},
		{"Debugf", logger.Debugf},
		{"Infof", logger.Infof},
		{"Warnf", logger.Warnf},
	}

	for _, m := range disabledMethods {
		t.Run(m.name+"_disabled", func(t *testing.T) {
			buf.Reset()
			m.fn("should not log %s", "arg")
			if buf.Len() != 0 {
				t.Errorf("%s should not log when disabled", m.name)
			}
		})
	}
}
