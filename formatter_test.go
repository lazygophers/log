package log

import (
	"path"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// stripANSICodes 用于从字符串中移除所有 ANSI 颜色转义码。
// 这在测试中很有用，可以只比较文本内容而不关心颜色。
func stripANSICodes(s string) string {
	re := regexp.MustCompile(`\x1B\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(s, "")
}

// TestFormatter_Format 测试 Formatter 的 Format 方法。
// 它使用表格驱动测试来覆盖不同的场景。
func TestFormatter_Format(t *testing.T) {
	// 定义了一系列测试用例
	testCases := []struct {
		name           string   // 测试用例名称
		entry          Entry    // 输入的日志条目
		disableParsing bool     // 是否禁用消息解析
		disableCaller  bool     // 是否禁用调用者信息
		expected       []string // 预期输出中应包含的字符串片段
	}{
		{
			name: "单行消息",
			entry: Entry{
				Message: "测试单行日志",
				Level:   InfoLevel,
			},
			expected: []string{"测试单行日志"},
		},
		{
			name: "多行消息",
			entry: Entry{
				Message: "第一行\n第二行",
				Level:   InfoLevel,
			},
			expected: []string{"第一行", "第二行"},
		},
		{
			name: "禁用消息解析",
			entry: Entry{
				Message: "带\n换行符",
				Level:   InfoLevel,
			},
			disableParsing: true,
			expected:       []string{"带\n换行符"},
		},
		{
			name: "带调用者信息",
			entry: Entry{
				Message:    "带调用者",
				Level:      InfoLevel,
				File:       "example.go",
				CallerDir:  "example/pkg",
				CallerLine: 42,
				CallerFunc: "TestFunction",
			},
			expected: []string{"带调用者", "example/pkg/example.go:42", "TestFunction"},
		},
		{
			name: "空消息",
			entry: Entry{
				Message: "",
				Level:   InfoLevel,
			},
			expected: []string{""},
		},
		{
			name: "仅包含换行符的消息",
			entry: Entry{
				Message: "\n\n",
				Level:   InfoLevel,
			},
			expected: []string{"", ""},
		},
		{
			name: "带前缀和后缀的消息",
			entry: Entry{
				Message:   "带前后缀",
				Level:     InfoLevel,
				PrefixMsg: []byte("前缀:"),
				SuffixMsg: []byte(":后缀"),
			},
			expected: []string{"前缀:", "带前后缀", ":后缀"},
		},
	}

	// 遍历所有测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 初始化 Formatter
			f := &Formatter{
				DisableParsingAndEscaping: tc.disableParsing,
				DisableCaller:             tc.disableCaller,
			}
			// 调用 Format 方法并获取原始输出
			rawResult := string(f.Format(&tc.entry))
			// 去除 ANSI 颜色代码，以便进行内容比较
			result := stripANSICodes(rawResult)

			// 验证格式化后的基础消息内容是否符合预期。
			// 遍历预期的结果切片，检查每一项是否存在于格式化后的字符串中。
			for _, expected := range tc.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("输出中未包含预期内容: %q", expected)
				}
			}

			// 如果测试用例中包含了调用者信息，并且没有禁用调用者信息，则进行验证。
			if !tc.disableCaller && tc.entry.CallerFunc != "" {
				// 构造预期的调用者路径格式，例如 "example/pkg/example.go:42"。
				callerPath := path.Join(tc.entry.CallerDir, path.Base(tc.entry.File)) + ":" + strconv.Itoa(tc.entry.CallerLine)
				// 检查格式化结果中是否包含完整的调用者路径。
				if !strings.Contains(result, callerPath) {
					t.Errorf("输出中未包含调用者路径: %q", callerPath)
				}
				// 检查格式化结果中是否包含调用者函数名。
				if !strings.Contains(result, tc.entry.CallerFunc) {
					t.Errorf("输出中未包含函数名: %q", tc.entry.CallerFunc)
				}
			}
		})
	}
}

// TestGetColorByLevel 测试根据日志级别获取对应 ANSI 颜色代码的函数。
func TestGetColorByLevel(t *testing.T) {
	// 定义不同日志级别及其预期的颜色代码
	colorTests := []struct {
		level    Level
		expected []byte
	}{
		{DebugLevel, colorGreen},
		{TraceLevel, colorGreen},
		{WarnLevel, colorYellow},
		{ErrorLevel, colorRed},
		{FatalLevel, colorRed},
		{PanicLevel, colorRed},
		{InfoLevel, colorGreen}, // InfoLevel 作为默认情况
	}

	// 遍历所有颜色测试用例
	for _, tt := range colorTests {
		t.Run(tt.level.String(), func(t *testing.T) {
			// 调用 getColorByLevel 获取实际的颜色代码
			color := getColorByLevel(tt.level)
			// 验证返回的颜色是否与预期一致
			if string(color) != string(tt.expected) {
				t.Errorf("预期颜色 %s, 实际 %s", tt.expected, color)
			}
		})
	}
}

// TestSplitPackageName 测试用于分割包路径和函数名的函数。
func TestSplitPackageName(t *testing.T) {
	// 定义输入路径和预期的目录、函数名
	splitTests := []struct {
		input    string // 完整的函数路径
		dir      string // 预期的目录部分
		function string // 预期的函数名部分
	}{
		{
			input:    "github.com/user/pkg.Function",
			dir:      "user/pkg",
			function: "Function",
		},
		{
			input:    "main.init",
			dir:      "main",
			function: "init",
		},
		{
			input:    "singlepackage",
			dir:      "singlepackage",
			function: "", // 没有函数名
		},
	}

	// 遍历所有分割测试用例
	for _, tt := range splitTests {
		t.Run(tt.input, func(t *testing.T) {
			// 调用被测试的 SplitPackageName 函数
			dir, fn := SplitPackageName(tt.input)
			// 验证分割后的目录和函数名是否与预期一致
			if dir != tt.dir || fn != tt.function {
				t.Errorf("预期 (%s, %s), 实际 (%s, %s)", tt.dir, tt.function, dir, fn)
			}
		})
	}
}
