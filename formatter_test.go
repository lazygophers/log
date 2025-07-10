package log

import (
	"path"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// 去除ANSI颜色代码
func stripANSICodes(s string) string {
	re := regexp.MustCompile(`\x1B\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(s, "")
}

func TestFormatter_Format(t *testing.T) {
	// 测试表格驱动测试用例
	testCases := []struct {
		name           string
		entry          Entry
		disableParsing bool
		disableCaller  bool
		expected       []string
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := &Formatter{
				DisableParsingAndEscaping: tc.disableParsing,
				DisableCaller:             tc.disableCaller,
			}
			rawResult := string(f.Format(&tc.entry))
			result := stripANSICodes(rawResult) // 去除颜色代码

			// 验证基础消息内容
			for _, expected := range tc.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("输出中未包含预期内容: %q", expected)
				}
			}

			// 验证调用者信息
			if !tc.disableCaller && tc.entry.CallerFunc != "" {
				// 根据 formatter.go 中的格式验证调用者信息
				callerPath := path.Join(tc.entry.CallerDir, path.Base(tc.entry.File)) + ":" + strconv.Itoa(tc.entry.CallerLine)
				if !strings.Contains(result, callerPath) {
					t.Errorf("输出中未包含调用者路径: %q", callerPath)
				}
				if !strings.Contains(result, tc.entry.CallerFunc) {
					t.Errorf("输出中未包含函数名: %q", tc.entry.CallerFunc)
				}
			}
		})
	}
}

func TestGetColorByLevel(t *testing.T) {
	// 测试不同日志级别的颜色映射
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
		{InfoLevel, colorGreen}, // 默认
	}

	for _, tt := range colorTests {
		t.Run(tt.level.String(), func(t *testing.T) {
			color := getColorByLevel(tt.level)
			if string(color) != string(tt.expected) {
				t.Errorf("预期颜色 %s, 实际 %s", tt.expected, color)
			}
		})
	}
}

func TestSplitPackageName(t *testing.T) {
	// 测试包路径分割
	splitTests := []struct {
		input    string
		dir      string
		function string
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
			function: "",
		},
	}

	for _, tt := range splitTests {
		t.Run(tt.input, func(t *testing.T) {
			dir, fn := SplitPackageName(tt.input)
			if dir != tt.dir || fn != tt.function {
				t.Errorf("预期 (%s, %s), 实际 (%s, %s)", tt.dir, tt.function, dir, fn)
			}
		})
	}
}
