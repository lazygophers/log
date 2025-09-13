package log

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

// TestNewLoggerWithCtx tests newLoggerWithCtx function
func TestNewLoggerWithCtx(t *testing.T) {
	logger := newLoggerWithCtx()
	
	require.NotNil(t, logger, "newLoggerWithCtx should return non-nil")
	require.NotNil(t, logger.Logger, "embedded Logger should be initialized")
	assert.IsType(t, &LoggerWithCtx{}, logger, "should return LoggerWithCtx type")
}

// TestLoggerWithCtx_Clone tests LoggerWithCtx Clone method
func TestLoggerWithCtx_Clone(t *testing.T) {
	// Create original logger and set some properties
	original := newLoggerWithCtx()
	original.SetLevel(WarnLevel)
	original.SetPrefixMsg("[ORIGINAL]")
	original.SetSuffixMsg("[END]")
	
	// Clone
	cloned := original.Clone()
	
	require.NotNil(t, cloned, "Clone should return non-nil")
	assert.IsType(t, &LoggerWithCtx{}, cloned, "cloned result should be LoggerWithCtx type")
	
	// Verify properties are correctly copied
	assert.Equal(t, original.Level(), cloned.Level(), "level should be copied")
	
	// Verify independent instances
	original.SetLevel(ErrorLevel)
	assert.Equal(t, WarnLevel, cloned.Level(), "modifying original logger should not affect cloned logger")
}

// TestLoggerWithCtx_SetOutput tests SetOutput method various scenarios
func TestLoggerWithCtx_SetOutput(t *testing.T) {
	logger := newLoggerWithCtx()
	
	t.Run("single output", func(t *testing.T) {
		buf := &bytes.Buffer{}
		result := logger.SetOutput(buf)
		
		// Verify returns self to support chaining
		assert.Same(t, logger, result, "SetOutput should return self")
	})
	
	t.Run("multiple outputs", func(t *testing.T) {
		buf1 := &bytes.Buffer{}
		buf2 := &bytes.Buffer{}
		logger.SetOutput(buf1, buf2)
		
		// Test log writes to all outputs
		ctx := context.Background()
		logger.Info(ctx, "test message")
		
		assert.Contains(t, buf1.String(), "test message", "first output should contain message")
		assert.Contains(t, buf2.String(), "test message", "second output should contain message")
	})
	
	t.Run("empty output list", func(t *testing.T) {
		logger.SetOutput()
		// Empty output will cause nil pointer, this is expected behavior
		ctx := context.Background()
		assert.Panics(t, func() {
			logger.Info(ctx, "this message will panic")
		}, "empty output list will cause panic")
	})
	
	t.Run("nil output", func(t *testing.T) {
		logger.SetOutput(nil)
		ctx := context.Background()
		assert.Panics(t, func() {
			logger.Info(ctx, "nil output test")
		}, "nil output will cause panic")
	})
	
	t.Run("mixed nil and valid output", func(t *testing.T) {
		buf := &bytes.Buffer{}
		logger.SetOutput(nil, buf, nil)
		
		ctx := context.Background()
		logger.Info(ctx, "mixed output test")
		assert.Contains(t, buf.String(), "mixed output test", "valid output should contain message")
	})
}

// TestLoggerWithCtx_Panic tests Panic method
func TestLoggerWithCtx_Panic(t *testing.T) {
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	
	ctx := context.Background()
	
	// Test Panic indeed triggers panic
	require.Panics(t, func() {
		logger.Panic(ctx, "panic test message")
	}, "Panic method should trigger panic")
	
	// Verify log was recorded
	output := buf.String()
	assert.Contains(t, output, "[panic]", "should contain panic level")
	assert.Contains(t, output, "panic test message", "should contain message")
}

// TestLoggerWithCtx_Fatal tests Fatal method
func TestLoggerWithCtx_Fatal(t *testing.T) {
	// Since Fatal calls os.Exit(1), we avoid actual exit by setting high level
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(PanicLevel) // Higher than Fatal level, avoid actual logging and exit
	
	ctx := context.Background()
	
	// When level is set higher than Fatal, won't log or exit
	assert.NotPanics(t, func() {
		logger.Fatal(ctx, "this won't be logged")
	}, "Fatal should not panic when level doesn't match")
	
	// Verify no output indeed
	assert.Empty(t, buf.String(), "should have no Fatal output at high level")
}

// TestLoggerWithCtx_Panicf tests Panicf method
func TestLoggerWithCtx_Panicf(t *testing.T) {
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	
	ctx := context.Background()
	
	// Test Panicf indeed triggers panic
	require.Panics(t, func() {
		logger.Panicf(ctx, "formatted panic: %s %d", "test", 123)
	}, "Panicf method should trigger panic")
	
	// Verify formatted log was recorded
	output := buf.String()
	assert.Contains(t, output, "[panic]", "should contain panic level")
	assert.Contains(t, output, "formatted panic: test 123", "should contain formatted message")
}

// TestLoggerWithCtx_Fatalf tests Fatalf method
func TestLoggerWithCtx_Fatalf(t *testing.T) {
	// Since Fatalf calls os.Exit(1), we avoid actual exit by setting high level
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(PanicLevel) // Higher than Fatal level, avoid actual logging and exit
	
	ctx := context.Background()
	
	// When level is set higher than Fatal, won't log or exit
	assert.NotPanics(t, func() {
		logger.Fatalf(ctx, "formatted fatal: %s", "test")
	}, "Fatalf should not panic when level doesn't match")
	
	// Verify no output indeed
	assert.Empty(t, buf.String(), "should have no Fatalf output at high level")
}

// TestLoggerWithCtx_CallerMethod tests Caller method
func TestLoggerWithCtx_CallerMethod(t *testing.T) {
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	
	ctx := context.Background()
	
	t.Run("enable caller info", func(t *testing.T) {
		buf.Reset()
		result := logger.Caller(true)
		
		// Verify returns self to support chaining
		assert.Same(t, logger, result, "Caller should return self")
		
		// Log to test caller info
		logger.Info(ctx, "test caller info")
		output := buf.String()
		
		// Should contain file name and line number
		assert.Contains(t, output, ".go:", "should contain file info when caller enabled")
	})
	
	t.Run("disable caller info", func(t *testing.T) {
		buf.Reset()
		logger.Caller(false)
		
		// Log to test caller info disabled
		logger.Info(ctx, "test no caller info")
		output := buf.String()
		
		// Based on formatter implementation, disabling caller may still have some format
		// We mainly ensure the method call doesn't error
		assert.Contains(t, output, "test no caller info", "should contain message")
	})
}

// TestLoggerWithCtx_ChainedCalls tests chained method calls
func TestLoggerWithCtx_ChainedCalls(t *testing.T) {
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	
	// Test all methods that support chaining
	result := logger.
		SetLevel(InfoLevel).
		SetOutput(buf).
		SetCallerDepth(3).
		SetPrefixMsg("[CHAIN]").
		AppendPrefixMsg("[TEST]").
		SetSuffixMsg("[END]").
		AppendSuffixMsg("[DONE]").
		ParsingAndEscaping(false).
		Caller(true)
	
	// Verify returns same instance
	assert.Same(t, logger, result, "chained calls should return same instance")
	
	// Test settings take effect
	ctx := context.Background()
	logger.Info(ctx, "chained call test")
	
	output := buf.String()
	assert.Contains(t, output, "[CHAIN][TEST]", "should contain chained prefix")
	assert.Contains(t, output, "[END][DONE]", "should contain chained suffix")
	assert.Contains(t, output, "chained call test", "should contain message")
}

// TestLoggerWithCtx_Integration tests LoggerWithCtx integration functionality
func TestLoggerWithCtx_Integration(t *testing.T) {
	// Save original state
	originalStd := std
	defer func() { std = originalStd }()
	
	// Create test LoggerWithCtx
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf).SetLevel(TraceLevel) // Set to lowest level to include all messages
	
	ctx := context.Background()
	
	t.Run("all log levels", func(t *testing.T) {
		buf.Reset()
		
		logger.Trace(ctx, "trace message")
		logger.Debug(ctx, "debug message")
		logger.Print(ctx, "print message")
		logger.Info(ctx, "info message")
		logger.Warn(ctx, "warn message")
		logger.Warning(ctx, "warning message")
		logger.Error(ctx, "error message")
		
		output := buf.String()
		assert.Contains(t, output, "trace message", "should contain trace message")
		assert.Contains(t, output, "debug message", "should contain debug message")
		assert.Contains(t, output, "print message", "should contain print message")
		assert.Contains(t, output, "info message", "should contain info message")
		assert.Contains(t, output, "warn message", "should contain warn message")
		assert.Contains(t, output, "warning message", "should contain warning message")
		assert.Contains(t, output, "error message", "should contain error message")
	})
	
	t.Run("formatted log levels", func(t *testing.T) {
		buf.Reset()
		
		logger.Tracef(ctx, "trace formatted: %d", 1)
		logger.Debugf(ctx, "debug formatted: %d", 2)
		logger.Printf(ctx, "printf formatted: %d", 3)
		logger.Infof(ctx, "info formatted: %d", 4)
		logger.Warnf(ctx, "warn formatted: %d", 5)
		logger.Warningf(ctx, "warning formatted: %d", 6)
		logger.Errorf(ctx, "error formatted: %d", 7)
		
		output := buf.String()
		assert.Contains(t, output, "trace formatted: 1", "should contain formatted trace")
		assert.Contains(t, output, "debug formatted: 2", "should contain formatted debug")
		assert.Contains(t, output, "printf formatted: 3", "should contain formatted printf")
		assert.Contains(t, output, "info formatted: 4", "should contain formatted info")
		assert.Contains(t, output, "warn formatted: 5", "should contain formatted warn")
		assert.Contains(t, output, "warning formatted: 6", "should contain formatted warning")
		assert.Contains(t, output, "error formatted: 7", "should contain formatted error")
	})
}

// TestLoggerWithCtx_ContextHandling tests context handling
func TestLoggerWithCtx_ContextHandling(t *testing.T) {
	logger := newLoggerWithCtx()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	
	t.Run("normal context", func(t *testing.T) {
		ctx := context.Background()
		logger.Info(ctx, "normal context test")
		
		assert.Contains(t, buf.String(), "normal context test", "should handle normal context")
	})
	
	t.Run("cancelled context", func(t *testing.T) {
		buf.Reset()
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel context immediately
		
		// Even if context is cancelled, logging should still work
		logger.Info(ctx, "cancelled context test")
		
		assert.Contains(t, buf.String(), "cancelled context test", "cancelled context should not affect logging")
	})
}