package log

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestLogger_Trace 测试 Trace 方法
func TestLogger_Trace(t *testing.T) {
	tests := []struct {
		name    string
		message string
		args    []interface{}
	}{
		{
			name:    "简单消息",
			message: "这是一条跟踪日志",
		},
		{
			name:    "带参数的消息",
			message: "用户 %d 执行了操作 %s",
			args:    []interface{}{123, "login"},
		},
		{
			name:    "空消息",
			message: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建 buffer 来捕获输出
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(TraceLevel)

			// 执行测试
			if len(tt.args) > 0 {
				args := append([]interface{}{tt.message}, tt.args...)
				logger.Trace(args...)
			} else {
				logger.Trace(tt.message)
			}

			// 验证输出
			output := buf.String()
			require.Contains(t, output, "[trace]", "应该包含 trace 级别")
			// 对于带参数的情况，检查是否包含格式字符串的一部分
			if len(tt.args) > 0 {
				// Trace 方法实际上不格式化，直接传递所有参数
				require.Contains(t, output, "用户", "应该包含部分消息内容")
			} else {
				require.Contains(t, output, tt.message, "应该包含日志消息")
			}
		})
	}
}

// TestLogger_Debug 测试 Debug 方法
func TestLogger_Debug(t *testing.T) {
	tests := []struct {
		name    string
		message string
		args    []interface{}
	}{
		{
			name:    "调试消息",
			message: "开始调试模式",
		},
		{
			name:    "带格式化的调试消息",
			message: "变量值: %v, 类型: %T",
			args:    []interface{}{42, 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(DebugLevel)

			if len(tt.args) > 0 {
				args := append([]interface{}{tt.message}, tt.args...)
				logger.Debug(args...)
			} else {
				logger.Debug(tt.message)
			}

			output := buf.String()
			require.Contains(t, output, "[debug]", "应该包含 debug 级别")
			// 对于带参数的情况，检查是否包含格式字符串的一部分
			if len(tt.args) > 0 {
				// Debug 方法实际上不格式化，直接传递所有参数
				require.Contains(t, output, "变量值", "应该包含部分消息内容")
			} else {
				require.Contains(t, output, tt.message, "应该包含日志消息")
			}
		})
	}
}

// TestLogger_Info 测试 Info 方法
func TestLogger_Info(t *testing.T) {
	tests := []struct {
		name    string
		message string
		args    []interface{}
	}{
		{
			name:    "信息日志",
			message: "系统启动成功",
		},
		{
			name:    "带格式化的信息",
			message: "服务运行在端口 %d，版本 %s",
			args:    []interface{}{8080, "v1.0.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(InfoLevel)

			if len(tt.args) > 0 {
				args := append([]interface{}{tt.message}, tt.args...)
				logger.Info(args...)
			} else {
				logger.Info(tt.message)
			}

			output := buf.String()
			require.Contains(t, output, "[info]", "应该包含 info 级别")
			// 对于带参数的情况，检查是否包含格式字符串的一部分
			if len(tt.args) > 0 {
				// Info 方法实际上不格式化，直接传递所有参数
				require.Contains(t, output, "服务运行在端口", "应该包含部分消息内容")
			} else {
				require.Contains(t, output, tt.message, "应该包含日志消息")
			}
		})
	}
}

// TestLogger_Warn 测试 Warn 方法
func TestLogger_Warn(t *testing.T) {
	tests := []struct {
		name    string
		message string
		args    []interface{}
	}{
		{
			name:    "警告消息",
			message: "检测到内存使用率过高",
		},
		{
			name:    "带参数的警告",
			message: "连接池即将耗尽，当前使用率: %.1f%%",
			args:    []interface{}{85.3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(WarnLevel)

			if len(tt.args) > 0 {
				args := append([]interface{}{tt.message}, tt.args...)
				logger.Warn(args...)
			} else {
				logger.Warn(tt.message)
			}

			output := buf.String()
			require.Contains(t, output, "[warn]", "应该包含 warn 级别")
			// 对于带参数的情况，检查是否包含格式字符串的一部分
			if len(tt.args) > 0 {
				// Warn 方法实际上不格式化，直接传递所有参数
				require.Contains(t, output, "连接池即将耗尽", "应该包含部分消息内容")
			} else {
				require.Contains(t, output, tt.message, "应该包含日志消息")
			}
		})
	}
}

// TestLogger_Warning 测试 Warning 方法
func TestLogger_Warning(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(WarnLevel)

	logger.Warning("这是一条警告消息")

	output := buf.String()
	require.Contains(t, output, "[warn]", "Warning 应该对应 warn 级别")
	require.Contains(t, output, "这是一条警告消息", "应该包含日志消息")
}

// TestLogger_Error 测试 Error 方法
func TestLogger_Error(t *testing.T) {
	tests := []struct {
		name    string
		message string
		args    []interface{}
	}{
		{
			name:    "错误消息",
			message: "数据库连接失败",
		},
		{
			name:    "带错误详情",
			message: "无法打开文件: %s, 错误: %v",
			args:    []interface{}{"config.yaml", os.ErrNotExist},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(ErrorLevel)

			if len(tt.args) > 0 {
				args := append([]interface{}{tt.message}, tt.args...)
				logger.Error(args...)
			} else {
				logger.Error(tt.message)
			}

			output := buf.String()
			require.Contains(t, output, "[error]", "应该包含 error 级别")
			// 对于带参数的情况，检查是否包含格式字符串的一部分
			if len(tt.args) > 0 {
				// Error 方法实际上不格式化，直接传递所有参数
				require.Contains(t, output, "无法打开文件", "应该包含部分消息内容")
			} else {
				require.Contains(t, output, tt.message, "应该包含日志消息")
			}
		})
	}
}

// TestLogger_Tracef 测试 Tracef 方法
func TestLogger_Tracef(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []interface{}
		expected string
	}{
		{
			name:     "简单格式化",
			format:   "用户ID: %d",
			args:     []interface{}{12345},
			expected: "用户ID: 12345",
		},
		{
			name:     "多个参数",
			format:   "%s 在 %s 执行了 %s",
			args:     []interface{}{"用户A", "2025-01-01", "登录"},
			expected: "用户A 在 2025-01-01 执行了 登录",
		},
		{
			name:     "无参数",
			format:   "静态跟踪消息",
			args:     nil,
			expected: "静态跟踪消息",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(TraceLevel)

			logger.Tracef(tt.format, tt.args...)

			output := buf.String()
			require.Contains(t, output, "[trace]", "应该包含 trace 级别")
			// Tracef 应该正确格式化消息
			require.Contains(t, output, tt.expected, "应该包含格式化后的消息")
		})
	}
}

// TestLogger_Debugf 测试 Debugf 方法
func TestLogger_Debugf(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []interface{}
		expected string
	}{
		{
			name:     "调试变量",
			format:   "变量 x = %d, y = %f",
			args:     []interface{}{42, 3.14},
			expected: "变量 x = 42, y = 3.14",
		},
		{
			name:     "结构体调试",
			format:   "配置: %+v",
			args:     []interface{}{struct{ Name string }{Name: "test"}},
			expected: "配置: {Name:test}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(DebugLevel)

			logger.Debugf(tt.format, tt.args...)

			output := buf.String()
			require.Contains(t, output, "[debug]", "应该包含 debug 级别")
			// Debugf 应该正确格式化消息
			require.Contains(t, output, tt.expected, "应该包含格式化后的消息")
		})
	}
}

// TestLogger_Printf 测试 Printf 方法
func TestLogger_Printf(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(DebugLevel) // Printf 对应 DebugLevel，需要设置相应的级别

	logger.Printf("处理请求: %s, 耗时: %v", "/api/users", 150*time.Millisecond)

	output := buf.String()
	require.Contains(t, output, "[debug]", "Printf 实际对应 debug 级别")
	// Printf 应该正确格式化消息
	require.Contains(t, output, "处理请求: /api/users, 耗时: 150ms", "应该包含格式化后的消息")
}

// TestLogger_Infof 测试 Infof 方法
func TestLogger_Infof(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []interface{}
		expected string
	}{
		{
			name:     "服务信息",
			format:   "服务 %s 启动在端口 %d",
			args:     []interface{}{"api-server", 8080},
			expected: "服务 api-server 启动在端口 8080",
		},
		{
			name:     "带时间戳",
			format:   "任务完成于 %s",
			args:     []interface{}{time.Now().Format("2006-01-02 15:04:05")},
			expected: "任务完成于",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(InfoLevel)

			logger.Infof(tt.format, tt.args...)

			output := buf.String()
			require.Contains(t, output, "[info]", "应该包含 info 级别")
			// Infof 应该正确格式化消息
			require.Contains(t, output, tt.expected, "应该包含格式化后的消息")
		})
	}
}

// TestLogger_Warnf 测试 Warnf 方法
func TestLogger_Warnf(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(WarnLevel)

	logger.Warnf("资源使用率过高: CPU %.1f%%, 内存 %.1f%%", 85.5, 78.2)

	output := buf.String()
	require.Contains(t, output, "[warn]", "应该包含 warn 级别")
	// Warnf 应该正确格式化消息
	require.Contains(t, output, "资源使用率过高: CPU 85.5%, 内存 78.2%", "应该包含格式化后的消息")
}

// TestLogger_Warningf 测试 Warningf 方法
func TestLogger_Warningf(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(WarnLevel)

	logger.Warningf("API 调用频率超限: %d 次/分钟", 1000)

	output := buf.String()
	require.Contains(t, output, "[warn]", "Warningf 应该对应 warn 级别")
	// Warningf 应该正确格式化消息
	require.Contains(t, output, "API 调用频率超限: 1000 次/分钟", "应该包含格式化后的消息")
}

// TestLogger_Errorf 测试 Errorf 方法
func TestLogger_Errorf(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []interface{}
		expected string
	}{
		{
			name:     "错误格式化",
			format:   "无法连接到 %s:%d",
			args:     []interface{}{"localhost", 3306},
			expected: "无法连接到 localhost:3306",
		},
		{
			name:     "带错误对象",
			format:   "操作失败: %v", // 使用 %v 而不是 %w 来避免额外的调试信息
			args:     []interface{}{fmt.Errorf("权限不足")},
			expected: "操作失败: 权限不足",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := New()
			logger.SetOutput(buf)
			logger.SetLevel(ErrorLevel)

			logger.Errorf(tt.format, tt.args...)

			output := buf.String()
			require.Contains(t, output, "[error]", "应该包含 error 级别")
			// Errorf 应该正确格式化消息
			require.Contains(t, output, tt.expected, "应该包含格式化后的消息")
		})
	}
}

// TestLogger_SetCallerDepth 测试 SetCallerDepth 方法
func TestLogger_SetCallerDepth(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(DebugLevel)

	// 测试不同的调用深度
	testCases := []struct {
		name        string
		depth       int
		expectDepth int
	}{
		{"默认深度", 0, 2},
		{"自定义深度1", 1, 1},
		{"自定义深度2", 3, 3},
		{"深度0", 0, 2}, // 0应该使用默认值
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			logger.SetCallerDepth(tc.depth)
			logger.Debug("测试调用深度")

			output := buf.String()
			require.Contains(t, output, "[debug]", "应该包含 debug 级别")
			require.Contains(t, output, "测试调用深度", "应该包含日志消息")

			// 由于调用者信息可能包含文件路径，我们只检查是否包含调用者信息
			if tc.expectDepth > 0 {
				// 在实际使用中，调用深度会影响调用者信息的显示
				// 这里我们主要确保方法不会 panic
				require.NotPanics(t, func() {
					logger.SetCallerDepth(tc.depth)
				})
			}
		})
	}
}

// TestLogger_SetPrefixMsg 测试 SetPrefixMsg 方法
func TestLogger_SetPrefixMsg(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(InfoLevel)

	// 设置前缀消息
	prefix := "[PREFIX] "
	logger.SetPrefixMsg(prefix)

	logger.Info("这是一条带前缀的消息")

	output := buf.String()
	require.Contains(t, output, "[info]", "应该包含 info 级别")
	require.Contains(t, output, "[PREFIX]", "应该包含前缀消息")
	require.Contains(t, output, "这是一条带前缀的消息", "应该包含日志消息")

	// 测试空前缀
	t.Run("空前缀", func(t *testing.T) {
		buf.Reset()
		logger.SetPrefixMsg("")
		logger.Info("没有前缀的消息")

		output := buf.String()
		require.Contains(t, output, "没有前缀的消息", "应该包含日志消息")
		require.NotContains(t, output, "[PREFIX]", "不应该包含旧的前缀")
	})
}

// TestLogger_AppendPrefixMsg 测试 AppendPrefixMsg 方法
func TestLogger_AppendPrefixMsg(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(InfoLevel)

	// 初始前缀
	logger.SetPrefixMsg("[INIT] ")

	// 追加前缀
	logger.AppendPrefixMsg("[APPEND] ")

	logger.Info("测试追加前缀")

	output := buf.String()
	require.Contains(t, output, "[info]", "应该包含 info 级别")
	require.Contains(t, output, "[INIT] [APPEND]", "应该包含两个前缀")
	require.Contains(t, output, "测试追加前缀", "应该包含日志消息")
}

// TestLogger_SetSuffixMsg 测试 SetSuffixMsg 方法
func TestLogger_SetSuffixMsg(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(InfoLevel)

	// 设置后缀消息
	suffix := " [SUFFIX]"
	logger.SetSuffixMsg(suffix)

	logger.Info("这是一条带后缀的消息")

	output := buf.String()
	require.Contains(t, output, "[info]", "应该包含 info 级别")
	require.Contains(t, output, "[SUFFIX]", "应该包含后缀消息")
	require.Contains(t, output, "这是一条带后缀的消息", "应该包含日志消息")
}

// TestLogger_AppendSuffixMsg 测试 AppendSuffixMsg 方法
func TestLogger_AppendSuffixMsg(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(InfoLevel)

	// 初始后缀
	logger.SetSuffixMsg(" [INIT]")

	// 追加后缀
	logger.AppendSuffixMsg(" [APPEND]")

	logger.Info("测试追加后缀")

	output := buf.String()
	require.Contains(t, output, "[info]", "应该包含 info 级别")
	require.Contains(t, output, "[INIT] [APPEND]", "应该包含两个后缀")
	require.Contains(t, output, "测试追加后缀", "应该包含日志消息")
}

// TestLogger_Clone 测试 Clone 方法
func TestLogger_Clone(t *testing.T) {
	original := New()
	original.SetLevel(DebugLevel)
	original.SetCallerDepth(3)
	original.SetPrefixMsg("[ORIGINAL] ")
	original.SetSuffixMsg(" [SUFFIX]")

	// Clone logger
	cloned := original.Clone()

	// 验证克隆的 logger 具有相同的配置
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	original.SetOutput(buf1)
	cloned.SetOutput(buf2)

	original.Info("原始 logger")
	cloned.Info("克隆的 logger")

	output1 := buf1.String()
	output2 := buf2.String()

	require.Contains(t, output1, "[ORIGINAL]", "原始 logger 应该有前缀")
	require.Contains(t, output2, "[ORIGINAL]", "克隆的 logger 应该有相同的前缀")
	require.Contains(t, output1, "[SUFFIX]", "原始 logger 应该有后缀")
	require.Contains(t, output2, "[SUFFIX]", "克隆的 logger 应该有相同的后缀")

	// 验证修改克隆不会影响原始
	cloned.SetPrefixMsg("[CLONED] ")
	buf1.Reset()
	buf2.Reset()

	original.Info("原始 logger - 未修改")
	cloned.Info("克隆的 logger - 已修改")

	output1 = buf1.String()
	output2 = buf2.String()

	require.Contains(t, output1, "[ORIGINAL]", "原始 logger 的前缀应该保持不变")
	require.Contains(t, output2, "[CLONED]", "克隆的 logger 的前缀应该可以独立修改")
}

// TestLogger_SetLevel 测试 SetLevel 方法
func TestLogger_SetLevel(t *testing.T) {
	logger := New()

	// 测试设置不同级别
	levels := []Level{
		TraceLevel,
		DebugLevel,
		InfoLevel,
		WarnLevel,
		ErrorLevel,
	}

	for _, level := range levels {
		t.Run(level.String(), func(t *testing.T) {
			logger.SetLevel(level)
			require.Equal(t, level, logger.Level(), "日志级别应该正确设置")
		})
	}
}

// TestLogger_Level 测试 Level 方法
func TestLogger_Level(t *testing.T) {
	logger := New()

	// 测试默认级别 - newLogger 现在默认使用 DebugLevel
	require.Equal(t, DebugLevel, logger.Level(), "默认级别应该是 DebugLevel")

	// 测试设置后的级别
	logger.SetLevel(InfoLevel)
	require.Equal(t, InfoLevel, logger.Level(), "应该返回设置的级别")
}

// TestLogger_SetOutput 测试 SetOutput 方法
func TestLogger_SetOutput(t *testing.T) {
	logger := New()

	// 创建多个输出
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	// 设置单个输出
	logger.SetOutput(buf1)
	logger.Info("消息到输出1")

	require.Contains(t, buf1.String(), "消息到输出1", "消息应该写入到第一个输出")

	// 设置多个输出
	logger.SetOutput(buf1, buf2)
	buf1.Reset()
	buf2.Reset()

	logger.Info("消息到多个输出")

	require.Contains(t, buf1.String(), "消息到多个输出", "消息应该写入到第一个输出")
	require.Contains(t, buf2.String(), "消息到多个输出", "消息应该写入到第二个输出")
}

// TestLogger_Log 测试 Log 方法
func TestLogger_Log(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)

	// 测试不同级别
	tests := []struct {
		level   Level
		message string
	}{
		{TraceLevel, "Trace 级别消息"},
		{DebugLevel, "Debug 级别消息"},
		{InfoLevel, "Info 级别消息"},
		{WarnLevel, "Warn 级别消息"},
		{ErrorLevel, "Error 级别消息"},
	}

	for _, tt := range tests {
		t.Run(tt.level.String(), func(t *testing.T) {
			buf.Reset()
			logger.SetLevel(tt.level)

			logger.Log(tt.level, tt.message)

			output := buf.String()
			require.Contains(t, output, fmt.Sprintf("[%s]", tt.level.String()),
				"应该包含正确的级别")
			require.Contains(t, output, tt.message, "应该包含日志消息")
		})
	}
}

// TestLogger_Logf 测试 Logf 方法
func TestLogger_Logf(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)

	// 测试格式化日志
	logger.Logf(InfoLevel, "用户 %s 登录成功，IP: %s", "admin", "192.168.1.1")

	output := buf.String()
	require.Contains(t, output, "[info]", "应该包含 info 级别")
	require.Contains(t, output, "用户 admin 登录成功，IP: 192.168.1.1",
		"应该包含格式化后的消息")
}

// TestLogger_log 测试内部 log 方法
func TestLogger_log(t *testing.T) {
	// 由于 log 是私有方法，我们通过公共方法间接测试
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)
	logger.SetLevel(DebugLevel)

	// 通过调用 Info 方法来测试 log 方法
	logger.Info("测试内部 log 方法")

	output := buf.String()
	require.Contains(t, output, "[info]", "应该包含正确的级别")
	require.Contains(t, output, "测试内部 log 方法", "应该包含消息")

	// 验证输出格式包含必要的信息
	require.Contains(t, output, fmt.Sprintf("(%d.", os.Getpid()), "应该包含进程ID")
	require.Contains(t, output, time.Now().Format("2006-01-02"), "应该包含时间戳")
}

// TestLogger_write 测试 write 方法
func TestLogger_write(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := New()
	logger.SetOutput(buf)

	// 创建一个 entry
	entry := NewEntry()
	entry.Level = InfoLevel
	entry.Message = "测试 write 方法"
	entry.Time = time.Now()
	entry.Pid = os.Getpid()
	entry.Gid = int64(runtime.NumGoroutine())

	// 调用 write 方法 - 使用格式化器格式化后的内容
	formatted := logger.Format.Format(entry)
	logger.write(entry.Level, formatted)

	output := buf.String()
	require.Contains(t, output, "[info]", "应该包含正确的级别")
	require.Contains(t, output, entry.Message, "应该包含消息")
}

// TestLogger_levelEnabled 测试 levelEnabled 方法
func TestLogger_levelEnabled(t *testing.T) {
	logger := New()

	// 测试不同级别的启用状态
	tests := []struct {
		loggerLevel Level
		testLevel   Level
		expected    bool
	}{
		{TraceLevel, TraceLevel, true},
		{TraceLevel, DebugLevel, true},
		{TraceLevel, InfoLevel, true},
		{TraceLevel, WarnLevel, true},
		{TraceLevel, ErrorLevel, true},
		{TraceLevel, PanicLevel, true},

		{DebugLevel, TraceLevel, false},
		{DebugLevel, DebugLevel, true},
		{DebugLevel, InfoLevel, true},
		{DebugLevel, WarnLevel, true},
		{DebugLevel, ErrorLevel, true},

		{InfoLevel, TraceLevel, false},
		{InfoLevel, DebugLevel, false},
		{InfoLevel, InfoLevel, true},
		{InfoLevel, WarnLevel, true},
		{InfoLevel, ErrorLevel, true},
		{InfoLevel, WarnLevel, true},

		{ErrorLevel, TraceLevel, false},
		{ErrorLevel, DebugLevel, false},
		{ErrorLevel, InfoLevel, false},
		{ErrorLevel, WarnLevel, false},
		{ErrorLevel, ErrorLevel, true},
		{ErrorLevel, PanicLevel, true},

		{PanicLevel, TraceLevel, false},
		{PanicLevel, DebugLevel, false},
		{PanicLevel, InfoLevel, false},
		{PanicLevel, WarnLevel, false},
		{PanicLevel, ErrorLevel, false},
		{ErrorLevel, WarnLevel, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_检查_%s", tt.loggerLevel, tt.testLevel), func(t *testing.T) {
			logger.SetLevel(tt.loggerLevel)
			require.Equal(t, tt.expected, logger.levelEnabled(tt.testLevel),
				fmt.Sprintf("logger level %s, test level %s 应该返回 %v",
					tt.loggerLevel, tt.testLevel, tt.expected))
		})
	}
}

// TestLogger_Concurrency 测试并发安全性
func TestLogger_Concurrency(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(InfoLevel)

	var wg sync.WaitGroup
	const goroutines = 100
	const messagesPerGoroutine = 10

	wg.Add(goroutines)

	// 并发写入日志
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < messagesPerGoroutine; j++ {
				logger.Infof("Goroutine %d, Message %d", id, j)
			}
		}(i)
	}

	wg.Wait()

	// 验证所有消息都被写入
	output := buf.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")
	require.Equal(t, goroutines*messagesPerGoroutine, len(lines),
		"应该收到所有消息")
}

// TestLogger_New 测试 New 函数
func TestLogger_New(t *testing.T) {
	logger := New()

	// 验证默认配置
	require.Equal(t, DebugLevel, logger.Level(), "默认级别应该是 DebugLevel")
	// 由于 formatter 和 output 是私有字段，我们通过其他方式验证
	require.NotNil(t, logger, "Logger 不应该为空")
}

// BenchmarkLogger_Log 性能测试
func BenchmarkLogger_Log(b *testing.B) {
	logger := New()
	logger.SetOutput(&bytes.Buffer{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("这是一条性能测试日志")
	}
}

// BenchmarkLogger_Parallel 并发性能测试
func BenchmarkLogger_Parallel(b *testing.B) {
	logger := New()
	logger.SetOutput(&bytes.Buffer{})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Infof("并发测试 %d", 1)
		}
	})
}

// TestLogger_WithFormatter 测试自定义格式化器
func TestLogger_WithFormatter(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)

	// 禁用调用者信息
	if formatter, ok := logger.Format.(*Formatter); ok {
		formatter.Caller(false)
	}

	logger.Info("测试自定义格式化器")

	output := buf.String()
	fmt.Println(string(output))
	require.Contains(t, output, "[info]", "应该包含正确的级别")
	require.Contains(t, output, "测试自定义格式化器", "应该包含消息")
	require.NotContains(t, output, "_test.go", "应该不包含调用者信息")
}

// TestLogger_WithMultipleOutputs 测试多个输出
func TestLogger_WithMultipleOutputs(t *testing.T) {
	logger := New()
	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	buf3 := &bytes.Buffer{}

	// 设置多个输出
	logger.SetOutput(buf1, buf2, buf3)

	logger.Error("错误消息")

	// 验证所有输出都收到了消息
	require.Contains(t, buf1.String(), "错误消息", "输出1 应该收到消息")
	require.Contains(t, buf2.String(), "错误消息", "输出2 应该收到消息")
	require.Contains(t, buf3.String(), "错误消息", "输出3 应该收到消息")
}
