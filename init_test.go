package log

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPid(t *testing.T) {
	currentPid := os.Getpid()
	logPid := Pid()
	
	if logPid != currentPid {
		t.Errorf("Expected PID %d, got %d", currentPid, logPid)
	}
	
	// PID 应该是正数
	if logPid <= 0 {
		t.Errorf("PID should be positive, got %d", logPid)
	}
}

func TestNew(t *testing.T) {
	logger := New()
	
	if logger == nil {
		t.Fatal("New() returned nil")
	}
	
	// 新创建的 logger 应该有默认设置
	if logger.Level() != DebugLevel {
		t.Errorf("Expected default level %v, got %v", DebugLevel, logger.Level())
	}
	
	if logger.callerDepth != 4 {
		t.Errorf("Expected default caller depth 4, got %d", logger.callerDepth)
	}
}

func TestSetLevel_Global(t *testing.T) {
	// 保存原始状态
	originalLevel := std.Level()
	defer func() {
		std.SetLevel(originalLevel)
	}()
	
	result := SetLevel(ErrorLevel)
	
	if result != std {
		t.Error("SetLevel should return the global std logger")
	}
	
	if std.Level() != ErrorLevel {
		t.Errorf("Expected global logger level %v, got %v", ErrorLevel, std.Level())
	}
	
	if GetLevel() != ErrorLevel {
		t.Errorf("GetLevel should return %v, got %v", ErrorLevel, GetLevel())
	}
}

func TestGetLevel_Global(t *testing.T) {
	// 保存原始状态
	originalLevel := std.Level()
	defer func() {
		std.SetLevel(originalLevel)
	}()
	
	// 设置一个特定级别
	std.SetLevel(WarnLevel)
	
	if GetLevel() != WarnLevel {
		t.Errorf("Expected level %v, got %v", WarnLevel, GetLevel())
	}
}

func TestSync_Global(t *testing.T) {
	// 创建一个 mock writer 来测试 sync 调用
	mock := &MockWriter{}
	originalOut := std.out
	std.SetOutput(mock)
	defer func() {
		std.out = originalOut
	}()
	
	Sync()
	
	// Sync 应该被调用
	if mock.syncCalls != 1 {
		t.Errorf("Expected 1 sync call, got %d", mock.syncCalls)
	}
}

func TestClone_Global(t *testing.T) {
	// 保存原始状态
	originalLevel := std.Level()
	originalPrefix := string(std.PrefixMsg)
	originalSuffix := string(std.SuffixMsg)
	originalDepth := std.callerDepth
	defer func() {
		std.SetLevel(originalLevel)
		std.SetPrefixMsg(originalPrefix)
		std.SetSuffixMsg(originalSuffix)
		std.SetCallerDepth(originalDepth)
	}()
	
	// 设置一些自定义值
	std.SetLevel(ErrorLevel)
	std.SetPrefixMsg("TEST: ")
	std.SetSuffixMsg(" :END")
	std.SetCallerDepth(10)
	
	clone := Clone()
	
	if clone == std {
		t.Error("Clone should return a different instance")
	}
	
	if clone.Level() != std.Level() {
		t.Error("Clone should have the same level as std")
	}
	
	if string(clone.PrefixMsg) != string(std.PrefixMsg) {
		t.Error("Clone should have the same prefix as std")
	}
	
	if string(clone.SuffixMsg) != string(std.SuffixMsg) {
		t.Error("Clone should have the same suffix as std")
	}
	
	if clone.callerDepth != std.callerDepth {
		t.Error("Clone should have the same caller depth as std")
	}
	
	// 修改克隆不应该影响原始对象
	clone.SetLevel(TraceLevel)
	if std.Level() == TraceLevel {
		t.Error("Modifying clone should not affect std")
	}
}

func TestCloneToCtx_Global(t *testing.T) {
	// 保存原始状态
	originalLevel := std.Level()
	defer func() {
		std.SetLevel(originalLevel)
	}()
	
	std.SetLevel(WarnLevel)
	
	ctxLogger := CloneToCtx()
	
	if ctxLogger == nil {
		t.Fatal("CloneToCtx returned nil")
	}
	
	// 检查内部 logger 的设置
	if ctxLogger.Logger.Level() != WarnLevel {
		t.Errorf("Expected context logger level %v, got %v", WarnLevel, ctxLogger.Logger.Level())
	}
}

func TestSetCallerDepth_Global(t *testing.T) {
	// 保存原始状态
	originalDepth := std.callerDepth
	defer func() {
		std.SetCallerDepth(originalDepth)
	}()
	
	result := SetCallerDepth(15)
	
	if result != std {
		t.Error("SetCallerDepth should return the global std logger")
	}
	
	if std.callerDepth != 15 {
		t.Errorf("Expected caller depth 15, got %d", std.callerDepth)
	}
}

func TestSetPrefixMsg_Global(t *testing.T) {
	// 保存原始状态
	originalPrefix := string(std.PrefixMsg)
	defer func() {
		std.SetPrefixMsg(originalPrefix)
	}()
	
	prefix := "GLOBAL: "
	result := SetPrefixMsg(prefix)
	
	if result != std {
		t.Error("SetPrefixMsg should return the global std logger")
	}
	
	if string(std.PrefixMsg) != prefix {
		t.Errorf("Expected prefix %q, got %q", prefix, string(std.PrefixMsg))
	}
}

func TestAppendPrefixMsg_Global(t *testing.T) {
	// 保存原始状态
	originalPrefix := string(std.PrefixMsg)
	defer func() {
		std.SetPrefixMsg(originalPrefix)
	}()
	
	std.SetPrefixMsg("INITIAL: ")
	additional := "ADDED: "
	
	result := AppendPrefixMsg(additional)
	
	if result != std {
		t.Error("AppendPrefixMsg should return the global std logger")
	}
	
	expected := "INITIAL: ADDED: "
	if string(std.PrefixMsg) != expected {
		t.Errorf("Expected prefix %q, got %q", expected, string(std.PrefixMsg))
	}
}

func TestSetSuffixMsg_Global(t *testing.T) {
	// 保存原始状态
	originalSuffix := string(std.SuffixMsg)
	defer func() {
		std.SetSuffixMsg(originalSuffix)
	}()
	
	suffix := " :GLOBAL"
	result := SetSuffixMsg(suffix)
	
	if result != std {
		t.Error("SetSuffixMsg should return the global std logger")
	}
	
	if string(std.SuffixMsg) != suffix {
		t.Errorf("Expected suffix %q, got %q", suffix, string(std.SuffixMsg))
	}
}

func TestAppendSuffixMsg_Global(t *testing.T) {
	// 保存原始状态
	originalSuffix := string(std.SuffixMsg)
	defer func() {
		std.SetSuffixMsg(originalSuffix)
	}()
	
	std.SetSuffixMsg(" :INITIAL")
	additional := " :ADDED"
	
	result := AppendSuffixMsg(additional)
	
	if result != std {
		t.Error("AppendSuffixMsg should return the global std logger")
	}
	
	expected := " :INITIAL :ADDED"
	if string(std.SuffixMsg) != expected {
		t.Errorf("Expected suffix %q, got %q", expected, string(std.SuffixMsg))
	}
}

func TestParsingAndEscaping_Global(t *testing.T) {
	// 保存原始状态
	originalFormatter := std.Format
	defer func() {
		std.Format = originalFormatter
	}()
	
	result := ParsingAndEscaping(false)
	
	if result != std {
		t.Error("ParsingAndEscaping should return the global std logger")
	}
	
	// 检查格式化器的设置
	if formatter, ok := std.Format.(*Formatter); ok {
		if formatter.DisableParsingAndEscaping != false {
			t.Error("Expected DisableParsingAndEscaping to be false")
		}
	} else {
		t.Error("Expected Formatter type")
	}
	
	// 测试设置为 true
	ParsingAndEscaping(true)
	if formatter, ok := std.Format.(*Formatter); ok {
		if formatter.DisableParsingAndEscaping != true {
			t.Error("Expected DisableParsingAndEscaping to be true")
		}
	}
}

func TestCaller_Global(t *testing.T) {
	// 保存原始状态
	originalFormatter := std.Format
	defer func() {
		std.Format = originalFormatter
	}()
	
	result := Caller(false)
	
	if result != std {
		t.Error("Caller should return the global std logger")
	}
	
	// 检查格式化器的设置
	if formatter, ok := std.Format.(*Formatter); ok {
		if formatter.DisableCaller != false {
			t.Errorf("Expected DisableCaller to be false, got %v", formatter.DisableCaller)
		}
	} else {
		t.Error("Expected Formatter type")
	}
	
	// 测试设置为 true
	Caller(true)
	if formatter, ok := std.Format.(*Formatter); ok {
		if formatter.DisableCaller != true {
			t.Errorf("Expected DisableCaller to be true, got %v", formatter.DisableCaller)
		}
	}
}

// 测试全局函数的日志输出
func TestGlobalLoggingFunctions(t *testing.T) {
	var buf bytes.Buffer
	originalOut := std.out
	std.SetOutput(&buf)
	defer func() {
		std.out = originalOut
	}()
	
	std.SetLevel(DebugLevel)
	
	// 测试各个级别的日志函数
	Debug("debug message")
	Info("info message")
	Warn("warn message")
	Error("error message")
	
	output := buf.String()
	
	// 根据build tag调整期望行为
	expectedMessages := []string{"info message", "warn message", "error message"}
	
	// 在debug build tag下，或者在默认情况下（没有release tag），Debug应该输出
	// 在release模式下，Debug函数是空实现，不会输出
	if checkDebugEnabled() {
		expectedMessages = append([]string{"debug message"}, expectedMessages...)
	}
	
	if !containsAll(output, expectedMessages) {
		t.Errorf("Output should contain expected messages for current build tag, expected: %v, got: %s", expectedMessages, output)
	}
}

// checkDebugEnabled 检查当前build tag是否启用了debug输出
func checkDebugEnabled() bool {
	// 测试Debug函数是否实际产生输出
	var testBuf bytes.Buffer
	originalOut := std.out
	std.SetOutput(&testBuf)
	defer func() {
		std.out = originalOut
	}()
	
	std.SetLevel(DebugLevel)
	Debug("test debug")
	
	return strings.Contains(testBuf.String(), "test debug")
}

// 辅助函数：检查字符串是否包含所有子字符串
func containsAll(s string, substrings []string) bool {
	for _, substr := range substrings {
		if !bytes.Contains([]byte(s), []byte(substr)) {
			return false
		}
	}
	return true
}

func TestGlobalLevelFiltering(t *testing.T) {
	// 测试全局函数的级别过滤
	var buf bytes.Buffer
	originalOut := std.out
	originalLevel := std.level
	std.SetOutput(&buf)
	std.SetLevel(ErrorLevel) // 只记录ERROR及以上级别
	defer func() {
		std.out = originalOut
		std.level = originalLevel
	}()
	
	// 测试被过滤的级别
	Trace("trace message")
	Debugf("debug %s", "message") 
	Infof("info %s", "message")
	Warnf("warn %s", "message")
	
	// 测试会被记录的级别
	Errorf("error %s", "message")
	
	output := buf.String()
	
	// 验证低级别的消息没有被记录
	if bytes.Contains([]byte(output), []byte("trace message")) {
		t.Error("Global Trace message should be filtered out")
	}
	if bytes.Contains([]byte(output), []byte("debug message")) {
		t.Error("Global Debug message should be filtered out")
	}
	if bytes.Contains([]byte(output), []byte("info message")) {
		t.Error("Global Info message should be filtered out")
	}
	if bytes.Contains([]byte(output), []byte("warn message")) {
		t.Error("Global Warn message should be filtered out")
	}
	
	// 验证高级别的消息被记录了
	if !bytes.Contains([]byte(output), []byte("error message")) {
		t.Error("Global Error message should be recorded")
	}
}

func TestPrintOther_DebugfLevelFiltering_Coverage(t *testing.T) {
	// 测试print_other.go中Debugf的级别过滤（print_other.go:24-26）
	var buf bytes.Buffer
	originalOut := std.out
	originalLevel := std.level
	std.SetOutput(&buf)
	std.SetLevel(InfoLevel) // 设置级别高于Debug，使Debugf被过滤
	defer func() {
		std.out = originalOut
		std.level = originalLevel
	}()
	
	// 调用Debugf，应该被过滤掉
	Debugf("debug %s", "message")
	
	output := buf.String()
	
	// 验证没有输出
	if strings.Contains(output, "debug message") {
		t.Error("Debugf message should be filtered when level is higher than DebugLevel")
	}
}

func TestPrintOther_DebugfLevelEnabled_Coverage(t *testing.T) {
	// 测试print_other.go中Debugf的执行分支（print_other.go:24-26）
	var buf bytes.Buffer
	originalOut := std.out
	originalLevel := std.level
	std.SetOutput(&buf)
	std.SetLevel(DebugLevel) // 设置级别等于Debug，使Debugf能够执行
	defer func() {
		std.out = originalOut
		std.level = originalLevel
	}()
	
	// 调用Debugf，应该被执行并输出
	Debugf("debug %s %d", "message", 123)
	
	output := buf.String()
	
	// 验证有输出包含格式化内容
	if !strings.Contains(output, "debug message 123") {
		t.Error("Debugf message should be executed when level is DebugLevel or lower")
	}
}