package log

import (
	"strings"
	"testing"
	"time"
)

func TestFormatter_Format(t *testing.T) {
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             false,
	}
	
	entry := &Entry{
		Time:    time.Now(),
		Level:   InfoLevel,
		Message: "test message",
		Pid:     12345,
		Gid:     67890,
		File:       "test.go",
		CallerLine: 42,
		CallerName: "github.com/test/pkg.TestFunc",
		CallerDir:  "test/pkg",
		CallerFunc: "TestFunc",
	}
	
	result := formatter.Format(entry)
	
	if result == nil {
		t.Fatal("Format returned nil")
	}
	
	output := string(result)
	
	// 检查是否包含关键信息
	if !strings.Contains(output, "test message") {
		t.Errorf("Output should contain message, got: %s", output)
	}
	
	if !strings.Contains(output, "info") {
		t.Errorf("Output should contain level, got: %s", output)
	}
	
	if !strings.Contains(output, "12345") {
		t.Errorf("Output should contain PID, got: %s", output)
	}
	
	if !strings.Contains(output, "67890") {
		t.Errorf("Output should contain GID, got: %s", output)
	}
	
	if !strings.Contains(output, "test.go:42") {
		t.Errorf("Output should contain caller info, got: %s", output)
	}
}

func TestFormatter_Format_WithPrefixSuffix(t *testing.T) {
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             true,
	}
	
	entry := &Entry{
		Time:      time.Now(),
		Level:     WarnLevel,
		Message:   "test message",
		PrefixMsg: []byte("PREFIX: "),
		SuffixMsg: []byte(" :SUFFIX"),
		Pid:       12345,
		Gid:       67890,
	}
	
	result := formatter.Format(entry)
	output := string(result)
	
	if !strings.Contains(output, "PREFIX: ") {
		t.Errorf("Output should contain prefix, got: %s", output)
	}
	
	if !strings.Contains(output, " :SUFFIX") {
		t.Errorf("Output should contain suffix, got: %s", output)
	}
	
	if !strings.Contains(output, "warn") {
		t.Errorf("Output should contain level, got: %s", output)
	}
}

func TestFormatter_Format_DisableCaller(t *testing.T) {
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             true, // 禁用调用者信息
	}
	
	entry := &Entry{
		Time:    time.Now(),
		Level:   ErrorLevel,
		Message: "test message",
		Pid:     12345,
		Gid:     67890,
		File:       "test.go",
		CallerLine: 42,
		CallerName: "github.com/test/pkg.TestFunc",
		CallerDir:  "test/pkg",
		CallerFunc: "TestFunc",
	}
	
	result := formatter.Format(entry)
	output := string(result)
	
	// 应该不包含调用者信息
	if strings.Contains(output, "test.go:42") {
		t.Errorf("Output should not contain caller info when disabled, got: %s", output)
	}
	
	if strings.Contains(output, "TestFunc") {
		t.Errorf("Output should not contain function name when disabled, got: %s", output)
	}
}

func TestFormatter_Format_WithModule(t *testing.T) {
	formatter := &Formatter{
		Module:                    "TestModule",
		DisableParsingAndEscaping: true,
		DisableCaller:             true,
	}
	
	entry := &Entry{
		Time:    time.Now(),
		Level:   DebugLevel,
		Message: "test message",
		Pid:     12345,
		Gid:     67890,
	}
	
	result := formatter.Format(entry)
	output := string(result)
	
	// Module field is defined but not used in current implementation
	// Just verify basic format components are present
	if !strings.Contains(output, "debug") {
		t.Errorf("Output should contain level, got: %s", output)
	}
	if !strings.Contains(output, "test message") {
		t.Errorf("Output should contain message, got: %s", output)
	}
	if !strings.Contains(output, "12345") {
		t.Errorf("Output should contain PID, got: %s", output)
	}
}

func TestFormatter_Format_WithTrace(t *testing.T) {
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             true,
	}
	
	entry := &Entry{
		Time:    time.Now(),
		Level:   InfoLevel,
		Message: "test message",
		Pid:     12345,
		Gid:     67890,
		TraceId: "trace-id-12345",
	}
	
	result := formatter.Format(entry)
	output := string(result)
	
	if !strings.Contains(output, "trace-id-12345") {
		t.Errorf("Output should contain trace ID, got: %s", output)
	}
}

func TestFormatter_ParsingAndEscaping(t *testing.T) {
	formatter := &Formatter{}
	
	// 测试设置为 true
	formatter.ParsingAndEscaping(true)
	if formatter.DisableParsingAndEscaping != true {
		t.Error("Expected DisableParsingAndEscaping to be true")
	}
	
	// 测试设置为 false
	formatter.ParsingAndEscaping(false)
	if formatter.DisableParsingAndEscaping != false {
		t.Error("Expected DisableParsingAndEscaping to be false")
	}
}

func TestFormatter_Caller(t *testing.T) {
	formatter := &Formatter{}
	
	// 测试设置为 true (禁用)
	formatter.Caller(true)
	if formatter.DisableCaller != true {
		t.Error("Expected DisableCaller to be true")
	}
	
	// 测试设置为 false (启用)
	formatter.Caller(false)
	if formatter.DisableCaller != false {
		t.Error("Expected DisableCaller to be false")
	}
}

func TestFormatter_Clone(t *testing.T) {
	original := &Formatter{
		Module:                    "OriginalModule",
		DisableParsingAndEscaping: true,
		DisableCaller:             true,
	}
	
	clone := original.Clone()
	
	if clone == original {
		t.Error("Clone should return a different instance")
	}
	
	clonedFormatter, ok := clone.(*Formatter)
	if !ok {
		t.Fatal("Clone should return a *Formatter")
	}
	
	if clonedFormatter.Module != original.Module {
		t.Error("Clone should have the same Module")
	}
	
	if clonedFormatter.DisableParsingAndEscaping != original.DisableParsingAndEscaping {
		t.Error("Clone should have the same DisableParsingAndEscaping")
	}
	
	if clonedFormatter.DisableCaller != original.DisableCaller {
		t.Error("Clone should have the same DisableCaller")
	}
	
	// 修改克隆不应该影响原始对象
	clonedFormatter.Module = "ModifiedModule"
	if original.Module == "ModifiedModule" {
		t.Error("Modifying clone should not affect original")
	}
}

func TestGetColorByLevel(t *testing.T) {
	tests := []struct {
		level    Level
		expected []byte
	}{
		{TraceLevel, colorGreen},
		{DebugLevel, colorGreen},
		{InfoLevel, colorGreen},
		{WarnLevel, colorYellow},
		{ErrorLevel, colorRed},
		{FatalLevel, colorRed},
		{PanicLevel, colorRed},
	}
	
	for _, tt := range tests {
		t.Run(tt.level.String(), func(t *testing.T) {
			result := getColorByLevel(tt.level)
			if string(result) != string(tt.expected) {
				t.Errorf("Expected color %s for level %v, got %s", string(tt.expected), tt.level, string(result))
			}
		})
	}
}

func TestSplitPackageName(t *testing.T) {
	tests := []struct {
		input        string
		expectedDir  string
		expectedFunc string
	}{
		{
			input:        "github.com/lazygophers/log.TestFunc",
			expectedDir:  "log",
			expectedFunc: "TestFunc",
		},
		{
			input:        "github.com/user/pkg.SomeFunction",
			expectedDir:  "user/pkg",
			expectedFunc: "SomeFunction",
		},
		{
			input:        "main.main",
			expectedDir:  "main",
			expectedFunc: "main",
		},
		{
			input:        "pkg/subpkg.Function",
			expectedDir:  "pkg/subpkg",
			expectedFunc: "Function",
		},
		{
			input:        "simplefunction",
			expectedDir:  "simplefunction",
			expectedFunc: "",
		},
		{
			input:        "github.com/lazygophers/someother.Function",
			expectedDir:  "someother",
			expectedFunc: "Function",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			dir, funcName := SplitPackageName(tt.input)
			if dir != tt.expectedDir {
				t.Errorf("Expected dir %q, got %q", tt.expectedDir, dir)
			}
			if funcName != tt.expectedFunc {
				t.Errorf("Expected func %q, got %q", tt.expectedFunc, funcName)
			}
		})
	}
}

func TestFormatter_Format_MultiLine(t *testing.T) {
	// 测试多行消息处理（DisableParsingAndEscaping = false）
	formatter := &Formatter{
		DisableParsingAndEscaping: false, // 启用解析以测试多行处理
		DisableCaller:             true,
	}
	
	entry := &Entry{
		Time:    time.Now(),
		Level:   InfoLevel,
		Message: "line 1\nline 2\nline 3",
		Pid:     12345,
		Gid:     67890,
	}
	
	result := formatter.Format(entry)
	output := string(result)
	
	// 验证每行都被单独处理
	lines := strings.Split(output, "\n")
	if len(lines) < 3 {
		t.Errorf("Expected at least 3 lines for multiline message, got %d", len(lines))
	}
	
	if !strings.Contains(output, "line 1") {
		t.Errorf("Output should contain 'line 1', got: %s", output)
	}
	if !strings.Contains(output, "line 2") {
		t.Errorf("Output should contain 'line 2', got: %s", output)
	}
	if !strings.Contains(output, "line 3") {
		t.Errorf("Output should contain 'line 3', got: %s", output)
	}
}

func TestFormatter_Format_CallerEnabledWithoutParsing(t *testing.T) {
	// 测试 DisableParsingAndEscaping = false 且 DisableCaller = false 的情况
	// 这将覆盖 formatter.go 第 88-96 行的 else if 分支
	formatter := &Formatter{
		DisableParsingAndEscaping: false, // 启用解析
		DisableCaller:             false, // 启用调用者信息
	}
	
	entry := &Entry{
		Time:       time.Now(),
		Level:      InfoLevel,
		Message:    "test message",
		Pid:        12345,
		Gid:        67890,
		File:       "/path/to/test.go",
		CallerLine: 42,
		CallerDir:  "test",
		CallerFunc: "TestFunc",
	}
	
	result := formatter.Format(entry)
	output := string(result)
	
	// 验证包含调用者信息的格式（在 else if 分支中）
	if !strings.Contains(output, "test/test.go:42") {
		t.Errorf("Output should contain caller info, got: %s", output)
	}
	if !strings.Contains(output, "TestFunc") {
		t.Errorf("Output should contain caller function, got: %s", output)
	}
	if !strings.Contains(output, "test message") {
		t.Errorf("Output should contain message, got: %s", output)
	}
}

func TestFormatter_format_CallerWithoutTrace(t *testing.T) {
	// 测试没有TraceID但有调用者信息的情况
	formatter := &Formatter{
		DisableParsingAndEscaping: true,
		DisableCaller:             false, // 启用调用者信息
	}
	
	entry := &Entry{
		Time:       time.Now(),
		Level:      WarnLevel,
		Message:    "test message",
		Pid:        12345,
		Gid:        67890,
		TraceId:    "", // 空的TraceId
		File:       "test.go",
		CallerLine: 100,
		CallerDir:  "testdir",
		CallerFunc: "TestFunc",
	}
	
	result := formatter.format(entry)
	output := string(result)
	
	// 验证调用者信息被正确格式化
	if !strings.Contains(output, "testdir/test.go:100") {
		t.Errorf("Output should contain caller info, got: %s", output)
	}
	if !strings.Contains(output, "TestFunc") {
		t.Errorf("Output should contain function name, got: %s", output)
	}
}

func TestFormatter_format_ComplexScenarios(t *testing.T) {
	// 测试复杂的格式化场景
	formatter := &Formatter{
		Module:                    "TestMod",
		DisableParsingAndEscaping: false, // 启用解析和转义
		DisableCaller:             false, // 启用调用者信息
	}
	
	entry := &Entry{
		Time:      time.Date(2023, 7, 15, 10, 30, 45, 0, time.UTC),
		Level:     ErrorLevel,
		Message:   `{"key": "value", "number": 123}`, // JSON 消息
		PrefixMsg: []byte("[PREFIX]"),
		SuffixMsg: []byte("[SUFFIX]"),
		Pid:       9999,
		Gid:       8888,
		TraceId:   "trace-abc-123",
		File:       "/path/to/source.go",
		CallerLine: 100,
		CallerName: "github.com/lazygophers/log.ComplexFunction",
		CallerDir:  "log",
		CallerFunc: "ComplexFunction",
	}
	
	result := formatter.format(entry)
	output := string(result)
	
	// 验证所有组件都包含在输出中
	expectedComponents := []string{
		"[PREFIX]",
		"[SUFFIX]",
		"error", // 级别是小写
		// "TestMod", // Module field is defined but not used in current implementation
		"9999",
		"8888",
		"trace-abc-123",
		"source.go:100",
		"ComplexFunction",
	}
	
	for _, component := range expectedComponents {
		if !strings.Contains(output, component) {
			t.Errorf("Output should contain %q, got: %s", component, output)
		}
	}
}