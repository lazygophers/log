package log

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	std.Log(InfoLevel, "This is an info message")
	if !bytes.Contains(buf.Bytes(), []byte("info")) {
		t.Errorf("Expected INFO level in output, got %q", buf.String())
	}
}

// 测试 Info 函数
func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	std.Info("This is an info message")
	if !bytes.Contains(buf.Bytes(), []byte("info")) {
		t.Errorf("Expected INFO level in output, got %q", buf.String())
	}
}

// 测试 Warn 函数
func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	std.Warn("This is a warning message")
	if !bytes.Contains(buf.Bytes(), []byte("warn")) {
		t.Errorf("Expected WARN level in output, got %q", buf.String())
	}
}

// 测试 Error 函数
func TestError(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	std.Error("This is an error message")
	if !bytes.Contains(buf.Bytes(), []byte("error")) {
		t.Errorf("Expected ERROR level in output, got %q", buf.String())
	}
}

// 测试 Panic 函数
func TestPanic(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()

	Panic("This is a panic message")
}

// 测试 Logf 函数
func TestLogf(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	Logf(InfoLevel, "Formatted message: %s", "test")
	expected := fmt.Sprintf("Formatted message: test")
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("Expected formatted output, got %q", buf.String())
	}
}

// 测试 Infof 函数
func TestInfof(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	Infof("Formatted message: %s", "test")
	expected := fmt.Sprintf("Formatted message: test")
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("Expected formatted output, got %q", buf.String())
	}
}

// 测试 Warnf 函数
func TestWarnf(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	Warnf("Formatted message: %s", "test")
	expected := fmt.Sprintf("Formatted message: test")
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("Expected formatted output, got %q", buf.String())
	}
}

// 测试 Errorf 函数
func TestErrorf(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	Errorf("Formatted message: %s", "test")
	expected := fmt.Sprintf("Formatted message: test")
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("Expected formatted output, got %q", buf.String())
	}
}

// 测试 Panicf 函数
func TestPanicf(t *testing.T) {
	var buf bytes.Buffer
	std.SetOutput(&buf)
	defer std.SetOutput(nil)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()

	Panicf("Formatted message: %s", "test")
}
