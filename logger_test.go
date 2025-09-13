package log

import (
	"bytes"
	"strings"
	"testing"
)

// MockWriter 用于测试的模拟 Writer
type MockWriter struct {
	data        bytes.Buffer
	syncCalls   int
	syncError   error
	shouldError bool
}

func (m *MockWriter) Write(p []byte) (n int, err error) {
	if m.shouldError {
		return 0, &MockWriteError{"mock write error"}
	}
	return m.data.Write(p)
}

func (m *MockWriter) Sync() error {
	m.syncCalls++
	return m.syncError
}

func (m *MockWriter) String() string {
	return m.data.String()
}

// MockWriteError 是模拟的写入错误
type MockWriteError struct {
	msg string
}

func (e *MockWriteError) Error() string {
	return e.msg
}

func TestWriteSyncerWrapper_Write(t *testing.T) {
	mock := &MockWriter{}
	wrapper := &WriteSyncerWrapper{writer: mock}
	
	testData := []byte("test data")
	n, err := wrapper.Write(testData)
	
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	
	if n != len(testData) {
		t.Errorf("Expected to write %d bytes, wrote %d", len(testData), n)
	}
	
	if mock.String() != string(testData) {
		t.Errorf("Expected data %q, got %q", string(testData), mock.String())
	}
}

func TestWriteSyncerWrapper_Sync_WithSyncMethod(t *testing.T) {
	mock := &MockWriter{}
	wrapper := &WriteSyncerWrapper{writer: mock}
	
	err := wrapper.Sync()
	if err != nil {
		t.Errorf("Sync failed: %v", err)
	}
	
	if mock.syncCalls != 1 {
		t.Errorf("Expected 1 sync call, got %d", mock.syncCalls)
	}
}

func TestWriteSyncerWrapper_Sync_WithoutSyncMethod(t *testing.T) {
	// 使用标准的 bytes.Buffer，它没有 Sync 方法
	var buf bytes.Buffer
	wrapper := &WriteSyncerWrapper{writer: &buf}
	
	err := wrapper.Sync()
	if err != nil {
		t.Errorf("Sync should not fail for writers without Sync method: %v", err)
	}
}

func TestWrapWriter_AlreadyWriteSyncer(t *testing.T) {
	mock := &MockWriter{}
	wrapper := &WriteSyncerWrapper{writer: mock}
	
	// wrapWriter 应该识别已经是 WriteSyncer 的对象
	result := wrapWriter(wrapper)
	
	if result != wrapper {
		t.Error("wrapWriter should return the same instance if already a WriteSyncer")
	}
}

func TestWrapWriter_NeedsWrapping(t *testing.T) {
	var buf bytes.Buffer
	result := wrapWriter(&buf)
	
	wrapper, ok := result.(*WriteSyncerWrapper)
	if !ok {
		t.Error("wrapWriter should return a WriteSyncerWrapper")
	}
	
	if wrapper.writer != &buf {
		t.Error("Wrapper should contain the original writer")
	}
}

func TestNewLogger(t *testing.T) {
	logger := newLogger()
	
	if logger == nil {
		t.Fatal("newLogger returned nil")
	}
	
	if logger.level != DebugLevel {
		t.Errorf("Expected level %v, got %v", DebugLevel, logger.level)
	}
	
	if logger.callerDepth != 4 {
		t.Errorf("Expected callerDepth 4, got %d", logger.callerDepth)
	}
	
	if logger.Format == nil {
		t.Error("Format should not be nil")
	}
	
	if logger.out == nil {
		t.Error("Output should not be nil")
	}
}

func TestLogger_SetCallerDepth(t *testing.T) {
	logger := newLogger()
	
	result := logger.SetCallerDepth(10)
	
	if result != logger {
		t.Error("SetCallerDepth should return the same logger instance")
	}
	
	if logger.callerDepth != 10 {
		t.Errorf("Expected callerDepth 10, got %d", logger.callerDepth)
	}
}

func TestLogger_SetPrefixMsg(t *testing.T) {
	logger := newLogger()
	
	prefix := "PREFIX: "
	result := logger.SetPrefixMsg(prefix)
	
	if result != logger {
		t.Error("SetPrefixMsg should return the same logger instance")
	}
	
	if string(logger.PrefixMsg) != prefix {
		t.Errorf("Expected prefix %q, got %q", prefix, string(logger.PrefixMsg))
	}
}

func TestLogger_AppendPrefixMsg(t *testing.T) {
	logger := newLogger()
	
	// 设置初始前缀
	logger.SetPrefixMsg("INITIAL: ")
	
	// 追加前缀
	additional := "ADDED: "
	result := logger.AppendPrefixMsg(additional)
	
	if result != logger {
		t.Error("AppendPrefixMsg should return the same logger instance")
	}
	
	expected := "INITIAL: ADDED: "
	if string(logger.PrefixMsg) != expected {
		t.Errorf("Expected prefix %q, got %q", expected, string(logger.PrefixMsg))
	}
}

func TestLogger_SetSuffixMsg(t *testing.T) {
	logger := newLogger()
	
	suffix := " :SUFFIX"
	result := logger.SetSuffixMsg(suffix)
	
	if result != logger {
		t.Error("SetSuffixMsg should return the same logger instance")
	}
	
	if string(logger.SuffixMsg) != suffix {
		t.Errorf("Expected suffix %q, got %q", suffix, string(logger.SuffixMsg))
	}
}

func TestLogger_AppendSuffixMsg(t *testing.T) {
	logger := newLogger()
	
	// 设置初始后缀
	logger.SetSuffixMsg(" :INITIAL")
	
	// 追加后缀
	additional := " :ADDED"
	result := logger.AppendSuffixMsg(additional)
	
	if result != logger {
		t.Error("AppendSuffixMsg should return the same logger instance")
	}
	
	expected := " :INITIAL :ADDED"
	if string(logger.SuffixMsg) != expected {
		t.Errorf("Expected suffix %q, got %q", expected, string(logger.SuffixMsg))
	}
}

func TestLogger_SetLevel(t *testing.T) {
	logger := newLogger()
	
	result := logger.SetLevel(ErrorLevel)
	
	if result != logger {
		t.Error("SetLevel should return the same logger instance")
	}
	
	if logger.level != ErrorLevel {
		t.Errorf("Expected level %v, got %v", ErrorLevel, logger.level)
	}
}

func TestLogger_Level(t *testing.T) {
	logger := newLogger()
	logger.SetLevel(WarnLevel)
	
	if logger.Level() != WarnLevel {
		t.Errorf("Expected level %v, got %v", WarnLevel, logger.Level())
	}
}

func TestLogger_SetOutput(t *testing.T) {
	logger := newLogger()
	var buf1, buf2 bytes.Buffer
	
	result := logger.SetOutput(&buf1, &buf2)
	
	if result != logger {
		t.Error("SetOutput should return the same logger instance")
	}
	
	// 测试写入是否工作
	logger.Info("test message")
	
	// 两个 buffer 都应该收到数据
	if buf1.Len() == 0 {
		t.Error("First buffer should contain data")
	}
	
	if buf2.Len() == 0 {
		t.Error("Second buffer should contain data")
	}
}

func TestLogger_Sync(t *testing.T) {
	logger := newLogger()
	mock := &MockWriter{}
	logger.SetOutput(mock)
	
	logger.Sync()
	
	// 应该调用底层 writer 的 sync
	if mock.syncCalls != 1 {
		t.Errorf("Expected 1 sync call, got %d", mock.syncCalls)
	}
}

func TestLogger_Clone_SimpleFormat(t *testing.T) {
	// 测试Clone方法中Format不是FormatFull类型的情况
	logger := newLogger()
	
	// 设置一个不实现FormatFull的简单格式化器
	simpleFormat := &SimpleFormat{}
	logger.Format = simpleFormat
	
	cloned := logger.Clone()
	
	if cloned == logger {
		t.Error("Clone should return a different logger instance")
	}
	
	if cloned.Format != simpleFormat {
		t.Error("Clone should share the same format instance for non-FormatFull types")
	}
}

func TestLogger_Clone(t *testing.T) {
	original := newLogger()
	original.SetLevel(ErrorLevel)
	original.SetPrefixMsg("PREFIX: ")
	original.SetSuffixMsg(" :SUFFIX")
	original.SetCallerDepth(10)
	
	clone := original.Clone()
	
	if clone == original {
		t.Error("Clone should return a different instance")
	}
	
	if clone.Level() != original.Level() {
		t.Error("Clone should have the same level")
	}
	
	if string(clone.PrefixMsg) != string(original.PrefixMsg) {
		t.Error("Clone should have the same prefix")
	}
	
	if string(clone.SuffixMsg) != string(original.SuffixMsg) {
		t.Error("Clone should have the same suffix")
	}
	
	if clone.callerDepth != original.callerDepth {
		t.Error("Clone should have the same caller depth")
	}
	
	// 修改克隆不应该影响原始对象
	clone.SetLevel(TraceLevel)
	if original.Level() == TraceLevel {
		t.Error("Modifying clone should not affect original")
	}
}

// 测试日志级别检查
func TestLogger_LevelCheck(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(WarnLevel)
	
	// Debug 和 Info 消息应该被过滤
	logger.Debug("debug message")
	logger.Info("info message")
	
	// 应该没有输出
	if buf.Len() > 0 {
		t.Error("Debug and Info messages should be filtered when level is Warn")
	}
	
	// Warn 消息应该通过
	logger.Warn("warn message")
	
	if buf.Len() == 0 {
		t.Error("Warn message should not be filtered")
	}
	
	if !strings.Contains(buf.String(), "warn message") {
		t.Error("Output should contain warn message")
	}
}

// TestLogger_ErrorfLevelDisabled tests Errorf when Error level is disabled
func TestLogger_ErrorfLevelDisabled(t *testing.T) {
	var buf bytes.Buffer
	logger := newLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(FatalLevel) // Only fatal messages, error is disabled
	
	// Errorf should be filtered out
	logger.Errorf("error message: %s", "test")
	
	// Should not have any output
	if buf.Len() > 0 {
		t.Error("Errorf should be filtered when Error level is disabled")
	}
}

// TestNewLogger_ReleaseLogPath tests the newLogger function's ReleaseLogPath branch
func TestNewLogger_ReleaseLogPath(t *testing.T) {
	// Test newLogger with empty ReleaseLogPath (should use os.Stdout)
	originalPath := ReleaseLogPath
	ReleaseLogPath = ""
	defer func() {
		ReleaseLogPath = originalPath
	}()
	
	logger := newLogger()
	
	if logger == nil {
		t.Fatal("newLogger returned nil")
	}
	
	if logger.level != DebugLevel {
		t.Errorf("Expected DebugLevel, got %v", logger.level)
	}
	
	if logger.callerDepth != 4 {
		t.Errorf("Expected callerDepth 4, got %d", logger.callerDepth)
	}
	
	// Test newLogger with non-empty ReleaseLogPath (should use hourly rotator)
	tmpDir := t.TempDir()
	ReleaseLogPath = tmpDir + "/test.log"
	
	logger2 := newLogger()
	
	if logger2 == nil {
		t.Fatal("newLogger returned nil with ReleaseLogPath")
	}
	
	// Verify it creates a logger (we can't easily test the exact output type without exposing internals)
	if logger2.level != DebugLevel {
		t.Errorf("Expected DebugLevel with ReleaseLogPath, got %v", logger2.level)
	}
}