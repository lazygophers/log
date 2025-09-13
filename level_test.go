package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestLevel_MarshalText 测试 Level 的 MarshalText 方法
func TestLevel_MarshalText(t *testing.T) {
	tests := []struct {
		level    Level
		expected string
	}{
		{TraceLevel, "trace"},
		{DebugLevel, "debug"},
		{InfoLevel, "info"},
		{WarnLevel, "warning"}, // 注意：MarshalText 中 WarnLevel 使用 "warning"
		{ErrorLevel, "error"},
		{FatalLevel, "fatal"},
		{PanicLevel, "panic"},
	}

	for _, tt := range tests {
		t.Run(tt.level.String(), func(t *testing.T) {
			result, err := tt.level.MarshalText()
			require.NoError(t, err, "MarshalText 不应该返回错误")
			require.Equal(t, tt.expected, string(result), "应该返回正确的文本表示")
		})
	}
}

// TestLevel_MarshalTextInvalidLevel 测试 MarshalText 处理无效级别
func TestLevel_MarshalTextInvalidLevel(t *testing.T) {
	// 创建一个无效的级别
	invalidLevel := Level(999)
	
	result, err := invalidLevel.MarshalText()
	require.Error(t, err, "无效级别应该返回错误")
	require.Nil(t, result, "错误时应该返回 nil")
	require.Contains(t, err.Error(), "not a valid logrus level", "错误消息应该包含预期文本")
}

// TestLevel_String_DefaultCase 测试 String 方法的默认情况
func TestLevel_String_DefaultCase(t *testing.T) {
	// 测试所有已知级别
	knownLevels := []Level{
		TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel,
	}

	for _, level := range knownLevels {
		result := level.String()
		require.NotEmpty(t, result, "已知级别应该返回非空字符串")
	}

	// 测试未知级别，应该返回默认值 "trace"
	unknownLevel := Level(999)
	result := unknownLevel.String()
	require.Equal(t, "trace", result, "未知级别应该返回默认值 'trace'")
}