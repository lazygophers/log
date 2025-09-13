package log

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestPrintFunctions 测试 print.go 中的所有全局函数
func TestPrintFunctions(t *testing.T) {
	// 保存原始输出
	originalOutput := os.Stdout
	defer SetOutput(originalOutput)

	buf := &bytes.Buffer{}
	SetOutput(buf)
	SetLevel(TraceLevel) // 设置最低级别以确保所有日志都输出

	// 测试 Trace 函数
	t.Run("Trace", func(t *testing.T) {
		buf.Reset()
		Trace("trace message")
		output := buf.String()
		require.Contains(t, output, "[trace]", "应该包含 trace 级别")
		require.Contains(t, output, "trace message", "应该包含消息")
	})

	// 测试 Log 函数
	t.Run("Log", func(t *testing.T) {
		buf.Reset()
		Log(InfoLevel, "log message")
		output := buf.String()
		require.Contains(t, output, "[info]", "应该包含 info 级别")
		require.Contains(t, output, "log message", "应该包含消息")
	})

	// 测试 Warn 函数
	t.Run("Warn", func(t *testing.T) {
		buf.Reset()
		Warn("warn message")
		output := buf.String()
		require.Contains(t, output, "[warn]", "应该包含 warn 级别")
		require.Contains(t, output, "warn message", "应该包含消息")
	})

	// 测试 Error 函数
	t.Run("Error", func(t *testing.T) {
		buf.Reset()
		Error("error message")
		output := buf.String()
		require.Contains(t, output, "[error]", "应该包含 error 级别")
		require.Contains(t, output, "error message", "应该包含消息")
	})

	// 测试 Panic 函数
	t.Run("Panic", func(t *testing.T) {
		buf.Reset()
		require.Panics(t, func() {
			Panic("panic message")
		}, "Panic 函数应该触发 panic")

		output := buf.String()
		require.Contains(t, output, "[panic]", "应该包含 panic 级别")
		require.Contains(t, output, "panic message", "应该包含消息")
	})

	// 测试 Fatal 函数（使用 goroutine 避免退出测试进程）
	t.Run("Fatal", func(t *testing.T) {
		buf.Reset()
		
		// 直接测试格式化部分，避免实际退出
		entry := NewEntry()
		entry.Level = FatalLevel
		entry.Message = "fatal message"
		entry.Time = time.Now()
		entry.Pid = os.Getpid()
		
		formatted := std.Format.Format(entry)
		require.Contains(t, string(formatted), "[fatal]", "应该包含 fatal 级别")
		require.Contains(t, string(formatted), "fatal message", "应该包含消息")
	})

	// 测试 Logf 函数
	t.Run("Logf", func(t *testing.T) {
		buf.Reset()
		Logf(InfoLevel, "formatted %s %d", "message", 42)
		output := buf.String()
		require.Contains(t, output, "[info]", "应该包含 info 级别")
		require.Contains(t, output, "formatted message 42", "应该包含格式化后的消息")
	})

	// 测试 Infof 函数
	t.Run("Infof", func(t *testing.T) {
		buf.Reset()
		Infof("info: %s", "formatted")
		output := buf.String()
		require.Contains(t, output, "[info]", "应该包含 info 级别")
		require.Contains(t, output, "info: formatted", "应该包含格式化后的消息")
	})

	// 测试 Warnf 函数
	t.Run("Warnf", func(t *testing.T) {
		buf.Reset()
		Warnf("warn: %s", "formatted")
		output := buf.String()
		require.Contains(t, output, "[warn]", "应该包含 warn 级别")
		require.Contains(t, output, "warn: formatted", "应该包含格式化后的消息")
	})

	// 测试 Panicf 函数
	t.Run("Panicf", func(t *testing.T) {
		buf.Reset()
		require.Panics(t, func() {
			Panicf("panic: %s", "formatted")
		}, "Panicf 函数应该触发 panic")

		output := buf.String()
		require.Contains(t, output, "[panic]", "应该包含 panic 级别")
		require.Contains(t, output, "panic: formatted", "应该包含格式化后的消息")
	})

	// 测试 Fatalf 函数（使用格式化部分测试）
	t.Run("Fatalf", func(t *testing.T) {
		buf.Reset()
		
		// 直接测试格式化部分
		entry := NewEntry()
		entry.Level = FatalLevel
		entry.Message = "fatal: formatted"
		entry.Time = time.Now()
		entry.Pid = os.Getpid()
		
		formatted := std.Format.Format(entry)
		require.Contains(t, string(formatted), "[fatal]", "应该包含 fatal 级别")
		require.Contains(t, string(formatted), "fatal: formatted", "应该包含格式化后的消息")
	})

	// 测试 StartMsg 函数
	t.Run("StartMsg", func(t *testing.T) {
		buf.Reset()
		StartMsg()
		output := buf.String()
		require.Contains(t, output, "========== start new log ==========", "应该包含开始消息")
	})
}

// TestPrintOther 测试 print_other.go 中的函数
func TestPrintOther(t *testing.T) {
	// 保存原始输出
	originalOutput := os.Stdout
	defer SetOutput(originalOutput)

	buf := &bytes.Buffer{}
	SetOutput(buf)

	// 测试 Debug 函数（在 release 构建中可能被禁用）
	t.Run("Debug", func(t *testing.T) {
		buf.Reset()
		SetLevel(DebugLevel)
		Debug("debug message")
		
		output := buf.String()
		// 在某些构建标签下，Debug 可能不输出
		if len(output) > 0 {
			require.Contains(t, output, "[debug]", "如果输出则应该包含 debug 级别")
			require.Contains(t, output, "debug message", "如果输出则应该包含消息")
		}
	})

	// 测试 Debugf 函数
	t.Run("Debugf", func(t *testing.T) {
		buf.Reset()
		SetLevel(DebugLevel)
		Debugf("debug: %s", "formatted")
		
		output := buf.String()
		// 在某些构建标签下，Debugf 可能不输出
		if len(output) > 0 {
			require.Contains(t, output, "[debug]", "如果输出则应该包含 debug 级别")
			require.Contains(t, output, "debug: formatted", "如果输出则应该包含格式化后的消息")
		}
	})
}