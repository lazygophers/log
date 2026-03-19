package log

import (
	"context"
	"io"
	"os"
	"os/exec"
	"testing"
)

// runSubprocess runs the current test binary in a subprocess with the given env var set.
// It returns the exit code of the subprocess.
func runSubprocess(t *testing.T, testName, envKey string) int {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=^"+testName+"$")
	cmd.Env = append(os.Environ(), envKey+"=1")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	err := cmd.Run()
	if err == nil {
		return 0
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode()
	}
	t.Fatalf("unexpected error running subprocess: %v", err)
	return -1
}

// --- Logger.Fatal ---

func TestLogger_Fatal_ExitCode(t *testing.T) {
	if os.Getenv("TEST_FATAL_LOGGER") == "1" {
		logger := newLogger()
		logger.SetOutput(io.Discard)
		logger.Fatal("fatal error")
		return
	}

	exitCode := runSubprocess(t, "TestLogger_Fatal_ExitCode", "TEST_FATAL_LOGGER")
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}

func TestLogger_Fatalf_ExitCode(t *testing.T) {
	if os.Getenv("TEST_FATALF_LOGGER") == "1" {
		logger := newLogger()
		logger.SetOutput(io.Discard)
		logger.Fatalf("fatal error %s %d", "formatted", 42)
		return
	}

	exitCode := runSubprocess(t, "TestLogger_Fatalf_ExitCode", "TEST_FATALF_LOGGER")
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}

// --- Global Fatal (print.go) ---

func TestGlobal_Fatal_ExitCode(t *testing.T) {
	if os.Getenv("TEST_FATAL_GLOBAL") == "1" {
		std.SetOutput(io.Discard)
		Fatal("global fatal error")
		return
	}

	exitCode := runSubprocess(t, "TestGlobal_Fatal_ExitCode", "TEST_FATAL_GLOBAL")
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}

func TestGlobal_Fatalf_ExitCode(t *testing.T) {
	if os.Getenv("TEST_FATALF_GLOBAL") == "1" {
		std.SetOutput(io.Discard)
		Fatalf("global fatal %s %d", "formatted", 99)
		return
	}

	exitCode := runSubprocess(t, "TestGlobal_Fatalf_ExitCode", "TEST_FATALF_GLOBAL")
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}

// --- LoggerWithCtx Fatal ---

func TestLoggerWithCtx_Fatal_ExitCode(t *testing.T) {
	if os.Getenv("TEST_FATAL_CTX") == "1" {
		ctxLogger := newLoggerWithCtx()
		ctxLogger.SetOutput(io.Discard)
		ctxLogger.Fatal(context.Background(), "ctx fatal error")
		return
	}

	exitCode := runSubprocess(t, "TestLoggerWithCtx_Fatal_ExitCode", "TEST_FATAL_CTX")
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}

func TestLoggerWithCtx_Fatalf_ExitCode(t *testing.T) {
	if os.Getenv("TEST_FATALF_CTX") == "1" {
		ctxLogger := newLoggerWithCtx()
		ctxLogger.SetOutput(io.Discard)
		ctxLogger.Fatalf(context.Background(), "ctx fatal %s %d", "formatted", 77)
		return
	}

	exitCode := runSubprocess(t, "TestLoggerWithCtx_Fatalf_ExitCode", "TEST_FATALF_CTX")
	if exitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", exitCode)
	}
}

// --- LoggerWithCtx Fatal with cancelled context (should NOT exit) ---

func TestLoggerWithCtx_Fatal_CancelledContext(t *testing.T) {
	if os.Getenv("TEST_FATAL_CTX_CANCELLED") == "1" {
		ctxLogger := newLoggerWithCtx()
		ctxLogger.SetOutput(io.Discard)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		ctxLogger.Fatal(ctx, "should not exit")
		// If we reach here, the process exits with 0 (test passes)
		os.Exit(0)
		return
	}

	exitCode := runSubprocess(t, "TestLoggerWithCtx_Fatal_CancelledContext", "TEST_FATAL_CTX_CANCELLED")
	if exitCode != 0 {
		t.Errorf("Expected exit code 0 (cancelled context should skip fatal), got %d", exitCode)
	}
}

// --- Logger.Fatal level filtering (level too low to trigger fatal) ---

func TestLogger_Fatal_LevelFiltered(t *testing.T) {
	if os.Getenv("TEST_FATAL_FILTERED") == "1" {
		logger := newLogger()
		logger.SetOutput(io.Discard)
		logger.SetLevel(PanicLevel) // PanicLevel(0) < FatalLevel(1), so Fatal should be filtered
		logger.Fatal("should not exit")
		os.Exit(0)
		return
	}

	exitCode := runSubprocess(t, "TestLogger_Fatal_LevelFiltered", "TEST_FATAL_FILTERED")
	if exitCode != 0 {
		t.Errorf("Expected exit code 0 (fatal should be filtered by level), got %d", exitCode)
	}
}

func TestLogger_Fatalf_LevelFiltered(t *testing.T) {
	if os.Getenv("TEST_FATALF_FILTERED") == "1" {
		logger := newLogger()
		logger.SetOutput(io.Discard)
		logger.SetLevel(PanicLevel)
		logger.Fatalf("should not exit %s", "filtered")
		os.Exit(0)
		return
	}

	exitCode := runSubprocess(t, "TestLogger_Fatalf_LevelFiltered", "TEST_FATALF_FILTERED")
	if exitCode != 0 {
		t.Errorf("Expected exit code 0 (fatalf should be filtered by level), got %d", exitCode)
	}
}
