package log

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestFormatter_ComprehensiveFormat 测试 format 方法的所有分支
func TestFormatter_ComprehensiveFormat(t *testing.T) {
	formatter := &Formatter{}
	
	// 创建基础 entry
	entry := &Entry{
		Pid:        12345,
		Gid:        67890,
		Time:       time.Date(2023, 12, 25, 10, 30, 45, 123456000, time.UTC),
		Level:      InfoLevel,
		Message:    "测试消息",
		File:       "test.go",
		CallerName: "github.com/test/pkg.TestFunction",
		CallerLine: 100,
		CallerDir:  "github.com/test/pkg",
		CallerFunc: "TestFunction",
		TraceId:    "",
		PrefixMsg:  nil,
		SuffixMsg:  nil,
	}

	t.Run("基本格式化", func(t *testing.T) {
		result := formatter.format(entry)
		output := string(result)
		
		require.Contains(t, output, "(12345.67890)", "应该包含进程ID和协程ID")
		require.Contains(t, output, "2023-12-25 10:30:45", "应该包含格式化的时间")
		require.Contains(t, output, "[info]", "应该包含日志级别")
		require.Contains(t, output, "测试消息", "应该包含消息")
		require.Contains(t, output, "test.go:100", "应该包含文件名和行号")
		require.Contains(t, output, "TestFunction", "应该包含函数名")
	})

	t.Run("带前缀消息", func(t *testing.T) {
		entry.PrefixMsg = []byte("[PREFIX]")
		result := formatter.format(entry)
		output := string(result)
		
		require.Contains(t, output, "[PREFIX] ", "应该包含前缀消息")
		entry.PrefixMsg = nil // 重置
	})

	t.Run("带后缀消息", func(t *testing.T) {
		entry.SuffixMsg = []byte("[SUFFIX]")
		result := formatter.format(entry)
		output := string(result)
		
		require.Contains(t, output, "[SUFFIX]", "应该包含后缀消息")
		entry.SuffixMsg = nil // 重置
	})

	t.Run("带TraceId", func(t *testing.T) {
		entry.TraceId = "trace-12345"
		result := formatter.format(entry)
		output := string(result)
		
		require.Contains(t, output, "trace-12345", "应该包含TraceId")
		entry.TraceId = "" // 重置
	})

	t.Run("禁用调用者信息", func(t *testing.T) {
		formatter.DisableCaller = true
		result := formatter.format(entry)
		output := string(result)
		
		require.NotContains(t, output, "test.go:100", "禁用时不应该包含调用者信息")
		formatter.DisableCaller = false // 重置
	})

	t.Run("禁用调用者信息但有TraceId", func(t *testing.T) {
		formatter.DisableCaller = true
		entry.TraceId = "trace-67890"
		result := formatter.format(entry)
		output := string(result)
		
		require.NotContains(t, output, "test.go:100", "禁用时不应该包含调用者信息")
		require.Contains(t, output, "trace-67890", "应该包含TraceId")
		
		formatter.DisableCaller = false // 重置
		entry.TraceId = "" // 重置
	})

	t.Run("只有TraceId无调用者信息", func(t *testing.T) {
		// 这个测试覆盖 !p.DisableCaller || entry.TraceId != "" 条件的第二部分
		entry.TraceId = "only-trace"
		entry.CallerName = ""
		entry.CallerFunc = ""
		entry.File = ""
		entry.CallerLine = 0
		entry.CallerDir = ""
		
		result := formatter.format(entry)
		output := string(result)
		
		require.Contains(t, output, "only-trace", "应该包含TraceId")
		
		// 重置
		entry.TraceId = ""
		entry.CallerName = "github.com/test/pkg.TestFunction"
		entry.CallerFunc = "TestFunction"
		entry.File = "test.go"
		entry.CallerLine = 100
		entry.CallerDir = "github.com/test/pkg"
	})

	t.Run("不同级别的颜色", func(t *testing.T) {
		levels := []Level{TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel}
		
		for _, level := range levels {
			entry.Level = level
			result := formatter.format(entry)
			output := string(result)
			
			require.Contains(t, output, fmt.Sprintf("[%s]", level.String()), 
				"应该包含正确的级别: %s", level.String())
		}
		entry.Level = InfoLevel // 重置
	})

	t.Run("消息末尾空白处理", func(t *testing.T) {
		entry.Message = "  带空白的消息  \n\t"
		result := formatter.format(entry)
		output := string(result)
		
		require.Contains(t, output, "带空白的消息", "应该包含去除空白后的消息")
		require.NotContains(t, output, "  带空白的消息  ", "不应该包含原始带空白的消息")
		entry.Message = "测试消息" // 重置
	})
}

// TestFormatter_FormatMultilineHandling 测试多行消息处理
func TestFormatter_FormatMultilineHandling(t *testing.T) {
	formatter := &Formatter{}
	
	entry := &Entry{
		Pid:        os.Getpid(),
		Gid:        123,
		Time:       time.Now(),
		Level:      InfoLevel,
		Message:    "第一行\n第二行\n第三行",
		File:       "test.go",
		CallerName: "test.TestFunc",
		CallerLine: 50,
		CallerDir:  "test",
		CallerFunc: "TestFunc",
	}

	t.Run("禁用解析和转义-单行处理", func(t *testing.T) {
		formatter.DisableParsingAndEscaping = true
		result := formatter.Format(entry)
		output := string(result)
		
		// 当禁用解析时，消息被直接处理，不会按换行符分割
		require.Contains(t, output, "第一行", "应该包含第一行内容")
		require.Contains(t, output, "第二行", "应该包含第二行内容")  
		require.Contains(t, output, "第三行", "应该包含第三行内容")
		
		formatter.DisableParsingAndEscaping = false // 重置
	})

	t.Run("启用解析和转义-多行处理", func(t *testing.T) {
		formatter.DisableParsingAndEscaping = false
		result := formatter.Format(entry)
		output := string(result)
		
		// 当启用解析时，多行消息应该分别格式化
		lines := strings.Split(strings.TrimSpace(output), "\n")
		require.Equal(t, 3, len(lines), "启用解析时应该有三行输出")
		require.Contains(t, output, "第一行", "应该包含第一行")
		require.Contains(t, output, "第二行", "应该包含第二行")
		require.Contains(t, output, "第三行", "应该包含第三行")
	})
}

// TestFormatter_ParsingAndEscaping 测试 ParsingAndEscaping 方法
func TestFormatter_ParsingAndEscaping(t *testing.T) {
	formatter := &Formatter{}
	
	// 测试启用
	formatter.ParsingAndEscaping(false)
	require.False(t, formatter.DisableParsingAndEscaping, "应该启用解析和转义")
	
	// 测试禁用
	formatter.ParsingAndEscaping(true)
	require.True(t, formatter.DisableParsingAndEscaping, "应该禁用解析和转义")
}

// TestFormatter_Caller 测试 Caller 方法
func TestFormatter_Caller(t *testing.T) {
	formatter := &Formatter{}
	
	// 测试启用调用者信息 (Caller(true) 表示启用，DisableCaller = false)
	formatter.Caller(true)
	require.False(t, formatter.DisableCaller, "应该启用调用者信息")
	
	// 测试禁用调用者信息 (Caller(false) 表示禁用，DisableCaller = true)
	formatter.Caller(false)
	require.True(t, formatter.DisableCaller, "应该禁用调用者信息")
}

// TestFormatter_Clone 测试 Clone 方法
func TestFormatter_Clone(t *testing.T) {
	original := &Formatter{
		Module:                    "test-module",
		DisableParsingAndEscaping: true,
		DisableCaller:             true,
	}
	
	cloned := original.Clone()
	require.NotNil(t, cloned, "克隆结果不应该为空")
	
	clonedFormatter, ok := cloned.(*Formatter)
	require.True(t, ok, "克隆结果应该是 *Formatter 类型")
	
	// 验证字段被正确复制
	require.Equal(t, original.Module, clonedFormatter.Module, "Module 应该被复制")
	require.Equal(t, original.DisableParsingAndEscaping, clonedFormatter.DisableParsingAndEscaping, "DisableParsingAndEscaping 应该被复制")
	require.Equal(t, original.DisableCaller, clonedFormatter.DisableCaller, "DisableCaller 应该被复制")
	
	// 验证是独立的对象
	original.Module = "modified"
	require.NotEqual(t, original.Module, clonedFormatter.Module, "修改原始对象不应该影响克隆对象")
}

// TestFormatter_EdgeCases 测试边界情况
func TestFormatter_EdgeCases(t *testing.T) {
	formatter := &Formatter{}
	
	t.Run("空消息", func(t *testing.T) {
		entry := &Entry{
			Pid:        os.Getpid(),
			Gid:        123,
			Time:       time.Now(),
			Level:      InfoLevel,
			Message:    "",
			CallerFunc: "TestFunc",
		}
		
		result := formatter.format(entry)
		require.NotNil(t, result, "即使消息为空也应该返回结果")
		require.Contains(t, string(result), "[info]", "应该包含级别")
	})

	t.Run("只有空白的消息", func(t *testing.T) {
		entry := &Entry{
			Pid:        os.Getpid(),
			Gid:        123,
			Time:       time.Now(),
			Level:      InfoLevel,
			Message:    "   \n\t  ",
			CallerFunc: "TestFunc",
		}
		
		result := formatter.format(entry)
		output := string(result)
		require.Contains(t, output, "[info]", "应该包含级别")
		// TrimSpace 应该移除所有空白，留下空字符串
	})

	t.Run("长路径处理", func(t *testing.T) {
		entry := &Entry{
			Pid:        os.Getpid(),
			Gid:        123,
			Time:       time.Now(),
			Level:      InfoLevel,
			Message:    "测试",
			File:       "/very/long/path/to/the/source/file.go",
			CallerDir:  "/very/long/path/to/the/source",
			CallerFunc: "VeryLongFunctionNameThatShouldStillWork",
			CallerLine: 999,
		}
		
		result := formatter.format(entry)
		output := string(result)
		require.Contains(t, output, "file.go:999", "应该包含文件名和行号")
		require.Contains(t, output, "VeryLongFunctionNameThatShouldStillWork", "应该包含函数名")
	})
}