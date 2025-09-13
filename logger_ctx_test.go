package log

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestLoggerWithCtx_CloneToCtx(t *testing.T) {
	logger := newLogger()
	logger.SetLevel(WarnLevel)
	logger.SetPrefixMsg("TEST: ")

	ctxLogger := logger.CloneToCtx()

	if ctxLogger == nil {
		t.Fatal("CloneToCtx returned nil")
	}

	if ctxLogger.Logger.Level() != WarnLevel {
		t.Errorf("Expected level %v, got %v", WarnLevel, ctxLogger.Logger.Level())
	}

	if string(ctxLogger.Logger.PrefixMsg) != "TEST: " {
		t.Errorf("Expected prefix %q, got %q", "TEST: ", string(ctxLogger.Logger.PrefixMsg))
	}
}

func TestNewLoggerWithCtx(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	if ctxLogger == nil {
		t.Fatal("newLoggerWithCtx returned nil")
	}

	if ctxLogger.Logger == nil {
		t.Error("LoggerWithCtx should have a logger")
	}
}

func TestLoggerWithCtx_SetCallerDepth(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	result := ctxLogger.SetCallerDepth(15)

	if result != ctxLogger {
		t.Error("SetCallerDepth should return the same LoggerWithCtx instance")
	}

	if ctxLogger.Logger.callerDepth != 15 {
		t.Errorf("Expected caller depth 15, got %d", ctxLogger.Logger.callerDepth)
	}
}

func TestLoggerWithCtx_SetPrefixMsg(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	prefix := "CTX_PREFIX: "
	result := ctxLogger.SetPrefixMsg(prefix)

	if result != ctxLogger {
		t.Error("SetPrefixMsg should return the same LoggerWithCtx instance")
	}

	if string(ctxLogger.Logger.PrefixMsg) != prefix {
		t.Errorf("Expected prefix %q, got %q", prefix, string(ctxLogger.Logger.PrefixMsg))
	}
}

func TestLoggerWithCtx_AppendPrefixMsg(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	ctxLogger.SetPrefixMsg("INITIAL: ")
	additional := "ADDED: "

	result := ctxLogger.AppendPrefixMsg(additional)

	if result != ctxLogger {
		t.Error("AppendPrefixMsg should return the same LoggerWithCtx instance")
	}

	expected := "INITIAL: ADDED: "
	if string(ctxLogger.Logger.PrefixMsg) != expected {
		t.Errorf("Expected prefix %q, got %q", expected, string(ctxLogger.Logger.PrefixMsg))
	}
}

func TestLoggerWithCtx_SetSuffixMsg(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	suffix := " :CTX_SUFFIX"
	result := ctxLogger.SetSuffixMsg(suffix)

	if result != ctxLogger {
		t.Error("SetSuffixMsg should return the same LoggerWithCtx instance")
	}

	if string(ctxLogger.Logger.SuffixMsg) != suffix {
		t.Errorf("Expected suffix %q, got %q", suffix, string(ctxLogger.Logger.SuffixMsg))
	}
}

func TestLoggerWithCtx_AppendSuffixMsg(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	ctxLogger.SetSuffixMsg(" :INITIAL")
	additional := " :ADDED"

	result := ctxLogger.AppendSuffixMsg(additional)

	if result != ctxLogger {
		t.Error("AppendSuffixMsg should return the same LoggerWithCtx instance")
	}

	expected := " :INITIAL :ADDED"
	if string(ctxLogger.Logger.SuffixMsg) != expected {
		t.Errorf("Expected suffix %q, got %q", expected, string(ctxLogger.Logger.SuffixMsg))
	}
}

func TestLoggerWithCtx_Clone(t *testing.T) {
	original := newLoggerWithCtx()
	original.SetLevel(ErrorLevel)
	original.SetPrefixMsg("ORIGINAL: ")

	clone := original.Clone()

	if clone == original {
		t.Error("Clone should return a different LoggerWithCtx instance")
	}

	if clone.Logger == original.Logger {
		t.Error("Clone should have a different Logger instance")
	}

	if clone.Logger.Level() != original.Logger.Level() {
		t.Error("Clone should have the same level")
	}

	if string(clone.Logger.PrefixMsg) != string(original.Logger.PrefixMsg) {
		t.Error("Clone should have the same prefix")
	}

	// 修改克隆不应该影响原始对象
	clone.SetLevel(TraceLevel)
	if original.Logger.Level() == TraceLevel {
		t.Error("Modifying clone should not affect original")
	}
}

func TestLoggerWithCtx_SetLevel(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	result := ctxLogger.SetLevel(ErrorLevel)

	if result != ctxLogger {
		t.Error("SetLevel should return the same LoggerWithCtx instance")
	}

	if ctxLogger.Logger.Level() != ErrorLevel {
		t.Errorf("Expected level %v, got %v", ErrorLevel, ctxLogger.Logger.Level())
	}
}

func TestLoggerWithCtx_SetOutput(t *testing.T) {
	ctxLogger := newLoggerWithCtx()
	var buf1, buf2 bytes.Buffer

	result := ctxLogger.SetOutput(&buf1, &buf2)

	if result != ctxLogger {
		t.Error("SetOutput should return the same LoggerWithCtx instance")
	}

	// 测试写入
	ctx := context.Background()
	ctxLogger.Info(ctx, "test message")

	if buf1.Len() == 0 {
		t.Error("First buffer should contain data")
	}

	if buf2.Len() == 0 {
		t.Error("Second buffer should contain data")
	}
}

func TestLoggerWithCtx_Log(t *testing.T) {
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(DebugLevel)

	ctx := context.Background()
	ctxLogger.Log(ctx, InfoLevel, "test log message")

	output := buf.String()
	if !strings.Contains(output, "test log message") {
		t.Errorf("Output should contain log message, got: %s", output)
	}
}

func TestLoggerWithCtx_Logf(t *testing.T) {
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(DebugLevel)

	ctx := context.Background()
	ctxLogger.Logf(ctx, InfoLevel, "test %s message %d", "formatted", 123)

	output := buf.String()
	if !strings.Contains(output, "test formatted message 123") {
		t.Errorf("Output should contain formatted message, got: %s", output)
	}
}

func TestLoggerWithCtx_AllLevels(t *testing.T) {
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(TraceLevel)

	ctx := context.Background()

	// 测试所有级别的日志方法
	ctxLogger.Trace(ctx, "trace message")
	ctxLogger.Debug(ctx, "debug message")
	ctxLogger.Print(ctx, "print message")
	ctxLogger.Info(ctx, "info message")
	ctxLogger.Warn(ctx, "warn message")
	ctxLogger.Warning(ctx, "warning message")
	ctxLogger.Error(ctx, "error message")

	output := buf.String()
	expectedMessages := []string{
		"trace message", "debug message", "print message", "info message",
		"warn message", "warning message", "error message",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Output should contain %q, got: %s", msg, output)
		}
	}
}

func TestLoggerWithCtx_AllFormattedLevels(t *testing.T) {
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(TraceLevel)

	ctx := context.Background()

	// 测试所有格式化级别的日志方法
	ctxLogger.Tracef(ctx, "trace %s", "formatted")
	ctxLogger.Printf(ctx, "print %s", "formatted")
	ctxLogger.Debugf(ctx, "debug %s", "formatted")
	ctxLogger.Infof(ctx, "info %s", "formatted")
	ctxLogger.Warnf(ctx, "warn %s", "formatted")
	ctxLogger.Warningf(ctx, "warning %s", "formatted")
	ctxLogger.Errorf(ctx, "error %s", "formatted")

	output := buf.String()
	expectedMessages := []string{
		"trace formatted", "print formatted", "debug formatted", "info formatted",
		"warn formatted", "warning formatted", "error formatted",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Output should contain %q, got: %s", msg, output)
		}
	}
}

func TestLoggerWithCtx_PanicAndFatal(t *testing.T) {
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(TraceLevel)

	ctx := context.Background()

	// 测试 Panic (应该会 panic，所以用 recover)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic should have panicked")
		}
	}()

	ctxLogger.Panic(ctx, "panic message")
}

func TestLoggerWithCtx_PanicfAndFatalf(t *testing.T) {
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(TraceLevel)

	ctx := context.Background()

	// 测试 Panicf
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panicf should have panicked")
		}
	}()

	ctxLogger.Panicf(ctx, "panic %s", "formatted")
}

func TestLoggerWithCtx_ParsingAndEscaping(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	result := ctxLogger.ParsingAndEscaping(true)

	if result != ctxLogger {
		t.Error("ParsingAndEscaping should return the same LoggerWithCtx instance")
	}

	if formatter, ok := ctxLogger.Logger.Format.(*Formatter); ok {
		if formatter.DisableParsingAndEscaping != true {
			t.Error("Expected DisableParsingAndEscaping to be true")
		}
	}
}

func TestLoggerWithCtx_Caller(t *testing.T) {
	ctxLogger := newLoggerWithCtx()

	result := ctxLogger.Caller(true)

	if result != ctxLogger {
		t.Error("Caller should return the same LoggerWithCtx instance")
	}

	if formatter, ok := ctxLogger.Logger.Format.(*Formatter); ok {
		if formatter.DisableCaller != true {
			t.Error("Expected DisableCaller to be true")
		}
	}
}

func TestLoggerWithCtx_SetOutputNil_Coverage(t *testing.T) {
	// 测试LoggerWithCtx SetOutput中nil writer的处理（logger_ctx.go:88-89）
	ctxLogger := newLoggerWithCtx()

	// 测试传入nil writer
	result := ctxLogger.SetOutput(nil)

	if result != ctxLogger {
		t.Error("SetOutput should return the same LoggerWithCtx instance")
	}

	// out应该被设置为nil
	if ctxLogger.out != nil {
		t.Error("Output should be nil when no writers provided")
	}
}

func TestLoggerWithCtx_SetOutputEmpty_Coverage(t *testing.T) {
	// 测试LoggerWithCtx SetOutput中空writer列表的处理（logger_ctx.go:94-96）
	ctxLogger := newLoggerWithCtx()

	// 测试传入空列表
	result := ctxLogger.SetOutput()

	if result != ctxLogger {
		t.Error("SetOutput should return the same LoggerWithCtx instance")
	}

	// out应该被设置为nil
	if ctxLogger.out != nil {
		t.Error("Output should be nil when no writers provided")
	}
}

func TestLoggerWithCtx_LogLevelFiltering_Coverage(t *testing.T) {
	// 测试LoggerWithCtx Log方法的级别过滤（logger_ctx.go:107-109）
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(ErrorLevel) // 只记录ERROR及以上级别

	ctx := context.Background()

	// 测试被过滤的Log调用
	ctxLogger.Log(ctx, DebugLevel, "debug message")
	ctxLogger.Log(ctx, InfoLevel, "info message")
	ctxLogger.Log(ctx, WarnLevel, "warn message")

	output := buf.String()

	// 验证低级别的消息没有被记录
	if strings.Contains(output, "debug message") {
		t.Error("Debug Log message should be filtered out")
	}
	if strings.Contains(output, "info message") {
		t.Error("Info Log message should be filtered out")
	}
	if strings.Contains(output, "warn message") {
		t.Error("Warn Log message should be filtered out")
	}
}

func TestLoggerWithCtx_LogfLevelFiltering_Coverage(t *testing.T) {
	// 测试LoggerWithCtx Logf方法的级别过滤（logger_ctx.go:116-118）
	var buf bytes.Buffer
	ctxLogger := newLoggerWithCtx()
	ctxLogger.SetOutput(&buf)
	ctxLogger.SetLevel(ErrorLevel) // 只记录ERROR及以上级别

	ctx := context.Background()

	// 测试被过滤的Logf调用
	ctxLogger.Logf(ctx, DebugLevel, "debug %s", "message")
	ctxLogger.Logf(ctx, InfoLevel, "info %s", "message")
	ctxLogger.Logf(ctx, WarnLevel, "warn %s", "message")

	output := buf.String()

	// 验证低级别的消息没有被记录
	if strings.Contains(output, "debug message") {
		t.Error("Debug Logf message should be filtered out")
	}
	if strings.Contains(output, "info message") {
		t.Error("Info Logf message should be filtered out")
	}
	if strings.Contains(output, "warn message") {
		t.Error("Warn Logf message should be filtered out")
	}
}
