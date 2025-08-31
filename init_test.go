package log

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPid 测试 Pid 函数是否返回正确的进程 ID
func TestPid(t *testing.T) {
	// 获取当前进程 ID
	expectedPid := os.Getpid()
	
	// 调用 Pid 函数
	actualPid := Pid()
	
	// 验证返回的 PID 是否正确
	assert.Equal(t, expectedPid, actualPid, "Pid() 应该返回当前进程的 ID")
}

// TestNew 测试 New 函数是否返回一个新的 Logger 实例
func TestNew(t *testing.T) {
	// 调用 New 函数
	logger := New()
	
	// 验证返回的不是 nil
	require.NotNil(t, logger, "New() 不应该返回 nil")
	
	// 验证返回的是 Logger 类型
	assert.IsType(t, &Logger{}, logger, "New() 应该返回 *Logger 类型")
	
	// 验证新的 logger 有默认值
	assert.Equal(t, DebugLevel, logger.level, "新 logger 的默认级别应该是 DebugLevel")
}

// TestSetLevel 和 TestGetLevel 测试设置和获取日志级别
func TestSetLevelAndGetLevel(t *testing.T) {
	// 保存原始级别以便恢复
	originalLevel := GetLevel()
	defer SetLevel(originalLevel)
	
	// 测试设置和获取各种级别
	testLevels := []Level{
		DebugLevel,
		InfoLevel,
		WarnLevel,
		ErrorLevel,
		FatalLevel,
		PanicLevel,
	}
	
	for _, level := range testLevels {
		// 设置级别
		SetLevel(level)
		
		// 获取级别
		retrievedLevel := GetLevel()
		
		// 验证级别设置正确
		assert.Equal(t, level, retrievedLevel, "GetLevel() 应该返回最近设置的级别")
	}
}

// TestSync 测试 Sync 函数是否能正常工作
func TestSync(t *testing.T) {
	// 这个测试主要是验证 Sync 函数不会 panic
	// 由于 std 是一个全局变量，我们无法直接测试其输出
	// 所以我们只是确保函数调用不会出错
	
	assert.NotPanics(t, func() {
		Sync()
	}, "Sync() 不应该 panic")
}

// TestClone 测试 Clone 函数是否能正确复制 Logger
func TestClone(t *testing.T) {
	// 设置原始 logger 的一些属性
	SetLevel(InfoLevel)
	SetPrefixMsg("[TEST]")
	SetSuffixMsg("[END]")
	SetCallerDepth(5)
	
	// 克隆 logger
	cloned := Clone()
	
	require.NotNil(t, cloned, "Clone() 不应该返回 nil")
	
	// 验证克隆的 logger 具有相同的属性
	assert.Equal(t, GetLevel(), cloned.Level(), "克隆的 logger 应该有相同的级别")
	
	// 修改原始 logger 的级别
	SetLevel(DebugLevel)
	
	// 验证克隆的 logger 不受影响
	assert.Equal(t, InfoLevel, cloned.Level(), "修改原始 logger 不应该影响克隆的 logger")
}

// TestSetCallerDepth 测试设置调用者深度
func TestSetCallerDepth(t *testing.T) {
	// 保存原始深度
	originalDepth := 4 // 默认深度
	
	// 测试设置不同的深度
	testDepths := []int{2, 4, 6, 8}
	
	for _, depth := range testDepths {
		SetCallerDepth(depth)
		// 由于 callerDepth 是私有属性，我们无法直接验证
		// 但我们可以确保函数调用不会出错
		assert.NotPanics(t, func() {
			Info("测试日志")
		}, "SetCallerDepth() 后的日志记录不应该 panic")
	}
	
	// 恢复默认深度
	SetCallerDepth(originalDepth)
}

// TestSetPrefixMsg 测试设置前缀消息
func TestSetPrefixMsg(t *testing.T) {
	// 保存原始前缀
	originalPrefix := std.PrefixMsg
	defer func() {
		std.PrefixMsg = originalPrefix
	}()
	
	// 测试设置前缀
	testPrefix := "[TEST-PREFIX]"
	SetPrefixMsg(testPrefix)
	
	// 验证前缀已设置
	assert.Equal(t, []byte(testPrefix), std.PrefixMsg, "前缀应该被正确设置")
	
	// 测试空前缀
	SetPrefixMsg("")
	assert.Empty(t, std.PrefixMsg, "空字符串应该被正确设置为前缀")
}

// TestAppendPrefixMsg 测试追加前缀消息
func TestAppendPrefixMsg(t *testing.T) {
	// 保存原始前缀
	originalPrefix := std.PrefixMsg
	defer func() {
		std.PrefixMsg = originalPrefix
	}()
	
	// 设置初始前缀
	initialPrefix := "[INITIAL]"
	SetPrefixMsg(initialPrefix)
	
	// 追加前缀
	appendPrefix := "[APPENDED]"
	AppendPrefixMsg(appendPrefix)
	
	// 验证前缀被正确追加
	expectedPrefix := initialPrefix + appendPrefix
	assert.Equal(t, []byte(expectedPrefix), std.PrefixMsg, "前缀应该被正确追加")
}

// TestSetSuffixMsg 测试设置后缀消息
func TestSetSuffixMsg(t *testing.T) {
	// 保存原始后缀
	originalSuffix := std.SuffixMsg
	defer func() {
		std.SuffixMsg = originalSuffix
	}()
	
	// 测试设置后缀
	testSuffix := "[TEST-SUFFIX]"
	SetSuffixMsg(testSuffix)
	
	// 验证后缀已设置
	assert.Equal(t, []byte(testSuffix), std.SuffixMsg, "后缀应该被正确设置")
	
	// 测试空后缀
	SetSuffixMsg("")
	assert.Empty(t, std.SuffixMsg, "空字符串应该被正确设置为后缀")
}

// TestAppendSuffixMsg 测试追加后缀消息
func TestAppendSuffixMsg(t *testing.T) {
	// 保存原始后缀
	originalSuffix := std.SuffixMsg
	defer func() {
		std.SuffixMsg = originalSuffix
	}()
	
	// 设置初始后缀
	initialSuffix := "[INITIAL]"
	SetSuffixMsg(initialSuffix)
	
	// 追加后缀
	appendSuffix := "[APPENDED]"
	AppendSuffixMsg(appendSuffix)
	
	// 验证后缀被正确追加
	expectedSuffix := initialSuffix + appendSuffix
	assert.Equal(t, []byte(expectedSuffix), std.SuffixMsg, "后缀应该被正确追加")
}

// TestParsingAndEscaping 测试解析和转义设置
func TestParsingAndEscaping(t *testing.T) {
	// 由于 ParsingAndEscaping 操作的是 Formatter 的私有属性
	// 我们只能测试函数调用不会 panic，并且会正确处理非 FormatFull 的情况
	
	assert.NotPanics(t, func() {
		ParsingAndEscaping(true)
	}, "ParsingAndEscaping(true) 不应该 panic")
	
	assert.NotPanics(t, func() {
		ParsingAndEscaping(false)
	}, "ParsingAndEscaping(false) 不应该 panic")
}

// TestCaller 测试调用者信息控制
func TestCaller(t *testing.T) {
	// 由于 Caller 操作的是 Formatter 的私有属性
	// 我们只能测试函数调用不会 panic，并且会正确处理非 FormatFull 的情况
	
	assert.NotPanics(t, func() {
		Caller(true)
	}, "Caller(true) 不应该 panic")
	
	assert.NotPanics(t, func() {
		Caller(false)
	}, "Caller(false) 不应该 panic")
}

// TestIntegration 测试多个函数的集成使用
func TestIntegration(t *testing.T) {
	// 保存原始状态
	originalLevel := GetLevel()
	originalPrefix := std.PrefixMsg
	originalSuffix := std.SuffixMsg
	originalDepth := std.callerDepth
	
	defer func() {
		SetLevel(originalLevel)
		std.PrefixMsg = originalPrefix
		std.SuffixMsg = originalSuffix
		std.callerDepth = originalDepth
	}()
	
	// 创建一个缓冲区来捕获日志输出
	var buf bytes.Buffer
	SetOutput(&buf)
	defer SetOutput(os.Stdout)
	
	// 设置各种属性
	SetLevel(InfoLevel)
	SetPrefixMsg("[INTEGRATION]")
	SetSuffixMsg("[END]")
	SetCallerDepth(3)
	
	// 记录日志
	Info("集成测试消息")
	
	// 验证有输出（不验证具体内容，因为包含时间戳等动态信息）
	assert.True(t, buf.Len() > 0, "应该有日志输出")
	
	// 测试级别过滤
	buf.Reset()
	SetLevel(WarnLevel)
	Debug("这条调试消息不应该出现")
	assert.Equal(t, 0, buf.Len(), "Debug 消息不应该在 Warn 级别输出")
}

// TestConcurrentAccess 测试并发访问的安全性
func TestConcurrentAccess(t *testing.T) {
	// 这个测试验证多个 goroutine 同时访问全局函数时的安全性
	done := make(chan bool, 10)
	
	// 启动多个 goroutine 同时调用各种函数
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() { done <- true }()
			
			// 执行各种操作
			SetLevel(Level(id % 6))
			GetLevel()
			SetPrefixMsg("PREFIX")
			SetSuffixMsg("SUFFIX")
			New()
			Clone()
			Sync()
		}(i)
	}
	
	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}
	
	// 如果没有 panic，测试通过
}