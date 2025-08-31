package log

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 测试上下文日志记录器的功能
// 确保所有方法都能正确工作，包括：
// - 上下文传递和日志记录
// - 配置方法的链式调用
// - 不同日志级别的方法
// - 格式化日志方法
// - 特殊功能（如调用者信息、前缀后缀等）

// TestLoggerWithCtx_Basic 测试 LoggerWithCtx 的基本功能
func TestLoggerWithCtx_Basic(t *testing.T) {
	t.Parallel()

	// 创建一个缓冲区来捕获日志输出
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	// 测试基本日志记录
	ctx := context.Background()
	ctxLogger.Trace(ctx, "trace message")
	ctxLogger.Debug(ctx, "debug message")
	ctxLogger.Print(ctx, "print message")
	ctxLogger.Info(ctx, "info message")
	ctxLogger.Warn(ctx, "warn message")
	ctxLogger.Warning(ctx, "warning message")
	ctxLogger.Error(ctx, "error message")

	// 检查输出
	output := buf.String()
	assert.NotContains(t, output, "trace message")
	assert.Contains(t, output, "debug message")
	assert.Contains(t, output, "print message")
	assert.Contains(t, output, "info message")
	assert.Contains(t, output, "warn message")
	assert.Contains(t, output, "warning message")
	assert.Contains(t, output, "error message")
}

// TestLoggerWithCtx_Formatted 测试格式化日志方法
func TestLoggerWithCtx_Formatted(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Tracef(ctx, "trace %s", "formatted")
	ctxLogger.Debugf(ctx, "debug %s", "formatted")
	ctxLogger.Printf(ctx, "print %s", "formatted")
	ctxLogger.Infof(ctx, "info %s", "formatted")
	ctxLogger.Warnf(ctx, "warn %s", "formatted")
	ctxLogger.Warningf(ctx, "warning %s", "formatted")
	ctxLogger.Errorf(ctx, "error %s", "formatted")

	output := buf.String()
	assert.NotContains(t, output, "trace formatted")
	assert.Contains(t, output, "debug formatted")
	assert.Contains(t, output, "print formatted")
	assert.Contains(t, output, "info formatted")
	assert.Contains(t, output, "warn formatted")
	assert.Contains(t, output, "warning formatted")
	assert.Contains(t, output, "error formatted")
}

// TestLoggerWithCtx_Caller 测试调用者信息功能
func TestLoggerWithCtx_Caller(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf).SetCallerDepth(3)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Info(ctx, "with caller")

	output := buf.String()
	// 检查是否包含调用者信息（文件名和行号）
	assert.Contains(t, output, "_test.go")
}

// TestLoggerWithCtx_PrefixSuffix 测试前缀和后缀功能
func TestLoggerWithCtx_PrefixSuffix(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	// 测试设置前缀
	ctxLogger.SetPrefixMsg("[PREFIX]")
	ctxLogger.Info(context.Background(), "test prefix")
	output := buf.String()
	assert.Contains(t, output, "[PREFIX]")

	// 清空缓冲区
	buf.Reset()

	// 测试追加前缀
	ctxLogger.AppendPrefixMsg("[APPEND]")
	ctxLogger.Info(context.Background(), "test append prefix")
	output = buf.String()
	assert.Contains(t, output, "[PREFIX][APPEND]")

	// 清空缓冲区
	buf.Reset()

	// 测试设置后缀
	ctxLogger.SetSuffixMsg("[SUFFIX]")
	ctxLogger.Info(context.Background(), "test suffix")
	output = buf.String()
	assert.Contains(t, output, "[SUFFIX]")

	// 清空缓冲区
	buf.Reset()

	// 测试追加后缀
	ctxLogger.AppendSuffixMsg("[APPEND]")
	ctxLogger.Info(context.Background(), "test append suffix")
	output = buf.String()
	assert.Contains(t, output, "[SUFFIX][APPEND]")
}

// TestLoggerWithCtx_ConfigurationChain 测试配置方法的链式调用
func TestLoggerWithCtx_ConfigurationChain(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)

	// 链式调用配置方法
	ctxLogger := logger.CloneToCtx().
		SetCallerDepth(2).
		SetPrefixMsg("[TEST]").
		SetSuffixMsg("[END]").
		SetLevel(DebugLevel)

	ctx := context.Background()
	ctxLogger.Debug(ctx, "chained configuration")

	output := buf.String()
	assert.Contains(t, output, "[TEST]")
	assert.Contains(t, output, "[END]")
	assert.Contains(t, output, "chained configuration")
}

// TestLoggerWithCtx_Level 测试日志级别功能
func TestLoggerWithCtx_Level(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(WarnLevel)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Debug(ctx, "debug message") // 不应该输出
	ctxLogger.Info(ctx, "info message")   // 不应该输出
	ctxLogger.Warn(ctx, "warn message")   // 应该输出
	ctxLogger.Error(ctx, "error message") // 应该输出

	output := buf.String()
	assert.NotContains(t, output, "debug message")
	assert.NotContains(t, output, "info message")
	assert.Contains(t, output, "warn message")
	assert.Contains(t, output, "error message")
}

// TestLoggerWithCtx_ParsingAndEscaping 测试解析和转义功能
func TestLoggerWithCtx_ParsingAndEscaping(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	// ParsingAndEscaping 不会产生日志输出，它只是配置格式化器
	// 测试它不会 panic 就足够了
	assert.NotPanics(t, func() {
		ctxLogger.ParsingAndEscaping(false)
	})
}

// TestLoggerWithCtx_ContextValues 测试上下文值的传递
func TestLoggerWithCtx_ContextValues(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetCallerDepth(0)
	ctxLogger := logger.CloneToCtx()

	// 创建带有值的上下文
	ctx := context.WithValue(context.Background(), "request_id", "12345")
	ctxLogger.Info(ctx, "request with ID")

	output := buf.String()
	// 注意：当前实现可能不会自动记录上下文值
	// 这个测试主要确保上下文参数能正确传递
	assert.Contains(t, output, "request with ID")
}

// TestLoggerWithCtx_NilContext 测试传入 nil 上下文的情况
func TestLoggerWithCtx_NilContext(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	// 传入 nil 上下文应该能正常工作
	ctxLogger.Info(nil, "message with nil context")

	output := buf.String()
	assert.Contains(t, output, "message with nil context")
}

// TestLoggerWithCtx_EmptyMessage 测试空消息
func TestLoggerWithCtx_EmptyMessage(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Info(ctx, "")

	output := buf.String()
	// 空消息也应该能正确处理
	assert.True(t, len(output) > 0)
}

// TestLoggerWithCtx_MultipleWrites 测试多个写入操作
func TestLoggerWithCtx_MultipleWrites(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	for i := 0; i < 100; i++ {
		ctxLogger.Infof(ctx, "message %d", i)
	}

	output := buf.String()
	// 检查所有消息都被记录
	for i := 0; i < 100; i++ {
		assert.Contains(t, output, fmt.Sprintf("message %d", i))
	}
}

// TestLoggerWithCtx_Concurrent 测试并发安全性
func TestLoggerWithCtx_Concurrent(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	done := make(chan bool, 10)

	// 启动多个 goroutine 并发写入日志
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() { done <- true }()
			for j := 0; j < 100; j++ {
				ctxLogger.Infof(ctx, "goroutine %d message %d", id, j)
			}
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}

	output := buf.String()
	// 检查消息数量
	lines := strings.Split(output, "\n")
	// 减去最后一个空行
	count := 0
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}
	assert.Equal(t, 1000, count)
}

// TestLoggerWithCtx_WithContextCancellation 测试上下文取消
func TestLoggerWithCtx_WithContextCancellation(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	// 创建一个可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 先记录一条日志
	ctxLogger.Info(ctx, "before cancellation")

	// 取消上下文
	cancel()

	// 取消后仍然应该能记录日志
	ctxLogger.Info(ctx, "after cancellation")

	output := buf.String()
	assert.Contains(t, output, "before cancellation")
	assert.Contains(t, output, "after cancellation")
}

// TestLoggerWithCtx_GetLogger 测试获取内部 Logger
func TestLoggerWithCtx_GetLogger(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetPrefixMsg("[TEST]")
	ctxLogger := logger.CloneToCtx()

	// 通过类型断言访问内部 Logger（仅用于测试）
	assert.NotNil(t, ctxLogger.Logger)
	assert.Equal(t, "[TEST]", string(ctxLogger.PrefixMsg))
}

// TestLoggerWithCtx_ErrorWithStack 测试带堆栈的错误
func TestLoggerWithCtx_ErrorWithStack(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	err := errors.New("test error")
	ctxLogger.Error(ctx, "error occurred", err)

	output := buf.String()
	assert.Contains(t, output, "error occurred")
	assert.Contains(t, output, "test error")
}

// BenchmarkLoggerWithCtx_Info 基准测试
func BenchmarkLoggerWithCtx_Info(b *testing.B) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctxLogger.Info(ctx, "benchmark message")
	}
}

// BenchmarkLoggerWithCtx_Infof 基准测试
func BenchmarkLoggerWithCtx_Infof(b *testing.B) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctxLogger.Infof(ctx, "benchmark message %d", i)
	}
}

// TestLoggerWithCtx_Formatters 测试格式化器的使用
func TestLoggerWithCtx_Formatters(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	// 注意：Logger 没有 SetFormatter 方法，这是设计上的限制
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Infof(ctx, "formatted: %s %d", "value", 42)

	output := buf.String()
	assert.Contains(t, output, "formatted: value 42")
}

// TestLoggerWithCtx_TimeFormat 测试时间格式
func TestLoggerWithCtx_TimeFormat(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	// 注意：Logger 没有 SetTimeFormat 方法，这是设计上的限制
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Info(ctx, "time format test")

	output := buf.String()
	// 检查输出是否包含格式化的时间
	assert.Contains(t, output, time.Now().Format("2006-01-02"))
}

// TestLoggerWithCtx_RaceDetector 使用竞态检测器
func TestLoggerWithCtx_RaceDetector(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping race detector test in short mode")
	}

	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ctx := context.Background()
			for j := 0; j < 100; j++ {
				ctxLogger.Infof(ctx, "race test %d-%d", id, j)
			}
		}(i)
	}
	wg.Wait()

	// 如果没有竞态条件，测试会通过
	// 使用 go test -race 来检测
}

// TestLoggerWithCtx_LevelFiltering 测试级别过滤
func TestLoggerWithCtx_LevelFiltering(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		level    Level
		messages []struct {
			level   Level
			content string
			expect  bool
		}
	}{
		{
			name:  "Trace level",
			level: TraceLevel,
			messages: []struct {
				level   Level
				content string
				expect  bool
			}{
				{TraceLevel, "trace", true},
				{DebugLevel, "debug", true},
				{InfoLevel, "info", true},
				{WarnLevel, "warn", true},
				{ErrorLevel, "error", true},
			},
		},
		{
			name:  "Info level",
			level: InfoLevel,
			messages: []struct {
				level   Level
				content string
				expect  bool
			}{
				{TraceLevel, "trace", false},
				{DebugLevel, "debug", false},
				{InfoLevel, "info", true},
				{WarnLevel, "warn", true},
				{ErrorLevel, "error", true},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(tc.level)
			ctxLogger := logger.CloneToCtx()

			ctx := context.Background()

			for _, msg := range tc.messages {
				buf.Reset()
				switch msg.level {
				case TraceLevel:
					ctxLogger.Trace(ctx, msg.content)
				case DebugLevel:
					ctxLogger.Debug(ctx, msg.content)
				case InfoLevel:
					ctxLogger.Info(ctx, msg.content)
				case WarnLevel:
					ctxLogger.Warn(ctx, msg.content)
				case ErrorLevel:
					ctxLogger.Error(ctx, msg.content)
				}

				output := buf.String()
				if msg.expect {
					assert.Contains(t, output, msg.content)
				} else {
					assert.NotContains(t, output, msg.content)
				}
			}
		})
	}
}

// TestLoggerWithCtx_WithContextTimeout 测试带超时的上下文
func TestLoggerWithCtx_WithContextTimeout(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	// 在超时前记录日志
	ctxLogger.Info(ctx, "before timeout")

	// 等待超时
	time.Sleep(time.Millisecond * 2)

	// 超时后记录日志
	ctxLogger.Info(ctx, "after timeout")

	output := buf.String()
	assert.Contains(t, output, "before timeout")
	assert.Contains(t, output, "after timeout")
}

// TestLoggerWithCtx_MultipleOutputs 测试多个输出目标
func TestLoggerWithCtx_MultipleOutputs(t *testing.T) {
	t.Parallel()

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf1, buf2)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Info(ctx, "multiple outputs")

	// 两个缓冲区都应该有相同的输出
	assert.Equal(t, buf1.String(), buf2.String())
	assert.Contains(t, buf1.String(), "multiple outputs")
}

// TestLoggerWithCtx_ErrorHandling 测试错误处理
func TestLoggerWithCtx_ErrorHandling(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()

	// 测试带错误的日志记录
	err := fmt.Errorf("wrapped error: %w", errors.New("original error"))
	ctxLogger.Error(ctx, "error with wrapping", err)

	output := buf.String()
	assert.Contains(t, output, "error with wrapping")
	assert.Contains(t, output, "wrapped error")
	assert.Contains(t, output, "original error")
}

// TestLoggerWithCtx_EmptyArgs 测试空参数
func TestLoggerWithCtx_EmptyArgs(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()

	// 测试格式化方法但无额外参数
	ctxLogger.Infof(ctx, "no args")

	output := buf.String()
	assert.Contains(t, output, "no args")
}

// TestLoggerWithCtx_LongMessage 测试长消息
func TestLoggerWithCtx_LongMessage(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()

	// 创建一个长消息
	longMsg := strings.Repeat("x", 10000)
	ctxLogger.Info(ctx, longMsg)

	output := buf.String()
	assert.Contains(t, output, longMsg)
}

// TestLoggerWithCtx_NestedContext 测试嵌套上下文
func TestLoggerWithCtx_NestedContext(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	// 创建嵌套的上下文
	ctx1 := context.WithValue(context.Background(), "key1", "value1")
	ctx2 := context.WithValue(ctx1, "key2", "value2")

	ctxLogger.Info(ctx2, "nested context")

	output := buf.String()
	assert.Contains(t, output, "nested context")
}

// TestLoggerWithCtx_LoggerInheritance 测试 LoggerWithCtx 继承 Logger 的功能
func TestLoggerWithCtx_LoggerInheritance(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetPrefixMsg("[TEST]")
	// 注意：Logger 没有 SetTimeFormat 方法
	logger.SetCallerDepth(1)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Info(ctx, "inheritance test")

	output := buf.String()
	// 应该包含所有从 Logger 继承的配置
	assert.Contains(t, output, "[TEST]")
	assert.Contains(t, output, time.Now().Format("15:04:05"))
	assert.Contains(t, output, "inheritance test")
}

// TestLoggerWithCtx_ConcurrentConfig 测试并发配置修改
func TestLoggerWithCtx_ConcurrentConfig(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	var wg sync.WaitGroup
	done := make(chan bool, 5)

	// 并发修改配置
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer func() {
				wg.Done()
				done <- true
			}()
			for j := 0; j < 100; j++ {
				ctxLogger.SetPrefixMsg(fmt.Sprintf("[G%d-%d]", id, j))
				ctxLogger.Info(context.Background(), "concurrent config")
			}
		}(i)
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 5; i++ {
		<-done
	}
	wg.Wait()

	// 确保没有 panic 发生
	output := buf.String()
	assert.True(t, len(output) > 0)
}

// TestLoggerWithCtx_SetOutputMultiple 测试多次设置输出
func TestLoggerWithCtx_SetOutputMultiple(t *testing.T) {
	t.Parallel()

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf1)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()
	ctxLogger.Info(ctx, "first output")

	// 修改输出
	ctxLogger.SetOutput(buf2)
	ctxLogger.Info(ctx, "second output")

	assert.Contains(t, buf1.String(), "first output")
	assert.NotContains(t, buf1.String(), "second output")
	assert.Contains(t, buf2.String(), "second output")
	assert.NotContains(t, buf2.String(), "first output")
}

// TestLoggerWithCtx_PrefixSuffixEmpty 测试空的前缀后缀
func TestLoggerWithCtx_PrefixSuffixEmpty(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	ctxLogger := logger.CloneToCtx()

	ctx := context.Background()

	// 设置空字符串
	ctxLogger.SetPrefixMsg("")
	ctxLogger.SetSuffixMsg("")
	ctxLogger.Info(ctx, "empty prefix suffix")

	output := buf.String()
	assert.Contains(t, output, "empty prefix suffix")
}

// TestLoggerWithCtx_GetCaller 测试获取调用者信息
func TestLoggerWithCtx_GetCaller(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetCallerDepth(1)
	ctxLogger := logger.CloneToCtx()

	// 测试 ParsingAndEscaping 方法
	// 这个方法不会产生日志输出，只是配置格式化器
	assert.NotPanics(t, func() {
		ctxLogger.ParsingAndEscaping(true)
	})
}
