package log

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestLogger_AllLoggingMethods(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel) // 设置最低级别以捕获所有日志

	// 测试所有日志级别方法
	logger.Trace("trace message")
	logger.Debug("debug message")
	logger.Print("print message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Warning("warning message")
	logger.Error("error message")

	output := buf.String()

	expectedMessages := []string{
		"trace message",
		"debug message",
		"print message",
		"info message",
		"warn message",
		"warning message",
		"error message",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Output should contain %q, got: %s", msg, output)
		}
	}
}

func TestLogger_AllFormattedLoggingMethods(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)

	// 测试所有格式化日志方法
	logger.Tracef("trace %s %d", "formatted", 1)
	logger.Printf("print %s %d", "formatted", 2)
	logger.Debugf("debug %s %d", "formatted", 3)
	logger.Infof("info %s %d", "formatted", 4)
	logger.Warnf("warn %s %d", "formatted", 5)
	logger.Warningf("warning %s %d", "formatted", 6)
	logger.Errorf("error %s %d", "formatted", 7)

	output := buf.String()

	expectedMessages := []string{
		"trace formatted 1",
		"print formatted 2",
		"debug formatted 3",
		"info formatted 4",
		"warn formatted 5",
		"warning formatted 6",
		"error formatted 7",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Output should contain %q, got: %s", msg, output)
		}
	}
}

func TestLogger_PanicAndFatal(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)

	// 测试 Panic 方法
	t.Run("Panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic should have panicked")
			}
		}()
		logger.Panic("panic message")
	})

	// 验证 panic 消息是否被记录
	if !strings.Contains(buf.String(), "panic message") {
		t.Errorf("Output should contain panic message, got: %s", buf.String())
	}
}

func TestLogger_Panicf(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)

	t.Run("Panicf", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panicf should have panicked")
			}
		}()
		logger.Panicf("panic %s %d", "formatted", 123)
	})

	if !strings.Contains(buf.String(), "panic formatted 123") {
		t.Errorf("Output should contain formatted panic message, got: %s", buf.String())
	}
}

func TestLogger_Fatal(t *testing.T) {
	// Fatal 会调用 os.Exit，所以我们不能直接测试 Fatal 方法
	// 而是测试格式化行为来验证 Fatal 级别的输出格式
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             false,
	}

	entry := getEntry()
	defer putEntry(entry)
	entry.Time = time.Now()
	entry.Level = FatalLevel
	entry.Message = "fatal message"
	entry.Pid = 12345
	entry.Gid = 67890
	entry.File = "test.go"
	entry.CallerLine = 42

	result := formatter.Format(entry)
	output := string(result)

	if !strings.Contains(output, "fatal message") {
		t.Errorf("Output should contain fatal message, got: %s", output)
	}

	if !strings.Contains(output, "fatal") {
		t.Errorf("Output should contain fatal level, got: %s", output)
	}
}

func TestLogger_Fatalf(t *testing.T) {
	// 类似 Fatal，我们测试格式化部分而不是 os.Exit
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             false,
	}

	entry := getEntry()
	defer putEntry(entry)
	entry.Time = time.Now()
	entry.Level = FatalLevel
	entry.Message = "fatal formatted 456"
	entry.Pid = 12345
	entry.Gid = 67890
	entry.File = "test.go"
	entry.CallerLine = 42

	result := formatter.Format(entry)
	output := string(result)

	if !strings.Contains(output, "fatal formatted 456") {
		t.Errorf("Output should contain formatted fatal message, got: %s", output)
	}
}

func TestGlobalLoggingMethods(t *testing.T) {
	var buf bytes.Buffer
	originalOut := std.out
	std.SetOutput(&buf)
	std.SetLevel(TraceLevel)
	defer func() {
		std.out = originalOut
	}()

	// 测试全局日志方法
	Trace("global trace")
	Log(InfoLevel, "global log")
	Info("global info")
	Warn("global warn")
	Error("global error")

	output := buf.String()

	expectedMessages := []string{
		"global trace",
		"global log",
		"global info",
		"global warn",
		"global error",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Global output should contain %q, got: %s", msg, output)
		}
	}
}

func TestGlobalFormattedLoggingMethods(t *testing.T) {
	var buf bytes.Buffer
	originalOut := std.out
	std.SetOutput(&buf)
	std.SetLevel(TraceLevel)
	defer func() {
		std.out = originalOut
	}()

	// 测试全局格式化日志方法
	Logf(InfoLevel, "global %s %d", "logf", 1)
	Infof("global %s %d", "infof", 2)
	Warnf("global %s %d", "warnf", 3)
	Errorf("global %s %d", "errorf", 4)

	output := buf.String()

	expectedMessages := []string{
		"global logf 1",
		"global infof 2",
		"global warnf 3",
		"global errorf 4",
	}

	for _, msg := range expectedMessages {
		if !strings.Contains(output, msg) {
			t.Errorf("Global formatted output should contain %q, got: %s", msg, output)
		}
	}
}

func TestGlobalPanicMethods(t *testing.T) {
	var buf bytes.Buffer
	originalOut := std.out
	std.SetOutput(&buf)
	std.SetLevel(TraceLevel)
	defer func() {
		std.out = originalOut
	}()

	t.Run("GlobalPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Global Panic should have panicked")
			}
		}()
		Panic("global panic message")
	})

	if !strings.Contains(buf.String(), "global panic message") {
		t.Errorf("Output should contain global panic message, got: %s", buf.String())
	}
}

func TestGlobalPanicf(t *testing.T) {
	var buf bytes.Buffer
	originalOut := std.out
	std.SetOutput(&buf)
	std.SetLevel(TraceLevel)
	defer func() {
		std.out = originalOut
	}()

	t.Run("GlobalPanicf", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Global Panicf should have panicked")
			}
		}()
		Panicf("global panic %s %d", "formatted", 789)
	})

	if !strings.Contains(buf.String(), "global panic formatted 789") {
		t.Errorf("Output should contain global formatted panic message, got: %s", buf.String())
	}
}

func TestGlobalFatalMethods(t *testing.T) {
	// 由于 Fatal 和 Fatalf 会调用 os.Exit，我们测试格式化行为
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             false,
	}

	entry := getEntry()
	defer putEntry(entry)
	entry.Time = time.Now()
	entry.Level = FatalLevel
	entry.Message = "simulated fatal"
	entry.Pid = 12345
	entry.Gid = 67890
	entry.File = "test.go"
	entry.CallerLine = 42

	result := formatter.Format(entry)
	output := string(result)

	if !strings.Contains(output, "simulated fatal") {
		t.Errorf("Output should contain simulated fatal message, got: %s", output)
	}

	if !strings.Contains(output, "fatal") {
		t.Errorf("Output should contain fatal level, got: %s", output)
	}
}

func TestStartMsg(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)

	// 测试 StartMsg 方法
	logger.StartMsg()

	output := buf.String()
	if !strings.Contains(output, "start new log") {
		t.Errorf("Output should contain start message, got: %s", output)
	}
}

func TestLogger_LevelFiltering(t *testing.T) {
	// 测试各种级别的过滤功能，包括格式化方法
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(ErrorLevel) // 只记录ERROR及以上级别

	// 测试被过滤的级别 - 这些方法会return early，不会写入任何内容
	logger.Tracef("trace %s", "message")
	logger.Debugf("debug %s", "message")
	logger.Infof("info %s", "message")
	logger.Warnf("warn %s", "message")

	// 这些级别应该被记录
	logger.Errorf("error %s", "message")
	// 注意：我们不能测试Fatalf，因为它会调用os.Exit

	output := buf.String()

	// 验证低级别的消息没有被记录
	if strings.Contains(output, "trace message") {
		t.Error("Trace message should be filtered out")
	}
	if strings.Contains(output, "debug message") {
		t.Error("Debug message should be filtered out")
	}
	if strings.Contains(output, "info message") {
		t.Error("Info message should be filtered out")
	}
	if strings.Contains(output, "warn message") {
		t.Error("Warn message should be filtered out")
	}

	// 验证高级别的消息被记录了
	if !strings.Contains(output, "error message") {
		t.Error("Error message should be recorded")
	}
}

func TestLogger_FatalMethod(t *testing.T) {
	// 测试Fatal方法（不会调用os.Exit的版本）
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             false,
	}

	entry := getEntry()
	defer putEntry(entry)
	entry.Time = time.Now()
	entry.Level = FatalLevel
	entry.Message = "fatal test"
	entry.Pid = 12345
	entry.Gid = 67890
	entry.File = "test.go"
	entry.CallerLine = 42

	result := formatter.Format(entry)
	output := string(result)

	if !strings.Contains(output, "fatal test") {
		t.Errorf("Output should contain fatal message, got: %s", output)
	}
}

func TestGlobalStartMsg(t *testing.T) {
	var buf bytes.Buffer
	originalOut := std.out
	std.SetOutput(&buf)
	std.SetLevel(TraceLevel)
	defer func() {
		std.out = originalOut
	}()

	// 测试全局 StartMsg
	StartMsg()

	output := buf.String()
	if !strings.Contains(output, "start new log") {
		t.Errorf("Global output should contain start message, got: %s", output)
	}
}

func TestLogger_Logf(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(TraceLevel)

	// 测试 Logf 方法
	logger.Logf(WarnLevel, "formatted %s message with %d parameters", "log", 2)

	output := buf.String()
	if !strings.Contains(output, "formatted log message with 2 parameters") {
		t.Errorf("Output should contain formatted message, got: %s", output)
	}

	if !strings.Contains(output, "warn") {
		t.Errorf("Output should contain warn level, got: %s", output)
	}
}

func TestLogger_ParsingAndEscaping_NotFormatFull(t *testing.T) {
	logger := newLogger()

	// 设置一个不实现 FormatFull 的格式化器
	logger.Format = &SimpleFormat{}

	// 这应该会 panic，因为格式化器不是 FormatFull 类型
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when formatter is not FormatFull")
		}
	}()

	logger.ParsingAndEscaping(true)
}

func TestLogger_Caller_NotFormatFull(t *testing.T) {
	logger := newLogger()

	// 设置一个不实现 FormatFull 的格式化器
	logger.Format = &SimpleFormat{}

	// 这应该会 panic
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when formatter is not FormatFull")
		}
	}()

	logger.Caller(true)
}

// SimpleFormat 是一个简单的格式化器，不实现 FormatFull 接口
type SimpleFormat struct{}

func (sf *SimpleFormat) Format(entry *Entry) []byte {
	return []byte(entry.Message + "\n")
}

func TestLogger_Write_Method(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(InfoLevel)

	// 测试 logger 的内部 write 方法的不同分支
	entry := &Entry{
		Level:   InfoLevel,
		Message: "test write method",
		Pid:     12345,
		Gid:     67890,
	}

	// 这将测试 write 方法的实现
	formatted := logger.Format.Format(entry)
	logger.write(entry.Level, formatted)

	output := buf.String()
	if !strings.Contains(output, "test write method") {
		t.Errorf("Output should contain message, got: %s", output)
	}
}

func TestLogger_LevelFilteringFormatted_Coverage(t *testing.T) {
	// 测试格式化方法的级别过滤，覆盖更多未覆盖的条件分支
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(ErrorLevel) // 只记录ERROR及以上级别

	// 测试被过滤的Printf方法（logger.go:320-322）
	logger.Printf("printf %s", "message")

	// 测试被过滤的Warnf方法（logger.go:360-362）
	logger.Warnf("warn %s", "message")

	// 测试被过滤的Warningf方法（logger.go:370-372）
	logger.Warningf("warning %s", "message")

	output := buf.String()

	// 验证被过滤的消息
	if strings.Contains(output, "printf message") {
		t.Error("Printf message should be filtered out")
	}
	if strings.Contains(output, "warn message") {
		t.Error("Warnf message should be filtered out")
	}
	if strings.Contains(output, "warning message") {
		t.Error("Warningf message should be filtered out")
	}

	// 注意：Fatal和Panic的级别过滤无法安全测试，因为它们会调用os.Exit或panic
}

func TestLogger_SetOutputNil_Coverage(t *testing.T) {
	// 测试SetOutput中nil writer的处理（logger.go:161-162）
	logger := newLogger()

	// 测试传入nil writer
	result := logger.SetOutput(nil)

	if result != logger {
		t.Error("SetOutput should return the same logger instance")
	}

	// out应该被设置为nil
	if logger.out != nil {
		t.Error("Output should be nil when no writers provided")
	}
}

func TestLogger_SetOutputEmpty_Coverage(t *testing.T) {
	// 测试SetOutput中空writer列表的处理（logger.go:167-169）
	logger := newLogger()

	// 测试传入空列表
	result := logger.SetOutput()

	if result != logger {
		t.Error("SetOutput should return the same logger instance")
	}

	// out应该被设置为nil
	if logger.out != nil {
		t.Error("Output should be nil when no writers provided")
	}
}

func TestLogger_LogfLevelFiltering_Coverage(t *testing.T) {
	// 测试Logf方法的级别过滤（logger.go:193-195）
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(ErrorLevel) // 只记录ERROR及以上级别

	// 测试被过滤的Logf调用
	logger.Logf(DebugLevel, "debug %s", "message")
	logger.Logf(InfoLevel, "info %s", "message")
	logger.Logf(WarnLevel, "warn %s", "message")

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

func TestEntryPool_Coverage(t *testing.T) {
	// 测试 entry pool 的使用
	entry := getEntry()
	if entry == nil {
		t.Fatal("getEntry returned nil")
	}

	// 设置一些值
	entry.Level = InfoLevel
	entry.Message = "pool test"

	// 重置并放回池中
	entry.Reset()
	putEntry(entry)

	// 再次获取，应该是重置后的状态
	entry2 := getEntry()
	if entry2.Message != "" {
		t.Error("Entry from pool should be reset")
	}

	putEntry(entry2)
}

// Note: TestLogger_PanicfLevelFiltering_Coverage removed because PanicLevel=0 is the highest priority
// and there's no level higher than PanicLevel that can filter it out safely
