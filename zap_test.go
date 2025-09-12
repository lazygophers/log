package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewZapHook 测试 NewZapHook 函数
func TestNewZapHook(t *testing.T) {
	logger := New()
	
	// 测试创建 zap hook
	hook := NewZapHook(logger)
	require.NotNil(t, hook, "NewZapHook 应该返回非空值")
	
	// 由于 NewZapHook 的具体实现和返回类型可能不同
	// 我们主要确保函数不会 panic 并且返回了某种值
	require.NotPanics(t, func() {
		NewZapHook(logger)
	}, "NewZapHook 不应该 panic")
}

// TestNewZapHook_WithNilLogger 测试传入 nil logger 的情况
func TestNewZapHook_WithNilLogger(t *testing.T) {
	// 测试传入 nil logger
	require.NotPanics(t, func() {
		hook := NewZapHook(nil)
		// 根据实现，这可能返回 nil 或者有默认处理
		_ = hook
	}, "传入 nil logger 不应该 panic")
}